// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"fmt"
	"unicode/utf8"

	"github.com/cozely/cozely/color"
)

////////////////////////////////////////////////////////////////////////////////

// A Cursor is used to write text on the canvas.
type Cursor struct {
	Color         color.Index
	Font          FontID
	Margin        int16
	LetterSpacing int16
	Interline     int16
	Position      XY
	Layer         int16
}

////////////////////////////////////////////////////////////////////////////////

// Style configures the color and font used to display text on the canvas. It
// also sets the cursor Interline to a sensible default for the selected font.
func (a *Cursor) Style(c color.Index, f FontID) {
	a.Color = c
	a.Font = f
	a.Interline = int16(float32(a.Font.Height()) * 1.25)
}

// Locate moves the text cursor to a specific position. It also defines column
// p.X as the cursor margin, i.e. the column to which the cursor returns to
// start a new line.
func (a *Cursor) Locate(layer int16, p XY) {
	a.Layer = layer
	a.Position = XY{p.X, p.Y}
	a.Margin = a.Position.X
}

////////////////////////////////////////////////////////////////////////////////

// Print queues a command on the GPU to display text on the canvas (works like
// fmt.Print).
func (a *Cursor) Print(args ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(a, args...)
	return n, err
}

// Println queues a command on the GPU to display text on the canvas (works like
// fmt.Println).
func (a *Cursor) Println(args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(a, args...)
	return n, err
}

// Printf queues a command on the GPU to display text on the canvas (works like
// fmt.Printf).
func (a *Cursor) Printf(format string, args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(a, format, args...)
	return n, err
}

////////////////////////////////////////////////////////////////////////////////

// // Write asks the GPU to display p (interpreted as an UTF8 string) on the
// // canvas. This method implements the io.Writer interface.
// func (a *Cursor) Write(p []byte) (n int, err error) {
// 	return renderer.Write(p)
// }

// // WriteRune asks the GPU to display a single rune on the canvas.
// func (a *Cursor) WriteRune(r rune) {
// 	renderer.WriteRune(r)
// }

////////////////////////////////////////////////////////////////////////////////

// Write asks the GPU to display p (interpreted as an UTF8 string) on the
// canvas. This method implements the io.Writer interface.
func (a *Cursor) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 0 {
		r, s := utf8.DecodeRune(p)
		a.WriteRune(r)
		p = p[s:]
	}
	return n, nil
}

// WriteRune asks the GPU to display a single rune on the canvas.
func (a *Cursor) WriteRune(r rune) {
	if a.Font == 0 && a.Interline == 0 {
		a.Color = 7
		a.Interline = int16(float32(a.Font.Height()) * 1.25)
		if a.Position.X == 0 && a.Position.Y == 0 {
			a.Position.X = 4
			a.Position.Y = a.Interline
		}
	}
	if r == '\n' {
		a.Position.Y += a.Interline
		a.Position.X = a.Margin
		return
	}

	g := a.Font.glyph(r)
	renderer.command(cmdText, 4, 1,
		int16(a.Color-fonts[a.Font].basecolor),
		a.Layer,
		a.Position.Y-fonts[a.Font].baseline,
		int16(g), a.Position.X)
	a.Position.X += pictures.mapping[g].w + a.LetterSpacing
}
