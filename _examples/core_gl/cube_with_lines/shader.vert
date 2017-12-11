#version 450 core

layout(location = 0) in vec3 Position;
layout(location = 1) in vec3 Color;

layout(std140, binding = 0) uniform PerObject {
	mat4 ScreenFromObject;
} obj;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) out vec3 Color;
} vert;

void main(void) {
	vec3 p = Position;
	if (p.y > 0) {
		p.x *= 0.25;
		p.z *= 0.25;
	}
	gl_Position = obj.ScreenFromObject * vec4(p, 1);
	vert.Color = Color;
}
