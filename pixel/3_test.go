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
	pict1, pict2   pixel.BoxID
	position, size pixel.XY
}

////////////////////////////////////////////////////////////////////////////////

func TestTest3(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop3{}
		err := l.setup()
		if err != nil {
			t.Error(err)
			return
		}

		err = cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (l *loop3) setup() error {
	color.Load(&pico8.Palette)
	pixel.SetZoom(4)
	err := resource.Path("testdata/")
	if err != nil {
		return err
	}
	l.position = pixel.XY{48, 48}
	l.size = pixel.XY{60, 40}
	return nil
}

func (l *loop3) Enter() {
	l.pict1 = pixel.Box("graphics/box")
	l.pict2 = pixel.Box("graphics/button")
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
	l.pict2.Paint(pixel.XY{48, 4}, pixel.XY{8, 8}, -1, 0)
	l.pict2.Paint(pixel.XY{64, 4}, pixel.XY{32, 16}, -1, 0)
	l.pict2.Paint(pixel.XY{102, 4}, pixel.XY{16, 32}, -1, 0)
	l.pict2.Paint(pixel.XY{128, 4}, pixel.XY{32, 32}, -1, 0)
	l.pict1.Paint(l.position, l.size, 0, 0)
}

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
