// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"fmt"
	"unicode/utf8"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
)

////////////////////////////////////////////////////////////////////////////////

// Cursor holds the state used to write text on the canvas.
type cursor struct {
	Color         color.Index
	Font          FontID
	Margin        int16
	LetterSpacing int16
	Interline     int16
	Position      coord.CR
}

// Cursor holds the state used to write text on the canvas.
var Cursor cursor

////////////////////////////////////////////////////////////////////////////////

// Text configures the color and font used to display text on the canvas.
//
// Note that you can also directly change the Cursor attributes.
func Text(c color.Index, f FontID) {
	Cursor.Color = c
	Cursor.Font = f
	if Cursor.Interline == 0 {
		Cursor.Interline = int16(float32(Cursor.Font.Height()) * 1.25)
	}
}

// Locate moves the text cursor to a specific position. It also defines column
// p.C as the cursor margin, i.e. the column to which the cursor returns to
// start a new line.
//
// Note that you can also directly change the TextCursor attributes.
func Locate(p coord.CR) {
	Cursor.Position = coord.CR{p.C, p.R}
	Cursor.Margin = Cursor.Position.C
}

////////////////////////////////////////////////////////////////////////////////

// Print queues a command on the GPU to display text on the canvas (works like
// fmt.Print).
func Print(args ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(&canvas, args...)
	return n, err
}

// Println queues a command on the GPU to display text on the canvas (works like
// fmt.Println).
func Println(args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(&canvas, args...)
	return n, err
}

// Printf queues a command on the GPU to display text on the canvas (works like
// fmt.Printf).
func Printf(format string, args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(&canvas, format, args...)
	return n, err
}

////////////////////////////////////////////////////////////////////////////////

// Write asks the GPU to display p (interpreted as an UTF8 string) on the
// canvas. This method implements the io.Writer interface.
func (a cursor)Write(p []byte) (n int, err error) {
	return canvas.Write(p)
}

// Write asks the GPU to display p (interpreted as an UTF8 string) on the
// canvas. This method implements the io.Writer interface.
func (a *cmdQueue) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 0 {
		r, s := utf8.DecodeRune(p)
		a.WriteRune(r)
		p = p[s:]
	}
	return n, nil
}


// WriteRune asks the GPU to display a single rune on the canvas.
func (a cursor) WriteRune(r rune) {
	canvas.WriteRune(r)
}

// WriteRune asks the GPU to display a single rune on the canvas.
func (a *cmdQueue) WriteRune(r rune) {
	if r == '\n' {
		if Cursor.Interline == 0 {
			Cursor.Position.R += int16(float32(Cursor.Font.Height()) * 1.25)
		} else {
			Cursor.Position.R += Cursor.Interline
		}
		Cursor.Position.C = Cursor.Margin
		return
	}

	g := Cursor.Font.glyph(r)
	a.command(cmdText, 4, 1,
		int16(Cursor.Color-fonts[Cursor.Font].basecolor),
		Cursor.Position.R-fonts[Cursor.Font].baseline,
		int16(g), Cursor.Position.C)
	Cursor.Position.C += pictures.mapping[g].w + Cursor.LetterSpacing
}

////////////////////////////////////////////////////////////////////////////////
