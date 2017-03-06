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

void main(void) {
  int x = int(gl_FragCoord.x - screen.Left) / screen.PixelSize;
  if (gl_FragCoord.x < screen.Left) {discard;}
  int y = int(gl_FragCoord.y - screen.Top)  / screen.PixelSize;
	if (gl_FragCoord.y < screen.Top) {discard;}
  int col = x / charWidth;
	if (col >= screen.NbCols) {discard;}
  int row = y / charHeight;
	if (row >= screen.NbRows) {discard;}

	// 1: Lookup the char in text SSBO

	// First, find the index of desired byte
  int chrI = col + row * screen.NbCols;
	// Fetch the word
  uint chrB = screen.Chars[chrI >> 2];
	// Isolate the correct byte inside the word
  chrB = chrB >> (8 * (chrI & 0x3));
  uint chr = chrB & 0xFF;

	if (chr == 0) {discard;}

	vec4 fg = vec4(screen.Foreground, 1.0);
	vec4 bg = screen.Background;

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
	bg.a *= float(fntB & 0x01);
	Color = fnt * fg + (1 - fnt) * bg;

	if (Color.a == 0) {
		discard;
	}
}
`
