// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var boxScreen = pixel.NewCanvas(pixel.Zoom(3))

//------------------------------------------------------------------------------

func TestPaint_box(t *testing.T) {
	do(func() {
		err := glam.Run(boxLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type boxLoop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (boxLoop) Enter() error {
	palette.Load("graphics/shape1")
	return nil
}

//------------------------------------------------------------------------------

func (boxLoop) Draw() error {
	boxScreen.Clear(0)

	const corner = 3

	const w = 24
	const dx = 32

	for i := int16(0); i < 13; i++ {
		boxScreen.Box(6, 0, i, 0, 8+i*dx, 8, 8+w+i*dx, 8+w)
	}

	for i := int16(0); i < 13; i++ {
		boxScreen.Box(0, 4, i, 0, 8+i*dx, 8+dx, 8+w+i*dx, 8+dx+w)
	}

	for i := int16(0); i < 13; i++ {
		boxScreen.Box(6, 4, i, 0, 8+i*dx, 8+2*dx, 8+w+i*dx, 8+2*dx+w)
	}

	for i := int16(0); i < 13; i++ {
		boxScreen.Box(4, 4, i, 0, 8+i*dx, 8+3*dx, 8+w+i*dx, 8+3*dx+w)
	}

	// boxScreen.Box(6, 0, 1, 2, 1, 32, 110, 32+6, 110+6)
	// boxScreen.Box(6, 4, 1, corner, 0, 8, 100, 120, 100+32)

	m := boxScreen.Mouse()
	boxScreen.Point(18, m.X, m.Y, 2)
	boxScreen.Display()
	return nil
}

//------------------------------------------------------------------------------
