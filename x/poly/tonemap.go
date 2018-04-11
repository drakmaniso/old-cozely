// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

import (
	"strings"

	"github.com/drakmaniso/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

// Main sources are:
// http://filmicworlds.com/blog/filmic-tonemapping-operators/
// https://www.shadertoy.com/view/lslGzl
// http://duikerresearch.com/2015/09/filmic-tonemapping-for-real-time-rendering/

////////////////////////////////////////////////////////////////////////////////

// Reinhardt Operator.
// (white-preserving luma based version)
const tmReinhardt = `
#version 450 core
vec3 glam_ToneMap(vec3 color) {
	const float white = 2.0;

	float luma = dot(color, vec3(0.2126, 0.7152, 0.0722));
	float toneMappedLuma = luma * (1. + luma / (white*white)) / (1. + luma);
	return color * toneMappedLuma / luma;
}
`

// ToneMapReinhardt provides a fragment shader function to apply tone mapping at
// the end of the pipeline. In order to use it, add the following declaration in
// your fragment shader:
//
//   vec3 glam_ToneMap(vec3 Color);
func ToneMapReinhardt() gl.PipelineConfig {
	return gl.FragmentShader(strings.NewReader(tmReinhardt))
}

////////////////////////////////////////////////////////////////////////////////

// Jim Hejl approximation of a filmic curve.
// https://twitter.com/jimhejl/status/633777619998130176
const tmHejl = `
#version 450 core
vec3 glam_ToneMap(vec3 color) {
	const float w = 1.2;
	vec4 h = vec4(color, w);
	vec4 a = (1.425 * h) + 0.05;
	vec4 f = ((h * a * 0.004) / ((h * (a + 0.55) + 0.0491))) - 0.0821;
	return f.rgb / f.www;
}
`

// ToneMapHejl provides a fragment shader function to apply tone mapping at
// the end of the pipeline. In order to use it, add the following declaration in
// your fragment shader:
//
//   vec3 glam_ToneMap(vec3 Color);
func ToneMapHejl() gl.PipelineConfig {
	return gl.FragmentShader(strings.NewReader(tmHejl))
}

////////////////////////////////////////////////////////////////////////////////

// John Hable approximation of a filmic curve.
// (aka "uncharted 2")
const tmHable = `
#version 450 core
vec3 glam_ToneMap(vec3 color) {
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
`

// ToneMapHable provides a fragment shader function to apply tone mapping at
// the end of the pipeline. In order to use it, add the following declaration in
// your fragment shader:
//
//   vec3 glam_ToneMap(vec3 Color);
func ToneMapHable() gl.PipelineConfig {
	return gl.FragmentShader(strings.NewReader(tmHable))
}

////////////////////////////////////////////////////////////////////////////////

// ACES filmic tone mapping curve (simple) approximation
// Source:
// https://knarkowicz.wordpress.com/2016/01/06/aces-filmic-tone-mapping-curve/
const tmSimpleACES = `
#version 450 core
vec3 glam_ToneMap(vec3 x)
{
    float a = 2.51f;
    float b = 0.03f;
    float c = 2.43f;
    float d = 0.59f;
    float e = 0.14f;
    return clamp((x*(a*x+b))/(x*(c*x+d)+e), vec3(0), vec3(1));
}
`

// ToneMapSimpleACES provides a fragment shader function to apply tone mapping at
// the end of the pipeline. In order to use it, add the following declaration in
// your fragment shader:
//
//   vec3 glam_ToneMap(vec3 Color);
func ToneMapSimpleACES() gl.PipelineConfig {
	return gl.FragmentShader(strings.NewReader(tmSimpleACES))
}

////////////////////////////////////////////////////////////////////////////////--------------------

// Stephen Hill approximation of ACES filmic curve
const tmACES = `
#version 450 core
// Source:
// https://github.com/TheRealMJP/BakingLab/blob/master/BakingLab/ACES.hlsl
//
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

vec3 glam_ToneMap(vec3 color)
{
    color = color * ACESInputMat;

    // Apply RRT and ODT
    color = RRTAndODTFit(color);

    color = color * ACESOutputMat;

    // Clamp to [0, 1]
    color = clamp(color, vec3(0.0), vec3(1.0));

    return color;
}
`

// ToneMapACES provides a fragment shader function to apply tone mapping at
// the end of the pipeline. In order to use it, add the following declaration in
// your fragment shader:
//
//   vec3 glam_ToneMap(vec3 Color);
func ToneMapACES() gl.PipelineConfig {
	return gl.FragmentShader(strings.NewReader(tmACES))
}

////////////////////////////////////////////////////////////////////////////////
