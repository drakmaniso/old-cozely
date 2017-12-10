#version 450 core

out gl_PerVertex {
	vec4 gl_Position;
};

void main(void)
{
	const vec4 triangle[3] = vec4[3](
		vec4(0, 0.65, 0.5, 1),
		vec4(-0.65, -0.475, 0.5, 1),
		vec4(0.65, -0.475, 0.5, 1)
	);
	gl_Position = triangle[gl_VertexID];
}
