// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/plane"
)

////////////////////////////////////////////////////////////////////////////////

var cnvContext = input.Context("TestCanvas", quit)

var cnvBindings = input.Bindings{
	"TestCanvas": {
		"Quit": {"Escape"},
	},
}

////////////////////////////////////////////////////////////////////////////////

var cnvScreen = pixel.Canvas(pixel.Zoom(2))

var shapePictures = []pixel.PictureID{
	pixel.Picture("graphics/shape1"),
	pixel.Picture("graphics/shape2"),
	pixel.Picture("graphics/shape3"),
	pixel.Picture("graphics/shape4"),
}

type shape struct {
	pict  pixel.PictureID
	pos   plane.CR
	depth int16
}

var shapes [2048]shape

////////////////////////////////////////////////////////////////////////////////

func TestCanvas_depth(t *testing.T) {
	do(func() {
		cozely.Configure(
			cozely.UpdateStep(1 / 60.0),
		)
		cozely.Events.Resize = resize
		err := cozely.Run(cnvLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type cnvLoop struct{}

////////////////////////////////////////////////////////////////////////////////

func (cnvLoop) Enter() error {
	input.Load(testBindings)
	testContext.Activate(1)
	palette.Load("graphics/shape1")
	return nil
}

func (cnvLoop) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (cnvLoop) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (cnvLoop) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (cnvLoop) Render() error {
	cnvScreen.Clear(0)
	for i, o := range shapes {
		if float64(i)/32 > cozely.GameTime() {
			break
		}
		cnvScreen.Picture(o.pict, o.depth, o.pos)
	}
	cnvScreen.Display()
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func resize() {
	s := cnvScreen.Size()
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
