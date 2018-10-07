package pixel

const drawVertexShader = "\n" + `#version 460 core

////////////////////////////////////////////////////////////////////////////////

const uint cmdPicture    = 1;
const uint cmdTriangle   = 2;
const uint cmdLine       = 3;
const uint cmdBox        = 4;
const uint cmdPoint      = 5;

const vec2 corners[4] = vec2[4](
	vec2(0, 0),
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
	layout(location=3) flat uint ColorIndex;
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
	int param = gl_InstanceID * 8;
	int vertex = gl_VertexID;

	uint c = texelFetch(parameters, param+0).r;
	Command = c >> 12;
	ColorIndex = c & 0xFF;
	int z = texelFetch(parameters, param+1).r;
	int x = texelFetch(parameters, param+2).r;
	int y = texelFetch(parameters, param+3).r;
	int p4 = texelFetch(parameters, param+4).r;
	int p5 = texelFetch(parameters, param+5).r;
	int p6 = texelFetch(parameters, param+6).r;
	int p7 = texelFetch(parameters, param+7).r;

	int dx, dy;
	vec2 p, wh;
	vec2 t, n, pts[4];
	switch (Command) {
	case cmdPicture:
		// Mapping of the picture
		p6 *= 5;
		Bin = texelFetch(pictureMap, p6+0).r;
		UV = vec2(texelFetch(pictureMap, p6+1).r, texelFetch(pictureMap, p6+2).r);
		wh = vec2(texelFetch(pictureMap, p6+3).r, texelFetch(pictureMap, p6+4).r);
		// Picture quad
		p = (CanvasMargin + vec2(x, y) + corners[vertex] * wh) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		UV += corners[vertex] * wh;
		break;

	// case cmdText:
	// 	// Mapping of the current character
	// 	p6 *= 5;
	// 	Bin = texelFetch(pictureMap, p6+0).r;
	// 	UV = vec2(texelFetch(pictureMap, p6+1).r, texelFetch(pictureMap, p6+2).r);
	// 	wh = vec2(texelFetch(pictureMap, p6+3).r, texelFetch(pictureMap, p6+4).r);
	// 	// Character quad
	// 	p = (CanvasMargin + vec2(x, y) + corners[vertex] * wh) * PixelSize;
	// 	gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
	// 	UV += corners[vertex] * wh;
	// 	ColorIndex = uint(c&0xFFFF);
	// 	break;

	case cmdPoint:
		// Position
		p = (CanvasMargin + vec2(x, y) + corners[vertex]) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		break;

	case cmdLine:
		// Position
		Box = vec4(x+CanvasMargin.x, y+CanvasMargin.y, p4+CanvasMargin.x, p5+CanvasMargin.y);
		dx = p4-x;
		dy = p5-y;
		Flags = uint(abs(dx) < abs(dy)) * steep;
		t = 0.25*normalize(vec2(dx, dy));
		n = 0.75*normalize(vec2(-dy, dx));
		pts = vec2[4](
			vec2(x, y)+n-t,
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
		// Parameters
		switch (gl_VertexID) {
		case 0:
			break;
		case 1:
			x = p4;
			y = p5;
			break;
		case 2:
		case 3:
			x = p6;
			y = p7;
			break;
		}
		p = (CanvasMargin + vec2(0.5,0.5) + vec2(x, y)) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		break;

	case cmdBox:
		// Parameters
		wh = vec2(p4+1, p5+1);
		Flags = p6;
		// Position
		Box = vec4(x+CanvasMargin.x, y+CanvasMargin.y, x+p4+CanvasMargin.x, y+p5+CanvasMargin.y);
		p = (CanvasMargin + vec2(x, y) + corners[vertex] * wh) * PixelSize;
		gl_Position = vec4(p * vec2(2, -2) + vec2(-1,1), floatZ(z), 1);
		ColorIndex |= (p7 & 0xFF) << 8;
		break;
	}
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
////////////////////////////////////////////////////////////////////////////////
