// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

type loop5 struct {
	points []pixel.XY
}

////////////////////////////////////////////////////////////////////////////////

func TestTest5(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop5{}
		l.declare()

		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (a *loop5) declare() {
	pixel.SetResolution(pixel.XY{128, 128})

	a.points = []pixel.XY{
		{4, 4},
		{4 + 1, 4 + 20},
		{4 + 1 + 20, 4 + 20 - 1},
		{16, 32},
	}
}

func (a *loop5) Enter() {
	input.ShowMouse(false)
}

func (loop5) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (a *loop5) React() {
	if quit.Pushed() {
		cozely.Stop(nil)
	}

	if next.Pushed() {
		m := pixel.XYof(cursor.XY())
		a.points = append(a.points, m)
	}

	if previous.Pushed() {
		if len(a.points) > 0 {
			a.points = a.points[:len(a.points)-1]
		}
	}
}

func (loop5) Update() {
}

func (a *loop5) Render() {
	pixel.Clear(1)
	m := pixel.XYof(cursor.XY())
	if !scenes[3].Pressed() {
		for i := 0; i < len(a.points) - 2; i++ {
			pixel.Triangle(a.points[i], a.points[i+1], a.points[i+2], 0, 2)
		}
	}
	if !scenes[2].Pressed() {
		for i := 0; i < len(a.points) - 1; i++ {
			pixel.Line(a.points[i], a.points[i+1], 0, 14)
		}
		pixel.Line(a.points[len(a.points)-1], m, 0, 13)
	}
	if !scenes[1].Pressed() {
		for _, p := range a.points {
			pixel.Point(p, 0, 8)
		}
		pixel.Point(m, 0, 7)
	}
}
