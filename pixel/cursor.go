// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"fmt"
	"unicode/utf8"

	"github.com/cozely/cozely/palette"
)

////////////////////////////////////////////////////////////////////////////////

// A Cursor is used to write text on the canvas.
type Cursor struct {
	Position      XY    // Current position (updated by Print and Write)
	Margin        int16 // X coordinate for new lines
	Layer         int16
	Font          FontID
	Color         palette.Index
	LetterSpacing int16 // Additional space between letters, in pixels
	LineSpacing   int16 // Additional space between lines, in pixels
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
	if r == '\n' {
		a.Position.Y += a.LineSpacing + a.Font.Interline()
		a.Position.X = a.Margin
		return
	}

	g := a.Font.glyph(r)
	renderer.command(cmdPicture,
		int16(a.Color-fonts[a.Font].basecolor),
		a.Layer,
		a.Position.X, a.Position.Y-fonts[a.Font].baseline,
		0, 0, //TODO
		int16(g),
		0,
	)
	a.Position.X += pictures.mapping[g].w + a.LetterSpacing
}
