package overl

const fragmentShaderFake = `
#version 450 core

out vec4 Color;

void main(void) {
	Color = vec4(1.0, 0.0, 0.0, 0.5);
}
`

const fragmentShader = `
#version 450 core

const int charWidth = 7;
const int charHeight = 11;

layout(std430, binding = 0) buffer Font {
  uint Data[2816 / 4];
} font;

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

layout(origin_upper_left) in vec4 gl_FragCoord;

out vec4 Color;

void main(void) {
  int x = int(gl_FragCoord.x - overlay.X) / overlay.PixelSize;
  int y = int(gl_FragCoord.y - overlay.Y)  / overlay.PixelSize;
  int col = x / charWidth;
	// if (col >= overlay.Columns) {discard;}
  int row = y / charHeight;
	// if (row >= overlay.Rows) {discard;}

	// 1: Lookup the char in text SSBO

	// First, find the index of desired byte
  int chrI = col + row * overlay.Columns;
	// Fetch the word
  uint chrB = overlay.Chars[chrI >> 2];
	// Isolate the correct byte inside the word
  chrB = chrB >> (8 * (chrI & 0x3));
  uint chr = chrB & 0xFF;

	if (chr == 0) { discard;}

  int dx = x - col*charWidth;
  int dy = y - row*charHeight;

	// 2: Lookup the char bitmap in font SSBO

	// First, find the index of desired byte
  int fntI = int(chr) * charHeight + dy;
	// Fetch the word
  uint fntB = font.Data[fntI >> 2];
	// Isolate the correct byte inside the word
  fntB = fntB >> (8 * (fntI & 0x3));
	fntB &= 0xFF;
	// Isolate the correct bit inside the byte
	uint fnt = (fntB >> (7 - dx)) & 0x1;

	// Calculate the color

	vec4 fg = vec4(1.0, 1.0, 1.0, 1.0);
	vec4 bg = vec4(0.0, 0.0, 0.0, 0.66);
	if (overlay.Flags != 0) {
		fg = vec4(0.0, 0.0, 0.0, 1.0);
		bg = vec4(1.0, 1.0, 1.0, 0.33);
	}
	if (chr == 0) {
		bg.a = 0.0;
	}
	Color = fnt * fg + (1 - fnt) * bg;

	if (Color.a == 0.0) {
		Color = vec4(1, 0.5, 0.5, 1.0);//discard;
	}
}
`
