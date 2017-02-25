#version 450 core

const int nbPoints = 512;

layout(location = 0) in vec2 Position;
layout(location = 1) in float Size;
layout(location = 2) in int Numerator;
layout(location = 3) in int Denominator;
layout(location = 4) in float Offset;
layout(location = 5) in float Speed;

layout(std140, binding = 0) uniform PerFrame {
	float Ratio;
	float Time;
} frame;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) out vec3 Color;
} vert;

void main(void) {
  // Calculate the rose
  vec2 p = Position;
  float k = float(Numerator) / float(Denominator);
  float theta = float(gl_VertexID) * 2 * 3.14159 / (-1 + float(nbPoints) / float(Denominator));
  float r = (cos(k*theta) + Offset) / (1.0 + Offset);
	p.x = r * cos(theta + frame.Time*Speed);
	p.y = r * sin(theta + frame.Time*Speed);

  // Position and size
  p *= 0.11;
  p.x *= frame.Ratio;
  p += vec2(
    1.0/8.0 + (gl_InstanceID % 8) / 4.0 - 1.0,
    1.0/8.0 + (gl_InstanceID / 8) / 4.0 - 1.0
  );

	gl_Position = vec4(p, 0.5, 1);
	vert.Color = vec3(0.2, 0.1, 0.05);
}
