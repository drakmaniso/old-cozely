// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/drakmaniso/glam/mouse"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
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

type picLoop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (picLoop) Enter() error {
	palette.Load("graphics/mire")
	return nil
}

//------------------------------------------------------------------------------

func (picLoop) Draw() error {
	picScreen.Clear(0)
	s := picScreen.Size()
	switch picMode {
	case 0:
		ps := mire.Size()
		picScreen.Picture(mire, 0, 0, 0)
		picScreen.Picture(mire, 0, s.Y-ps.Y, 0)
		picScreen.Picture(mire, s.X-ps.X, 0, 0)
		picScreen.Picture(mire, s.X-ps.X, s.Y-ps.Y, 0)
	case 1:
		ps := srgbGray.Size()
		picScreen.Picture(srgbGray, s.X/2-ps.X/2, 32, 0)
		picScreen.Picture(srgbRed, s.X/4-ps.X/2, 96, 0)
		picScreen.Picture(srgbGreen, s.X/2-ps.X/2, 96, 0)
		picScreen.Picture(srgbBlue, 3*s.X/4-ps.X/2, 96, 0)
	}
	picScreen.Display()
	return pixel.Err()
}

//------------------------------------------------------------------------------

func (picLoop) MouseButtonDown(_ mouse.Button, _ int) {
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

//------------------------------------------------------------------------------
