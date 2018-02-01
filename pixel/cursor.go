// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"fmt"
	"unicode/utf8"

	"github.com/drakmaniso/glam/palette"
)

//------------------------------------------------------------------------------

type Cursor struct {
	canvas   *ScreenCanvas
	font     Font
	color    palette.Index
	ox, oy   int16
	x, y, dx int16
	params   []int16
}

//------------------------------------------------------------------------------

func (s *ScreenCanvas) NewCursor() *Cursor {
	return &Cursor{canvas: s}
}

//------------------------------------------------------------------------------

func (c *Cursor) Font(f Font) {
	c.font = f
}

func (c *Cursor) ColorShift(s palette.Index) {
	c.color = s
}

//------------------------------------------------------------------------------

func (c *Cursor) Locate(x, y int16) {
	c.ox, c.oy = x, y
	c.x, c.y = x, y
	c.dx = 0
}

//------------------------------------------------------------------------------

func (c *Cursor) Position() Coord {
	return Coord{c.x + c.dx, c.y}
}

//------------------------------------------------------------------------------

func (c *Cursor) Print(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(c, a...)
	c.Flush()
	return n, err
}

func (c *Cursor) Println(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(c, a...)
	c.Flush()
	return n, err
}

func (c *Cursor) Printf(format string, a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(c, format, a...)
	c.Flush()
	return n, err
}

//------------------------------------------------------------------------------

func (c *Cursor) Write(p []byte) (n int, err error) {
	for len(p) > 0 {
		r, s := utf8.DecodeRune(p)
		c.WriteRune(r)
		p = p[s:]
		n++
	}
	return n, nil
}

//------------------------------------------------------------------------------

func (c *Cursor) WriteRune(r rune) {
	if r == '\n' {
		c.x = c.ox
		c.y += c.font.Height() + 2 //TODO:
		c.dx = 0
		c.Flush()
		return
	}

	if len(c.params) == 0 {
		c.x = c.x + c.dx
		c.dx = 0
		c.params = append(c.params, int16(c.font), int16(c.color), c.x, c.y)
	}

	g := c.font.getGlyph(r)
	c.params = append(c.params, g, c.dx)
	c.dx += glyphsMap[g].w + 0 //TODO:
}

//------------------------------------------------------------------------------

func (c *Cursor) Flush() {
	if len(c.params) > 0 {
		c.canvas.appendCommand(cmdText, 4, uint32(len(c.params)-4)/2, c.params...)
		c.params = c.params[:0]
	}
}

//------------------------------------------------------------------------------
