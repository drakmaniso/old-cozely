// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane_test

import (
	"math/rand"
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
)

//------------------------------------------------------------------------------

var screen = pixel.NewCanvas(pixel.Zoom(4))

var cursor = pixel.NewCursor()

func init() {
	cursor.Canvas(screen)
}

var (
	seeds []plane.Coord
)

//------------------------------------------------------------------------------

func newSeeds() {
	for i := range seeds {
		seeds[i] = plane.Coord{X: rand.Float32(), Y: rand.Float32()}
	}
}

//------------------------------------------------------------------------------

func TestPlane_triangulation(t *testing.T) {
	do(func() {
		err := glam.Run(triLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type triLoop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (triLoop) Enter() error {
	seeds = make([]plane.Coord, 3)
	newSeeds()

	palette.Clear()
	palette.Index(1).SetColour(colour.LRGB{1, 1, 1})
	palette.Index(2).SetColour(colour.LRGB{1, 0, 0})
	palette.Index(3).SetColour(colour.LRGB{0, 1, 0})
	palette.Index(4).SetColour(colour.LRGB{0, 0, 1})
	return nil
}

//------------------------------------------------------------------------------

func (triLoop) Draw() error {
	screen.Clear(0)
	r := float32(screen.Size().Y)
	o := (float32(screen.Size().X) - r)
	for i, sd := range seeds {
		p := pixel.Coord{
			X: int16(o + sd.X*r),
			Y: int16(sd.Y * r),
		}
		screen.Point(2+palette.Index(i), p.X, p.Y, 1)
	}

	m := screen.Mouse()
	cursor.Locate(2, 8, 0x7FFF)
	cursor.Printf("Pos: %3d, %3d\n", m.X, m.Y)
	screen.Point(1, m.X, m.Y, 1)

	t := plane.Triangle{0, 1, 2}.CounterClockwise(seeds)
	if t[1] == 1 {
		cursor.Print("CounterCW")
	} else {
		cursor.Print("CW")
	}

	screen.Display()
	return nil
}

//------------------------------------------------------------------------------

func (triLoop) MouseButtonDown(b mouse.Button, _ int) {
	switch b {
	case mouse.Left:
		newSeeds()
	case mouse.Right:
	}
}

//------------------------------------------------------------------------------
