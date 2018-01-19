package pixel

const vertexShader = "\n" + `#version 460 core

//------------------------------------------------------------------------------

const uint cmdIndexed = 1;
const uint cmdFullColor = 3;

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
};

//------------------------------------------------------------------------------

void main(void)
{
	Command = gl_VertexID >> 2;
	int param = gl_BaseInstance;

	switch (Command) {
	case cmdIndexed:
	case cmdFullColor:
		int m = texelFetch(parameters, param+0).r;
		int x = texelFetch(parameters, param+1).r;
		int y = texelFetch(parameters, param+2).r;

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
		break;
	}
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
