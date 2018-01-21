package pixel

const vertexShader = "\n" + `#version 460 core

//------------------------------------------------------------------------------

const uint cmdIndexed = 1;
const uint cmdFullColor = 3;
const uint cmdPoint = 5;
const uint cmdPointList = 6;
const uint cmdLine = 7;
const uint cmdLineAA = 8;

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

layout(binding = 5) uniform isamplerBuffer mappings;
layout(binding = 6) uniform isamplerBuffer parameters;

//------------------------------------------------------------------------------

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location=0) flat uint Command;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) flat vec4 Color;
	layout(location=4) flat vec2 Orig;
	layout(location=5) flat float Slope;
	layout(location=6) flat bool Steep;
};

//------------------------------------------------------------------------------

void main(void)
{
	Command = gl_VertexID >> 2;
	int param = gl_BaseInstance;
	int vertex = gl_VertexID & 0x3;

	int x, y, x1, y1, x2, y2, dx, dy;
	vec2 p, v, ov, pts[4];
	uint rg, ba;
	switch (Command) {
	case cmdIndexed:
	case cmdFullColor:
		// Picture Paramters
		param += 3*gl_InstanceID;
		int m = texelFetch(parameters, param+0).r;
		x = texelFetch(parameters, param+1).r;
		y = texelFetch(parameters, param+2).r;
		// Picture Mapping in Atlas
		m *= 5;
		Bin = texelFetch(mappings, m+0).r;
		UV = vec2(texelFetch(mappings, m+1).r, texelFetch(mappings, m+2).r);
		vec2 WH = vec2(texelFetch(mappings, m+3).r, texelFetch(mappings, m+4).r);
		// Picture Position and Corner
		vec2 p = (vec2(x, y) + corners[vertex] * WH) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		UV += corners[vertex] * WH;
		break;

	case cmdPoint:
		param += 4*gl_InstanceID;
		// Point Parameters
		rg = texelFetch(parameters, param+0).r;
		ba = texelFetch(parameters, param+1).r;
		x = texelFetch(parameters, param+2).r;
		y = texelFetch(parameters, param+3).r;
		// Position
		p = (vec2(x, y) + corners[vertex] * vec2(1.5,1.5)) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		// Color
		Color = vec4(
			float(rg>>8)/float(0xFF),
			float(rg&0xFF)/float(0xFF),
			float(ba>>8)/float(0xFF),
			float(ba&0xFF)/float(0xFF)
		);
		break;

	case cmdPointList:
		// Point Parameters
		rg = texelFetch(parameters, param+0).r;
		ba = texelFetch(parameters, param+1).r;
		param += 2 + 2*gl_InstanceID;
		x = texelFetch(parameters, param+0).r;
		y = texelFetch(parameters, param+1).r;
		// Position
		p = (vec2(x, y) + corners[vertex] * vec2(1.5,1.5)) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		// Color
		Color = vec4(
			float(rg>>8)/float(0xFF),
			float(rg&0xFF)/float(0xFF),
			float(ba>>8)/float(0xFF),
			float(ba&0xFF)/float(0xFF)
		);
		break;

	case cmdLine:
		param += 6*gl_InstanceID;
		// Point Parameters
		rg = texelFetch(parameters, param+0).r;
		ba = texelFetch(parameters, param+1).r;
		x1 = texelFetch(parameters, param+2).r;
		y1 = texelFetch(parameters, param+3).r;
		x2 = texelFetch(parameters, param+4).r;
		y2 = texelFetch(parameters, param+5).r;
		// Position
		dx = x2-x1;
		dy = y2-y1;
		v = 0.5*normalize(vec2(x2-x1, y2-y1));
		ov = 0.5*normalize(vec2(-y2+y1, x2-x1));
		pts = vec2[4](
			vec2(x1, y1)-v-ov,
			vec2(x2, y2)+v-ov,
			vec2(x1, y1)-v+ov,
			vec2(x2, y2)+v+ov
		);
		p = (vec2(0.5,0.5) + pts[vertex]) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		Orig = vec2(x1, y1);
		Steep = abs(dx) < abs(dy);
		if (Steep) {
			Slope = float(dx)/float(dy);
		} else {
			Slope = float(dy)/float(dx);
		}
		// Color
		Color = vec4(
			float(rg>>8)/float(0xFF),
			float(rg&0xFF)/float(0xFF),
			float(ba>>8)/float(0xFF),
			float(ba&0xFF)/float(0xFF)
		);
		break;

	case cmdLineAA:
		param += 6*gl_InstanceID;
		// Point Parameters
		rg = texelFetch(parameters, param+0).r;
		ba = texelFetch(parameters, param+1).r;
		x1 = texelFetch(parameters, param+2).r;
		y1 = texelFetch(parameters, param+3).r;
		x2 = texelFetch(parameters, param+4).r;
		y2 = texelFetch(parameters, param+5).r;
		// Position
		dx = x2-x1;
		dy = y2-y1;
		v = 0.5*normalize(vec2(x2-x1, y2-y1));
		ov = 1.0*normalize(vec2(-y2+y1, x2-x1));
		pts = vec2[4](
			vec2(x1, y1)-v-ov,
			vec2(x2, y2)+v-ov,
			vec2(x1, y1)-v+ov,
			vec2(x2, y2)+v+ov
		);
		p = (vec2(0.5,0.5) + pts[vertex]) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		Orig = vec2(x1, y1);
		Steep = abs(dx) < abs(dy);
		if (Steep) {
			Slope = float(dx)/float(dy);
		} else {
			Slope = float(dy)/float(dx);
		}
		// Color
		Color = vec4(
			float(rg>>8)/float(0xFF),
			float(rg&0xFF)/float(0xFF),
			float(ba>>8)/float(0xFF),
			float(ba&0xFF)/float(0xFF)
		);
		break;
	}
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
