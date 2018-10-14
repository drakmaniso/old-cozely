package pixel_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/pico8"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/window"
)

////////////////////////////////////////////////////////////////////////////////

type loop2 struct {
	txtcol color.Index
	picts  []pixel.PictureID
	shapes []shape
}

type shape struct {
	pict pixel.PictureID
	pos  pixel.XY
}

////////////////////////////////////////////////////////////////////////////////

func TestTest2(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop2{}
		l.setup()

		window.Events.Resize = l.resize
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (l *loop2) setup() {
	color.Load(&pico8.Palette)
	pixel.SetZoom(2)

	l.txtcol = 7

	l.picts = []pixel.PictureID{
		pixel.Picture("graphics/shape1"),
		pixel.Picture("graphics/shape2"),
		pixel.Picture("graphics/shape3"),
		pixel.Picture("graphics/shape4"),
	}
	l.shapes = make([]shape, 400000)
}

func (l *loop2) Enter() {
}

func (loop2) Leave() {
}

func (l *loop2) resize() {
	s := len(l.shapes)
	l.shapes = l.shapes[:0]
	for i := 0; i < s; i++ {
		l.addShape()
	}
}

func (l *loop2) addShape() {
	r := pixel.Resolution()
	j := rand.Intn(len(l.picts))
	p := l.picts[j]
	s := shape{
		pict: p,
		pos: pixel.XY{
			int16(rand.Intn(int(r.X - p.Size().X))),
			int16(rand.Intn(int(r.Y - p.Size().Y))),
		},
	}
	l.shapes = append(l.shapes, s)
}

////////////////////////////////////////////////////////////////////////////////

func (l *loop2) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}

	s := len(l.shapes)
	if input.Up.Pressed() {
		for i := 0; i < 50000; i++ {
			l.addShape()
		}
	}
	if input.Down.Pressed() {
		s -= 50000
		if s < 0 {
			s = 0
		}
		l.shapes = l.shapes[:s]
	}
	if input.Right.Pressed() {
		for i := 0; i < 1000; i++ {
			l.addShape()
		}
	}
	if input.Left.Pressed() {
		s -= 1000
		if s < 0 {
			s = 0
		}
		l.shapes = l.shapes[:s]
	}

	if input.Select.Pressed() {
		l.resize()
	}

	if input.Click.Pressed() {
		l.shapes = append(l.shapes, shape{})
		i := len(l.shapes) - 1
		j := rand.Intn(len(l.picts))
		p := l.picts[j]
		l.shapes[i].pict = p
		//TODO:
		l.shapes[i].pos = pixel.XYof(input.Pointer.XY()).Minus(p.Size().Slash(2))
	}
}

func (loop2) Update() {
}

func (l *loop2) Render() {
	pixel.Clear(0)
	for i, o := range l.shapes {
		l := i - (0xFFFF / 2)
		if l > 0xFFFF/2 {
			l = 0xFFFF / 2
		}
		o.pict.Paint(o.pos, pixel.Layer(l))
		// o.pict.Tile(o.pos, o.pict.Size(), pixel.Layer(l))
	}
	cur := pixel.Cursor{
		Position: pixel.XY{8, 16},
		Layer:    0xFFFF / 2,
		Color:    l.txtcol,
	}
	ft, ov := cozely.RenderStats()
	cur.Printf("%dk pictures: %6.2f", len(l.shapes)/1000, ft*1000)
	if ov > 0 {
		cur.Printf(" (%d)", ov)
	}
}

// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
