#version 450 core

layout(location = 0) in vec2 Position;
layout(location = 1) in vec3 Color;

layout(std140, binding = 0) uniform PerObject {
	mat3 Transform;
} obj;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) out vec3 Color;
} vert;

void main(void) {
	vec3 p = obj.Transform * vec3(Position, 1);
	gl_Position = vec4(p.xy, 0.5, 1);
	vert.Color = Color;
}
