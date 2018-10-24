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
	mode                                   int
}

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop1{}
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

func (l *loop1) setup() error {
	pixel.SetResolution(pixel.XY{320, 180})
	err := resource.Path("testdata/")
	if err != nil {
		return err
	}
	l.srgbGray = pixel.Picture("graphics/srgb-gray")
	l.srgbRed = pixel.Picture("graphics/srgb-red")
	l.srgbGreen = pixel.Picture("graphics/srgb-green")
	l.srgbBlue = pixel.Picture("graphics/srgb-blue")
	l.mire = pixel.Picture("graphics/mire")
	return nil
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

	s := l.mire.Size()
	l.mire.Paint(pixel.XY{0, 0}, 0)
	l.mire.Paint(pixel.XY{0, r.Y - s.Y}, 0)
	l.mire.Paint(pixel.XY{r.X - s.X, 0}, 0)
	l.mire.Paint(r.Minus(s), 0)

	pixel.FilledRectangle.TileMod(s, r.Minus(s.Times(2)), -1, color.Black)
	pixel.Rectangle.TileMod(s, r.Minus(s.Times(2)), -1, color.DarkGray)

	for i := int16(0); i < 6; i++ {
		pixel.Rectangle.Tile(
			pixel.XY{40 + i*10, r.Y - 28},
			pixel.XY{3 + i, 3 + i},
			0,
		)
		pixel.FilledRectangle.Tile(
			pixel.XY{40 + i*10, r.Y - 14},
			pixel.XY{3 + i, 3 + i},
			0,
		)
		pixel.RectangleR1.Tile(
			pixel.XY{130 + i*10, r.Y - 28},
			pixel.XY{3 + i, 3 + i},
			0,
		)
		pixel.FilledRectangleR1.Tile(
			pixel.XY{130 + i*10, r.Y - 14},
			pixel.XY{3 + i, 3 + i},
			0,
		)
		pixel.RectangleR2.Tile(
			pixel.XY{220 + i*10, r.Y - 28},
			pixel.XY{5 + i, 5 + i},
			0,
		)
		pixel.FilledRectangleR2.Tile(
			pixel.XY{220 + i*10, r.Y - 14},
			pixel.XY{5 + i, 5 + i},
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
	l.srgbGray.Paint(pixel.XY{s.X + 44 + 1*(s.X+4), 96}, 0)
	l.srgbRed.Paint(pixel.XY{s.X + 44 + 2*(s.X+4), 96}, 0)
	l.srgbGreen.Paint(pixel.XY{s.X + 44 + 3*(s.X+4), 96}, 0)
	l.srgbBlue.Paint(pixel.XY{s.X + 44 + 4*(s.X+4), 96}, 0)

	pixel.MouseCursor.Paint(pixel.XYof(input.Pointer.XY()), 0)
}

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
