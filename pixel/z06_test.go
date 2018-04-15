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

////////////////////////////////////////////////////////////////////////////////

func TestTest6(t *testing.T) {
	do(func() {
		err := cozely.Run(test6{})
		if err != nil {
			t.Error(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type test6 struct{}

////////////////////////////////////////////////////////////////////////////////

func (test6) Enter() error {
	bindings.Load()
	context.Activate(1)
	palette2.Activate()
	return nil
}

func (test6) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (test6) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (test6) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (test6) Render() error {
	canvas6.Clear(0)

	const corner = 3

	o := coord.CR{8, 8}
	s := coord.CR{24, 24}
	dx := coord.CR{32, 0}
	dy := coord.CR{0, 32}

	for i := int16(0); i < 13; i++ {
		canvas6.Box(6, 0, i, 0, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		canvas6.Box(0, 4, i, 0, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		canvas6.Box(6, 4, i, 0, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	o = o.Plus(dy)
	for i := int16(0); i < 13; i++ {
		canvas6.Box(4, 4, i, 0, o.Plus(dx.Times(i)), o.Plus(dx.Times(i)).Plus(s))
	}

	m := canvas6.FromWindow(input.Cursor.Position())
	canvas6.Point(18, 2, m)
	canvas6.Display()
	return nil
}

////////////////////////////////////////////////////////////////////////////////
