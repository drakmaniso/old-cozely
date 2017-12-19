package pixel

const fragmentShader = "\n" + `#version 450 core

in PerVertex {
	layout(location=0) flat uint Mode;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) flat uint Depth;
	layout(location=4) flat uint Tint;
};

const uint modeIndexed = 1;
const uint modeRGBA = 2;

layout(binding = 0) uniform usampler2DArray IndexedSampler;
// layout(binding = 1) uniform samplerBuffer RGBASampler;
layout(binding = 1) uniform sampler2DArray RGBASampler;

layout(std430, binding = 2) buffer PaletteBuffer {
	vec4 Colours[256];
};

out vec4 color;

// int coordOf(uint addr, uint stride, uint x, uint y) {
// 	return int(addr + x + y*stride);
// }

void main(void)
{

	if (Mode == modeIndexed) {

		uint p = texelFetch(IndexedSampler, ivec3(UV.x, UV.y, 0), 0).x;
		//texelFetch(IndexedSampler, coordOf(Address, Stride, uint(UV.x), uint(UV.y))).x;
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
		// color += vec4(0.5, 0, 0.5, 0.25);
		// color = vec4(float(p)/8.0,float(p)/8.0,float(p)/8.0,1);

	} else {

		color = texelFetch(RGBASampler, ivec3(UV.x, UV.y, 0), 0);
		// color += vec4(0.5, 0, 0.5, 0.25);
		// color = vec4(UV.x/16.0, UV.y/16.0, 0, 1);
	}
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
