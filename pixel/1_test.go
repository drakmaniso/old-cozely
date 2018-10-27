package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/resource"
)

////////////////////////////////////////////////////////////////////////////////

var ()

type loop1 struct {
	mode int
}

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		// pixel.SetResolution(pixel.XY{320, 180})
		pixel.SetZoom(4)
		err := resource.Path("testdata/")
		if err != nil {
			t.Error(err)
			return
		}

		err = cozely.Run(&loop1{})
		if err != nil {
			t.Error(err)
		}
	})
}

func (l *loop1) Enter() {
	l.mode = 0
	input.ShowMouse(false)
}

func (l *loop1) Leave() {
	input.ShowMouse(false)
}

////////////////////////////////////////////////////////////////////////////////

func (l *loop1) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}
}

func (l *loop1) Update() {
}

func (l *loop1) Render() {
	pixel.Clear(0)
	r := pixel.Resolution()

	p := pixel.Picture("mire")
	s := p.Size()
	p.Paint(pixel.XY{0, 0}, 2, 0)
	p.Paint(pixel.XY{0, r.Y - s.Y}, 2, 0)
	p.Paint(pixel.XY{r.X - s.X, 0}, 2, 0)
	p.Paint(r.Minus(s), 2, 0)

	cur := pixel.Cursor{
		Position: pixel.XY{s.X + 8, 12},
	}
	cur.Margin = cur.Position.X
	cur.Color = color.MidGray
	cur.Println("  sRGB:")
	cur.Println("Linear:")
	pixel.Picture("srgb-gray").Paint(pixel.XY{s.X + 23 + 1*(s.X+4), 0}, 0, 0)
	pixel.Picture("srgb-red").Paint(pixel.XY{s.X + 23 + 2*(s.X+4), 0}, 0, 0)
	pixel.Picture("srgb-green").Paint(pixel.XY{s.X + 23 + 3*(s.X+4), 0}, 0, 0)
	pixel.Picture("srgb-blue").Paint(pixel.XY{s.X + 23 + 4*(s.X+4), 0}, 0, 0)

	p = pixel.Picture("origtest")
	p.Paint(pixel.XY{r.X - s.X - 32, s.Y / 2}, 0, 0)
	pixel.Point(pixel.XY{r.X - s.X - 32, s.Y / 2}, 0, 255)

	fill := pixel.Box("builtins/fill")
	rect := pixel.Box("builtins/rectangle")
	fill.Paint(s, r.Minus(s.Timess(2)), -1, color.Black)
	rect.Paint(s, r.Minus(s.Timess(2)), -1, color.DarkGray)

	for i := int16(0); i < 6; i++ {
		rect.Paint(
			pixel.XY{130 + i*10, r.Y - 28},
			pixel.XY{3 + i, 3 + i},
			0,
			0,
		)
		fill.Paint(
			pixel.XY{130 + i*10, r.Y - 14},
			pixel.XY{3 + i, 3 + i},
			0,
			0,
		)
	}

	pixel.Picture("builtins/cursor").Paint(pixel.XYof(input.Pointer.XY()), 0, 0)
}

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
