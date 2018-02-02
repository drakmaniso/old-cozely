// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"fmt"
	"unicode/utf8"

	"github.com/drakmaniso/glam/palette"
)

//------------------------------------------------------------------------------

// A Cursor holds the state necessary to write text on a canvas.
//
// It maintains the current position, but also the x coordinate used for new
// lines of text.
type Cursor struct {
	canvas   *ScreenCanvas
	font     Font
	color    palette.Index
	tracking int16
	leading  int16
	x, y, dx int16
	params   []int16
}

//------------------------------------------------------------------------------

// NewCursor returns a new cursor that can be used to write text on the canvas.
func (s *ScreenCanvas) NewCursor() *Cursor {
	c := Cursor{canvas: s}
	c.params = make([]int16, 0, 128)
	return &c
}

//------------------------------------------------------------------------------

// Font changes the current font.
//
// Note: Flush is automatically called before the change.
func (c *Cursor) Font(f Font) {
	c.Flush()
	c.font = f
}

// Interline defines the vertical distance between two lines of text (from
// baseline to baseline). If set to 0, the distance will be computed on the fly
// as 125% of the current font height. Default is 0.
func (c *Cursor) Interline(dy int16) {
	c.leading = dy
}

// LetterSpacing sets an offset to the default space between characters.
func (c *Cursor) LetterSpacing(dx int16) {
	c.tracking = dx
}

// ColorShift sets an offset to all colors used to write text.
func (c *Cursor) ColorShift(s palette.Index) {
	c.color = s
}

//------------------------------------------------------------------------------

// Locate moves the cursor to a specific position. It also defines column x as
// the starting point for new line of text: i.e. when writing a newline, the
// cursor will be set to the coordinate (x, current y + interline).
//
// Note: Flush is automatically called before the relocation. See also Move and
// Moveto.
func (c *Cursor) Locate(x, y int16) {
	c.Flush()
	c.x, c.y = x, y
	c.dx = 0
}

// Move moves the cursor relatively to its current position.
//
// Note: it does not change the starting point for new lines, and only Flush the
// cursor when dy is not null. See also MoveTo and Locate.
func (c *Cursor) Move(dx, dy int16) {
	if dy != 0 {
		c.Flush()
		c.y += dy
	}
	c.dx += dx
}

// MoveTo changes the cursor position.
//
// Note: it does not change the starting point for new lines, and only Flush the
// cursor when dy is not null. See also Move and Locate.
func (c *Cursor) MoveTo(x, y int16) {
	if y != c.y {
		c.Flush()
		c.y = y
	}
	c.dx = (x - c.x)
}

// Position returns the current cursor position.
func (c *Cursor) Position() Coord {
	return Coord{c.x + c.dx, c.y}
}

//------------------------------------------------------------------------------

// Print displays text on the canvas; it works like fmt.Print.
//
// Note: Flush is automatically called at the end of the text.
func (c *Cursor) Print(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(c, a...)
	c.Flush()
	return n, err
}

// Println displays text on the canvas; it works like fmt.Println.
//
// Note: Flush is automatically called at the end of the text.
func (c *Cursor) Println(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(c, a...)
	c.Flush()
	return n, err
}

// Printf displays text on the canvas; it works like fmt.Printf.
//
// Note: Flush is automatically called at the end of the text.
func (c *Cursor) Printf(format string, a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(c, format, a...)
	c.Flush()
	return n, err
}

//------------------------------------------------------------------------------

// Write implements the io.Writer interface. It is a low-level method used to
// display p (interpreted as an UTF8 string) on the canvas.
//
// Note that you need to call Flush to ensure that the text is actually
// displayed; this is because consecutive calls to Write and WriteRune happening
// on the same line are merged into a single draw command. See also the more
// convenient Print, Println and Printf methods.
func (c *Cursor) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 0 {
		r, s := utf8.DecodeRune(p)
		c.WriteRune(r)
		p = p[s:]
	}
	return n, nil
}

// WriteRune is a low-level method used to display a single rune on the canvas.
//
// Note that you need to call Flush to ensure that the text is actually
// displayed; this is because consecutive calls to Write and WriteRune happening
// on the same line are merged into a single draw command.
func (c *Cursor) WriteRune(r rune) {
	if r == '\n' {
		if c.leading == 0 {
			c.y += int16(float32(c.font.Height()) * 1.25)
		} else {
			c.y += c.leading
		}
		c.dx = 0
		c.Flush()
		return
	}

	if len(c.params) == 0 {
		c.params = append(c.params, int16(c.font), int16(c.color), c.x, c.y)
	}

	g := c.font.getGlyph(r)
	c.params = append(c.params, g, c.dx)
	c.dx += glyphsMap[g].w + c.tracking
}

// Flush ensures that all text written by the cursor through Write and Writerune
// is immediately displayed.
func (c *Cursor) Flush() {
	if len(c.params) > 0 {
		c.canvas.appendCommand(cmdText, 4, uint32(len(c.params)-4)/2, c.params...)
		c.params = c.params[:0]
	}
}

//------------------------------------------------------------------------------
