package overl

const vertexShader = `
#version 450 core

layout(location = 0) in vec3 Position;

out gl_PerVertex {
	vec4 gl_Position;
};

layout(std430, binding = 1) buffer Overlay {
  float Left;
  float Top;
	float Right;
  float Bottom;
	//
	int X;
	int Y;
  int Columns;
	int Rows;
	//
	int PixelSize;
	uint Flags;
	int unused1;
	int unused2;
  //
	uint Chars[];
} overlay;

const

void main(void) {
	vec2 positions[4] = vec2[4](
		vec2(overlay.Left, overlay.Top),
		vec2(overlay.Right, overlay.Top),
		vec2(overlay.Left, overlay.Bottom),
		vec2(overlay.Right, overlay.Bottom)
	);
	gl_Position = vec4(positions[gl_VertexID], 0.5, 1.0);
}
`
