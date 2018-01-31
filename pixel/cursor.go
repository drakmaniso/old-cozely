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
	newline  bool
	cmd      [][]int16
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
	return fmt.Fprint(c, a...)
}

func (c *Cursor) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c, a...)
}

func (c *Cursor) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c, format, a...)
}

//------------------------------------------------------------------------------

func (c *Cursor) Write(p []byte) (n int, err error) {
	for len(p) > 0 {
		r, s := utf8.DecodeRune(p)
		c.put(r)
		p = p[s:]
		n++
	}
	c.flush()
	return n, nil
}

//------------------------------------------------------------------------------

func (c *Cursor) put(r rune) {
	if r == '\n' {
		c.x = c.ox
		c.y += c.font.Height() + 4 //TODO:
		c.dx = 0
		c.newline = true
		return
	}

	n := len(c.cmd) - 1
	if n < 0 || c.dx > 0x1FF || c.newline {
		c.newline = false
		c.x = c.x + c.dx
		c.dx = 0
		c.cmd = append(c.cmd, make([]int16, 0, 32))
		n = len(c.cmd) - 1
		c.cmd[n] = append(c.cmd[n], int16(c.font), int16(c.color), c.x, c.y)
	}

	rr := uint16(0x7F)
	if r <= 0x7F {
		rr = uint16(r)
	}
	_, _, _, rw, _ := c.font.getMap(rune(rr))
	rr |= uint16(c.dx) << 7
	c.cmd[n] = append(c.cmd[n], int16(rr))
	c.dx += int16(rw) + 0 //TODO:
}

//------------------------------------------------------------------------------

func (c *Cursor) flush() {
	for n := range c.cmd {
		c.canvas.appendCommand(cmdPrint, 4, uint32(len(c.cmd[n])-4), c.cmd[n]...)
		c.cmd[n] = c.cmd[n][:0]
	}
	c.cmd = c.cmd[:0]
	c.newline = false
}

//------------------------------------------------------------------------------
