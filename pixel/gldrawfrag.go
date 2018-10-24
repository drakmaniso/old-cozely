package pixel

const drawFragmentShader = "\n" + `#version 460 core

////////////////////////////////////////////////////////////////////////////////

const uint cmdPicture    = 1;
const uint cmdTile       = 2;
const uint cmdSubpicture = 3;
const uint cmdTriangle   = 4;
const uint cmdLine       = 5;
const uint cmdPoint      = 6;

////////////////////////////////////////////////////////////////////////////////

in PerVertex {
	layout(location=0) flat uint Command;
	layout(location=1) flat uint ColorParam;
	layout(location=2) flat uint Param1;
	layout(location=3) flat float Param2;
	layout(location=4) flat ivec2 UV;
	layout(location=5) flat ivec2 PictSize;
	layout(location=6) flat ivec2 TileSize;
	layout(location=7) flat ivec2 Borders;
	layout(location=8) vec2 Position;
	layout(location=9) flat vec4 Box;
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
		c = texelFetch(Pictures, ivec3(Position.x, Position.y, Param1), 0).x;
		if (c != 0) {
			c += ColorParam;
			if (c > 255) {
				c -= 255;
			}
		}
		break;

	case cmdTile:
		ivec2 uv = ivec2(UV);
		ivec2 p = ivec2(Position);

		int l = Borders.x>>8;
		int r = Borders.x&0xFF;
		if (Position.x < l) {
			uv.x += p.x;
		} else if (p.x < TileSize.x - r) {
			uv.x += l + (p.x - l) % (PictSize.x - r - l);
		} else {
			uv.x += PictSize.x - TileSize.x + p.x;
		}

		int t = Borders.y>>8;
		int b = Borders.y&0xFF;
		if (Position.y < t) {
			uv.y += p.y;
		} else if (p.y < TileSize.y - b) {
			uv.y += t + (p.y - t) % (PictSize.y - t - b);
		} else {
			uv.y += PictSize.y - TileSize.y + p.y;
		}

		c = texelFetch(Pictures, ivec3(uv.x, uv.y, Param1), 0).x;
		if (c != 0) {
			c += ColorParam;
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
		if (Param1 == steep) {
			if (x == round(Param2*y)) {
				c = ColorParam;
			}
		} else {
			if (y == round(Param2*x)) {
				c = ColorParam;
			}
		}
		break;
	}

	if (c == 0) {
		discard;
	}

	Color = c;
}
`

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
