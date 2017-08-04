package overl

const vertexShader = `
#version 450 core

layout(location = 0) in vec3 Position;

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
	uint PixelSize;
	uint Flags;
	int unused1;
	int unused2;
  //
	uint Chars[];
} overlay;

out gl_PerVertex {
	vec4 gl_Position;
};

out Header {
  layout(location = 0) flat int X;
	layout(location = 1) flat int Y;
  layout(location = 2) flat int Columns;
	layout(location = 3) flat int Rows;
	layout(location = 4) flat uint PixelSize;
	layout(location = 5) flat vec4 Foreground;
	layout(location = 6) flat vec4 Background;
};

const

void main(void) {
	vec2 positions[4] = vec2[4](
		vec2(overlay.Left, overlay.Top),
		vec2(overlay.Right, overlay.Top),
		vec2(overlay.Left, overlay.Bottom),
		vec2(overlay.Right, overlay.Bottom)
	);
	gl_Position = vec4(positions[gl_VertexID], 0.5, 1.0);

	// Avoid SSBO reads in fragment shader
	X = overlay.X;
	Y = overlay.Y;
	Columns = overlay.Columns;
	Rows = overlay.Rows;
	PixelSize = overlay.PixelSize;
	if (overlay.Flags != 0) {
		Foreground = vec4(1.0, 1.0, 1.0, 1.0);
		Background = vec4(0.0, 0.0, 0.0, 0.66);
	} else {
		Foreground = vec4(0.0, 0.0, 0.0, 1.0);
		Background = vec4(1.0, 1.0, 1.0, 0.33);
	}
}
`
