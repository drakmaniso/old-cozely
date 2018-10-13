// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/pico8"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

type loop6 struct {
}

////////////////////////////////////////////////////////////////////////////////

func TestTest6(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop6{}
		l.setup()

		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (loop6) setup() {
	color.Load(&pico8.Palette)
	pixel.SetZoom(3)
}

func (loop6) Enter() {
	input.ShowMouse(false)
}

func (loop6) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop6) React() {
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
	}
}

func (loop6) Update() {
}

func (loop6) Render() {
	pixel.Clear(1)

	const corner = 3

	o := pixel.XY{8, 8}
	s := pixel.XY{16, 16}
	dx := pixel.XY{20, 0}
	dy := pixel.XY{0, 20}

	for i := int16(0); i < 16; i++ {
		pixel.Box(o.Plus(dx.Times(i)), s, 0, i, 7, 0)
	}

	o = o.Plus(dy)
	for i := int16(0); i < 16; i++ {
		pixel.Box(o.Plus(dx.Times(i)), s, 0, i, 13, 0)
	}

	o = o.Plus(dy)
	for i := int16(0); i < 16; i++ {
		pixel.Box(o.Plus(dx.Times(i)), s, 0, i, 7, 7)
	}

	o = o.Plus(dy)
	for i := int16(0); i < 16; i++ {
		pixel.Box(o.Plus(dx.Times(i)), s, 0, i, 7, 13)
	}

	m := pixel.XYof(input.MenuPointer.XY())
	pixel.Point(m, 8, 0)
}
