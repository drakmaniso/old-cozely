package pixel

const vertexShader = "\n" + `#version 460 core

//------------------------------------------------------------------------------

const uint cmdPicture    = 1;
const uint cmdPictureExt = 2;
const uint cmdText       = 3;
const uint cmdPoint      = 4;
const uint cmdLines      = 5;
const uint cmdTriangles  = 6;

const vec2 corners[4] = vec2[4](
	vec2(0, 0),
	vec2(1, 0),
	vec2(0, 1),
	vec2(1, 1)
);

//------------------------------------------------------------------------------

layout(std140, binding = 0) uniform ScreenUBO {
	vec2 PixelSize;
};

//------------------------------------------------------------------------------

layout(binding = 0) uniform isamplerBuffer parameters;
layout(binding = 1) uniform isamplerBuffer pictureMap;
layout(binding = 3) uniform isamplerBuffer glyphMap;

//------------------------------------------------------------------------------

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location=0) flat uint Command;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) flat uint ColorIndex;
	layout(location=4) flat vec2 Orig;
	layout(location=5) flat float Slope;
	layout(location=6) flat bool Steep;
};

//------------------------------------------------------------------------------

float floatZ(int z) {
	return float(z)/float(0x7FFF);
}

//------------------------------------------------------------------------------

void main(void)
{
	Command = gl_BaseInstance >> 24;
	int param = gl_BaseInstance & 0xFFFFFF;
	int offset;
	int instance = gl_InstanceID;
	int vertex = gl_VertexID;

	int x, y, z, x2, y2, x3, y3, dx, dy;
	uint c;
	vec2 p, wh;
	vec2 t, n, pts[4];
	switch (Command) {
	case cmdPicture:
		// Parameters
		offset = 4*instance;
		int m = texelFetch(parameters, param+0+offset).r;
		x = texelFetch(parameters, param+1+offset).r;
		y = texelFetch(parameters, param+2+offset).r;
		z = texelFetch(parameters, param+3+offset).r;
		// Mapping of the picture
		m *= 5;
		Bin = texelFetch(pictureMap, m+0).r;
		UV = vec2(texelFetch(pictureMap, m+1).r, texelFetch(pictureMap, m+2).r);
		wh = vec2(texelFetch(pictureMap, m+3).r, texelFetch(pictureMap, m+4).r);
		// Picture quad
		p = (vec2(x, y) + corners[vertex] * wh) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		UV += corners[vertex] * wh;
		break;

	case cmdText:
	  // Parameters of the whole Print command
		uint f = texelFetch(parameters, param+0).r;
		c = texelFetch(parameters, param+1).r;
		x = texelFetch(parameters, param+2).r;
		y = texelFetch(parameters, param+3).r;
		z = texelFetch(parameters, param+4).r;
		// Parameter for the current character
		offset = 2*instance;
		int r = texelFetch(parameters, param+5+offset).r;
		dx = texelFetch(parameters, param+6+offset).r;
		// Mapping of the current character
		r *= 5;
		Bin = texelFetch(glyphMap, r+0).r;
		UV = vec2(texelFetch(glyphMap, r+1).r, texelFetch(glyphMap, r+2).r);
		wh = vec2(texelFetch(glyphMap, r+3).r, texelFetch(glyphMap, r+4).r);
		// Character quad
		p = (vec2(x+dx, y) + corners[vertex] * wh) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		UV += corners[vertex] * wh;
		ColorIndex = uint(c);
		break;

	case cmdPoint:
		offset = 4*instance;
		// Parameters
		c = texelFetch(parameters, param+0+offset).r;
		x = texelFetch(parameters, param+1+offset).r;
		y = texelFetch(parameters, param+2+offset).r;
		z = texelFetch(parameters, param+3+offset).r;
		// Position
		p = (vec2(x, y) + corners[vertex] * vec2(1.5,1.5)) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		// Color
		ColorIndex = uint(c);
		break;

	case cmdLines:
		offset = 2*instance;
		// Parameters
		c = texelFetch(parameters, param+0).r;
		z = texelFetch(parameters, param+1).r;
		x = texelFetch(parameters, param+2+offset).r;
		y = texelFetch(parameters, param+3+offset).r;
		x2 = texelFetch(parameters, param+4+offset).r;
		y2 = texelFetch(parameters, param+5+offset).r;
		// Position
		Orig = vec2(x, y);
		dx = x2-x;
		dy = y2-y;
		Steep = abs(dx) < abs(dy);
		t = 0.25*normalize(vec2(dx, dy));
		n = 0.75*normalize(vec2(-dy, dx));
		pts = vec2[4](
			vec2(x, y)+n-t,
			vec2(x, y)-n-t,
			vec2(x2, y2)+n+t,
			vec2(x2, y2)-n+t
		);
		p = (vec2(0.5,0.5) + pts[vertex].xy) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		if (Steep) {
			Slope = float(dx)/float(dy);
		} else {
			Slope = float(dy)/float(dx);
		}
		// Color
		ColorIndex = uint(c);
		break;

	case cmdTriangles:
		offset = 2*gl_VertexID;
		// Parameters
		c = texelFetch(parameters, param+0).r;
		z = texelFetch(parameters, param+1).r;
		x = texelFetch(parameters, param+2+offset).r;
		y = texelFetch(parameters, param+3+offset).r;
		p = (vec2(0.5,0.5) + vec2(x, y)) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		// Color
		ColorIndex = uint(c);
		break;

	}
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
