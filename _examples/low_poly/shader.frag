#version 450 core

in PerVertex {
	layout(location = 0) flat in vec3 Normal;
  layout(location = 1) flat in uint Material;
} vertex;

out vec4 Color;

const vec4 palette[] = {
  {1, 0.5, 0, 1},
  {0.5, 0, 1, 1},
  {0, 0.5, 1, 1},
  {1, 0, 0, 1},
  {0, 1, 0, 1},
  {0, 0, 1, 1},
};

void main(void) {
	Color = palette[vertex.Material];
  // Color = vec4(0, 0, 0, 1);
}
