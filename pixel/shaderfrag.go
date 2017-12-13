package pixel

const fragmentShader = "\n" + `#version 450 core

in PerVertex {
	layout(location=0) vec2 UV;
	layout(location=1) flat uint Address;
	layout(location=2) flat uint Stride;
	layout(location=3) flat uint Depth;
	layout(location=4) flat uint Tint;
};


layout(binding = 1) uniform usamplerBuffer PictureSampler;

layout(std430, binding = 2) buffer PaletteBuffer {
	vec4 Colours[256];
};

out vec4 color;

uint getPixel(uint addr, uint stride, uint x, uint y) {
	return texelFetch(PictureSampler, int(addr + x + y*stride)).x;
}

float rand(vec2 c){
	return fract(sin(dot(c ,vec2(12.9898,78.233))) * 43758.5453);
}

void main(void)
{
	uint p = getPixel(Address, Stride, uint(UV.x), uint(UV.y));

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
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
