#version 450 core

in PerVertex {
	layout(location = 0) in vec3 Normal;
  layout(location = 1) flat in uint Material;
  layout(location = 2) in vec3 SurfaceToCamera;
} vertex;

out vec4 Color;

//--------------------------------------------------------------------------------------------------

// const float exposure = 0.005787;
//const float exposure = 0.00000083333;
const float exposure = 0.00004;
const vec3 sun_illuminance = vec3(37051.21);

//--------------------------------------------------------------------------------------------------

const float PI = 3.14159265358979323846;

const vec3 ambient_luminance = vec3(99.0/256.0, 155.0/256.0, 196.0/256.0) * 29000.0 / PI;
//const vec3 ambient_luminance = vec3(29000.0) / PI;


//--------------------------------------------------------------------------------------------------

const vec4 palette[] = {
  {0.6, 0.4, 0.2, 1.0},
  {0.6, 0.2, 0.4, 1.0},
  {0.4, 0.5, 0.2, 1.0},
  {0.2, 0.5, 0.4, 1.0},
  {0.4, 0.2, 0.6, 1.0},
  {0.2, 0.35, 0.55, 1.0},
};

//--------------------------------------------------------------------------------------------------

vec3 phong_lighting (vec3 L, vec3 V, vec3 N, vec3 diffuse_color, vec3 specular_color)
{
    vec3 diffuse = diffuse_color * max (0.0, dot (N, L));

    vec3 R = reflect (-L, N);
    float RdotV = max (0.0, dot (R, V));
    vec3 specular = specular_color * pow (RdotV, 10.0);

    return diffuse + specular;
}


vec3 normalized_blinn_phong_lighting (vec3 L, vec3 V, vec3 N, vec3 diffuse_color, vec3 specular_color, float roughness)
{
    vec3 H = normalize (V + L);
    float NdotH = max (0.0, dot (N, H));

    vec3 diffuse = diffuse_color * max (0.0, dot (N, L));

    vec3 specular = ((roughness + 8.0) / 8.0) * specular_color * pow (NdotH, roughness);

    return diffuse + specular;
}


vec3 minimalist_cook_torrance_lighting (vec3 L, vec3 V, vec3 N, vec3 diffuse_color, vec3 specular_color, float roughness)
{
    vec3 H = normalize (V + L);
    float NdotH = max (0.0, dot (N, H));
    float NdotL = max (0.0, dot (N, L));
    float LdotH = max (0.0, dot (L, H));

    vec3 diffuse = diffuse_color * NdotL;

    vec3 specular = specular_color * (roughness + 1) * pow (NdotH, roughness) / (8.0 * pow (LdotH, 3));

    return diffuse + specular;
}


//--------------------------------------------------------------------------------------------------


float schlick_fresnel (float f0 , float f90, float u)
{
    return f0 + (f90 - f0) * pow (1.0 - u , 5.0);
}


vec3 schlick_fresnel_rgb (vec3 f0 , float f90, float u)
{
    return f0 + (f90 - f0) * pow (1.0 - u , 5.0);
}


vec3 burley_diffuse_brdf (float NdotV, float NdotL, float LdotH, vec3 diffuse_color, float linear_roughness)
{
    // Renormalization to keep total energy (diffuse + specular) below 1.0 [LaDe14]
    float normalization_bias = mix (0.0 , 0.5 , linear_roughness);
    float normalization_factor = mix (1.0, 1.0 / 1.51, linear_roughness);

    // Burley's diffuse BRDF, aka Disney Diffuse [Burley12]
    float fd90 = normalization_bias + 2.0 * LdotH * LdotH * linear_roughness;
    float f0 = 1.0;
    float light_scatter = schlick_fresnel (f0 , fd90 , NdotL);
    float view_scatter = schlick_fresnel (f0 , fd90 , NdotV);

    return normalization_factor * diffuse_color * light_scatter * view_scatter; // Division by PI omitted (it's factored in light intensity)
}


float smith_ggx_height_correlated_visibility (float NdotL , float NdotV, float roughness)
{
    // First, reduce specular "hotness" for small roughness values [Burley12]
    float alpha = 0.5 + roughness / 2.0;

    // Height-correlated masking and shadowing function [Heitz14]
    // (Optimized version from [LaDe14])
    float alpha2 = alpha * alpha;
    float lambda_V = NdotL * sqrt ((-NdotV * alpha2 + NdotV) * NdotV + alpha2);
    float lambda_L = NdotV * sqrt ((-NdotL * alpha2 + NdotL) * NdotL + alpha2);

    return 0.5 / (lambda_V + lambda_L);
}


float ggx_ndf (float NdotH , float m)
{
    // Microfacet normal distribution function: GGX [WMLT07] aka Trowbridge-Reitz
    float m2 = m * m ;
    float f = (NdotH * m2 - NdotH) * NdotH + 1;

    return m2 / (f * f); // Division by PI omitted (it's factored in light intensity)
}


vec3 specular_brdf (float NdotV, float NdotL, float NdotH, float LdotH, vec3 f0, float roughness)
{
    float ggx_roughness = roughness * roughness;
    vec3 fresnel = schlick_fresnel_rgb (f0, 1.0, LdotH);
    float visibility = smith_ggx_height_correlated_visibility (NdotV, NdotL, ggx_roughness);
    float ndf = ggx_ndf (NdotH, ggx_roughness);

    return fresnel * visibility * ndf;
}


vec3 light_luminance (vec3 illuminance, vec3 L, vec3 V, vec3 N, vec3 base_color, vec3 f0, float roughness)
{
    float NdotV = abs (dot (N , V)) + 0.000000001;

    vec3 H = normalize (V + L);
    float LdotH = max (0.0, dot (L, H));
    float NdotH = max (0.0, dot (N, H));
    float NdotL = max (0.0, dot (N, L));

    vec3 specular = specular_brdf (NdotV, NdotL, NdotH, LdotH, f0, roughness);

    vec3 diffuse = burley_diffuse_brdf (NdotV, NdotL, LdotH, base_color, roughness);

    return illuminance * max (0.0, dot (N, L)) * (specular + diffuse); // Division by PI omitted (it's factored in light intensity)
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

  // // Simplistic diffuse lighting
  // const vec3 L = normalize(vec3(0.4, 0.6, 0.8));
  // float NdotL = dot(vertex.Normal, L);
  // float diff = clamp(NdotL, 0.2, 1.0);

	// Color = diff * palette[vertex.Material];


    vec3 N = vertex.Normal;

    // Material

    vec3 base_color;
    float smoothness  = 0.35;
    float metal_mask  = 0.00;
    float reflectance = 0.04;

    base_color = palette[vertex.Material].rgb;

    // base_color = vec3 (1.000000, 0.765557, 0.336057); // Gold
    // base_color = vec3 (0.971519, 0.959915, 0.915324); // Silver
    // base_color = vec3 (0.913183, 0.921494, 0.924524); // Aluminium
    // base_color = vec3 (0.955008, 0.637427, 0.538163); // Copper
    // base_color = vec3 (0.549585, 0.556114, 0.554256); // Chromium
    //base_color = vec3 (0.659777, 0.608679, 0.525649); // Nickel
    //base_color = vec3 (0.541931, 0.496791, 0.449419); // Titanium
    //base_color = vec3 (0.662124, 0.654864, 0.633732); // Cobalt
    //base_color = vec3 (0.672411, 0.637331, 0.585456); // Platinum
    // base_color = vec3 (0.56, 0.57, 0.58); // Iron
    //base_color = vec3 (1.00, 0.71, 0.29); // Gold
    //base_color = vec3 (0.95, 0.93, 0.88); // Silver

    float roughness = 1 - smoothness;
    vec3 f0 = mix (vec3 (reflectance), base_color, metal_mask);
    base_color = (1.0 - metal_mask) * base_color;


    // Lighting

    vec3 V = normalize (vertex.SurfaceToCamera);
    vec3 L = normalize(vec3(0.4, 0.6, 0.8)); //w_direction_to_sun;

    //vec3 luminance = (N + 1.0) / 2.0 + 0.000000001 * base_color;
    // vec3 luminance = sun_illuminance * phong_lighting (L, V, N, base_color, mix (vec3 (0.0), vec3 (0.9), smoothness * smoothness)) + base_color * ambient_luminance;
    //vec3 luminance = sun_illuminance * normalized_blinn_phong_lighting (L, V, N, base_color, f0, 1000.0 * smoothness * smoothness * smoothness * smoothness) + base_color * ambient_luminance;
    //vec3 luminance = sun_illuminance * minimalist_cook_torrance_lighting (L, V, N, base_color, f0, 500.0 * smoothness * smoothness * smoothness * smoothness) + base_color * ambient_luminance;
    vec3 luminance = light_luminance (sun_illuminance, L, V, N, base_color, f0, roughness) + base_color * ambient_luminance;

    // Dithering

    // vec3 dither = vec3 (dot (vec2 (171.0, 231.0), gl_FragCoord.xy + time));
    // dither = fract (dither / vec3(103.0, 71.0, 97.0)) - vec3(0.5, 0.5, 0.5);
    // dither = 0.75 + dither * 0.25;
    // luminance *= dither;

    Color = vec4 (luminance * exposure, 1.0);

}

//--------------------------------------------------------------------------------------------------
