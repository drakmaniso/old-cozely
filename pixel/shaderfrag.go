package pixel

const fragmentShader = "\n" + `#version 450 core

in PerVertex {
	layout(location=0) flat uint Mode;
	layout(location=1) vec2 UV;
	layout(location=2) flat uint Address;
	layout(location=3) flat uint Stride;
	layout(location=4) flat uint Depth;
	layout(location=5) flat uint Tint;
};

const uint modeIndexed = 1;
const uint modeRGBA = 2;

layout(binding = 0) uniform usamplerBuffer IndexedSampler;
layout(binding = 1) uniform samplerBuffer RGBASampler;

layout(std430, binding = 2) buffer PaletteBuffer {
	vec4 Colours[256];
};

out vec4 color;

int coordOf(uint addr, uint stride, uint x, uint y) {
	return int(addr + x + y*stride);
}

void main(void)
{

	if (Mode == modeIndexed) {

		uint p = texelFetch(IndexedSampler, coordOf(Address, Stride, uint(UV.x), uint(UV.y))).x;
		uint c;
		if (p == 0) {
			c = 0;
		} else {
			c = p + Tint;
			if (c > 255) {
				c -= 255;
			}
		}
		color = Colours[c];

	} else {

		color = texelFetch(RGBASampler, coordOf(Address, Stride, uint(UV.x), uint(UV.y)));

	}
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
