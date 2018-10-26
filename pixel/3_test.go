package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/pico8"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/resource"
)

////////////////////////////////////////////////////////////////////////////////

var ()

type loop3 struct {
	boxtest, button pixel.BoxID
	position, size  pixel.XY
}

////////////////////////////////////////////////////////////////////////////////

func TestTest3(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		color.Load(&pico8.Palette)
		pixel.SetZoom(4)
		err := resource.Path("testdata/")
		if err != nil {
			t.Error(err)
			return
		}

		err = cozely.Run(&loop3{})
		if err != nil {
			t.Error(err)
		}
	})
}

func (l *loop3) Enter() {
	l.boxtest = pixel.Box("box")
	l.button = pixel.Box("button")
	l.position = pixel.XY{48, 48}
	l.size = pixel.XY{60, 40}
}

func (loop3) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (l *loop3) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}

	if input.Click.Ongoing() {
		l.size = pixel.XYof(input.Pointer.XY()).Minus(l.position)
	}
}

func (loop3) Update() {
}

func (l *loop3) Render() {
	pixel.Clear(pico8.White)
	l.button.Paint(pixel.XY{40, 10}, pixel.XY{8, 8}, -1, 0)
	l.button.Paint(pixel.XY{64, 10}, pixel.XY{16, 8}, -1, 0)
	l.button.Paint(pixel.XY{102, 10}, pixel.XY{8, 16}, -1, 0)
	l.button.Paint(pixel.XY{128, 10}, pixel.XY{16, 16}, -1, 0)
	l.boxtest.Paint(l.position, l.size, 0, 0)
	pixel.Point(l.position, 0, pico8.White)
}

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
