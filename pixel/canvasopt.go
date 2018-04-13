// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// A CanvasOption represents a configuration option used to change some
// parameters of a canvas see NewCanvas.
type CanvasOption = func(CanvasID) error

////////////////////////////////////////////////////////////////////////////////

// Resolution defines a target resolution for the automatic resizing of
// the canvas.
//
// It guarantees that:
// - the canvas will never be smaller than the target resolution,
// - the target resolution will occupy as much screen as possible.
func Resolution(w, h int16) CanvasOption {
	return func(cv CanvasID) error {
		s := &canvases[cv]
		s.resolution.C, s.resolution.R = w, h
		s.fixedres = true
		return nil
	}
}

// Zoom sets the pixel size used to display the canvas.
func Zoom(z int16) CanvasOption {
	return func(cv CanvasID) error {
		s := &canvases[cv]
		if z < 1 {
			return errors.New("pixel zoom null or negative")
		}
		s.pixel = z
		s.fixedres = false
		if internal.Running {
			CanvasID(0).autoresize()
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////
