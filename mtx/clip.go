// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mtx

//------------------------------------------------------------------------------

import (
	"fmt"

	micro "github.com/drakmaniso/glam/internal/microtext"
)

//------------------------------------------------------------------------------

// A Clip is used to output text to a part of the MTX screen.
//
// It implements the io.Writer interface.
type Clip struct {
	// Bounds in screen coordinates (i.e. 0,0 is the top left of the screen; -1,-1
	// is at bottom right). Both corners are included in the resulting rectangle.
	Left, Top     int
	Right, Bottom int

	// Cursor position, relative to top-left corner
	x, y int

	// Should characters be drawn in reverse video?
	ReverseVideo bool

	// Character used to clear the clip
	ClearChar byte

	// Should space characters be opaque or "see-through"?
	TransparentSpace bool

	// Whether vertical and horizontal auto-scrolling are active
	HScroll bool
	VScroll bool
}

//------------------------------------------------------------------------------

// Clamp returns x and y clipped to the bounds of the Clip. Both the input and
// output are relative to the Clip bounds. If an input is negative, it is
// interpreted relative to the bottom right corner. The output is always
// positive.
func (cl *Clip) Clamp(x, y int) (int, int) {
	l, t, r, b := cl.Bounds()
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

// Bounds returns the top left and bottom right corners of the Clip, in
// screen coordinates.
func (cl *Clip) Bounds() (left, top int, right, bottom int) {
	left, top = Clamp(cl.Left, cl.Top)
	right, bottom = Clamp(cl.Right, cl.Bottom)
	if cl.Right == 0 && cl.Bottom == 0 && cl.Left == 0 && cl.Top == 0 {
		right, bottom = micro.Size()
		right--
		bottom--
	}
	return left, top, right, bottom
}

//------------------------------------------------------------------------------

// Locate positions the Clip cursor, in coordinates relative to its bounds.
// Positive coordinates are interpreted from the top left corner, while negative
// coordinates are interpreted from the bottom-right corner.
func (cl *Clip) Locate(x, y int) {
	cl.x, cl.y = cl.Clamp(x, y)
}

//------------------------------------------------------------------------------

// Clear erases all text from the clip.
//
// Note: you can customize the character used to erase by setting the ClearChar
// field of the clip.
func (cl *Clip) Clear() {
	l, t, r, b := cl.Bounds()
	cl.x, cl.y = 0, 0

	for y := t; y <= b; y++ {
		for x := l; x <= r; x++ {
			micro.Poke(x, y, cl.ClearChar)
		}
	}

	micro.Touch()
}

//------------------------------------------------------------------------------

// Write outputs text to the clip, starting at the cursor position. Special
// characters such as newline and tabs are recognized. It always returns the
// total number of bytes in the slice, even if some characters are out-of-bounds
// and clipped.
func (cl *Clip) Write(p []byte) (n int, err error) {
	l, t, r, b := cl.Bounds()
	sx, sy := r-l, b-t

	colour := byte(0x00)
	// Prepare reverse video mask
	if cl.ReverseVideo {
		colour = byte(0x80)
	}

	// Prepare space character
	spCh := byte(' ') | colour
	if cl.TransparentSpace {
		spCh = '\000'
	}

	x, y := cl.Clamp(cl.x, cl.y)
	for _, c := range p {
		switch {
		case ' ' < c && c <= '~':
			c |= colour

		case c == ' ':
			c = spCh

		case c == '\n':
			x = 0
			if y == sy && cl.VScroll {
				cl.Scroll(0, -1)
			} else {
				y++
			}
			continue

		case c == '\r':
			x = 0
			continue

		case c == '\f':
			cl.Clear()
			x, y = 0, 0
			continue

		case c == '\v':
			i := x
			if i < l {
				i = l
			}
			if y >= 0 && y <= sy {
				for ; i <= r; i++ {
					micro.Poke(l+i, l+y, cl.ClearChar)
				}
				micro.Touch()
			}
			continue

		case c == '\t':
			n := ((x/8)+1)*8 - x
			for i := 0; i < n; i++ {
				if x >= 0 && x <= sx && y >= 0 && y <= sy {
					micro.Poke(l+x+i, t+y, spCh)
				}
			}
			micro.Touch()
			x += n
			continue

		case c == '\b':
			x--
			continue

		case c == '\a':
			if colour != 0 {
				colour = 0
			} else {
				colour = 0x80
			}
			continue

		default:
			c = '\x7F' | colour
		}

		// Write the character
		if x >= 0 && x <= sx && y >= 0 && y <= sy {
			oc := micro.Peek(l+x, t+y)
			if oc != c {
				micro.Poke(l+x, t+y, c)
				micro.Touch()
			}
		}

		// Either scroll or advance the cursor
		if x == sx && cl.HScroll {
			cl.Scroll(-1, 0)
		} else {
			x++
		}
	}

	cl.x, cl.y = x, y

	return len(p), nil
}

//------------------------------------------------------------------------------

// Print writes formatted text on the clip. It is equivalent to a call to
// fmt.Fprintf on the clip.
func (cl *Clip) Print(format string, a ...interface{}) {
	fmt.Fprintf(cl, format, a...)
}

//------------------------------------------------------------------------------

// Scroll moves the content of the clip by a specified amount. Everything that
// move out of bounds is discarded, and the liberated space is cleared.
func (cl *Clip) Scroll(dx, dy int) {
	l, t, r, b := cl.Bounds()

	var x1, x2, x3, incX, y1, y2, y3, incY int
	var cmpX, cmpY func(int, int) bool
	if dx >= 0 {
		x1 = r
		if dx > r-l {
			dx = r - l + 1
		}
		x2 = l + dx
		x3 = l
		incX = -1
		cmpX = func(a, b int) bool { return a >= b }
	} else {
		x1 = l
		if dx < l-r {
			dx = l - r - 1
		}
		x2 = r + dx
		x3 = r
		incX = +1
		cmpX = func(a, b int) bool { return a <= b }
	}
	if dy >= 0 {
		y1 = b
		if dy > b-t {
			dy = b - t + 1
		}
		y2 = t + dy
		y3 = t
		incY = -1
		cmpY = func(a, b int) bool { return a >= b }
	} else {
		y1 = t
		if dy < t-b {
			dy = t - b - 1
		}
		y2 = b + dy
		y3 = b
		incY = +1
		cmpY = func(a, b int) bool { return a <= b }
	}

	for y := y1; cmpY(y, y2); y += incY {
		for x := x1; cmpX(x, x2); x += incX {
			Poke(x, y, Peek(x-dx, y-dy))
		}
		for x := x2 + incX; cmpX(x, x3); x += incX {
			Poke(x, y, cl.ClearChar)
		}
	}
	for y := y2 + incY; cmpY(y, y3); y += incY {
		for x := l; x <= r; x++ {
			Poke(x, y, cl.ClearChar)
		}
	}

	micro.Touch()
}

//------------------------------------------------------------------------------
