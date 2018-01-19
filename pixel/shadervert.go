package pixel

const vertexShader = "\n" + `#version 460 core

layout(std140, binding = 0) uniform ScreenUBO {
	vec2 PixelSize;
};

layout(binding = 5) uniform isamplerBuffer mappings;
layout(binding = 6) uniform isamplerBuffer parameters;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location=0) flat uint Mode;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
};

const uint modeIndexed = 1;
const uint modeRGBA = 2;

void main(void)
{
	Mode = gl_VertexID >> 2;
	int prm = gl_BaseInstance;

	int m = texelFetch(parameters, prm+0).r;
	int x = texelFetch(parameters, prm+1).r;
	int y = texelFetch(parameters, prm+2).r;

	// Mode = int(s.ModeMapping & 0xFFFF);

	// Picture Mapping in Atlas
	m *= 5;
	Bin = texelFetch(mappings, m+0).r;
	UV = vec2(texelFetch(mappings, m+1).r, texelFetch(mappings, m+2).r);
	vec2 WH = vec2(texelFetch(mappings, m+3).r, texelFetch(mappings, m+4).r);

	// Picture Position
	vec2 XY = vec2(x, y);

	// Determine which corner of the stamp this is
	const vec2 corners[4] = vec2[4](
		vec2(0, 0),
		vec2(1, 0),
		vec2(0, 1),
		vec2(1, 1)
	);
	int vrt = gl_VertexID & 0x3;
	vec2 p = (XY + corners[vrt] * WH) * PixelSize;
	gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);

	UV += corners[vrt] * WH;
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
