// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var canvas1 = pixel.Canvas(pixel.TargetResolution(320, 180))

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
	input.Load(bindings)
	context.Activate(1)

	mode = 0

	palette.Load("graphics/mire")
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
			palette.Load("graphics/mire")
		case 1:
			palette.Load("graphics/srgb-gray")
		}
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
