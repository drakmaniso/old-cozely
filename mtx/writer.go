// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mtx

//------------------------------------------------------------------------------

import (
	"fmt"

	micro "github.com/drakmaniso/glam/internal/microtext"
)

//------------------------------------------------------------------------------

// A Writer is to output text to a part of the MTX screen.
type Writer struct {
	// Bounds in screen coordinates (i.e. 0,0 is the top left of the screen; -1,-1
	// is at bottom right). Both corners are included in the resulting rectangle.
	Left, Top     int
	Right, Bottom int

	// Cursor position, relative to top-left corner
	x, y int

	// ReverseVideo mask
	colour byte

	// Character used to clear
	clear byte

	// Whether whitespace is opaque or transparent
	overwrite bool
}

//------------------------------------------------------------------------------

// Clamp returns x and y clipped to the bounds of the Writer. Both the input and
// output are relative to the Writer bounds. If an input is negative, it is
// interpreted relative to the bottom right corner. The output is always
// positive.
func (w *Writer) Clamp(x, y int) (int, int) {
	l, t, r, b := w.Bounds()
	sx, sy := r-l, b-t

	if x < 0 {
		x += sx
	}
	if x < 0 {
		x = 0
	}
	if x > sx {
		x = sx
	}

	if y < 0 {
		y += sy
	}
	if y < 0 {
		y = 0
	}
	if y > sy {
		y = sy
	}

	return x, y
}

// Bounds returns the top left and bottom right corners of the Writer, in
// screen coordinates.
func (w *Writer) Bounds() (left, top int, right, bottom int) {
	left, top = Clamp(w.Left, w.Top)
	right, bottom = Clamp(w.Right, w.Bottom)
	if w.Right == 0 && w.Bottom == 0 && w.Left == 0 && w.Top == 0 {
		right, bottom = micro.Size()
		right--
		bottom--
	}
	return left, top, right, bottom
}

//------------------------------------------------------------------------------

// Locate positions the Writer cursor, in coordinates relative to the Writer
// bounds. Positive coordinates are interpreted from the top left corner, while
// negative coordinates are interpreted from the bottom-right corner.
func (w *Writer) Locate(x, y int) {
	w.x, w.y = w.Clamp(x, y)
}

//------------------------------------------------------------------------------

func (w *Writer) ReverseVideo(r bool) {
	if r {
		w.colour = 0x80
	} else {
		w.colour = 0x00
	}
}

//------------------------------------------------------------------------------

func (w *Writer) Clear() {
	l, t, r, b := w.Bounds()
	w.x, w.y = 0, 0

	for y := t; y <= b; y++ {
		for x := l; x <= r; x++ {
			micro.Poke(x, y, w.clear)
		}
	}

	micro.TextUpdated = true
}

func (w *Writer) SetClearChar(c byte) {
	w.clear = c
}

//------------------------------------------------------------------------------

func (w *Writer) Write(p []byte) (n int, err error) {
	l, t, r, b := w.Bounds()
	sx, sy := r-l, b-t

	x, y := w.x, w.y
	for _, c := range p {
		switch {
		case c <= ' ':
			switch c {
			case ' ':
				if w.overwrite {
					x++
					continue
				}

			case '\n':
				y++
				x = 0
				continue

			case '\r':
				x = 0
				continue

			case '\f':
				w.Clear()
				continue

			case '\v':
				i := x
				if i < l {
					i = l
				}
				if y >= t && y <= b {
					for ; i <= r; i++ {
						micro.Poke(i, y, w.clear)
					}
				}
				continue

			case '\t':
				//TODO: should insert spaces if no overwrite
				x = ((x / 8) + 1) * 8
				continue

			case '\b':
				x--
				continue

			case '\a':
				if w.colour != 0 {
					w.colour = 0
				} else {
					w.colour = 0x80
				}
				continue

			default:
				c = '\x7F'
			}

		case c > '~':
			c = '\x7F'
		}

		if x >= 0 && x <= sx && y >= 0 && y <= sy {
			micro.Poke(l+x, t+y, c|w.colour)
		}
		x++
	}

	w.x, w.y = x, y

	micro.TextUpdated = true

	return len(p), nil
}

func (w *Writer) SetOverwrite(o bool) {
	w.overwrite = o
}

//------------------------------------------------------------------------------

func (w *Writer) Print(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}

//------------------------------------------------------------------------------
