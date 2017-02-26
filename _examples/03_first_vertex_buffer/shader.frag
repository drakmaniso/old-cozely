#version 450 core

in PerVertex {
	layout(location = 0) in vec3 Color;
} vertex;

out vec4 Color;

void main(void) {
	Color = vec4(vertex.Color, 1);
}
