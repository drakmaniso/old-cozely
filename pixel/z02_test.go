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
	context2 = input.Context("TestCanvas", quit)
	palette2 = color.PaletteFrom("graphics/shape1")
)

var bindings2 = input.Bindings{
	"TestCanvas": {
		"Quit": {"Escape"},
	},
}

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

var shapes [2048]shape

////////////////////////////////////////////////////////////////////////////////

func TestTest2(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		cozely.Configure(
			cozely.UpdateStep(1 / 60.0),
		)
		cozely.Events.Resize = resize
		err := cozely.Run(loop2{})
		if err != nil {
			t.Error(err)
		}
	})
}

func (loop2) Enter() {
	input.Load(bindings2)
	context2.Activate(0)
	palette2.Activate()
}

func (loop2) Leave() {
}

func resize() {
	s := canvas2.Size()
	for i := range shapes {
		j := rand.Intn(len(shapePictures))
		shapes[i].depth = int16(j)
		p := shapePictures[j]
		shapes[i].pict = p
		shapes[i].pos.C = int16(rand.Intn(int(s.C - p.Size().C)))
		shapes[i].pos.R = int16(rand.Intn(int(s.R - p.Size().R)))
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop2) React() {
	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop2) Update() {

}

func (loop2) Render() {
	canvas2.Clear(0)
	for i, o := range shapes {
		if float64(i)/32 > cozely.GameTime() {
			break
		}
		canvas2.Picture(o.pict, o.depth, o.pos)
	}
	canvas2.Display()
}
