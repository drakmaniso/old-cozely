package pixel

const vertexShader = "\n" + `#version 460 core

//------------------------------------------------------------------------------

const uint cmdPicture    = 1;
const uint cmdPictureExt = 2;
const uint cmdText       = 3;
const uint cmdPoint      = 4;
const uint cmdLine       = 5;

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
	Command = gl_VertexID >> 2;
	int param = gl_BaseInstance;
	int vertex = gl_VertexID & 0x3;

	int x, y, z, x1, y1, z1, x2, y2, z2, dx, dy;
	uint c;
	vec2 p, wh;
	vec3 v, ov, pts[4];
	switch (Command) {
	case cmdPicture:
		// Parameters
		param += 4*gl_InstanceID;
		int m = texelFetch(parameters, param+0).r;
		x = texelFetch(parameters, param+1).r;
		y = texelFetch(parameters, param+2).r;
		z = texelFetch(parameters, param+3).r;
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
		int r = texelFetch(parameters, param+5 + 2*gl_InstanceID+0).r;
		dx = texelFetch(parameters, param+5 + 2*gl_InstanceID+1).r;
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
		param += 4*gl_InstanceID;
		// Parameters
		c = texelFetch(parameters, param+0).r;
		x = texelFetch(parameters, param+1).r;
		y = texelFetch(parameters, param+2).r;
		z = texelFetch(parameters, param+3).r;
		// Position
		p = (vec2(x, y) + corners[vertex] * vec2(1.5,1.5)) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		// Color
		ColorIndex = uint(c);
		break;

	case cmdLine:
		param += 7*gl_InstanceID;
		// Parameters
		c = texelFetch(parameters, param+0).r;
		x1 = texelFetch(parameters, param+1).r;
		y1 = texelFetch(parameters, param+2).r;
		z1 = texelFetch(parameters, param+3).r;
		x2 = texelFetch(parameters, param+4).r;
		y2 = texelFetch(parameters, param+5).r;
		z2 = texelFetch(parameters, param+6).r;
		// Position
		dx = x2-x1;
		dy = y2-y1;
		v = vec3(0.5*normalize(vec2(x2-x1, y2-y1)), 0.0);
		ov = vec3(0.5*normalize(vec2(-y2+y1, x2-x1)), 0.0);
		pts = vec3[4](
			vec3(x1, y1, floatZ(z1))-v-ov,
			vec3(x2, y2, floatZ(z2))+v-ov,
			vec3(x1, y1, floatZ(z1))-v+ov,
			vec3(x2, y2, floatZ(z2))+v+ov
		);
		p = (vec2(0.5,0.5) + pts[vertex].xy) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), pts[vertex].z, 1);
		Orig = vec2(x1, y1);
		Steep = abs(dx) < abs(dy);
		if (Steep) {
			Slope = float(dx)/float(dy);
		} else {
			Slope = float(dy)/float(dx);
		}
		// Color
		ColorIndex = uint(c);
		break;

	}
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
