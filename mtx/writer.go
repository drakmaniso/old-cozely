// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mtx

//------------------------------------------------------------------------------

import (
	"fmt"

	micro "github.com/drakmaniso/glam/internal/microtext"
)

//------------------------------------------------------------------------------

type Writer struct {
	left, top     int
	right, bottom int
	x, y          int
	colour        byte
	clear         byte
	overwrite     bool
}

//------------------------------------------------------------------------------

func (w *Writer) Clip(left, top int, right, bottom int) {
	w.left, w.top = Clamp(left, top)
	w.right, w.bottom = clampBR(right, bottom)
	w.x, w.y = w.left, w.top
}

func (w *Writer) Clamp(x, y int) (int, int) {
	if x < 0 {
		x += w.right
	}
	if x < w.left {
		x = w.left
	}
	if x >= w.right {
		x = w.right - 1
	}

	if y < 0 {
		y += w.bottom
	}
	if y < w.top {
		y = w.top
	}
	if y >= w.bottom {
		y = w.bottom - 1
	}

	return x, y
}

//------------------------------------------------------------------------------

func (w *Writer) Locate(x, y int) {
	if w.right == 0 && w.bottom == 0 && w.left == 0 && w.top == 0 {
		w.right, w.bottom = micro.Size()
	}
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
	if w.right == 0 && w.bottom == 0 && w.left == 0 && w.top == 0 {
		w.right, w.bottom = micro.Size()
	}
	w.left, w.top = Clamp(w.left, w.top)
	w.right, w.bottom = clampBR(w.right, w.bottom)
	w.x, w.y = w.left, w.top

	for y := w.top; y < w.bottom; y++ {
		for x := w.left; x < w.right; x++ {
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
	if w.right == 0 && w.bottom == 0 && w.left == 0 && w.top == 0 {
		w.right, w.bottom = micro.Size()
	}
	x, y := w.x, w.y
	for _, b := range p {
		switch {
		case b <= ' ':
			switch b {
			case ' ':
				if w.overwrite {
					x++
					continue
				}

			case '\n':
				y++
				x = w.left
				continue

			case '\r':
				x = w.left
				continue

			case '\f':
				w.Clear()
				continue

			case '\v':
				x = w.left
				y = w.top
				continue

			case '\t':
				//TODO: should insert spaces if no overwrite
				x = w.left + (((x-w.left)/8)+1)*8
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
				b = '\x7F'
			}

		case b > '~':
			b = '\x7F'
		}

		if x >= w.left && x < w.right && y >= w.top && y < w.bottom {
			micro.Poke(x, y, b|w.colour)
		}
		x++
	}

	w.x, w.y = w.Clamp(x, y)

	micro.TextUpdated = true

	return len(p), nil
}

func (w *Writer) SetOverwrite(o bool) {
	w.overwrite = o
}

//------------------------------------------------------------------------------

func (w *Writer) Print(format string, a ...interface{}) {
	if w.right == 0 && w.bottom == 0 && w.left == 0 && w.top == 0 {
		w.right, w.bottom = micro.Size()
	} else {
		w.left, w.top = Clamp(w.left, w.top)
		w.right, w.bottom = clampBR(w.right, w.bottom)
	}
	w.x, w.y = w.Clamp(w.x, w.y)

	fmt.Fprintf(w, format, a...)
}

//------------------------------------------------------------------------------
