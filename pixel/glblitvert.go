package pixel

const blitVertexShader = "\n" + `#version 460 core

////////////////////////////////////////////////////////////////////////////////

layout(std140, binding = 0) uniform BlitUniforms {
	vec2 WindowSize;
	vec2 ScreenSize;
	vec2 ScreenZoom;
};

const vec2 corners[4] = vec2[4](
	vec2(0, 0),
	vec2(1, 0),
	vec2(0, 1),
	vec2(1, 1)
);

////////////////////////////////////////////////////////////////////////////////

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location=0) vec2 ScreenPosition;
};

////////////////////////////////////////////////////////////////////////////////

void main(void) {
	ScreenPosition = corners[gl_VertexID] * ScreenSize;
	vec2 b = floor((WindowSize - ScreenSize * ScreenZoom));
	b = 2.0 * b / WindowSize;
	vec2 w[4] = vec2[4] (
		vec2(-1, -1 + b.y),
		vec2(1 - b.x, -1 + b.y),
		vec2(-1, 1),
		vec2(1 - b.x, 1)
	);
	gl_Position = vec4(
		w[gl_VertexID],
		0.5,
		1.0
	);
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
////////////////////////////////////////////////////////////////////////////////
