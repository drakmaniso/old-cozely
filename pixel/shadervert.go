package pixel

const vertexShader = "\n" + `#version 450 core

// const vec2 PixelSize = vec2(1.0/320.0, 1.0/180.0);

layout(std140, binding = 0) uniform ScreenUBO {
	vec2 PixelSize;
};

struct Stamp {
	uint Address;
	uint WH;
	uint XY;
	uint DepthTintTrans;
	uint Mode;
};
layout(std430, binding = 0) buffer StampBuffer {
	Stamp []Stamps;
};

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location=0) flat uint Mode;
	layout(location=1) vec2 UV;
	layout(location=2) flat uint Address;
	layout(location=3) flat uint Stride;
	layout(location=4) flat uint Depth;
	layout(location=5) flat uint Tint;
};

const uint modeIndexed = 1;
const uint modeRGBA = 2;

void main(void)
{
	// Calculate index in face buffer
	uint stampIndex = gl_InstanceID;

	int w = int(Stamps[stampIndex].WH & 0xFFFF);
	int h = int(Stamps[stampIndex].WH >> 16);
	//TODO: is it useful to extend sign on width and height?
	vec2 WH = vec2(
		w | (((w & 0x8000) >> 15) * 0xFFFF0000),
		h | (((h & 0x8000) >> 15) * 0xFFFF0000)
	);
	int x = int(Stamps[stampIndex].XY & 0xFFFF);
	int y = int(Stamps[stampIndex].XY >> 16);
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

	UV = corners[gl_VertexID] * WH;
	Address = Stamps[stampIndex].Address;
	Stride = uint(WH.x);
	Depth = Stamps[stampIndex].DepthTintTrans & 0xFFFF;
	Tint = (Stamps[stampIndex].DepthTintTrans >> 16) & 0xFF;
	Mode = Stamps[stampIndex].Mode;
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
