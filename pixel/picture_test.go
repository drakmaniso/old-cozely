// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/input"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
)

//------------------------------------------------------------------------------

var picScreen = pixel.NewCanvas(pixel.TargetResolution(320, 180))

var (
	mire      = pixel.NewPicture("graphics/mire")
	srgbGray  = pixel.NewPicture("graphics/srgb-gray")
	srgbRed   = pixel.NewPicture("graphics/srgb-red")
	srgbGreen = pixel.NewPicture("graphics/srgb-green")
	srgbBlue  = pixel.NewPicture("graphics/srgb-blue")
)

var picMode int

//------------------------------------------------------------------------------

func TestPicture_basic(t *testing.T) {
	do(func() {
		err := glam.Run(picLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type picLoop struct{}

//------------------------------------------------------------------------------

func (picLoop) Enter() error {
	input.Load(testBindings)
	testContext.Activate(1)

	palette.Load("graphics/mire")
	return nil
}

func (picLoop) Leave() error { return nil }

//------------------------------------------------------------------------------

func (picLoop) React() error {
	if quit.JustPressed(1) {
		glam.Stop()
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

//------------------------------------------------------------------------------

func (picLoop) Update() error { return nil }

//------------------------------------------------------------------------------

func (picLoop) Draw() error {
	picScreen.Clear(0)
	sz := picScreen.Size()
	switch picMode {
	case 0:
		pz := mire.Size()
		picScreen.Picture(mire, 0, plane.Pixel{0, 0})
		picScreen.Picture(mire, 0, plane.Pixel{0, sz.Y - pz.Y})
		picScreen.Picture(mire, 0, plane.Pixel{sz.X - pz.X, 0})
		picScreen.Picture(mire, 0, sz.Minus(pz))
	case 1:
		pz := srgbGray.Size()
		picScreen.Picture(srgbGray, 0, plane.Pixel{sz.X/2 - pz.X/2, 32})
		picScreen.Picture(srgbRed, 0, plane.Pixel{sz.X/4 - pz.X/2, 96})
		picScreen.Picture(srgbGreen, 0, plane.Pixel{sz.X/2 - pz.X/2, 96})
		picScreen.Picture(srgbBlue, 0, plane.Pixel{3*sz.X/4 - pz.X/2, 96})
	}
	picScreen.Display()
	return pixel.Err()
}

//------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------

func (picLoop) Resize()  {}
func (picLoop) Show()    {}
func (picLoop) Hide()    {}
func (picLoop) Focus()   {}
func (picLoop) Unfocus() {}
func (picLoop) Quit() {
	glam.Stop()
}

//------------------------------------------------------------------------------
