// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// A CanvasOption represents a configuration option used to change some
// parameters of a canvas see NewCanvas.
type CanvasOption = func(Canvas) error

//------------------------------------------------------------------------------

// TargetResolution defines a target resolution for the virtual screen.
func TargetResolution(w, h int16) CanvasOption {
	return func(cv Canvas) error {
		s := &canvases[cv]
		s.target.X, s.target.Y = w, h
		s.autozoom = true
		return nil
	}
}

// Zoom sets the size of the screen pixels (in window pixels).
func Zoom(z int32) CanvasOption {
	return func(cv Canvas) error {
		s := &canvases[cv]
		if z < 1 {
			return errors.New("pixel zoom null or negative")
		}
		s.pixel = z
		s.autozoom = false
		if internal.Running {
			Canvas(0).autoresize()
		}
		return nil
	}
}

//------------------------------------------------------------------------------
