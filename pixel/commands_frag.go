package pixel

const fragmentShader = "\n" + `#version 460 core

////////////////////////////////////////////////////////////////////////////////

const uint cmdPicture    = 1;
const uint cmdPictureExt = 2;
const uint cmdText       = 3;
const uint cmdPoint      = 4;
const uint cmdLines      = 5;
const uint cmdTriangles  = 6;
const uint cmdBox        = 7;

////////////////////////////////////////////////////////////////////////////////

in PerVertex {
	layout(location=0) flat uint Command;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) flat uint ColorIndex;
	layout(location=4) flat vec4 Box;
	layout(location=5) flat float Slope;
	layout(location=6) flat uint Flags;
};

const uint steep = 0x01;

layout(origin_upper_left, pixel_center_integer) in vec4 gl_FragCoord;

////////////////////////////////////////////////////////////////////////////////

layout(binding = 4) uniform usampler2DArray Pictures;
layout(binding = 6) uniform usampler2DArray Glyphs;

////////////////////////////////////////////////////////////////////////////////

layout(location = 0) out vec4 out_color;
layout(location = 1) out vec4 out_filter;

////////////////////////////////////////////////////////////////////////////////

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
	case cmdTriangles:
		c = ColorIndex;
		break;

	case cmdLines:
		x = gl_FragCoord.x - Box.x;
		y = gl_FragCoord.y - Box.y;
		if (Flags == steep) {
			if (x == round(Slope*y)) {
				c = ColorIndex;
			}
		} else {
			if (y == round(Slope*x)) {
				c = ColorIndex;
			}
		}
		break;

	case cmdBox:
		x = gl_FragCoord.x;
		y = gl_FragCoord.y;
		uint cor = Flags;
		float dx = min(x-Box.x, Box.z-x);
		float dy = min(y-Box.y, Box.w-y);
		if (dx + dy < cor) {
			c = 0;
		}	else if (dx + dy == cor || dx < 1 || dy < 1) {
			c = ColorIndex>>8;
		} else {
			c = ColorIndex&0xFF;
		}

		break;
	}

	if (c == 0) {
		discard;
	}

	if (c != 999) {
		out_color = vec4(float(c)/255.0, 0, 0, 1);
		out_filter = vec4(0, 0, 0, 1);
	} else {
		out_color = vec4(float(c)/255.0, 0, 0, 0);
		out_filter = vec4(1, 0, 0, 1);
	}
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
////////////////////////////////////////////////////////////////////////////////
