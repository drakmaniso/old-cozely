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

var picScreen = pixel.Canvas(pixel.TargetResolution(320, 180))

var (
	mire      = pixel.Picture("graphics/mire")
	srgbGray  = pixel.Picture("graphics/srgb-gray")
	srgbRed   = pixel.Picture("graphics/srgb-red")
	srgbGreen = pixel.Picture("graphics/srgb-green")
	srgbBlue  = pixel.Picture("graphics/srgb-blue")
)

var picMode int

////////////////////////////////////////////////////////////////////////////////

func TestPicture_basic(t *testing.T) {
	do(func() {
		err := cozely.Run(picLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type picLoop struct{}

////////////////////////////////////////////////////////////////////////////////

func (picLoop) Enter() error {
	input.Load(testBindings)
	testContext.Activate(1)

	palette.Load("graphics/mire")
	return nil
}

func (picLoop) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (picLoop) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}

	if next.JustPressed(1) {
		picMode++
		if picMode > 1 {
			picMode = 0
		}
		switch picMode {
		case 0:
			palette.Load("graphics/mire")
		case 1:
			palette.Load("graphics/srgb-gray")
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (picLoop) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (picLoop) Render() error {
	picScreen.Clear(0)
	sz := picScreen.Size()
	switch picMode {
	case 0:
		pz := mire.Size()
		picScreen.Picture(mire, 0, coord.CR{0, 0})
		picScreen.Picture(mire, 0, coord.CR{0, sz.R - pz.R})
		picScreen.Picture(mire, 0, coord.CR{sz.C - pz.C, 0})
		picScreen.Picture(mire, 0, sz.Minus(pz))
	case 1:
		pz := srgbGray.Size()
		picScreen.Picture(srgbGray, 0, coord.CR{sz.C/2 - pz.C/2, 32})
		picScreen.Picture(srgbRed, 0, coord.CR{sz.C/4 - pz.C/2, 96})
		picScreen.Picture(srgbGreen, 0, coord.CR{sz.C/2 - pz.C/2, 96})
		picScreen.Picture(srgbBlue, 0, coord.CR{3*sz.C/4 - pz.C/2, 96})
	}
	picScreen.Display()
	return pixel.Err()
}

////////////////////////////////////////////////////////////////////////////////

//TODO:
// func (picLoop) MouseButtonDown(_ mouse.Button, _ int) {
// 	picMode++
// 	if picMode > 1 {
// 		picMode = 0
// 	}
// 	switch picMode {
// 	case 0:
// 		palette.Load("graphics/mire")
// 	case 1:
// 		palette.Load("graphics/srgb-gray")
// 	}
// }

////////////////////////////////////////////////////////////////////////////////
