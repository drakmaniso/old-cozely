package pixel

const drawVertexShader = "\n" + `#version 460 core

////////////////////////////////////////////////////////////////////////////////

const uint cmdPicture    = 1;
const uint cmdTriangle   = 2;
const uint cmdLine       = 3;
const uint cmdBox        = 4;
const uint cmdPoint      = 5;

const vec2 corners[6] = vec2[6](
	vec2(0, 0),
	vec2(1, 0),
	vec2(0, 1),
	vec2(1, 0),
	vec2(0, 1),
	vec2(1, 1)
);

////////////////////////////////////////////////////////////////////////////////

layout(std140, binding = 0) uniform ScreenUBO {
	vec2 PixelSize;
	ivec2 CanvasMargin;
};

////////////////////////////////////////////////////////////////////////////////

layout(binding = 0) uniform isamplerBuffer parameters;
layout(binding = 1) uniform isamplerBuffer pictureMap;

////////////////////////////////////////////////////////////////////////////////

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location=0) flat uint Command;
	layout(location=1) flat uint Bin;
	layout(location=2) vec2 UV;
	layout(location=3) flat uint ColorParam;
	layout(location=4) flat vec4 Box;
	layout(location=5) flat float Slope;
	layout(location=6) flat uint Flags;
};

const uint steep = 0x01;

////////////////////////////////////////////////////////////////////////////////

float floatZ(int z) {
	return float(z)/float(0x7FFF);
}

////////////////////////////////////////////////////////////////////////////////

void main(void)
{
	int param = gl_VertexID / 6;
	int vertex = gl_VertexID - 6 * param;
	param *= 8;

	uint c = texelFetch(parameters, param+0).r;
	Command = c >> 12;
	ColorParam = c & 0xFF;
	int z = texelFetch(parameters, param+1).r;
	int x = texelFetch(parameters, param+2).r;
	int y = texelFetch(parameters, param+3).r;
	int p4, p5, p6, p7;

	vec2 p, wh;
	switch (Command) {
	case cmdPicture:
		// Mapping of the picture
		int m = 5 * texelFetch(parameters, param+6).r;
		Bin = texelFetch(pictureMap, m+0).r;
		UV = vec2(texelFetch(pictureMap, m+1).r, texelFetch(pictureMap, m+2).r);
		wh = vec2(texelFetch(pictureMap, m+3).r, texelFetch(pictureMap, m+4).r);
		// Picture quad
		p = (CanvasMargin + vec2(x, y) + corners[vertex] * wh) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		UV += corners[vertex] * wh;
		break;

	case cmdPoint:
		p = (CanvasMargin + vec2(x, y) + corners[vertex]) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		break;

	case cmdLine:
		p4 = texelFetch(parameters, param+4).r;
		p5 = texelFetch(parameters, param+5).r;
		Box = vec4(x+CanvasMargin.x, y+CanvasMargin.y, p4+CanvasMargin.x, p5+CanvasMargin.y);
		int dx = p4-x;
		int dy = p5-y;
		Flags = uint(abs(dx) < abs(dy)) * steep;
		vec2 t = 0.25*normalize(vec2(dx, dy));
		vec2 n = 0.75*normalize(vec2(-dy, dx));
		vec2 pts[6] = vec2[6](
			vec2(x, y)+n-t,
			vec2(x, y)-n-t,
			vec2(p4, p5)+n+t,
			vec2(x, y)-n-t,
			vec2(p4, p5)+n+t,
			vec2(p4, p5)-n+t
		);
		p = (CanvasMargin + vec2(0.5,0.5) + pts[vertex].xy) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		if (Flags == steep) {
			Slope = float(dx)/float(dy);
		} else {
			Slope = float(dy)/float(dx);
		}
		break;

	case cmdTriangle:
		p4 = texelFetch(parameters, param+4).r;
		p5 = texelFetch(parameters, param+5).r;
		p6 = texelFetch(parameters, param+6).r;
		p7 = texelFetch(parameters, param+7).r;
		switch (vertex) {
		case 0:
			break;
		case 1:
			x = p4;
			y = p5;
			break;
		case 2:
			x = p6;
			y = p7;
			break;
		case 3:
		case 4:
		case 5:
			break;
		}
		p = (CanvasMargin + vec2(0.5,0.5) + vec2(x, y)) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		break;

	case cmdBox:
		p4 = texelFetch(parameters, param+4).r;
		p5 = texelFetch(parameters, param+5).r;
		p6 = texelFetch(parameters, param+6).r;
		p7 = texelFetch(parameters, param+7).r;
		wh = vec2(p4+1, p5+1);
		Flags = p6;
		Box = vec4(x+CanvasMargin.x, y+CanvasMargin.y, x+p4+CanvasMargin.x, y+p5+CanvasMargin.y);
		p = (CanvasMargin + vec2(x, y) + corners[vertex] * wh) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		ColorParam |= (p7 & 0xFF) << 8;
		break;
	}
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
////////////////////////////////////////////////////////////////////////////////
