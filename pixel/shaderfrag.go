package pixel

const fragmentShader = "\n" + `#version 450 core

in PerVertex {
	layout(location=0) vec2 UV;
	layout(location=1) flat uint Address;
	layout(location=2) flat uint Stride;
	layout(location=3) flat uint Depth;
	layout(location=4) flat uint Tint;
};

layout(std430, binding = 1) buffer PictureBuffer {
	uint []Pixels;
};

layout(std430, binding = 2) buffer PaletteBuffer {
	vec4 Colours[256];
};

out vec4 color;

uint getByte(uint addr) {
	uint waddr = addr >> 2;
	uint w = Pixels[waddr];
	w = w >> (8 * (addr & 0x3));
	return w & 0xFF;
}

uint getPixel(uint addr, uint stride, uint x, uint y) {
	return getByte(addr + x + y*stride);
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
