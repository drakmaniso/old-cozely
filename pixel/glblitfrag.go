package pixel

const blitFragmentShader = "\n" + `#version 460 core

////////////////////////////////////////////////////////////////////////////////

layout(origin_upper_left, pixel_center_integer) in vec4 gl_FragCoord;

layout(binding = 0) uniform usampler2D screenTexture;
layout(binding = 1) uniform usampler2D filterTexture;

layout(std430, binding = 0) buffer Palette {
	vec4 Colours[256];
};

////////////////////////////////////////////////////////////////////////////////

out vec4 out_color;

in PerVertex {
	layout(location=0) vec2 ScreenPosition;
};

////////////////////////////////////////////////////////////////////////////////

void main(void) {
	// float cc = texelFetch(screenTexture, ivec2(int(ScreenPosition.x), int(ScreenPosition.y)), 0).r;
	// uint c = uint(cc*255);
	uint c = texelFetch(screenTexture, ivec2(int(ScreenPosition.x), int(ScreenPosition.y)), 0).r;

	// float ff = texelFetch(filterTexture, ivec2(int(ScreenPosition.x), int(ScreenPosition.y)), 0).r;
	// uint f = uint(ff*255);
	uint f = texelFetch(filterTexture, ivec2(int(ScreenPosition.x), int(ScreenPosition.y)), 0).r;

	if (c == 0) {
		discard;
	}

	out_color = Colours[c];
	// if (f != 0) {
	// 	out_color = Colours[c-1];
	// }
}

`

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
