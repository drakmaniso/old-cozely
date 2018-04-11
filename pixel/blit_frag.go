package pixel

const blitFragmentShader = "\n" + `#version 460 core

////////////////////////////////////////////////////////////////////////////////

layout(origin_upper_left, pixel_center_integer) in vec4 gl_FragCoord;

layout(binding = 0) uniform usampler2D screenTexture;

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
	uint c = texelFetch(screenTexture, ivec2(int(ScreenPosition.x), int(ScreenPosition.y)), 0).r;

	if (c == 0) {
		discard;
	}

	out_color = Colours[c];
}

////////////////////////////////////////////////////////////////////////////////

`
