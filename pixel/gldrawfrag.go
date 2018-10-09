package pixel

const drawFragmentShader = "\n" + `#version 460 core

////////////////////////////////////////////////////////////////////////////////

const uint cmdPicture    = 1;
const uint cmdTriangle   = 2;
const uint cmdLine       = 3;
const uint cmdBox        = 4;
const uint cmdPoint      = 5;

////////////////////////////////////////////////////////////////////////////////

in PerVertex {
	layout(location=0) flat uint Command;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) flat uint ColorParam;
	layout(location=4) flat vec4 Box;
	layout(location=5) flat float Slope;
	layout(location=6) flat uint Flags;
};

const uint steep = 0x01;

layout(origin_upper_left, pixel_center_integer) in vec4 gl_FragCoord;

////////////////////////////////////////////////////////////////////////////////

layout(binding = 3) uniform usampler2DArray Pictures;

////////////////////////////////////////////////////////////////////////////////

out uint Color;

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
			c = p + ColorParam;
			if (c > 255) {
				c -= 255;
			}
		}
		break;

	case cmdPoint:
	case cmdTriangle:
		c = ColorParam;
		break;

	case cmdLine:
		x = gl_FragCoord.x - Box.x;
		y = gl_FragCoord.y - Box.y;
		if (Flags == steep) {
			if (x == round(Slope*y)) {
				c = ColorParam;
			}
		} else {
			if (y == round(Slope*x)) {
				c = ColorParam;
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
			c = ColorParam>>8;
		} else {
			c = ColorParam&0xFF;
		}

		break;
	}

	if (c == 0) {
		discard;
	}

	Color = c;
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
////////////////////////////////////////////////////////////////////////////////
