package pixel

const vertexShader = "\n" + `#version 460 core

//------------------------------------------------------------------------------

const uint cmdIndexed = 1;
const uint cmdFullColor = 3;
const uint cmdPoint = 5;

const vec2 corners[4] = vec2[4](
	vec2(0, 0),
	vec2(1, 0),
	vec2(0, 1),
	vec2(1, 1)
);

//------------------------------------------------------------------------------

layout(std140, binding = 0) uniform ScreenUBO {
	vec2 PixelSize;
};

//------------------------------------------------------------------------------

layout(binding = 5) uniform isamplerBuffer mappings;
layout(binding = 6) uniform isamplerBuffer parameters;

//------------------------------------------------------------------------------

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location=0) flat uint Command;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) vec4 Color;
};

//------------------------------------------------------------------------------

void main(void)
{
	Command = gl_VertexID >> 2;
	int param = gl_BaseInstance;
	int vertex = gl_VertexID & 0x3;

	int x, y;
	vec2 p;
	switch (Command) {
	case cmdIndexed:
	case cmdFullColor:
		// Picture Paramters
		param += 3*gl_InstanceID;
		int m = texelFetch(parameters, param+0).r;
		x = texelFetch(parameters, param+1).r;
		y = texelFetch(parameters, param+2).r;
		// Picture Mapping in Atlas
		m *= 5;
		Bin = texelFetch(mappings, m+0).r;
		UV = vec2(texelFetch(mappings, m+1).r, texelFetch(mappings, m+2).r);
		vec2 WH = vec2(texelFetch(mappings, m+3).r, texelFetch(mappings, m+4).r);
		// Picture Position and Corner
		vec2 p = (vec2(x, y) + corners[vertex] * WH) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		UV += corners[vertex] * WH;
		break;

	case cmdPoint:
		param += 4*gl_InstanceID;
		// Point Parameters
		uint rg = texelFetch(parameters, param+0).r;
		uint ba = texelFetch(parameters, param+1).r;
		x = texelFetch(parameters, param+2).r;
		y = texelFetch(parameters, param+3).r;
		// Position
		p = (vec2(x, y) + corners[vertex] * vec2(1.5,1.5)) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		// Color
		Color = vec4(
			float(rg>>8)/float(0xFF),
			float(rg&0xFF)/float(0xFF),
			float(ba>>8)/float(0xFF),
			float(ba&0xFF)/float(0xFF)
		);
		break;
	}
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
