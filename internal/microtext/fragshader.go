package microtext

const fragmentShader = `
#version 450 core

const int charWidth = 8;
const int charHeight = 12;

layout(std430, binding = 0) buffer Font {
  uint Data[3072 / 4];
} font;

layout(std430, binding = 1) buffer Screen {
  int Left;
  int Top;
  int NbCols;
  int NbRows;
  //
  int PixelSize;
  float FgRed;
  float FgGreen;
  float FgBlue;
  //
  uint Opacity;
  float BgRed;
  float BgGreen;
  float BgBlue;
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

uint fontByte(uint c, int l) {
  int b = int(c) * 12 + l; // The byte we're looking for
  uint v = font.Data[b >> 2];
  v = v >> (8 * (b & 0x3));
  return v & 0xFF;
}

void main(void) {
  // TODO: use alpha and mix/step instead of lots of if statements?
  if (gl_FragCoord.x < screen.Left || gl_FragCoord.y < screen.Top) {
    discard;
  }
  int x = int(gl_FragCoord.x - screen.Left) / screen.PixelSize;
  int y = int(gl_FragCoord.y - screen.Top)  / screen.PixelSize;
  int col = x / charWidth;
  if (col >= screen.NbCols) {
    discard;
  }
  int row = y / charHeight;
  if (row >= screen.NbRows) {
    discard;
  }
  uint chr = screenChar(col, row);
  if (chr == 0xFF) {
    discard;
  }
  int dx = x - col*charWidth;
  int dy = y - row*charHeight;
  uint v;
  v = fontByte(chr, dy);
  bool fg = bool((v >> (7 - dx)) & 0x1); // != (chr >> 7);
  if (fg) {
	  Color = vec4(screen.FgRed, screen.FgGreen, screen.FgBlue, 1.0);
  } else {
    if (screen.Opacity == 0) { // && ((chr & 0x80) == 0)) {
      discard;
    }
    Color =  vec4(screen.BgRed, screen.BgGreen, screen.BgBlue, 1.0);
  }
}
`
