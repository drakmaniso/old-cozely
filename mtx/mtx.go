// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mtx

//------------------------------------------------------------------------------

import (
	"fmt"

	micro "github.com/drakmaniso/carol/internal/microtext"
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

// SetReverseVideo activates or deactivates reverse video.
func SetReverseVideo(t bool) {
	micro.SetReverseVideo(t)
}

// GetReverseVideo returns true if microtext is in reverse video.
func GetReverseVideo() bool {
	return micro.GetReverseVideo()
}

// ToggleReverseVideo toggles reverse video mode.
func ToggleReverseVideo() {
	micro.ToggleReverseVideo()
}

//------------------------------------------------------------------------------

// Clear removes all MTX content.
//
// Note that the text won't disappear from the screen until the screen itself is
// cleared or drawn over.
func Clear() {
	micro.Clear()
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
	micro.Poke(x, y, value)
}

//------------------------------------------------------------------------------

// Locate positions the cursor, in coordinates relative to screen bounds.
// Positive coordinates are interpreted from the top left corner, while negative
// coordinates are interpreted from the bottom-right corner.
func Locate(x, y int) {
	stdClip.Locate(x, y)
}

// Print writes formatted text to the screen, at the current cursor position. If
// the bottom of the screen is reached, the whole screen is scrolled updward.
func Print(format string, a ...interface{}) {
	fmt.Fprintf(&stdClip, format, a...)
}

var stdClip = Clip{
	Left: 0, Top: 0,
	Right: -1, Bottom: -1,
}

//------------------------------------------------------------------------------

// ShowFrameTime enable or disable a mini widget showing average frame time.
func ShowFrameTime(enable bool, x, y int) {
	micro.ShowFrameTime(enable, x, y)
}

//------------------------------------------------------------------------------
