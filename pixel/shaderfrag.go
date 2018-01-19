package pixel

const fragmentShader = "\n" + `#version 460 core

in PerVertex {
	layout(location=0) flat uint Mode;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
};

const uint Indexed = 1;
const uint FullColor = 2;

layout(binding = 1) uniform usampler2DArray IndexedSampler;
layout(binding = 2) uniform sampler2DArray RGBASampler;

layout(std430, binding = 0) buffer PaletteBuffer {
	vec4 Colours[256];
};

out vec4 color;

void main(void)
{

	if (Mode == Indexed) {

		uint p = texelFetch(IndexedSampler, ivec3(UV.x, UV.y, 0), 0).x;
		uint c;
		if (p == 0) {
			c = 0;
		} else {
			c = p;// + Tint;
			if (c > 255) {
				c -= 255;
			}
		}
		color = Colours[c];

	} else {

		color = texelFetch(RGBASampler, ivec3(UV.x, UV.y, 0), 0);
	}
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
