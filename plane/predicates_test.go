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

var screen = pixel.NewCanvas(pixel.Zoom(2))

var cursor = pixel.NewCursor()

func init() {
	cursor.Canvas(screen)
}

var (
	points []plane.Coord
)

//------------------------------------------------------------------------------

func newPoints() {
	for i := range points {
		points[i] = plane.Coord{X: rand.Float32(), Y: rand.Float32()}
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
	points = make([]plane.Coord, 3)
	newPoints()

	palette.Clear()
	palette.Index(1).SetColour(colour.LRGB{1, 1, 1})
	palette.Index(2).SetColour(colour.LRGB{1, 0.2, 0.1})
	palette.Index(3).SetColour(colour.LRGB{0.1, 1, 0.2})
	palette.Index(4).SetColour(colour.LRGB{0.1, 0.2, 1})
	palette.Index(5).SetColour(colour.LRGB{0.2, 0.2, 0.2})
	return nil
}

//------------------------------------------------------------------------------

func (triLoop) Draw() error {
	screen.Clear(0)
	r := float32(screen.Size().Y)
	ox := (float32(screen.Size().X) - r)
	oy := float32(screen.Size().Y)
	pt := make([]pixel.Coord, len(points))
	for i, sd := range points {
		pt[i] = pixel.Coord{
			X: int16(ox + sd.X*r),
			Y: int16(oy - sd.Y*r),
		}
		// screen.Point(2+palette.Index(i), pt[i].X, pt[i].Y, 1)
		cursor.Locate(pt[i].X-2, pt[i].Y-3, +2)
		cursor.ColorShift(1 + palette.Index(i))
		cursor.Print([]string{"A", "B", "C"}[i])
	}
	screen.Lines(1, 0, pt[0], pt[1], pt[2], pt[0])

	m := screen.Mouse()
	p := plane.Coord{X: (float32(m.X) - ox) / r, Y: (oy - float32(m.Y)) / r}
	cursor.Locate(2, 8, 0x7FFF)
	cursor.ColorShift(0)
	cursor.Printf("A: %.3f, %.3f\n", points[0].X, points[0].Y)
	cursor.Printf("B: %.3f, %.3f\n", points[1].X, points[1].Y)
	cursor.Printf("C: %.3f, %.3f\n", points[2].X, points[2].Y)
	if p.X >= 0 {
		cursor.Printf("   %.3f, %.3f\n", p.X, p.Y)
	} else {
		cursor.Println(" ")
	}
	screen.Point(1, m.X, m.Y, 1)

	cursor.Println()

	if plane.IsCCW(points[0], points[1], points[2]) {
		cursor.ColorShift(3)
		cursor.Println("IsCCW: TRUE")
	} else {
		cursor.ColorShift(0)
		cursor.Println("IsCCW: false")
	}

	if plane.InTriangle(points[0], points[1], points[2], p) {
		cursor.ColorShift(1)
		cursor.Println("InTriangle: TRUE")
	} else {
		cursor.ColorShift(0)
		cursor.Println("InTriangle: false")
	}

	a, b, c := 0, 1, 2
	if !plane.IsCCW(points[a], points[b], points[c]) {
		b, c = c, b
	}
	if plane.InTriangleCCW(points[a], points[b], points[c], p) {
		cursor.ColorShift(1)
		cursor.Println("InTriangleCCW: TRUE")
	} else {
		cursor.ColorShift(0)
		cursor.Println("InTriangleCCW: false")
	}

	if plane.InCircle(points[a], points[b], points[c], p) {
		cursor.ColorShift(2)
		cursor.Println("InCircle: TRUE")
	} else {
		cursor.ColorShift(0)
		cursor.Println("InCircle: false")
	}

	screen.Display()
	return nil
}

//------------------------------------------------------------------------------

func (triLoop) MouseButtonDown(b mouse.Button, _ int) {
	switch b {
	case mouse.Left:
		newPoints()
	case mouse.Right:
	}
}

//------------------------------------------------------------------------------
