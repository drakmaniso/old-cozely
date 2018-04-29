package pixel

const blitVertexShader = "\n" + `#version 460 core

////////////////////////////////////////////////////////////////////////////////

layout(std140, binding = 0) uniform BlitUniforms {
	vec2 ScreenSize;
};

const vec2 srcCorners[4] = vec2[4](
	vec2(0, 0),
	vec2(1, 0),
	vec2(0, 1),
	vec2(1, 1)
);

const vec2 winCorners[4] = vec2[4](
	vec2(-1, 1),
	vec2(1, 1),
	vec2(-1, -1),
	vec2(1, -1)
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
	ScreenPosition = srcCorners[gl_VertexID] * ScreenSize;
	gl_Position = vec4(winCorners[gl_VertexID], 0.5, 1.0);
}
`

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
////////////////////////////////////////////////////////////////////////////////
