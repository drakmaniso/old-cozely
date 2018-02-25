// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/drakmaniso/glam/palette"
)

//------------------------------------------------------------------------------

// A Cursor holds the state necessary to write text on a canvas.
//
// It maintains the current position, but also the x coordinate used for new
// lines of text.
type Cursor uint8

type cursor struct {
	canvas      Canvas
	font        Font
	color       palette.Index
	tracking    int16
	leading     int16
	x, y, z, dx int16
	params      []int16
}

var cursors []cursor

//------------------------------------------------------------------------------

// NewCursor returns a new cursor that can be used to write text. By default the
// text will be displayed on screen (see Canvas to change it).
func NewCursor() Cursor {
	if len(cursors) >= 255 {
		setErr("in NewCursor", errors.New("too many cursors"))
		return Cursor(0)
	}
	cu := cursor{canvas: Canvas(0)}
	cu.params = make([]int16, 0, 128)
	cursors = append(cursors, cu)
	return Cursor(len(cursors) - 1)
}

//------------------------------------------------------------------------------

// Canvas sets the cursor to display text on v.
//
// Note: Flush is automatically called before the change.
func (c Cursor) Canvas(cv Canvas) {
	c.Flush()
	cursors[c].canvas = cv
}

//------------------------------------------------------------------------------

// Font changes the current font.
//
// Note: Flush is automatically called before the change.
func (c Cursor) Font(f Font) {
	c.Flush()
	cursors[c].font = f
}

// Interline defines the vertical distance between two lines of text (from
// baseline to baseline). If set to 0, the distance will be computed on the fly
// as 125% of the current font height. Default is 0.
func (c Cursor) Interline(dy int16) {
	cursors[c].leading = dy
}

// LetterSpacing sets an offset to the default space between characters.
func (c Cursor) LetterSpacing(dx int16) {
	cursors[c].tracking = dx
}

// ColorShift sets an offset to all colors used to write text.
func (c Cursor) ColorShift(s palette.Index) {
	cursors[c].color = s
}

//------------------------------------------------------------------------------

// Locate moves the cursor to a specific position. It also defines column x as
// the starting point for new lines of text: i.e. when writing a newline, the
// cursor will be set to the coordinate (x, current y + interline).
//
// Note: Flush is automatically called before the relocation. See also Move and
// Moveto.
func (c Cursor) Locate(x, y, z int16) {
	cu := &cursors[c]
	c.Flush()
	cu.x, cu.y, cu.z = x, y, z
	cu.dx = 0
}

// Move moves the cursor relatively to its current position.
//
// Note: it does not change the starting point for new lines, and only Flush the
// cursor when dy or dz are not null. See also MoveTo and Locate.
func (c Cursor) Move(dx, dy, dz int16) {
	cu := &cursors[c]
	if dy != 0 || dz != 0 {
		c.Flush()
		cu.y += dy
		cu.z += dz
	}
	cu.dx += dx
}

// MoveTo changes the cursor position.
//
// Note: it does not change the starting point for new lines, and only Flush the
// cursor when y or z are different than the current position. See also Move and
// Locate.
func (c Cursor) MoveTo(x, y, z int16) {
	cu := &cursors[c]
	if y != cu.y || z != cu.z {
		c.Flush()
		cu.y = y
		cu.z = z
	}
	cu.dx = (x - cu.x)
}

// Position returns the current cursor position.
func (c Cursor) Position() (x, y, z int16) {
	cu := &cursors[c]
	return cu.x + cu.dx, cu.y, cu.z
}

//TODO: implement Link and Unlink

//------------------------------------------------------------------------------

// Print displays text on the canvas; it works like fmt.Print.
//
// Note: Flush is automatically called at the end of the text.
func (c Cursor) Print(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(c, a...)
	c.Flush()
	return n, err
}

// Println displays text on the canvas; it works like fmt.Println.
//
// Note: Flush is automatically called at the end of the text.
func (c Cursor) Println(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(c, a...)
	c.Flush()
	return n, err
}

// Printf displays text on the canvas; it works like fmt.Printf.
//
// Note: Flush is automatically called at the end of the text.
func (c Cursor) Printf(format string, a ...interface{}) (n int, err error) {
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
func (c Cursor) Write(p []byte) (n int, err error) {
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
func (c Cursor) WriteRune(r rune) {
	cu := &cursors[c]
	if r == '\n' {
		if cu.leading == 0 {
			cu.y += int16(float32(cu.font.Height()) * 1.25)
		} else {
			cu.y += cu.leading
		}
		cu.dx = 0
		c.Flush()
		return
	}

	if len(cu.params) == 0 {
		cu.params = append(
			cu.params,
			int16(cu.font),
			int16(cu.color),
			cu.x, cu.y-fonts[cu.font].baseline, cu.z,
		)
	}

	g := cu.font.glyph(r)
	cu.params = append(cu.params, int16(g), cu.dx)
	cu.dx += glyphMap[g].w + cu.tracking
}

// Flush ensures that all text written by the cursor through Write and Writerune
// is immediately displayed. Note that Flush is automatically called when the
// canvas is painted (or the screen blit); manually calling flush is only
// necessary when drawing order is important.
func (c Cursor) Flush() {
	cu := &cursors[c]
	if len(cu.params) > 0 {
		cu.canvas.appendCommand(cmdText, 4, uint32(len(cu.params)-4)/2, cu.params...)
		cu.params = cu.params[:0]
	}
}

//------------------------------------------------------------------------------
