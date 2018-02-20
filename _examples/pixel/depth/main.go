// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"math/rand"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var screen = pixel.NewCanvas(pixel.Zoom(2))

var shapes = []pixel.Picture{
	pixel.NewPicture("graphics/shape1"),
	pixel.NewPicture("graphics/shape2"),
	pixel.NewPicture("graphics/shape3"),
	pixel.NewPicture("graphics/shape4"),
}

type object struct {
	pict  pixel.Picture
	pos   pixel.Coord
	depth int16
}

var objects [1024]object

//------------------------------------------------------------------------------

func main() {

	glam.Configure(
		glam.TimeStep(1 / 60.0),
	)
	err := glam.Run(loop{})
	if err != nil {
		glam.ShowError(err)
	}
}

type loop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (loop) WindowResized(_, _ int32) {
	s := screen.Size()
	for i := range objects {
		j := rand.Intn(len(shapes))
		objects[i].depth = int16(j)
		p := shapes[j]
		objects[i].pict = p
		objects[i].pos.X = int16(rand.Intn(int(s.X - p.Size().X)))
		objects[i].pos.Y = int16(rand.Intn(int(s.Y - p.Size().Y)))
	}
}

//------------------------------------------------------------------------------

func (loop) Enter() error {
	palette.Load("graphics/shape1")
	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw() error {
	screen.Clear(0)
	for i, o := range objects {
		if float64(i)/32 > glam.GameTime() {
			break
		}
		screen.Picture(o.pict, o.pos.X, o.pos.Y, o.depth)
	}
	screen.Display()
	return nil
}

//------------------------------------------------------------------------------
