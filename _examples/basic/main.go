// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var logo, mire, midgray, midrgb *pixel.Picture

//------------------------------------------------------------------------------

func main() {
	err := glam.Run(loop{})
	if err != nil {
		glam.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
}

func (loop) Setup() error {
	palette.Change("MSX2")

	logo = pixel.GetPicture("pictures/logo")
	mire = pixel.GetPicture("pictures/mire")
	midgray = pixel.GetPicture("pictures/midgray")
	midrgb = pixel.GetPicture("pictures/midrgb")

	return pixel.Err()
}

func (loop) Update() error {
	x++
	if x >= pixel.ScreenSize().X {
		x = -64
	}
	return nil
}

var x = int16(0)

var (
	timer = 0.0
	count = 0
)

func (loop) Draw(delta, _ float64) error {
	timer += delta
	if timer > 0.25 {
		count++
		timer = 0.0
		if count%2 != 0 {
			palette.Index(2).Set(colour.RGBA{1, 1, 1, 1})
		} else {
			palette.Index(2).Set(colour.RGBA{1, 0, 0.5, 1})
		}
	}

	logo.Paint(x, 10)

	s := pixel.ScreenSize()

	mire.Paint(0, 0)
	mire.Paint(s.X-32, 0)
	mire.Paint(0, s.Y-32)
	mire.Paint(s.X-32, s.Y-32)

	logo.Paint(s.X/2-32, 20)

	midrgb.Paint(s.X/2-48, s.Y/2-20)
	midgray.Paint(s.X/2-16, s.Y/2+20+8)

	return pixel.Err()
}

//------------------------------------------------------------------------------
