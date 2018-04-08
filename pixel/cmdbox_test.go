// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
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

	o := plane.Pixel{8, 8}
	s := plane.Pixel{24, 24}
	dx := plane.Pixel{32, 0}
	dy := plane.Pixel{0, 32}

	for i := int16(0); i < 13; i++ {
		boxScreen.Box(6, 0, i, 0, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		boxScreen.Box(0, 4, i, 0, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		boxScreen.Box(6, 4, i, 0, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		boxScreen.Box(4, 4, i, 0, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	m := boxScreen.Mouse()
	boxScreen.Point(18, 2, m)
	boxScreen.Display()
	return nil
}

//------------------------------------------------------------------------------
