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
	mire                                   pixel.PictureID
	srgbGray, srgbRed, srgbGreen, srgbBlue pixel.PictureID
	rectangle, fill                        pixel.BoxID
	cursor                                 pixel.PictureID
	mode                                   int
}

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		pixel.SetResolution(pixel.XY{320, 180})
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
	l.srgbGray = pixel.Picture("graphics/srgb-gray")
	l.srgbRed = pixel.Picture("graphics/srgb-red")
	l.srgbGreen = pixel.Picture("graphics/srgb-green")
	l.srgbBlue = pixel.Picture("graphics/srgb-blue")
	l.mire = pixel.Picture("graphics/mire")
	l.cursor = pixel.Picture("builtins/cursor")
	l.rectangle = pixel.Box("builtins/rectangle")
	l.fill = pixel.Box("builtins/fill")
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

	s := l.mire.Size()
	l.mire.Paint(pixel.XY{0, 0}, 0, 0)
	l.mire.Paint(pixel.XY{0, r.Y - s.Y}, 0, 0)
	l.mire.Paint(pixel.XY{r.X - s.X, 0}, 0, 0)
	l.mire.Paint(r.Minus(s), 0, 0)

	l.fill.Paint(s, r.Minus(s.Times(2)), -1, color.Black)
	l.rectangle.Paint(s, r.Minus(s.Times(2)), -1, color.DarkGray)

	for i := int16(0); i < 6; i++ {
		l.rectangle.Paint(
			pixel.XY{130 + i*10, r.Y - 28},
			pixel.XY{3 + i, 3 + i},
			0,
			0,
		)
		l.fill.Paint(
			pixel.XY{130 + i*10, r.Y - 14},
			pixel.XY{3 + i, 3 + i},
			0,
			0,
		)
	}

	cur := pixel.Cursor{
		Position: pixel.XY{s.X + 28, 108},
	}
	cur.Margin = cur.Position.X
	cur.Color = color.MidGray
	cur.Println("  sRGB:")
	cur.Println("Linear:")

	s = l.srgbGray.Size()
	l.srgbGray.Paint(pixel.XY{s.X + 44 + 1*(s.X+4), 96}, 0, 0)
	l.srgbRed.Paint(pixel.XY{s.X + 44 + 2*(s.X+4), 96}, 0, 0)
	l.srgbGreen.Paint(pixel.XY{s.X + 44 + 3*(s.X+4), 96}, 0, 0)
	l.srgbBlue.Paint(pixel.XY{s.X + 44 + 4*(s.X+4), 96}, 0, 0)

	l.cursor.Paint(pixel.XYof(input.Pointer.XY()), 0, 0)
}

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
