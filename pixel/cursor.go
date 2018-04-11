// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"fmt"
	"unicode/utf8"

	"github.com/drakmaniso/cozely/palette"
	"github.com/drakmaniso/cozely/plane"
)

//------------------------------------------------------------------------------

// A Cursor holds the state necessary to write text on a canvas.
//
// It maintains the current position, but also the x coordinate used for new
// lines of text.
type Cursor struct {
	Canvas    CanvasID
	Font      FontID
	Color     palette.Index
	Spacing   int16
	Interline int16
	Depth     int16
	Origin    plane.Pixel
	Position  plane.Pixel
}

//------------------------------------------------------------------------------

// Locate moves the cursor to a specific position. It also defines column x as
// the starting point for new lines of text: i.e. when writing a newline, the
// cursor will be set to the coordinate (x, current y + interline).
//
// Note: Flush is automatically called before the relocation. See also Move and
// Moveto.
func (c *Cursor) Locate(x, y int16) {
	c.Position.X, c.Position.Y = x, y
	c.Origin = c.Position
}

//------------------------------------------------------------------------------

// Print displays text on the canvas; it works like fmt.Print.
//
// Note: Flush is automatically called at the end of the text.
func (c *Cursor) Print(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(c, a...)
	return n, err
}

// Println displays text on the canvas; it works like fmt.Println.
//
// Note: Flush is automatically called at the end of the text.
func (c *Cursor) Println(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(c, a...)
	return n, err
}

// Printf displays text on the canvas; it works like fmt.Printf.
//
// Note: Flush is automatically called at the end of the text.
func (c *Cursor) Printf(format string, a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(c, format, a...)
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
		if c.Interline == 0 {
			c.Position.Y += int16(float32(c.Font.Height()) * 1.25)
		} else {
			c.Position.Y += c.Interline
		}
		c.Position.X = c.Origin.X
		return
	}

	g := c.Font.glyph(r)
	c.Canvas.appendCommand(cmdText, 4, 1,
		int16(c.Color), c.Depth, c.Position.Y-fonts[c.Font].baseline,
		int16(g), c.Position.X)
	c.Position.X += glyphMap[g].w + c.Spacing
}

//------------------------------------------------------------------------------
