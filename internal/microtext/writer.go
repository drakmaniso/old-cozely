// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package microtext

//------------------------------------------------------------------------------

type Writer struct {
	Left, Top     int
	Right, Bottom int
	X, Y          int
	colour        byte
}

func (w *Writer) Write(p []byte) (n int, err error) {
	x, y := w.X, w.Y
	for _, b := range p {
		switch {
		case b < ' ':
			switch b {
			case '\n':
				y++
				x = w.Left
				continue

			case '\a':
				if w.colour != 0 {
					w.colour = 0
				} else {
					w.colour = 0x80
				}
				continue

			default:
				b = '\x7F'
			}

		case b > '~':
			b = '\x7F'
		}

		if x < w.Right && y < w.Bottom {
			Text[x+y*int(screen.nbCols)] = b | w.colour
		}
		x++
	}

	w.X, w.Y = w.Clamp(x, y)

	TextUpdated = true

	return len(p), nil
}

//------------------------------------------------------------------------------

func (w *Writer) Clamp(x, y int) (int, int) {
	if x < 0 {
		x += w.Right
		if x < w.Left {
			x = w.Left
		}
	}
	if x >= w.Right {
		x = w.Right - 1
	}

	if y < 0 {
		y += w.Bottom
		if y < w.Top {
			y = w.Top
		}
	}
	if y >= w.Bottom {
		y = w.Bottom - 1
	}

	return x, y
}

//------------------------------------------------------------------------------
