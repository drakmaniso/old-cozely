// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"math/rand"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var logo, mire, midgray, midrgb *pixel.Picture

var pts [1024]pixel.Coord

//------------------------------------------------------------------------------

func main() {
	glam.Configure(
		glam.TimeStep(1.0/60),
		pixel.TargetResolution(320, 180),
		pixel.AutoPalette(false),
	)

	err := glam.Run(setup, loop{})
	if err != nil {
		glam.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

func setup() error {
	palette.Change("MSX2")

	err := pixel.Load("graphics")
	if err != nil {
		return err
	}

	logo = pixel.GetPicture("logo")
	mire = pixel.GetPicture("mire")
	midgray = pixel.GetPicture("midgray")
	midrgb = pixel.GetPicture("midrgb")

	for i := range pts {
		pts[i].X = int16(rand.Intn(320))
		pts[i].Y = int16(rand.Intn(180))
	}

	return pixel.Err()
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
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

func (loop) Draw() error {
	timer += glam.FrameTime()
	if timer > 0.25 {
		count++
		timer = 0.0
		if count%4 != 0 {
			palette.Index(1).Set(colour.LRGBA{1, 1, 1, 1})
		} else {
			palette.Index(1).Set(colour.LRGBA{1, 0, 0.5, 1})
		}
	}

	if count%3 != 0 {
		pixel.PointList(colour.SRGB{0.25, 0.25, 0.25}, pts[:]...)
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

	pixel.Point(colour.LRGB{1, 0.5, 0}, s.X/2, 60)
	pixel.Point(colour.LRGB{0, 0.5, 1}, x, 100)
	m := pixel.Mouse()
	pixel.Point(colour.LRGB{1, 0, 0}, m.X, m.Y)

	return pixel.Err()
}

//------------------------------------------------------------------------------
