package vector

const fragmentShader = "\n" + `#version 460 core

//------------------------------------------------------------------------------

const uint cmdIndexed = 1;
const uint cmdFullColor = 3;
const uint cmdPoint = 5;
const uint cmdPointList = 6;
const uint cmdLine = 7;
const uint cmdLineAA = 8;

//------------------------------------------------------------------------------

in PerVertex {
	layout(location=0) flat uint Command;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) flat vec4 Color;
	layout(location=4) flat vec2 Orig;
	layout(location=5) flat float Slope;
	layout(location=6) flat bool Steep;
};

layout(origin_upper_left, pixel_center_integer) in vec4 gl_FragCoord;

//------------------------------------------------------------------------------

layout(std430, binding = 0) buffer Palette {
	vec4 Colours[256];
};

//------------------------------------------------------------------------------

out vec4 out_color;

//------------------------------------------------------------------------------

void main(void)
{
	float x, y;
	switch (Command) {
	case cmdPoint:
	case cmdPointList:
		out_color = Color;
		break;

	case cmdLine:
		x = gl_FragCoord.x - Orig.x;
		y = gl_FragCoord.y - Orig.y;
		if (
			(!Steep && y == round(Slope*x)) ||
			(Steep && x == round(Slope*y))
		) {
			out_color = Color;
		} else {
			out_color = vec4(0,0,0,0);
		}
		break;

	case cmdLineAA:
		x = gl_FragCoord.x - Orig.x;
		y = gl_FragCoord.y - Orig.y;
		float a;
		if (Steep) {
			a = 1 - abs(x - Slope*y);
		} else {
			a = 1 - abs(y - Slope*x);
		}
		// a = round(a*8)/8.0;
		out_color = Color * vec4(1,1,1,a);
		break;

	default:
		out_color = vec4(1,0,1,0.5);
	}
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
