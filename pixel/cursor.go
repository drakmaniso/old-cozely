package pixel

import (
	"fmt"
	"unicode/utf8"

	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// A Cursor is used to write text on the canvas.
type Cursor struct {
	Position      XY    // Current position (updated by Print and Write)
	Margin        int16 // X coordinate for new lines
	Layer         int16
	Font          FontID
	Color         color.Index
	LetterSpacing int16 // Additional space between letters, in pixels
	LineSpacing   int16 // Additional space between lines, in pixels
}

////////////////////////////////////////////////////////////////////////////////

// Print queues a command on the GPU to display text on the canvas (works like
// fmt.Print).
func (c *Cursor) Print(args ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(c, args...)
	return n, err
}

// Println queues a command on the GPU to display text on the canvas (works like
// fmt.Println).
func (c *Cursor) Println(args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(c, args...)
	return n, err
}

// Printf queues a command on the GPU to display text on the canvas (works like
// fmt.Printf).
func (c *Cursor) Printf(format string, args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(c, format, args...)
	return n, err
}

////////////////////////////////////////////////////////////////////////////////

// Write asks the GPU to display p (interpreted as an UTF8 string) on the
// canvas. This method implements the io.Writer interface.
func (c *Cursor) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 0 {
		r, s := utf8.DecodeRune(p)
		c.WriteRune(r)
		p = p[s:]
	}
	return n, nil
}

// WriteRune asks the GPU to display a single rune on the canvas.
func (c *Cursor) WriteRune(r rune) {
	if r == '\n' {
		c.Position.Y += c.LineSpacing + c.Font.Interline()
		c.Position.X = c.Margin
		return
	}

	g := c.Font.glyph(r)
	renderer.command(cmdPicture,
		int16(c.Color),
		c.Layer,
		c.Position.X, c.Position.Y-fonts.baseline[c.Font],
		0, 0, //TODO
		int16(g),
		0,
	)
	c.Position.X += pictures.mapping[g].w + c.LetterSpacing
}

//// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
