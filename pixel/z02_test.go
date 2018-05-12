// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

type loop2 struct {
	canvas  pixel.CanvasID
	palette color.PaletteID
	txtcol  color.Index
	picts   []pixel.PictureID
	shapes  []shape
}

type shape struct {
	pict pixel.PictureID
	pos  coord.CR
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
		cozely.Events.Resize = l.resize
		input.Load(bindings)
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (a *loop2) declare() {
	a.canvas = pixel.Canvas(pixel.Zoom(2))

	a.palette = color.PaletteFrom("graphics/shape1")
	a.txtcol = a.palette.Entry(color.LRGB{1, 1, 1})

	a.picts = []pixel.PictureID{
		pixel.Picture("graphics/shape1"),
		pixel.Picture("graphics/shape2"),
		pixel.Picture("graphics/shape3"),
		pixel.Picture("graphics/shape4"),
	}
	a.shapes = make([]shape, 350000)
}

func (a *loop2) Enter() {
	a.palette.Activate()
	a.canvas.Text(a.txtcol, pixel.Monozela10)
}

func (loop2) Leave() {
}

func (a *loop2) resize() {
	s := a.canvas.Size()
	for i := range a.shapes {
		j := rand.Intn(len(a.picts))
		p := a.picts[j]
		a.shapes[i].pict = p
		a.shapes[i].pos.C = int16(rand.Intn(int(s.C - p.Size().C)))
		a.shapes[i].pos.R = int16(rand.Intn(int(s.R - p.Size().R)))
	}
}

////////////////////////////////////////////////////////////////////////////////

func (a *loop2) React() {
	if scene1.Started(0) {
		a.shapes = make([]shape, 1000)
		a.resize()
	}
	if scene2.Started(0) {
		a.shapes = make([]shape, 10000)
		a.resize()
	}
	if scene3.Started(0) {
		a.shapes = make([]shape, 100000)
		a.resize()
	}
	if scene4.Started(0) {
		a.shapes = make([]shape, 200000)
		a.resize()
	}
	if scene5.Started(0) {
		a.shapes = make([]shape, 300000)
		a.resize()
	}
	if scene6.Started(0) {
		a.shapes = make([]shape, 350000)
		a.resize()
	}
	if scene7.Started(0) {
		a.shapes = make([]shape, 400000)
		a.resize()
	}
	if scene8.Started(0) {
		a.shapes = make([]shape, 450000)
		a.resize()
	}
	if scene9.Started(0) {
		a.shapes = make([]shape, 500000)
		a.resize()
	}
	if scene10.Started(0) {
		a.shapes = make([]shape, 10)
		a.resize()
	}
	if next.Started(0) {
		a.shapes = append(a.shapes, shape{})
		i := len(a.shapes) - 1
		j := rand.Intn(len(a.picts))
		p := a.picts[j]
		a.shapes[i].pict = p
		a.shapes[i].pos = a.canvas.FromWindow(cursor.XY(0).CR()).Minus(p.Size().Slash(2))
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
	a.canvas.Clear(0)
	for _, o := range a.shapes {
		a.canvas.Picture(o.pict, o.pos)
	}
	a.canvas.Locate(coord.CR{8, 16})
	ft, ov := cozely.RenderStats()
	a.canvas.Printf("%dk pictures: %6.2f", len(a.shapes)/1000, ft*1000)
	if ov > 0 {
		a.canvas.Printf(" (%d)", ov)
	}
	a.canvas.Display()
}
