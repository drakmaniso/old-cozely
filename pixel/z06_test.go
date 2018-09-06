// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/window"
)

////////////////////////////////////////////////////////////////////////////////

type loop6 struct {
	palette color.PaletteID
}

////////////////////////////////////////////////////////////////////////////////

func TestTest6(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop6{}
		l.declare()

		input.Load(bindings)
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (a *loop6) declare() {
	pixel.SetZoom(3)
	a.palette = color.PaletteFrom("graphics/shape1")
}

func (a *loop6) Enter() {
	input.ShowMouse(false)
	a.palette.Activate()
}

func (loop6) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop6) React() {
	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop6) Update() {
}

func (a *loop6) Render() {
	pixel.Clear(0)

	const corner = 3

	o := pixel.XY{8, 8}
	s := pixel.XY{24, 24}
	dx := pixel.XY{32, 0}
	dy := pixel.XY{0, 32}

	for i := int16(0); i < 13; i++ {
		pixel.Box(6, 0, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		pixel.Box(0, 4, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		pixel.Box(6, 4, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		pixel.Box(4, 4, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	m := pixel.ToCanvas(window.XYof(cursor.XY(0)))
	pixel.Point(18, m)
}

////////////////////////////////////////////////////////////////////////////////
