#version 450 core

layout(binding = 0) uniform sampler2D Diffuse;

in PerVertex {
	layout(location = 0) vec2 UV;
} vert;

out vec4 Color;

void main(void) {
	Color = texture(Diffuse, vert.UV);
}
