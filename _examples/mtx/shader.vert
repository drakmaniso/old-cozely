#version 450 core

layout(location = 0) in vec3 Position;
layout(location = 1) in vec3 Color;

layout(std140, binding = 0) uniform PerFrame {
	mat4 ViewProjection;
  float time;
} frame;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) out vec3 Color;
} vert;

mat4 rotation(float angle, vec3 axis) {
  float c = cos(angle);
  float s = sin(angle);
  return mat4(
		c + axis.x*axis.x*(1-c), -axis.z*s + axis.x*axis.y*(1-c), axis.y*s + axis.x*axis.z*(1-c), 0,
		axis.z*s + axis.y*axis.x*(1-c), c + axis.y*axis.y*(1-c), -axis.x*s + axis.y*axis.z*(1-c), 0,
		-axis.y*s + axis.z*axis.x*(1-c), axis.x*s + axis.z*axis.y*(1-c), c + axis.z*axis.z*(1-c), 0,
		0, 0, 0, 1
  );
}

void main(void) {
  float iid = float(gl_InstanceID % 28) + 1.0;
  float ix = float(gl_InstanceID % 28 % 7);
  float iy = float(gl_InstanceID % 28 / 7);
  mat4 r1 = rotation(frame.time / (29.0-ix-iy), vec3(1,0,0));
  mat4 r2 = rotation(frame.time/(3.0+ix), vec3(0,1,0));
  mat4 r3 = rotation(frame.time/(5.0+iy), vec3(0,0,1));
	gl_Position = r1 * r2 * r3 * vec4(Position, 1);
  gl_Position.x += -6.0 + 2.0 * ix;
  gl_Position.y += -3.0 + 2.0 * iy;
	gl_Position = frame.ViewProjection * gl_Position;
	vert.Color = Color;
}
