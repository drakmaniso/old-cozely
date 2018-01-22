// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
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

var logo, mire, midgray, midrgb pixel.Picture

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
	midgray = pixel.GetPicture("logo")
	midrgb = pixel.GetPicture("logo")

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
	if x >= pixel.Screen().Size().X {
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
	s := pixel.Screen()
	sz := s.Size()

	timer += glam.FrameTime()
	if timer > 0.25 {
		count++
		timer = 0.0
		if count%4 != 0 {
			palette.Index(1).Set(colour.LRGBA{1, 0, 0.5, 1})
		} else {
			palette.Index(1).Set(colour.LRGBA{1, 1, 1, 1})
		}
	}

	s.PointList(0x05, pts[:]...)

	s.Picture(logo, x, 10)

	s.Picture(mire, 0, 0)
	s.Picture(mire, sz.X-32, 0)
	s.Picture(mire, 0, sz.Y-32)
	s.Picture(mire, sz.X-32, sz.Y-32)

	s.Picture(logo, sz.X/2-32, 20)

	s.Picture(midrgb, sz.X/2-48, sz.Y/2-20)
	s.Picture(midgray, sz.X/2-16, sz.Y/2+20+8)

	s.Point(0x18, sz.X/2, 60)
	for xx := int16(4); xx < sz.X; xx += 8 {
		s.Point(0x18, xx, (x+xx)%sz.Y)
	}

	m := s.Mouse()

	s.Point(0x18, sz.X/2, sz.Y/2)
	s.Point(0x18, m.X, m.Y)
	s.Line(0xE8, sz.X/2, sz.Y/2, m.X, m.Y)
	s.Point(0xE0, sz.X/2, sz.Y/2)
	s.Point(0xE0, m.X, m.Y)

	s.Blit()

	return pixel.Err()
}

//------------------------------------------------------------------------------
