// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// A Config represents a configuration option used to change some parameters of
// the pixel package: see glam.Configure.
type Config = func() error

//------------------------------------------------------------------------------

// TargetResolution defines a target resolution for the virtual screen.
func TargetResolution(w, h int16) Config {
	return func() error {
		screen.target.X, screen.target.Y = w, h
		screen.autozoom = true
		return nil
	}
}

// Zoom sets the size of the screen pixels (in window pixels).
func Zoom(z int32) Config {
	return func() error {
		if z < 1 {
			return errors.New("pixel zoom null or negative")
		}
		screen.pixel = z
		screen.autozoom = false
		if internal.Running {
			internal.ResizeScreen()
		}
		return nil
	}
}

//------------------------------------------------------------------------------

// AutoPalette enables or disable the automatic addition of unkown colors when
// loading indexed images.
func AutoPalette(auto bool) Config {
	// TODO: automatically disable when using palette.Change or palette.New.
	return func() error {
		autoPalette = auto
		return nil
	}
}

//------------------------------------------------------------------------------
