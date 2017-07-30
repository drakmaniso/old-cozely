// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package overlay

//------------------------------------------------------------------------------

import (
	"fmt"

	"github.com/drakmaniso/glam/internal/overl"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

type Overlay struct {
	ovr *overl.Overlay

	//
	transparent bool

	// Cursor position
	x, y      int
	highlight bool
}

//------------------------------------------------------------------------------

func Create(position pixel.Coord, columns, rows int, transparent bool) *Overlay {
	var o Overlay

	o.ovr = overl.Create(position, columns, rows)
	o.transparent = transparent
	o.Clear()

	return &o
}

//------------------------------------------------------------------------------

// Poke sets the character at given coordinates.
func (o *Overlay) Poke(x, y int, c byte) {
	o.ovr.Poke(x, y, c)
}

//------------------------------------------------------------------------------

// Clamp takes a character posision, and returns it clipped to the size of the
// overlay.
//
// If x is negative, it is intepreted relative to the right border. If
// y is negative, it is interpreted relative to the bottom border. In other
// words: (0, 0) is the character at top left, and (-1, -1) the character at
// bottom right.
func (o *Overlay) Clamp(x, y int) (int, int) {
	sx, sy := o.ovr.Size()

	if x < 0 {
		x += sx
	}
	if x < 0 {
		x = 0
	}
	if x >= sx {
		x = sx - 1
	}

	if y < 0 {
		y += sy
	}
	if y < 0 {
		y = 0
	}
	if y >= sy {
		y = sy - 1
	}

	return x, y
}

//------------------------------------------------------------------------------

func (o *Overlay) Size() (columns, rows int) {
	return o.ovr.Size()
}

//------------------------------------------------------------------------------

// Clear erases all the overlay content.
func (o *Overlay) Clear() {
	sx, sy := o.ovr.Size()
	o.x, o.y = 0, 0

	clr := byte(' ')
	if o.transparent {
		clr = '\x00'
	}

	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			o.ovr.Poke(x, y, clr)
		}
	}
}

//------------------------------------------------------------------------------

// Locate positions the cursor.
//
// Note that the coordinates are clamped first.
func (o *Overlay) Locate(x, y int) {
	o.x, o.y = o.Clamp(x, y)
}

//------------------------------------------------------------------------------

// Scroll moves the content of the overlay by a specified amount. Everything
// that move out of bounds is discarded, and the liberated space is cleared.
func (o *Overlay) Scroll(dx, dy int) {
	sx, sy := o.ovr.Size()

	clr := byte(' ')
	if o.transparent {
		clr = '\x00'
	}

	var x1, x2, x3, incX, y1, y2, y3, incY int
	var cmpX, cmpY func(int, int) bool
	if dx >= 0 {
		x1 = sx - 1
		if dx > sx {
			dx = sx
		}
		x2 = dx
		x3 = 0
		incX = -1
		cmpX = func(a, b int) bool { return a >= b }
	} else {
		x1 = 0
		if dx < -sx {
			dx = -sx
		}
		x2 = sx - 1 + dx
		x3 = sx - 1
		incX = +1
		cmpX = func(a, b int) bool { return a <= b }
	}
	if dy >= 0 {
		y1 = sy - 1
		if dy > sy {
			dy = sy
		}
		y2 = dy
		y3 = 0
		incY = -1
		cmpY = func(a, b int) bool { return a >= b }
	} else {
		y1 = 0
		if dy < -sy {
			dy = -sy
		}
		y2 = sy - 1 + dy
		y3 = sy - 1
		incY = +1
		cmpY = func(a, b int) bool { return a <= b }
	}

	for y := y1; cmpY(y, y2); y += incY {
		for x := x1; cmpX(x, x2); x += incX {
			o.ovr.Poke(x, y, o.ovr.Peek(x-dx, y-dy))
		}
		for x := x2 + incX; cmpX(x, x3); x += incX {
			o.ovr.Poke(x, y, clr)
		}
	}
	for y := y2 + incY; cmpY(y, y3); y += incY {
		for x := (0); x <= (sx - 1); x++ {
			o.ovr.Poke(x, y, clr)
		}
	}
}

//------------------------------------------------------------------------------

// Write outputs text to the clip, starting at the cursor position. Special
// characters such as newline and tabs are recognized. It always returns the
// total number of bytes in the slice, even if some characters are out-of-bounds
// and clipped.
//
// Special characters:
// - '\a' toggle highlight
// - '\b' move cursor one character left
// - '\f' blank space (i.e. fully transparent)
// - '\n' newline
// - '\r' move cursor to beginning of line
// - '\t' tabulation
// - '\v' clear until end of line
func (o *Overlay) Write(p []byte) (n int, err error) {
	sx, sy := o.ovr.Size()

	colour := byte(0x00)
	// Prepare highlight mask
	if o.highlight {
		colour = byte(0x80)
	}

	clr := byte(' ')
	if o.transparent {
		clr = '\x00'
	}

	x, y := o.Clamp(o.x, o.y)

	for _, c := range p {
		switch {
		case ' ' <= c && c <= '~':
			c |= colour

		case c == '\r':
			x = 0
			continue

		case c == '\f':
			c = '\x00'

		case c == '\n':
			// First, clear to end of line
			if 0 <= y && y <= sy-1 {
				i := x
				if i < 0 {
					i = 0
				}
				for ; i <= sx-1; i++ {
					o.ovr.Poke(i, y, clr)
				}
			}
			// Go to next line
			x = 0
			if y == sy-1 {
				o.Scroll(0, -1)
			} else {
				y++
			}
			continue
		case c == '\v':
			// Clear to end of line
			if 0 <= y && y <= sy-1 {
				i := x
				if i < 0 {
					i = 0
				}
				for ; i <= sx-1; i++ {
					o.ovr.Poke(i, y, clr)
				}
			}
			continue

		case c == '\t':
			if 0 <= y && y <= sy-1 {
				n := ((x/8)+1)*8 - x
				for i := 0; i < n && x+i <= sy-1; i++ {
					o.ovr.Poke(x+i, y, ' ')
				}
				x += n
			}
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

		// Handle out of bounds
		var xx, yy int

		if x < 0 {
			c = '~' + 1
			xx = 0
		} else if x > sx-1 {
			c = '~' + 1
			xx = sx - 1
		} else {
			xx = x
		}

		if y < 0 {
			c = '~' + 1
			yy = 0
		} else if y > sy-1 {
			c = '~' + 1
			yy = sy - 1
		} else {
			yy = y
		}

		o.ovr.Poke(xx, yy, c)

		// Either scroll horizontally or move cursor
		if x == sx-1 {
			x = 0
			if y == sy-1 {
				o.Scroll(0, -1)
			} else {
				y++
			}
		} else {
			x++
		}
	}

	o.x, o.y = x, y

	return len(p), nil
}

//------------------------------------------------------------------------------

// Print writes formatted text on the clip. It is equivalent to a call to
// fmt.Fprintf on the clip.
func (o *Overlay) Print(format string, a ...interface{}) {
	fmt.Fprintf(o, format, a...)
}

//------------------------------------------------------------------------------
