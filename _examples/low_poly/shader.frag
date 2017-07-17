#version 450 core

//------------------------------------------------------------------------------

layout(std140, binding = 0) uniform glam_Camera {
	mat4  glam_ProjectionView;
	vec3  glam_CameraPosition;
  float glam_CameraExposure;
};

layout(std140, binding = 1) uniform Misc {
	mat4  glam_Model;
	vec3  glam_SunIlluminance;
  float glam_unused1;
};

//------------------------------------------------------------------------------

in glam_PerVertex {
	layout(location = 0)      vec3 glam_Normal;
  layout(location = 1)      vec3 glam_SurfaceToCamera;
  layout(location = 2) flat uint glam_Material;
};

//------------------------------------------------------------------------------

out vec3 Color;

//--------------------------------------------------------------------------------------------------

const float PI = 3.14159265358979323846;

const vec3 ambient_luminance = vec3(99.0/256.0, 155.0/256.0, 196.0/256.0) * 29000.0 / PI;
//const vec3 ambient_luminance = vec3(29000.0) / PI;
// const vec3 ambient_luminance = vec3(0.0);

//--------------------------------------------------------------------------------------------------

const vec4 palette[] = {
  {0.6, 0.4, 0.2, 1.0},
  {0.6, 0.2, 0.4, 1.0},
  {0.4, 0.5, 0.2, 1.0},
  {0.2, 0.5, 0.4, 1.0},
  {0.4, 0.2, 0.6, 1.0},
  {0.2, 0.35, 0.55, 1.0},
  {1.0, 1.0, 1.0, 1.0},
};

//--------------------------------------------------------------------------------------------------

// Material

vec3 glam_BaseColor;
float glam_Smoothness = 0.5;
float glam_Metallic  = 0.0;
float glam_Reflectance = 0.5;

float glam_Roughness;
vec3 glam_F0;

//--------------------------------------------------------------------------------------------------

vec3 phong_lighting (vec3 illuminance, vec3 L, vec3 V, vec3 N, vec3 glam_BaseColor, vec3 specular_color)
{
    vec3 diffuse = glam_BaseColor * max (0.0, dot (N, L));

    vec3 R = reflect (-L, N);
    float RdotV = max (0.0, dot (R, V));
    vec3 specular = specular_color * pow (RdotV, 10.0);

    return diffuse + specular;
}


vec3 normalized_blinn_phong_lighting (vec3 illuminance, vec3 L, vec3 V, vec3 N, vec3 glam_BaseColor, vec3 specular_color, float glam_Roughness)
{
    vec3 H = normalize (V + L);
    float NdotH = max (0.0, dot (N, H));

    vec3 diffuse = glam_BaseColor * max (0.0, dot (N, L));

    vec3 specular = ((glam_Roughness + 8.0) / 8.0) * specular_color * pow (NdotH, glam_Roughness);

    return diffuse + specular;
}


vec3 minimalist_cook_torrance_lighting (vec3 illuminance, vec3 L, vec3 V, vec3 N, vec3 glam_BaseColor, vec3 specular_color, float glam_Roughness)
{
    vec3 H = normalize (V + L);
    float NdotH = max (0.0, dot (N, H));
    float NdotL = max (0.0, dot (N, L));
    float LdotH = max (0.0, dot (L, H));

    vec3 diffuse = glam_BaseColor * NdotL;

    vec3 specular = specular_color * (glam_Roughness + 1) * pow (NdotH, glam_Roughness) / (8.0 * pow (LdotH, 3));

    return diffuse + specular;
}


//--------------------------------------------------------------------------------------------------

void glam_SetupPBR() {
  glam_Roughness = 1.0 - glam_Smoothness;
  glam_F0 = mix (vec3 (0.16*glam_Reflectance*glam_Reflectance), glam_BaseColor, glam_Metallic);
  glam_BaseColor = (1.0 - glam_Metallic) * glam_BaseColor;
}

//--------------------------------------------------------------------------------------------------

// Fresnel: Schlick approximation
float fresnel(float f0 , float f90, float u)
{
  return f0 + (f90 - f0) * pow (1.0 - u , 5.0);
}

vec3 fresnelRGB (vec3 f0 , float f90, float u)
{
  return f0 + (f90 - f0) * pow (1.0 - u , 5.0);
}

//--------------------------------------------------------------------------------------------------

// Burley Diffuse BRDF
vec3 diffuseBRDF (float NdotV, float NdotL, float LdotH)
{
  // Renormalization to keep total energy (diffuse + specular) below 1.0 [LaDe14]
  float normalization_bias = mix (0.0 , 0.5 , glam_Roughness);
  float normalization_factor = mix (1.0, 1.0 / 1.51, glam_Roughness);

  // Burley's diffuse BRDF, aka Disney Diffuse [Burley12]
  float fd90 = normalization_bias + 2.0 * LdotH * LdotH * glam_Roughness;
  float f0 = 1.0;
  float light_scatter = fresnel(f0 , fd90 , NdotL);
  float view_scatter = fresnel(f0 , fd90 , NdotV);

  return normalization_factor * glam_BaseColor * light_scatter * view_scatter;
  // (division by PI omitted, it's factored in light intensity)
}

//--------------------------------------------------------------------------------------------------

// Geometry function: Height-correlated Smith GGX visibility [Heitz14]
float geometry(float NdotL , float NdotV, float a)
{
  // Original formulation of Smith GGX height-correlated:
  // lambdaV = 0.5 * (-1 + sqrt(a2 * (1-NdotL2)/NdotL2 + 1))
  // lambdaL = 0.5 * (-1 + sqrt(a2 * (1-NdotV2)/NdotV2 + 1))
  // G = 1 / (1 + lambdaV + lambdaL)
  // V = G / (4 * NdotL * NdotV)
  //
  // Optimized version from [LaDe14]:
  float a2 = a * a;
  float lambdaV = NdotL * sqrt ((-NdotV * a2 + NdotV) * NdotV + a2);
  float lambdaL = NdotV * sqrt ((-NdotL * a2 + NdotL) * NdotL + a2);

  return 0.5 / (lambdaV + lambdaL);
}

//--------------------------------------------------------------------------------------------------

// Normal distribution function: GGX [WMLT07], aka Trowbridge-Reitz
float ndf(float NdotH, float a)
{
  float a2 = a * a ;
  float f = (NdotH * a2 - NdotH) * NdotH + 1;

  return a2 / (f * f);
  // (division by PI omitted, it's factored in light intensity)
}

//--------------------------------------------------------------------------------------------------

// Specular BRDF
vec3 specularBRDF (float NdotV, float NdotL, float NdotH, float LdotH)
{
  vec3 F = fresnelRGB(glam_F0, 1.0, LdotH);
  float a = glam_Roughness * glam_Roughness;
  float G = geometry(NdotV, NdotL, a);
  float D = ndf(NdotH, a);

  return F * G * D;
}

//--------------------------------------------------------------------------------------------------

vec3 lightLuminance(vec3 L, vec3 V, vec3 N)
{
  float NdotV = abs(dot(N , V)) + 0.000000001; //TODO: factorize out of function

  vec3 H = normalize(V + L);
  float LdotH = max(0.0, dot(L, H));
  float NdotH = max(0.0, dot(N, H));
  float NdotL = max(0.0, dot(N, L));

  vec3 specular = specularBRDF(NdotV, NdotL, NdotH, LdotH);

  vec3 diffuse = diffuseBRDF(NdotV, NdotL, LdotH);

  return NdotL * (specular + diffuse);
  // (division by PI omitted, it's factored in light intensity)
}


// [Burley12] "Physically Based Shading at Disney"
//            Brent Burley (SIGGRAPH 2012)
//            http://blog.selfshadow.com/publications/s2012-shading-course/burley/s2012_pbs_disney_brdf_notes_v3.pdf
//
// [Heitz14] "Understanding the Masking-Shadowing Function in Microfacet-Based BRDFs"
//           Eric Heitz (2014)
//
// [LaDe14]  "Moving Frostbite to Physically Based Rendering"
//           Sebastien Lagarde, Charles de Rousiers (SIGGRAPH 2014)
//           http://www.frostbite.com/2014/11/moving-frostbite-to-pbr/
//
// [WMLT07] "Microfacet Models for Refraction through Rough Surfaces"
//          Bruce Walter, Stephen R. Marschner, Hongsong Li, Kenneth E. Torrance (2007)


//--------------------------------------------------------------------------------------------------

// Main sources are:
// http://filmicworlds.com/blog/filmic-tonemapping-operators/
// https://www.shadertoy.com/view/lslGzl
// http://duikerresearch.com/2015/09/filmic-tonemapping-for-real-time-rendering/

// Reinhardt Operator.
// (luma based version)
vec3 tonemapReinhardt(vec3 color) {
	const float white = 2.0;

	float luma = dot(color, vec3(0.2126, 0.7152, 0.0722));
	float toneMappedLuma = luma * (1. + luma / (white*white)) / (1. + luma);
	return color * toneMappedLuma / luma;
}

// Jim Hejl and Richard Burgess-Dawson approximation of a filmic curve.
// Note that it incorporates the gamma correction (i.e. don't use with sRGB buffers)
vec3 tonemapHejlBurgessDawson(vec3 color) {
  color = max(vec3(0.0), color - vec3(0.004));
  return (color * (6.2 * color + 0.5)) / (color * (6.2 * color + 1.7) + 0.06);
}

// John Hable approximation of a filmic curve.
// aka "uncharted 2"
vec3 tonemapHable(vec3 color) {
	const float W = 11.2;

	const float A = 0.15;
	const float B = 0.50;
	const float C = 0.10;
	const float D = 0.20;
	const float E = 0.02;
	const float F = 0.30;

	color = ((color * (A * color + C * B) + D * E) / (color * (A * color + B) + D * F)) - E / F;
	float white = ((W * (A * W + C * B) + D * E) / (W * (A * W + B) + D * F)) - E / F;

  return color / white;
}

// aka RomBinDaHouse
vec3 tonemapGalashov(vec3 color) {
  const float exposureBias = 2.72;
  const float blackPoint = 0.15;

  return exp(-1.0 / (exposureBias*color + blackPoint));
}

// Found in the comments of:
// https://mynameismjp.wordpress.com/2010/04/30/a-closer-look-at-tone-mapping/
vec3 tonemapSteveM(vec3 color) {
  const float exposureBias = 1.0;

	float a = 1.8; /// Mid
  float b = 1.4; /// Toe
  float c = 0.5; /// Shoulder
  float d = 1.5; /// Mid

  color *= exposureBias;
  return (color * (a * color + b)) / (color * (a * color + c) + d);
}

//--------------------------------------------------------------------------------------------------

// From https://knarkowicz.wordpress.com/2016/01/06/aces-filmic-tone-mapping-curve/
vec3 ACESFilm( vec3 x )
{
    float a = 2.51f;
    float b = 0.03f;
    float c = 2.43f;
    float d = 0.59f;
    float e = 0.14f;
    return clamp((x*(a*x+b))/(x*(c*x+d)+e), vec3(0), vec3(1));
}

//--------------------------------------------------------------------------------------------------

vec3 ff_filmic_gamma3(vec3 linear) {
    vec3 x = max(vec3(0.0), linear-0.004);
    return (x*(x*6.2+0.5))/(x*(x*6.2+1.7)+0.06);
}

//--------------------------------------------------------------------------------------------------


// The code in this file was originally written by Stephen Hill (@self_shadow), who deserves all
// credit for coming up with this fit and implementing it. Buy him a beer next time you see him. :)

// sRGB => XYZ => D65_2_D60 => AP1 => RRT_SAT
const mat3 ACESInputMat =
{
    {0.59719, 0.35458, 0.04823},
    {0.07600, 0.90834, 0.01566},
    {0.02840, 0.13383, 0.83777}
};
const mat3 ACESInputMatGL =
{
    {0.59719, 0.07600, 0.02840},
    {0.35458, 0.90834, 0.13383},
    {0.04823, 0.01566, 0.83777}
};

// ODT_SAT => XYZ => D60_2_D65 => sRGB
const mat3 ACESOutputMat =
{
    { 1.60475, -0.53108, -0.07367},
    {-0.10208,  1.10813, -0.00605},
    {-0.00327, -0.07276,  1.07602}
};
const mat3 ACESOutputMatGL =
{
    { 1.60475, -0.10208, -0.00327},
    {-0.53108, 1.10813, -0.07276},
    {-0.07367, -0.00605, 1.07602}
};

vec3 RRTAndODTFit(vec3 v)
{
    vec3 a = v * (v + 0.0245786) - 0.000090537;
    vec3 b = v * (0.983729 * v + 0.4329510) + 0.238081;
    return a / b;
}

vec3 ACESFitted(vec3 color)
{
    color = color * ACESInputMat;

    // Apply RRT and ODT
    color = RRTAndODTFit(color);

    color = color * ACESOutputMat;

    // Clamp to [0, 1]
    color = clamp(color, vec3(0.0), vec3(1.0));

    return color;
}
//--------------------------------------------------------------------------------------------------

void main(void) {

  // Material

  glam_BaseColor = palette[glam_Material].rgb;

  // glam_BaseColor = vec3 (1.000000, 0.765557, 0.336057); // Gold
  // glam_BaseColor = vec3 (0.971519, 0.959915, 0.915324); // Silver
  // glam_BaseColor = vec3 (0.913183, 0.921494, 0.924524); // Aluminium
  // glam_BaseColor = vec3 (0.955008, 0.637427, 0.538163); // Copper
  // glam_BaseColor = vec3 (0.549585, 0.556114, 0.554256); // Chromium
  //glam_BaseColor = vec3 (0.659777, 0.608679, 0.525649); // Nickel
  //glam_BaseColor = vec3 (0.541931, 0.496791, 0.449419); // Titanium
  //glam_BaseColor = vec3 (0.662124, 0.654864, 0.633732); // Cobalt
  //glam_BaseColor = vec3 (0.672411, 0.637331, 0.585456); // Platinum
  // glam_BaseColor = vec3 (0.56, 0.57, 0.58); // Iron
  //glam_BaseColor = vec3 (1.00, 0.71, 0.29); // Gold
  //glam_BaseColor = vec3 (0.95, 0.93, 0.88); // Silver

  // glam_BaseColor = vec3 (0.42, 0.72, 0.86); // Cartoonish Iron

  glam_SetupPBR();

  // Lighting

  vec3 N = glam_Normal;
  vec3 V = glam_SurfaceToCamera;
  vec3 L = normalize(vec3(-0.4, 0.6, 0.8)); //w_direction_to_sun;

  // vec3 luminance = glam_SunIlluminance * phong_lighting (L, V, N, glam_BaseColor, mix (vec3 (0.9), vec3 (0.0), glam_Roughness * glam_Roughness)) + glam_BaseColor * ambient_luminance;
  // vec3 luminance = glam_SunIlluminance * normalized_blinn_phong_lighting (L, V, N, glam_BaseColor, glam_F0, 1000.0 * smoothness * smoothness * smoothness * smoothness) + glam_BaseColor * ambient_luminance;
  // vec3 luminance = glam_SunIlluminance * minimalist_cook_torrance_lighting (L, V, N, glam_BaseColor, glam_F0, 500.0 * smoothness * smoothness * smoothness * smoothness) + glam_BaseColor * ambient_luminance;
  vec3 luminance = glam_SunIlluminance * lightLuminance(L, V, N) + glam_BaseColor * ambient_luminance;


  // Dithering

  // vec3 dither = vec3 (dot (vec2 (171.0, 231.0), gl_FragCoord.xy + time));
  // dither = fract (dither / vec3(103.0, 71.0, 97.0)) - vec3(0.5, 0.5, 0.5);
  // dither = 0.75 + dither * 0.25;
  // luminance *= dither;

  Color = luminance * glam_CameraExposure;

  // Color = tonemapGalashov(Color);
  // Color = tonemapReinhardt(Color);
  // Color = tonemapHejlBurgessDawson(Color);
  // Color = tonemapHable(Color);
  // Color = tonemapSteveM(Color);
  Color = ACESFitted(Color);
  // Color = ACESFilm(Color);
  // Color = ff_filmic_gamma3(Color);
}

//--------------------------------------------------------------------------------------------------
