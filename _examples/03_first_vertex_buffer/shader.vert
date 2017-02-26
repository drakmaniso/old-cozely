#version 450 core

layout(location = 0) in vec2 Position;
layout(location = 1) in vec3 Color;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) out vec3 Color;
} vertex;

void main(void) {
	gl_Position = vec4(Position, 0.5, 1);
	vertex.Color = Color;
}
