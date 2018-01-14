package pixel

const vertexShader = "\n" + `#version 450 core

layout(std140, binding = 0) uniform ScreenUBO {
	vec2 PixelSize;
};

layout(binding = 5) uniform isamplerBuffer mappings;

struct Stamp {
	uint ModeMapping;
	uint XY;
};
layout(std430, binding = 2) buffer StampBuffer {
	Stamp []Stamps;
};

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
	Stamp s = Stamps[gl_InstanceID];

	Mode = int(s.ModeMapping & 0xFFFF);

	// Picture Mapping in Atlas
	int m = 5*int(s.ModeMapping >> 16);
	Bin = texelFetch(mappings, m+0).r;
	UV = vec2(texelFetch(mappings, m+1).r, texelFetch(mappings, m+2).r);
	vec2 WH = vec2(texelFetch(mappings, m+3).r, texelFetch(mappings, m+4).r);

	// Picture Position
	int x = int(s.XY & 0xFFFF);
	int y = int(s.XY >> 16);
	vec2 XY = vec2(
		x | (((x & 0x8000) >> 15) * 0xFFFF0000),
		y | (((y & 0x8000) >> 15) * 0xFFFF0000)
	);

	// Determine which corner of the stamp this is
	const vec2 corners[4] = vec2[4](
		vec2(0, 0),
		vec2(1, 0),
		vec2(0, 1),
		vec2(1, 1)
	);
	vec2 p = (XY + corners[gl_VertexID] * WH) * PixelSize;
	gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);

	UV += corners[gl_VertexID] * WH;
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
