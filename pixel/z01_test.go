// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/palettes/c64"
	"github.com/cozely/cozely/color/palettes/cpc"
	"github.com/cozely/cozely/color/palettes/msx"
	"github.com/cozely/cozely/color/palettes/msx2"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var (
	canvas1   = pixel.Canvas(pixel.Resolution(320, 180))
	palette1a = color.PaletteFrom("graphics/mire")
	palette1b = color.PaletteFrom("graphics/srgb-gray")
)

var (
	mire      = pixel.Picture("graphics/mire")
	srgbGray  = pixel.Picture("graphics/srgb-gray")
	srgbRed   = pixel.Picture("graphics/srgb-red")
	srgbGreen = pixel.Picture("graphics/srgb-green")
	srgbBlue  = pixel.Picture("graphics/srgb-blue")
)

var mode int

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		err := cozely.Run(loop1{})
		if err != nil {
			t.Error(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type loop1 struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Enter() error {
	bindings.Load()
	context.Activate(1)

	mode = 0

	palette1a.Activate()
	return nil
}

func (loop1) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop1) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}

	if next.JustPressed(1) {
		mode++
		if mode > 1 {
			mode = 0
		}
		switch mode {
		case 0:
			palette1a.Activate()
		case 1:
			palette1b.Activate()
		}
	}

	if scene1.JustPressed(1) {
		c64.Palette.Activate()
	}
	if scene2.JustPressed(1) {
		cpc.Palette.Activate()
	}
	if scene3.JustPressed(1) {
		msx.Palette.Activate()
	}
	if scene4.JustPressed(1) {
		msx2.Palette.Activate()
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop1) Render() error {
	canvas1.Clear(0)
	sz := canvas1.Size()
	switch mode {
	case 0:
		pz := mire.Size()
		canvas1.Picture(mire, 0, coord.CR{0, 0})
		canvas1.Picture(mire, 0, coord.CR{0, sz.R - pz.R})
		canvas1.Picture(mire, 0, coord.CR{sz.C - pz.C, 0})
		canvas1.Picture(mire, 0, sz.Minus(pz))
	case 1:
		pz := srgbGray.Size()
		canvas1.Picture(srgbGray, 0, coord.CR{sz.C/2 - pz.C/2, 32})
		canvas1.Picture(srgbRed, 0, coord.CR{sz.C/4 - pz.C/2, 96})
		canvas1.Picture(srgbGreen, 0, coord.CR{sz.C/2 - pz.C/2, 96})
		canvas1.Picture(srgbBlue, 0, coord.CR{3*sz.C/4 - pz.C/2, 96})
	}
	canvas1.Display()
	return pixel.Err()
}

////////////////////////////////////////////////////////////////////////////////
