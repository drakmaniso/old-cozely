#version 450 core

layout(location = 0) in vec2 Position;

out gl_PerVertex {
	vec4 gl_Position;
};

layout(std140, binding = 0) uniform PerFrame {
  vec2  Scale;
  float Angle;
} frame;

void main(void) {
  float x = Position.x * cos(frame.Angle) - Position.y * sin(frame.Angle);
  float y = Position.x * sin(frame.Angle) + Position.y * cos(frame.Angle);
  gl_Position = vec4(frame.Scale.x * x, frame.Scale.y * y, 0.5, 1.0);
}
