package microtext

const fragmentShader = `
#version 450 core

const int charWidth = 7;
const int charHeight = 11;

layout(std430, binding = 0) buffer Font {
  uint Data[2816 / 4];
} font;

layout(std430, binding = 1) buffer Screen {
  int Left;
  int Top;
  int NbCols;
  int NbRows;
  //
  vec3 Foreground;
  int PixelSize;
  //
	vec4 Background;
  //
	uint Chars[];
} screen;

layout(origin_upper_left) in vec4 gl_FragCoord;

out vec4 Color;

uint screenChar(int col, int row) {
  int b = col + row * screen.NbCols; // The byte we're looking for
  uint v = screen.Chars[b >> 2];
  v = v >> (8 * (b & 0x3));
  return v & 0xFF;
}

void main(void) {
	vec4 fg = vec4(screen.Foreground, 1.0);
	vec4 bg = screen.Background;

  int x = int(gl_FragCoord.x - screen.Left) / screen.PixelSize;
  int y = int(gl_FragCoord.y - screen.Top)  / screen.PixelSize;
  int col = x / charWidth;
  int row = y / charHeight;

  if (gl_FragCoord.x < screen.Left || gl_FragCoord.y < screen.Top ||
		col >= screen.NbCols || row >= screen.NbRows) {
    discard;
  }

  uint chr = screenChar(col, row);

  int dx = x - col*charWidth;
  int dy = y - row*charHeight;

	// Calculate color

	// First, find the desired byte in font SSBO
  int ib = int(chr) * charHeight + dy; // byte index
  uint b = font.Data[ib >> 2];
  b = b >> (8 * (ib & 0x3));
	b &= 0xFF;
	// Calculate the color
	uint v = (b >> (7 - dx)) & 0x1;
	bg.a *= float(b & 0x01);
	Color = v * fg + (1 - v) * bg;

	if (Color.a == 0) {
		discard;
	}
}
`
