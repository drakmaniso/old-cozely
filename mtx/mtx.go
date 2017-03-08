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

// Color sets the foreground and background color.
func Color(fg, bg color.RGB) {
	micro.SetColor(fg, bg)
}

// Opaque selects wether the background is drawn or not. When set to false, the
// text is drawn directly over the game screen. When set to true, it is are
// drawn over a colored background. Note that blank space (i.e. areas without
// MTX content, not white space) is always transparent.
//
// Note: It's possiblt to toggle this settings in game, using
// Control+Alt+NumPadEnter.
func Opaque(t bool) {
	micro.SetBgAlpha(t)
}

// IsOpaque returns wether the background is currently opaque or not.
func IsOpaque() bool {
	return micro.GetBgAlpha()
}

//------------------------------------------------------------------------------

// Clear removes all MTX content.
//
// Note that the text won't disappear from the screen until the screen itself is
// cleared or drawn over.
func Clear() {
	micro.Clear()
	micro.Touch()
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
		micro.Touch()
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

// ShowFrameTime enable or disable a mini widget showing average frame time.
func ShowFrameTime(enable bool, x, y int, opaque bool) {
	micro.ShowFrameTime(enable, x, y, opaque)
}

//------------------------------------------------------------------------------
