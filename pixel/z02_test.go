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

var (
	palette2 = color.PaletteFrom("graphics/shape1")
	txtColor = palette2.Entry(color.LRGB{1, 1, 1})
)

////////////////////////////////////////////////////////////////////////////////

var canvas2 = pixel.Canvas(pixel.Zoom(2))

var shapePictures = []pixel.PictureID{
	pixel.Picture("graphics/shape1"),
	pixel.Picture("graphics/shape2"),
	pixel.Picture("graphics/shape3"),
	pixel.Picture("graphics/shape4"),
}

type shape struct {
	pict  pixel.PictureID
	pos   coord.CR
	depth int16
}

type loop2 struct{}

var shapes = make([]shape, 350000)

////////////////////////////////////////////////////////////////////////////////

func TestTest2(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		cozely.Configure(
			cozely.UpdateStep(1 / 60.0),
		)
		cozely.Events.Resize = resize
		input.Load(bindings)
		err := cozely.Run(loop2{})
		if err != nil {
			t.Error(err)
		}
	})
}

func (loop2) Enter() {
	palette2.Activate()
	canvas2.Text(txtColor, pixel.Monozela10)
}

func (loop2) Leave() {
}

func resize() {
	s := canvas2.Size()
	for i := range shapes {
		j := rand.Intn(len(shapePictures))
		if i < 0x7FFF {
		shapes[i].depth = int16(i)
		} else {
			shapes[i].depth = 0x7FFF
		}
		p := shapePictures[j]
		shapes[i].pict = p
		shapes[i].pos.C = int16(rand.Intn(int(s.C - p.Size().C)))
		shapes[i].pos.R = int16(rand.Intn(int(s.R - p.Size().R)))
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop2) React() {
	if scene1.Started(0) {
		shapes = make([]shape, 1000)
		resize()
	}
	if scene2.Started(0) {
		shapes = make([]shape, 10000)
		resize()
	}
	if scene3.Started(0) {
		shapes = make([]shape, 100000)
		resize()
	}
	if scene4.Started(0) {
		shapes = make([]shape, 200000)
		resize()
	}
	if scene5.Started(0) {
		shapes = make([]shape, 300000)
		resize()
	}
	if scene6.Started(0) {
		shapes = make([]shape, 350000)
		resize()
	}
	if scene7.Started(0) {
		shapes = make([]shape, 400000)
		resize()
	}
	if scene8.Started(0) {
		shapes = make([]shape, 450000)
		resize()
	}
	if scene9.Started(0) {
		shapes = make([]shape, 500000)
		resize()
	}
	if scene10.Started(0) {
		shapes = make([]shape, 10)
		resize()
	}
	if next.Started(0) {
		shapes = append(shapes, shape{})
		i := len(shapes) - 1
		j := rand.Intn(len(shapePictures))
		if i < 0x7FFF {
			shapes[i].depth = int16(i)
			} else {
				shapes[i].depth = 0x7FFF
			}
			p := shapePictures[j]
		shapes[i].pict = p
		shapes[i].pos = canvas2.FromWindow(cursor.XY(0).CR()).Minus(p.Size().Slash(2))
	}
	if previous.Started(0) {
		shapes = shapes[:len(shapes)-1]
	}
	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop2) Update() {

}

func (loop2) Render() {
	canvas2.Clear(0)
	for _, o := range shapes {
		canvas2.Picture(o.pict, o.depth, o.pos)
	}
	canvas2.Locate(0x7FFF, coord.CR{8, 16})
	ft, ov := cozely.RenderStats()
	canvas2.Printf("%dk pictures: %6.2f", len(shapes)/1000, ft*1000)
	if ov > 0 {
		canvas2.Printf(" (%d)", ov)
	}
	canvas2.Display()
}
