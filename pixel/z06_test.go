// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

type loop6 struct {
	canvas  pixel.CanvasID
	scene pixel.SceneID
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
	a.canvas = pixel.Canvas(pixel.Zoom(3))
	a.scene = pixel.Scene()
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
	a.canvas.Clear(0)

	const corner = 3

	o := coord.CR{8, 8}
	s := coord.CR{24, 24}
	dx := coord.CR{32, 0}
	dy := coord.CR{0, 32}

	for i := int16(0); i < 13; i++ {
		a.scene.Box(6, 0, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		a.scene.Box(0, 4, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		a.scene.Box(6, 4, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		a.scene.Box(4, 4, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	m := a.canvas.FromWindow(cursor.XY(0).CR())
	a.scene.Point(18, m)

	a.canvas.Display(a.scene)
}

////////////////////////////////////////////////////////////////////////////////
