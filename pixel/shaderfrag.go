package pixel

const fragmentShader = "\n" + `#version 460 core

//------------------------------------------------------------------------------

const uint cmdIndexed = 1;
const uint cmdFullColor = 3;
const uint cmdPoint = 5;

//------------------------------------------------------------------------------

in PerVertex {
	layout(location=0) flat uint Command;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) vec4 Color;
};

//------------------------------------------------------------------------------

layout(binding = 1) uniform usampler2DArray IndexedTextures;
layout(binding = 2) uniform sampler2DArray FullColorTextures;

layout(std430, binding = 0) buffer Palette {
	vec4 Colours[256];
};

//------------------------------------------------------------------------------

out vec4 out_color;

//------------------------------------------------------------------------------

void main(void)
{

	switch (Command) {
	case cmdIndexed:
		uint p = texelFetch(IndexedTextures, ivec3(UV.x, UV.y, 0), 0).x;
		uint c;
		if (p == 0) {
			c = 0;
		} else {
			c = p;// + Tint;
			if (c > 255) {
				c -= 255;
			}
		}
		out_color = Colours[c];
		break;

	case cmdFullColor:
		out_color = texelFetch(FullColorTextures, ivec3(UV.x, UV.y, 0), 0);
		break;

	case cmdPoint:
		out_color = Color;
		break;

	default:
		out_color = vec4(1,0,1,0.5);
	}
}
`

// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
