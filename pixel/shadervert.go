package pixel

const vertexShader = "\n" + `#version 460 core

//------------------------------------------------------------------------------

const uint cmdPicture    = 1;
const uint cmdPictureExt = 2;
const uint cmdPrint      = 3;
const uint cmdPoint      = 4;
const uint cmdPointList  = 5;
const uint cmdLine       = 6;

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
layout(binding = 7) uniform isamplerBuffer fontMap;
layout(binding = 6) uniform isamplerBuffer parameters;

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

void main(void)
{
	Command = gl_VertexID >> 2;
	int param = gl_BaseInstance;
	int vertex = gl_VertexID & 0x3;

	int x, y, x1, y1, x2, y2, dx, dy;
	uint c;
	vec2 p, wh, v, ov, pts[4];
	switch (Command) {
	case cmdPicture:
		// Parameters
		param += 3*gl_InstanceID;
		int m = texelFetch(parameters, param+0).r;
		x = texelFetch(parameters, param+1).r;
		y = texelFetch(parameters, param+2).r;
		// Mapping of the picture
		m *= 5;
		Bin = texelFetch(mappings, m+0).r;
		UV = vec2(texelFetch(mappings, m+1).r, texelFetch(mappings, m+2).r);
		wh = vec2(texelFetch(mappings, m+3).r, texelFetch(mappings, m+4).r);
		// Picture quad
		p = (vec2(x, y) + corners[vertex] * wh) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		UV += corners[vertex] * wh;
		break;

	case cmdPrint:
	  // Parameters of the whole Print command
		uint f = texelFetch(parameters, param+0).r;
		c = texelFetch(parameters, param+1).r;
		x = texelFetch(parameters, param+2).r;
		y = texelFetch(parameters, param+3).r;
		// Parameter for the current character
		uint r = texelFetch(parameters, param+4 + gl_InstanceID).r;
		int cr = int(r & 0x7F);
		dx = int((r >> 7)&0x1FF);
		// Mapping of the current character
		int fm = int(f) * (2 + 128*4);
		Bin = texelFetch(fontMap, fm+2+cr*4+0).r;
		wh = vec2(texelFetch(fontMap, fm+2+cr*4+1).r, texelFetch(fontMap, fm+0).r);
		UV = vec2(texelFetch(fontMap, fm+2+cr*4+2).r, texelFetch(fontMap, fm+2+cr*4+3).r);
		// Character quad
		p = (vec2(x+dx, y) + corners[vertex] * wh) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		UV += corners[vertex] * wh;
		ColorIndex = uint(c);
		break;

	case cmdPoint:
		param += 3*gl_InstanceID;
		// Parameters
		c = texelFetch(parameters, param+0).r;
		x = texelFetch(parameters, param+1).r;
		y = texelFetch(parameters, param+2).r;
		// Position
		p = (vec2(x, y) + corners[vertex] * vec2(1.5,1.5)) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		// Color
		ColorIndex = uint(c);
		break;

	case cmdPointList:
		// Parameters
		c = texelFetch(parameters, param+0).r;
		param += 1 + 2*gl_InstanceID;
		x = texelFetch(parameters, param+0).r;
		y = texelFetch(parameters, param+1).r;
		// Position
		p = (vec2(x, y) + corners[vertex] * vec2(1.5,1.5)) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), 0.5, 1);
		// Color
		ColorIndex = uint(c);
		break;

	case cmdLine:
		param += 5*gl_InstanceID;
		// Parameters
		c = texelFetch(parameters, param+0).r;
		x1 = texelFetch(parameters, param+1).r;
		y1 = texelFetch(parameters, param+2).r;
		x2 = texelFetch(parameters, param+3).r;
		y2 = texelFetch(parameters, param+4).r;
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
		ColorIndex = uint(c);
		break;

	}
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
//------------------------------------------------------------------------------
