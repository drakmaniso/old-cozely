package pixel

const vertexShader = "\n" + `#version 450 core

// const vec2 PixelSize = vec2(1.0/320.0, 1.0/180.0);

layout(std140, binding = 0) uniform ScreenUBO {
	vec2 PixelSize;
};

struct Stamp {
	uint ModeBin;
	uint XY;
	uint WH;
	uint UV;
};
layout(std430, binding = 0) buffer StampBuffer {
	Stamp []Stamps;
};

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location=0) flat uint Mode;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) flat uint Depth;
	layout(location=4) flat uint Tint;
};

const uint modeIndexed = 1;
const uint modeRGBA = 2;

void main(void)
{
	Stamp s = Stamps[gl_InstanceID];

	int x = int(s.XY & 0xFFFF);
	int y = int(s.XY >> 16);
	vec2 XY = vec2(
		x | (((x & 0x8000) >> 15) * 0xFFFF0000),
		y | (((y & 0x8000) >> 15) * 0xFFFF0000)
	);
	int w = int(s.WH & 0xFFFF);
	int h = int(s.WH >> 16);
	//TODO: is it useful to extend sign on width and height?
	vec2 WH = vec2(
		w | (((w & 0x8000) >> 15) * 0xFFFF0000),
		h | (((h & 0x8000) >> 15) * 0xFFFF0000)
	);
	int u = int(s.UV & 0xFFFF);
	int v = int(s.UV >> 16);
	//TODO: is it useful to extend sign on width and height?
	UV = vec2(
		u | (((u & 0x8000) >> 15) * 0xFFFF0000),
		v | (((v & 0x8000) >> 15) * 0xFFFF0000)
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
	Depth = 0;
	Tint = 0; //(s.DepthTintTrans >> 16) & 0xFF;
	Mode = int(s.ModeBin & 0xFFFF);
	Bin = int(s.ModeBin >> 16);
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
