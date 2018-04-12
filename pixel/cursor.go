// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"fmt"
	"unicode/utf8"

	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/palette"
)

////////////////////////////////////////////////////////////////////////////////

// A Cursor holds the state necessary to write text on a canvas.
//
// It maintains the current position, but also the x coordinate used for new
// lines of text.
type Cursor struct {
	Color         palette.Index
	Font          FontID
	Margin        int16
	LetterSpacing int16
	Interline     int16
	Position      coord.CRD
}

////////////////////////////////////////////////////////////////////////////////

// Text configures the color and font used to display text on the canvas.
//
// Note that you can also directly change the attributes of the cursor: see
// Cursor.
func (a CanvasID) Text(c palette.Index, f FontID) {
	cu := &canvases[a].cursor
	cu.Color = c
	cu.Font = f
	if cu.Interline == 0 {
		cu.Interline = int16(float32(cu.Font.Height()) * 1.25)
	}
}

// Locate moves the text cursor to a specific position. It also defines column c
// as the starting point for new lines of text.
//
// Note that you can also directly change these attributes: see Cursor.
func (a CanvasID) Locate(c, r, d int16) {
	cu := &canvases[a].cursor
	cu.Position = coord.CRD{c, r, d}
	cu.Margin = cu.Position.C
}

// Cursor gives access to the attributes used to display text on the canvas.
// These attributes can be changed at anytime.
func (a CanvasID) Cursor() *Cursor {
	return &canvases[a].cursor
}

////////////////////////////////////////////////////////////////////////////////

// Print displays text on the canvas; it works like fmt.Print.
func (a CanvasID) Print(args ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(a, args...)
	return n, err
}

// Println displays text on the canvas; it works like fmt.Println.
func (a CanvasID) Println(args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(a, args...)
	return n, err
}

// Printf displays text on the canvas; it works like fmt.Printf.
func (a CanvasID) Printf(format string, args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(a, format, args...)
	return n, err
}

////////////////////////////////////////////////////////////////////////////////

// Write implements the io.Writer interface. It is a low-level method used to
// display p (interpreted as an UTF8 string) on the canvas.
func (a CanvasID) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 0 {
		r, s := utf8.DecodeRune(p)
		a.WriteRune(r)
		p = p[s:]
	}
	return n, nil
}

// WriteRune is a low-level method used to display a single rune on the canvas.
func (a CanvasID) WriteRune(r rune) {
	cu := &canvases[a].cursor
	if r == '\n' {
		if cu.Interline == 0 {
			cu.Position.R += int16(float32(cu.Font.Height()) * 1.25)
		} else {
			cu.Position.R += cu.Interline
		}
		cu.Position.C = cu.Margin
		return
	}

	g := cu.Font.glyph(r)
	a.command(cmdText, 4, 1,
		int16(cu.Color), cu.Position.D, cu.Position.R-fonts[cu.Font].baseline,
		int16(g), cu.Position.C)
	cu.Position.C += glyphMap[g].w + cu.LetterSpacing
}

////////////////////////////////////////////////////////////////////////////////
