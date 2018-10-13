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

type loop5 struct {
	points []pixel.XY
}

////////////////////////////////////////////////////////////////////////////////

func TestTest5(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop5{}
		l.setup()

		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (l *loop5) setup() {
	color.Load(pico8.Palette)
	pixel.SetResolution(pixel.XY{128, 128})

	l.points = []pixel.XY{
		{4, 4},
		{4 + 1, 4 + 20},
		{4 + 1 + 20, 4 + 20 - 1},
		{16, 32},
	}
}

func (l *loop5) Enter() {
	input.ShowMouse(false)
}

func (l *loop5) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (l *loop5) React() {
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
	}

	if input.MenuClick.Pushed() {
		m := pixel.XYof(input.MenuPointer.XY())
		l.points = append(l.points, m)
	}

	if input.MenuLeft.Pushed() {
		if len(l.points) > 0 {
			l.points = l.points[:len(l.points)-1]
		}
	}
}

func (loop5) Update() {
}

func (l *loop5) Render() {
	pixel.Clear(1)
	m := pixel.XYof(input.MenuPointer.XY())
	if !input.MenuUp.Pressed() {
		for i := 0; i < len(l.points)-2; i++ {
			pixel.Triangle(l.points[i], l.points[i+1], l.points[i+2], 0, 2)
		}
	}
	if !input.MenuRight.Pressed() {
		for i := 0; i < len(l.points)-1; i++ {
			pixel.Line(l.points[i], l.points[i+1], 0, 14)
		}
		pixel.Line(l.points[len(l.points)-1], m, 0, 13)
	}
	if !input.MenuDown.Pressed() {
		for _, p := range l.points {
			pixel.Point(p, 0, 8)
		}
		pixel.Point(m, 0, 7)
	}
}
