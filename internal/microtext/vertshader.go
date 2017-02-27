package microtext

const vertexShader = `
#version 450 core

layout(location = 0) in vec3 Position;

out gl_PerVertex {
	vec4 gl_Position;
};

const vec2 positions[4] = vec2[4](
  vec2(-1, -1),
  vec2(1, -1),
  vec2(-1, 1),
  vec2(1, 1)
);

void main(void) {
	gl_Position = vec4(positions[gl_VertexID], 0.5, 1.0);
}
`
