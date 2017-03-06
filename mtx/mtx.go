// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mtx

//------------------------------------------------------------------------------

import (
	"fmt"
	"github.com/drakmaniso/glam/color"
	micro "github.com/drakmaniso/glam/internal/microtext"
)

//------------------------------------------------------------------------------

// Size returns the size of the MTX screen.
func Size() (x, y int) {
	return micro.Size()
}

// Clamp restricts the given coordinates to the screen boundaries. Negative
// coordinates are anchored to the bottom right instead of the top left.
func Clamp(x, y int) (int, int) {
	sx, sy := micro.Size()

	if x < 0 {
		x += sx
		if x < 0 {
			x = 0
		}
	}
	if x >= sx {
		x = sx - 1
	}

	if y < 0 {
		y += sy
		if y < 0 {
			y = 0
		}
	}
	if y >= sy {
		y = sy - 1
	}

	return x, y
}

//------------------------------------------------------------------------------

// Clear erases the MTX screen.
func Clear() {
	for i := range micro.Text {
		micro.Text[i] = '\x00'
	}
	micro.TextUpdated = true
}

//------------------------------------------------------------------------------

// Peek returns the character at given coordinates.
func Peek(x, y int) byte {
	x, y = Clamp(x, y)
	return micro.Peek(x, y)
}

// Poke sets the character at given coordinates.
func Poke(x, y int, value byte) {
	x, y = Clamp(x, y)
	ov := micro.Peek(x, y)
	if value != ov {
		micro.Poke(x, y, value)
		micro.TextUpdated = true
	}
}

//------------------------------------------------------------------------------

// Print writes formatted text to the screen, at given coordinates.
func Print(x, y int, format string, a ...interface{}) {
	stdClip.Locate(x, y)

	fmt.Fprintf(&stdClip, format, a...)
}

var stdClip = Clip{
	Left: 0, Top: 0,
	Right: -1, Bottom: -1,
}

//------------------------------------------------------------------------------

// Color sets the foreground and background color.
func Color(fg, bg color.RGB) {
	micro.SetColor(fg, bg)
}

// Opaque selects wether the background is drawn or not. If set to false,
// letters are drawn without background. If set to true, letters are drawn on
// colored background. In all cases but blank space is always transparent.
//
// Note: You can toggle this settings in game, using Control+Alt+NumPadEnter.
func Opaque(t bool) {
	micro.SetBgAlpha(t)
}

// IsOpaque returns wether the background is currently opaque or not.
func IsOpaque() bool {
	return micro.GetBgAlpha()
}

//------------------------------------------------------------------------------
