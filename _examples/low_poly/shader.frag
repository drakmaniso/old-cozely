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

out vec4 Color;

//------------------------------------------------------------------------------

vec3 glam_ToneMap(vec3 Color);

//--------------------------------------------------------------------------------------------------

const float PI = 3.14159265358979323846;

const vec3 ambient_luminance = vec3(99.0/256.0, 155.0/256.0, 196.0/256.0) * 29000.0 / PI;
//const vec3 ambient_luminance = vec3(29000.0) / PI;
// const vec3 ambient_luminance = vec3(0.0);

//--------------------------------------------------------------------------------------------------

const vec4 palette[] = {
  {0.4, 0.3, 0.2, 1.0},
  {0.6, 0.26, 0.38, 1.0},
  {0.27, 0.07, 0.12, 1.0},
  {0.07, 0.07, 0.12, 1.0},
  {1.0, 1.0, 1.0, 1.0},
  {0.4, 0.5, 0.2, 1.0},
  {0.2, 0.5, 0.4, 1.0},
};

//--------------------------------------------------------------------------------------------------

// Material

vec3 glam_BaseColor;
float glam_Smoothness = 0.5;
float glam_Metallic  = 0.0;
float glam_Reflectance = 0.5;

float glam_Subsurface = 0.0;

float glam_Roughness;
vec3 glam_F0;

//--------------------------------------------------------------------------------------------------

vec3 phong_lighting (vec3 L, vec3 V, vec3 N, vec3 glam_BaseColor, vec3 specular_color)
{
    vec3 diffuse = glam_BaseColor * max (0.0, dot (N, L));

    vec3 R = reflect (-L, N);
    float RdotV = max (0.0, dot (R, V));
    vec3 specular = specular_color * pow (RdotV, 10.0);

    return diffuse + specular;
}


vec3 normalized_blinn_phong_lighting (vec3 L, vec3 V, vec3 N, vec3 glam_BaseColor, vec3 specular_color, float glam_Roughness)
{
    vec3 H = normalize (V + L);
    float NdotH = max (0.0, dot (N, H));

    vec3 diffuse = glam_BaseColor * max (0.0, dot (N, L));

    vec3 specular = ((glam_Roughness + 8.0) / 8.0) * specular_color * pow (NdotH, glam_Roughness);

    return diffuse + specular;
}


vec3 minimalist_cook_torrance_lighting (vec3 L, vec3 V, vec3 N, vec3 glam_BaseColor, vec3 specular_color, float glam_Roughness)
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
vec3 diffuseBurleyBRDF (float NdotV, float NdotL, float LdotH)
{
  // Burley's diffuse BRDF, aka Disney Diffuse [Burley12]
  float fd90 = 0.5 + 2.0 * LdotH * LdotH * glam_Roughness;
  float f0 = 1.0;
  float light_scatter = fresnel(f0 , fd90 , NdotL);
  float view_scatter = fresnel(f0 , fd90 , NdotV);

  return glam_BaseColor * light_scatter * view_scatter;
  // (division by PI omitted, it's factored in light intensity)
}

// Normalized Burley Diffuse BRDF
vec3 diffuseNormalizedBurleyBRDF (float NdotV, float NdotL, float LdotH)
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
vec3 specularBRDF (float NdotV, float NdotL, float NdotH, float LdotH, out vec3 specF)
{
  vec3 F = fresnelRGB(glam_F0, 1.0, LdotH);
  specF = F;
  float a = glam_Roughness * glam_Roughness;
  float G = geometry(NdotV, NdotL, a);
  float D = ndf(NdotH, a);

  return F * G * D;
}

//--------------------------------------------------------------------------------------------------

vec3 pbr_lightingSimple(vec3 L, vec3 V, vec3 N)
{
  float NdotV = abs(dot(N , V)) + 0.000000001; //TODO: factorize out of function

  vec3 H = normalize(V + L);
  float LdotH = max(0.0, dot(L, H));
  float NdotH = max(0.0, dot(N, H));
  float NdotL = max(0.0, dot(N, L));

  vec3 specF;
  vec3 specular = specularBRDF(NdotV, NdotL, NdotH, LdotH, specF);


  return NdotL * (specular + (1 - specF)*glam_BaseColor);
  // (division by PI omitted, it's factored in light intensity)
}

vec3 pbr_lightingBurley(vec3 L, vec3 V, vec3 N)
{
  float NdotV = abs(dot(N , V)) + 0.000000001; //TODO: factorize out of function

  vec3 H = normalize(V + L);
  float LdotH = max(0.0, dot(L, H));
  float NdotH = max(0.0, dot(N, H));
  float NdotL = max(0.0, dot(N, L));

  vec3 specF;
  vec3 specular = specularBRDF(NdotV, NdotL, NdotH, LdotH, specF);

  vec3 diffuse = diffuseBurleyBRDF(NdotV, NdotL, LdotH);

  return NdotL * (specular + diffuse);
  // (division by PI omitted, it's factored in light intensity)
}

vec3 pbr_lightingBurleySubsurface(vec3 L, vec3 V, vec3 N)
{
  float NdotV = abs(dot(N , V)) + 0.000000001; //TODO: factorize out of function

  vec3 H = normalize(V + L);
  float LdotH = max(0.0, dot(L, H));
  float NdotH = max(0.0, dot(N, H));
  float NdotL = max(0.0, dot(N, L));

  vec3 specF;
  vec3 specular = specularBRDF(NdotV, NdotL, NdotH, LdotH, specF);

  vec3 diffuse = diffuseBurleyBRDF(NdotV, NdotL, LdotH);

	float Fss90 = LdotH*LdotH*glam_Roughness;
	float FL = fresnel(1, Fss90, NdotL);
	float FV = fresnel(1, Fss90, NdotV);
	float Fss = FL*FV;
	float ss = 1.25 * (Fss * (1 / (NdotL + NdotV) - .5) + .5);

  return NdotL * (specular + mix(diffuse, ss*glam_BaseColor, glam_Subsurface));
  // (division by PI omitted, it's factored in light intensity)
}

vec3 pbr_lightingNormalizedBurley(vec3 L, vec3 V, vec3 N)
{
  float NdotV = abs(dot(N , V)) + 0.000000001; //TODO: factorize out of function

  vec3 H = normalize(V + L);
  float LdotH = max(0.0, dot(L, H));
  float NdotH = max(0.0, dot(N, H));
  float NdotL = max(0.0, dot(N, L));

  vec3 specF;
  vec3 specular = specularBRDF(NdotV, NdotL, NdotH, LdotH, specF);

  vec3 diffuse = diffuseNormalizedBurleyBRDF(NdotV, NdotL, LdotH);

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

void main(void) {

  // Material

  glam_BaseColor = palette[glam_Material].rgb;

  // glam_BaseColor = vec3 (1.000000, 0.765557, 0.336057); // Gold
  // glam_BaseColor = vec3 (0.971519, 0.959915, 0.915324); // Silver
  // glam_BaseColor = vec3 (0.913183, 0.921494, 0.924524); // Aluminium
  // glam_BaseColor = vec3 (0.955008, 0.637427, 0.538163); // Copper
  // glam_BaseColor = vec3 (0.549585, 0.556114, 0.554256); // Chromium
  // glam_BaseColor = vec3 (0.659777, 0.608679, 0.525649); // Nickel
  //glam_BaseColor = vec3 (0.541931, 0.496791, 0.449419); // Titanium
  //glam_BaseColor = vec3 (0.662124, 0.654864, 0.633732); // Cobalt
  // glam_BaseColor = vec3 (0.672411, 0.637331, 0.585456); // Platinum
  // glam_BaseColor = vec3 (0.56, 0.57, 0.58); // Iron
  //glam_BaseColor = vec3 (1.00, 0.71, 0.29); // Gold
  //glam_BaseColor = vec3 (0.95, 0.93, 0.88); // Silver

  // glam_BaseColor = vec3 (0.42, 0.72, 0.86); // Cartoonish Iron
  // glam_BaseColor = vec3 (0.99, 0.76, 0.073); // Cartoonish Gold

  glam_SetupPBR();

  // Lighting

  vec3 N = glam_Normal;
  vec3 V = glam_SurfaceToCamera;
  vec3 L = normalize(vec3(-0.4, 0.6, 0.8)); //w_direction_to_sun;

  // vec3 luminance = glam_SunIlluminance * phong_lighting (L, V, N, glam_BaseColor, mix (vec3 (0.9), vec3 (0.0), glam_Roughness * glam_Roughness)) + glam_BaseColor * ambient_luminance;
  // vec3 luminance = glam_SunIlluminance * normalized_blinn_phong_lighting (L, V, N, glam_BaseColor, glam_F0, 1000.0 * glam_Smoothness * glam_Smoothness * glam_Smoothness * glam_Smoothness) + glam_BaseColor * ambient_luminance;
  // vec3 luminance = glam_SunIlluminance * minimalist_cook_torrance_lighting (L, V, N, glam_BaseColor, glam_F0, 500.0 * glam_Smoothness * glam_Smoothness * glam_Smoothness * glam_Smoothness) + glam_BaseColor * ambient_luminance;
  // vec3 luminance = glam_SunIlluminance * pbr_lightingSimple(L, V, N) + glam_BaseColor * ambient_luminance;
  // vec3 luminance = glam_SunIlluminance * pbr_lightingBurley(L, V, N) + glam_BaseColor * ambient_luminance;
  vec3 luminance = glam_SunIlluminance * pbr_lightingBurleySubsurface(L, V, N) + glam_BaseColor * ambient_luminance;
  // vec3 luminance = glam_SunIlluminance * pbr_lightingNormalizedBurley(L, V, N) + glam_BaseColor * ambient_luminance;


  // Dithering

  // vec3 dither = vec3 (dot (vec2 (171.0, 231.0), gl_FragCoord.xy + time));
  // dither = fract (dither / vec3(103.0, 71.0, 97.0)) - vec3(0.5, 0.5, 0.5);
  // dither = 0.75 + dither * 0.25;
  // luminance *= dither;

  Color = vec4(glam_ToneMap(luminance * glam_CameraExposure), 1.0);
}

//--------------------------------------------------------------------------------------------------
