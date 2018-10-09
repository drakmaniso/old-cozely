// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
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
		l.declare()

		window.Events.Resize = l.resize
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (l *loop2) declare() {
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
	s := pixel.Resolution()
	for i := range l.shapes {
		j := rand.Intn(len(l.picts))
		p := l.picts[j]
		l.shapes[i].pict = p
		l.shapes[i].pos.X = int16(rand.Intn(int(s.X - p.Size().X)))
		l.shapes[i].pos.Y = int16(rand.Intn(int(s.Y - p.Size().Y)))
	}
}

////////////////////////////////////////////////////////////////////////////////

func (l *loop2) React() {
	if scenes[1].Pushed() {
		l.shapes = make([]shape, 1000)
		l.resize()
	}
	if scenes[2].Pushed() {
		l.shapes = make([]shape, 10000)
		l.resize()
	}
	if scenes[3].Pushed() {
		l.shapes = make([]shape, 100000)
		l.resize()
	}
	if scenes[4].Pushed() {
		l.shapes = make([]shape, 200000)
		l.resize()
	}
	if scenes[5].Pushed() {
		l.shapes = make([]shape, 300000)
		l.resize()
	}
	if scenes[6].Pushed() {
		l.shapes = make([]shape, 350000)
		l.resize()
	}
	if scenes[7].Pushed() {
		l.shapes = make([]shape, 400000)
		l.resize()
	}
	if scenes[8].Pushed() {
		l.shapes = make([]shape, 450000)
		l.resize()
	}
	if scenes[9].Pushed() {
		l.shapes = make([]shape, 500000)
		l.resize()
	}
	if scenes[0].Pushed() {
		l.shapes = make([]shape, 10)
		l.resize()
	}
	if scrollup.Pushed() {
		l.shapes = make([]shape, len(l.shapes)+1000)
		l.resize()
	}
	if scrolldown.Pushed() && len(l.shapes) > 1000 {
		l.shapes = make([]shape, len(l.shapes)-1000)
		l.resize()
	}
	if next.Pushed() {
		l.shapes = append(l.shapes, shape{})
		i := len(l.shapes) - 1
		j := rand.Intn(len(l.picts))
		p := l.picts[j]
		l.shapes[i].pict = p
		//TODO:
		l.shapes[i].pos = pixel.XYof(cursor.XY()).Minus(p.Size().Slash(2))
	}
	if previous.Pushed() && len(l.shapes) > 0 {
		l.shapes = l.shapes[:len(l.shapes)-1]
	}
	if quit.Pushed() {
		cozely.Stop(nil)
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
	}
	cur := pixel.Cursor{
		Position: pixel.XY{8, 16},
		Layer: 0xFFFF/2,
		Color: l.txtcol,
	}
	ft, ov := cozely.RenderStats()
	cur.Printf("%dk pictures: %6.2f", len(l.shapes)/1000, ft*1000)
	if ov > 0 {
		cur.Printf(" (%d)", ov)
	}
}
