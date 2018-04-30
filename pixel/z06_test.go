// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var canvas6 = pixel.Canvas(pixel.Zoom(3))

type test6 struct{}

////////////////////////////////////////////////////////////////////////////////

func TestTest6(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		input.Load(bindings)
		err := cozely.Run(test6{})
		if err != nil {
			t.Error(err)
		}
	})
}

func (test6) Enter() {
	input.ShowMouse(false)
	palette2.Activate()
}

func (test6) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (test6) React() {
	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (test6) Update() {
}

func (test6) Render() {
	canvas6.Clear(0)

	const corner = 3

	o := coord.CR{8, 8}
	s := coord.CR{24, 24}
	dx := coord.CR{32, 0}
	dy := coord.CR{0, 32}

	for i := int16(0); i < 13; i++ {
		canvas6.Box(6, 0, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		canvas6.Box(0, 4, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		canvas6.Box(6, 4, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		canvas6.Box(4, 4, i, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	m := canvas6.FromWindow(cursor.XY(0).CR())
	canvas6.Point(18, m)
	canvas6.Display()
}

////////////////////////////////////////////////////////////////////////////////
