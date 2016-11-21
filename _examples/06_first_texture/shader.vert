#version 450 core

layout(location = 0) in vec3 Position;
layout(location = 1) in vec2 UV;

layout(std140, binding = 0) uniform PerObject {
	mat4 Transform;
} obj;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) out vec2 UV;
} vert;

void main(void) {
	gl_Position = obj.Transform * vec4(Position, 1);
	vert.UV = UV;
}
