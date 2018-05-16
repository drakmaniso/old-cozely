package pixel

const blitFragmentShader = "\n" + `#version 460 core

////////////////////////////////////////////////////////////////////////////////

layout(origin_upper_left, pixel_center_integer) in vec4 gl_FragCoord;

layout(std140, binding = 0) uniform BlitUniforms {
	vec2 ScreenSize;
};

layout(binding = 0) uniform usampler2D canvas;
layout(binding = 1) uniform sampler2D depth;
layout(binding = 7, r32ui) uniform uimage2D filtert;

layout(std430, binding = 1) buffer Heap {
	uint data[];
} heap;

layout(std430, binding = 0) buffer Palette {
	vec4 Colours[256];
};

////////////////////////////////////////////////////////////////////////////////

out vec4 out_color;

in PerVertex {
	layout(location=0) vec2 ScreenPosition;
};

////////////////////////////////////////////////////////////////////////////////

#define MAX_FRAGMENTS 75

struct fragment {
	uint color;
	uint order;
};

struct fragment frags[MAX_FRAGMENTS];
int count = 0;

void main(void) {
	ivec2 p = ivec2(int(ScreenPosition.x), int(ScreenSize.y - ScreenPosition.y));
	uint c = texelFetch(canvas, p, 0).r;
	uint order = uint(float(0xFFFFFF)*texelFetch(depth, p, 0).r);

	uint n = imageLoad(filtert, ivec2(int(ScreenPosition.x), int(ScreenPosition.y))).r;

	while (n != 0 && count < MAX_FRAGMENTS) {
		uint o = heap.data[2*n+1];
		if (o > order) {
			frags[count].color = heap.data[2*n] & 0xFF;
			frags[count].order = o;
			count++;
		}
		n = heap.data[2*n] >> 8;
	}

	for (int i = 1; i < count; i++) {
		struct fragment f = frags[i];
		int j = i;
		while (j > 0 && f.order <= frags[j-1].order) {
			frags[j] = frags[j-1];
			j--;
		}
		frags[j] = f;
	}

	for (int i = 0; i < count; i++) {
		switch(frags[i].color) {
		case 1:
			if (c < 8) {
				c = 1;
			} else if (c < 16) {
				c = 1;
			} else if (c < 24) {
				c = 9;
			} else {
				c = 17;
			}
			break;
		case 2:
			if (c < 8) {
				c = 2;
			} else if (c < 16) {
				c = 2;
			} else if (c < 24) {
				c = 10;
			} else {
				c = 18;
			}
			break;
		case 3:
			if (c < 8) {
				c = 3;
			} else if (c < 16) {
				c = 3;
			} else if (c < 24) {
				c = 11;
			} else {
				c = 19;
			}
			break;
		case 4:
			if (c < 8) {
				c = 4;
			} else if (c < 16) {
				c = 4;
			} else if (c < 24) {
				c = 12;
			} else {
				c = 20;
			}
			break;
		default:
			c = 8;
		}
	}

	if (c == 0) {
		discard; //TODO:
	}

	out_color = Colours[c];
}

////////////////////////////////////////////////////////////////////////////////

`
