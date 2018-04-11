// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mouse

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/plane"
)

//------------------------------------------------------------------------------

// Position returns the current mouse position, relative to the game window.
// Updated at the start of each game loop iteration.
func Position() (x, y int32) {
	return int32(internal.MousePositionX), int32(internal.MousePositionY)
}

// Delta returns the mouse position relative to the last call of Delta.
func Delta() (dx, dy int32) {
	dx, dy = int32(internal.MouseDeltaX), int32(internal.MouseDeltaY)
	internal.MouseDeltaX, internal.MouseDeltaY = 0, 0
	return dx, dy
}

// SetRelative enables or disables the relative mode, where the mouse is
// hidden and mouse motions are continuously reported.
func SetRelative(enabled bool) error {
	return internal.MouseSetRelative(enabled)
}

// Relative returns true if the relative mode is enabled.
func RelativeMode() bool {
	return internal.MouseRelative()
}

//------------------------------------------------------------------------------

// SetSmoothing sets the smoothing factor for SmoothDelta.
func SetSmoothing(s float32) {
	smoothing = s
}

// SmoothDelta returns relative to the last call of SmoothDelta (or Delta), but
// smoothed to avoid jitter. The is best used with a fixed timestep (see
// glam.LoopStable).
func SmoothDelta() plane.Coord {
	dx, dy := Delta()
	d := plane.Coord{float32(dx), float32(dy)}
	smoothed = smoothed.Plus(d.Minus(smoothed).Times(smoothing))
	return smoothed
}

var smoothed plane.Coord
var smoothing = float32(0.4)

//------------------------------------------------------------------------------
