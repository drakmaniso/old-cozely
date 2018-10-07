// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/window"
)

////////////////////////////////////////////////////////////////////////////////

// screen is the virtual screen. It is mad of a canvas and a filter, with depth
// information.
var screen = struct {
	resolution XY    // fixed resolution, or {0, 0} for fixed zoom
	zoom       int16 // in window pixels

	size   XY        // size of the canvas
	margin XY        // for fixed resolution only, = size - resolution
}{
	resolution: XY{},
	zoom:       2,
}

////////////////////////////////////////////////////////////////////////////////

// SetResolution defines a target resolution for the automatic resizing of
// the canvas.
//
// It guarantees that:
// - the canvas will never be smaller than the target resolution,
// - the target resolution will occupy as much screen as possible.
func SetResolution(r XY) {
	//TODO: allow runtime changes (defered to render?)
	if internal.Running {
		setErr(errors.New("Resolution must be called before starting the framework"))
		return
	}
	screen.resolution = r
	// if internal.Running {
	// 	resize()
	// }
}

// SetZoom sets the pixel size used to display the canvas.
func SetZoom(z int16) {
	//TODO: allow runtime changes (defered to render?)
	if internal.Running {
		setErr(errors.New("Resolution must be called before starting the framework"))
		return
	}
	if z < 1 {
		z = 1
	}
	screen.zoom = z
	screen.resolution = XY{}
	screen.margin = XY{}
	// if internal.Running {
	// 	resize()
	// }
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PixelResize = resize
}

func resize() {
	//TODO: use window.XY
	win := window.XY{internal.Window.Width, internal.Window.Height}

	if !screen.resolution.Null() {
		// Find best fit for pixel size
		p := win.SlashXY(window.XYof(screen.resolution))
		if p.X < p.Y {
			screen.zoom = p.X
		} else {
			screen.zoom = p.Y
		}
		if screen.zoom < 1 {
			screen.zoom = 1
		}
	}

	// Extend the screen to cover the window
	screen.size = XY(win.Slash(screen.zoom))
	adjustScreenTextures()

	// For fixed resolution, compute the margin and fix the size
	if (screen.resolution == XY{}) {
		screen.margin = XY{}
	} else {
		screen.margin = screen.size.Minus(screen.resolution).Slash(2)
	}
}

////////////////////////////////////////////////////////////////////////////////

// Clear sets the color of all pixels on the canvas; it also resets the filter
// of all pixels.
func Clear(c color.Index) {
	renderer.clear(c)
}

////////////////////////////////////////////////////////////////////////////////

// Resolution returns the current dimension of the canvas (in *canvas* pixels).
func Resolution() XY {
	if !screen.resolution.Null() {
		return screen.resolution
	}
	return screen.size
}

// Zoom returns the size of one canvas pixel, in *window* pixels.
func Zoom() int16 {
	return screen.zoom
}
