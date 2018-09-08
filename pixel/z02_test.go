// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
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
		l.declare()

		cozely.Configure(
			cozely.UpdateStep(1 / 60.0),
		)
		window.Events.Resize = l.resize
		input.Load(bindings)
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (a *loop2) declare() {
	pixel.SetZoom(2)

	a.txtcol = 8

	a.picts = []pixel.PictureID{
		pixel.Picture("graphics/shape1"),
		pixel.Picture("graphics/shape2"),
		pixel.Picture("graphics/shape3"),
		pixel.Picture("graphics/shape4"),
	}
	a.shapes = make([]shape, 200000)
}

func (a *loop2) Enter() {
	pixel.Text(a.txtcol, pixel.Monozela10)
}

func (loop2) Leave() {
}

func (a *loop2) resize() {
	s := pixel.Resolution()
	for i := range a.shapes {
		j := rand.Intn(len(a.picts))
		p := a.picts[j]
		a.shapes[i].pict = p
		a.shapes[i].pos.X = int16(rand.Intn(int(s.X - p.Size().X)))
		a.shapes[i].pos.Y = int16(rand.Intn(int(s.Y - p.Size().Y)))
	}
}

////////////////////////////////////////////////////////////////////////////////

func (a *loop2) React() {
	if scenes[1].Started(0) {
		a.shapes = make([]shape, 1000)
		a.resize()
	}
	if scenes[2].Started(0) {
		a.shapes = make([]shape, 10000)
		a.resize()
	}
	if scenes[3].Started(0) {
		a.shapes = make([]shape, 100000)
		a.resize()
	}
	if scenes[4].Started(0) {
		a.shapes = make([]shape, 200000)
		a.resize()
	}
	if scenes[5].Started(0) {
		a.shapes = make([]shape, 300000)
		a.resize()
	}
	if scenes[6].Started(0) {
		a.shapes = make([]shape, 350000)
		a.resize()
	}
	if scenes[7].Started(0) {
		a.shapes = make([]shape, 400000)
		a.resize()
	}
	if scenes[8].Started(0) {
		a.shapes = make([]shape, 450000)
		a.resize()
	}
	if scenes[9].Started(0) {
		a.shapes = make([]shape, 500000)
		a.resize()
	}
	if scenes[0].Started(0) {
		a.shapes = make([]shape, 10)
		a.resize()
	}
	if scrollup.Started(0) {
		a.shapes = make([]shape, len(a.shapes)+1000)
		a.resize()
	}
	if scrolldown.Started(0) && len(a.shapes) > 1000 {
		a.shapes = make([]shape, len(a.shapes)-1000)
		a.resize()
	}
	if next.Started(0) {
		a.shapes = append(a.shapes, shape{})
		i := len(a.shapes) - 1
		j := rand.Intn(len(a.picts))
		p := a.picts[j]
		a.shapes[i].pict = p
		//TODO:
		a.shapes[i].pos = pixel.ToCanvas(window.XYof(cursor.XY(0))).Minus(p.Size().Slash(2))
	}
	if previous.Started(0) && len(a.shapes) > 0 {
		a.shapes = a.shapes[:len(a.shapes)-1]
	}
	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop2) Update() {

}

func (a *loop2) Render() {
	pixel.Clear(0)
	for _, o := range a.shapes {
		pixel.Paint(o.pict, o.pos)
	}
	pixel.Locate(pixel.XY{8, 16})
	ft, ov := cozely.RenderStats()
	pixel.Printf("%dk pictures: %6.2f", len(a.shapes)/1000, ft*1000)
	if ov > 0 {
		pixel.Printf(" (%d)", ov)
	}
}
