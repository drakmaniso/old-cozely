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

var (
	logo = pixel.NewPicture("graphics/logo")
	mire = pixel.NewPicture("graphics/mire")
	midgray = pixel.NewPicture("graphics/logo")
	midrgb = pixel.NewPicture("graphics/logo")
)

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

	err := pixel.LoadAssets()
	if err != nil {
		return err
	}

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
	scr := pixel.Screen()
	w, h := scr.Size().X, scr.Size().Y

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

	scr.PointList(0x05, pts[:]...)

	scr.Picture(logo, x, 10)

	scr.Picture(mire, 0, 0)
	scr.Picture(mire, w-32, 0)
	scr.Picture(mire, 0, h-32)
	scr.Picture(mire, w-32, h-32)

	scr.Picture(logo, w/2-32, 20)

	scr.Picture(midrgb, w/2-48, h/2-20)
	scr.Picture(midgray, w/2-16, h/2+20+8)

	scr.Point(0x18, w/2, 60)
	for xx := int16(4); xx < w; xx += 8 {
		scr.Point(0x18, xx, (x+xx)%h)
	}

	m := scr.Mouse()

	scr.Point(0x18, w/2, h/2)
	scr.Point(0x18, m.X, m.Y)
	scr.Line(0xE8, w/2, h/2, m.X, m.Y)
	scr.Point(0xE0, w/2, h/2)
	scr.Point(0xE0, m.X, m.Y)

	scr.Blit()

	return pixel.Err()
}

//------------------------------------------------------------------------------
