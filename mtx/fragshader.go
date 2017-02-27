package mtx

const fragmentShader = `
#version 450 core

const uint charWidth = 8;
const uint charHeight = 12;
const uint nbCols = 120;
const uint nbRows = 45;

layout(std430, binding = 0) buffer Font {
  uint Data[1536 / 4];
} font;

layout(std430, binding = 1) buffer Screen {
  uint Top;
  uint Left;
  uint unused1;
  uint unused2;
  vec4 Color;
	uint Chars[(nbCols * nbRows) / 4];
} screen;

layout(origin_upper_left) in vec4 gl_FragCoord;

out vec4 Color;

uint screenChar(uint col, uint row) {
  uint b = col + row * nbCols; // The byte we're looking for
  uint v = screen.Chars[b / 4];
  v = v >> (8 * (b % 4));
  v &= 0xFF;
  return v;
}

uint fontByte(uint c, uint l) {
  uint b = c * 12 + l; // The byte we're looking for
  uint v = font.Data[b / 4];
  v = v >> (8 * (b % 4));
  v &= 0xFF;
  return v;
}

void main(void) {
  if (gl_FragCoord.x < screen.Left) {
    discard;
  }
  uint x = uint(gl_FragCoord.x) >> 1;
  uint y = uint(gl_FragCoord.y) >> 1;
  uint col = x / charWidth;
  if (col > nbCols) {
    discard;
  }
  uint row = y / charHeight;
  if (row > nbRows) {
    discard;
  }
  uint chr = screenChar(col, row);
  uint dx = x - col*charWidth;
  uint dy = y - row*charHeight;
  uint v;
  v = fontByte(chr, dy);
  if (((v >> (7 - dx)) & 0x1) != 0) {
	  Color = screen.Color;
  } else {
    discard;
  }
}
`
