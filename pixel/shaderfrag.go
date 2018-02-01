package pixel

const fragmentShader = "\n" + `#version 460 core

//------------------------------------------------------------------------------

const uint cmdPicture    = 1;
const uint cmdPictureExt = 2;
const uint cmdText       = 3;
const uint cmdPoint      = 4;
const uint cmdPointList  = 5;
const uint cmdLine       = 6;

//------------------------------------------------------------------------------

in PerVertex {
	layout(location=0) flat uint Command;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) flat uint ColorIndex;
	layout(location=4) flat vec2 Orig;
	layout(location=5) flat float Slope;
	layout(location=6) flat bool Steep;
};

layout(origin_upper_left, pixel_center_integer) in vec4 gl_FragCoord;

//------------------------------------------------------------------------------

layout(binding = 1) uniform usampler2DArray Pictures;
layout(binding = 2) uniform usampler2DArray Glyphs;

//------------------------------------------------------------------------------

out uint out_color;

//------------------------------------------------------------------------------

void main(void)
{
	float x, y;
	uint c = 0;
	switch (Command) {
	case cmdPicture:
		uint p = texelFetch(Pictures, ivec3(UV.x, UV.y, Bin), 0).x;
		if (p == 0) {
			c = 0;
		} else {
			c = p;// + Tint;
			if (c > 255) {
				c -= 255;
			}
		}
		break;

	case cmdText:
		p = texelFetch(Glyphs, ivec3(UV.x, UV.y, Bin), 0).x;
		if (p == 0) {
			c = 0;
		} else {
			c = p + ColorIndex;
			if (c > 255) {
				c -= 255;
			}
		}
		break;

	case cmdPoint:
	case cmdPointList:
		c = ColorIndex;
		break;

	case cmdLine:
		x = gl_FragCoord.x - Orig.x;
		y = gl_FragCoord.y - Orig.y;
		if (Steep) {
			if (x == round(Slope*y)) {
				c = ColorIndex;
			}
		} else {
			if (y == round(Slope*x)) {
				c = ColorIndex;
			}
		}
		break;
	}

	if (c == 0) {
		discard;
	}

	out_color = c;
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
