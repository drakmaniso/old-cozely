// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"math/rand"
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var cnvScreen = pixel.NewCanvas(pixel.Zoom(2))

var shapePictures = []pixel.Picture{
	pixel.NewPicture("graphics/shape1"),
	pixel.NewPicture("graphics/shape2"),
	pixel.NewPicture("graphics/shape3"),
	pixel.NewPicture("graphics/shape4"),
}

type shape struct {
	pict  pixel.Picture
	pos   pixel.Coord
	depth int16
}

var shapes [2048]shape

//------------------------------------------------------------------------------

func TestCanvas_depth(t *testing.T) {
	do(func() {
		glam.Configure(
			glam.TimeStep(1 / 60.0),
		)
		err := glam.Run(cnvLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type cnvLoop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (cnvLoop) WindowResized(_, _ int32) {
	s := cnvScreen.Size()
	for i := range shapes {
		j := rand.Intn(len(shapePictures))
		shapes[i].depth = int16(j)
		p := shapePictures[j]
		shapes[i].pict = p
		shapes[i].pos.X = int16(rand.Intn(int(s.X - p.Size().X)))
		shapes[i].pos.Y = int16(rand.Intn(int(s.Y - p.Size().Y)))
	}
}

//------------------------------------------------------------------------------

func (cnvLoop) Enter() error {
	palette.Load("graphics/shape1")
	return nil
}

//------------------------------------------------------------------------------

func (cnvLoop) Draw() error {
	cnvScreen.Clear(0)
	for i, o := range shapes {
		if float64(i)/32 > glam.GameTime() {
			break
		}
		cnvScreen.Picture(o.pict, o.pos.X, o.pos.Y, o.depth)
	}
	cnvScreen.Display()
	return nil
}

//------------------------------------------------------------------------------
