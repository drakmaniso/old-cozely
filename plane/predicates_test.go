// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane_test

import (
	"github.com/drakmaniso/cozely/input"
	"math/rand"
	"testing"

	"github.com/drakmaniso/cozely"
	"github.com/drakmaniso/cozely/colour"
	"github.com/drakmaniso/cozely/palette"
	"github.com/drakmaniso/cozely/pixel"
	"github.com/drakmaniso/cozely/plane"
	"github.com/drakmaniso/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

var (
	screen = pixel.Canvas(pixel.Zoom(2))
	cursor = pixel.Cursor{Canvas: screen}
)

var (
	points []plane.Coord
)

var (
	ratio  float32
	offset plane.Coord
)

////////////////////////////////////////////////////////////////////////////////

func newPoints() {
	for i := range points {
		points[i] = plane.Coord{X: rand.Float32(), Y: rand.Float32()}
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestPlane_predicates(t *testing.T) {
	do(func() {
		err := cozely.Run(triLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type triLoop struct{}

////////////////////////////////////////////////////////////////////////////////

func (triLoop) Enter() error {
	input.Load(testBindings)
	testContext.Activate(1)

	points = make([]plane.Coord, 3)
	newPoints()

	palette.Clear()
	palette.Index(1).SetColour(colour.LRGB{1, 1, 1})
	palette.Index(2).SetColour(colour.LRGB{0.4, 0.05, 0.0})
	palette.Index(3).SetColour(colour.LRGB{0.0, 0.4, 0.05})
	palette.Index(4).SetColour(colour.LRGB{0.0, 0.05, 0.45})
	palette.Index(5).SetColour(colour.LRGB{0.1, 0.0, 0.15})
	palette.Index(6).SetColour(colour.LRGB{0.25, 0.25, 0.25})
	palette.Index(7).SetColour(colour.LRGB{0.025, 0.025, 0.025})
	return nil
}

func (triLoop) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (triLoop) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}

	if next.JustPressed(1) {
		newPoints()
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (triLoop) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (triLoop) Render() error {
	screen.Clear(0)
	cursor.Depth = 0x7FFF
	ratio = float32(screen.Size().Y)
	offset = plane.Coord{
		X: (float32(screen.Size().X) - ratio),
		Y: float32(screen.Size().Y),
	}
	pt := make([]plane.Pixel, len(points))
	s := plane.Pixel{5, 5}
	for i, sd := range points {
		pt[i] = toScreen(sd)
		screen.Box(2+palette.Index(i), 2+palette.Index(i), 2, 2,
			pt[i].Minus(s), pt[i].Plus(s))
		cursor.Locate(pt[i].X-2, pt[i].Y+3)
		cursor.Color = 0
		cursor.Print([]string{"A", "B", "C"}[i])
	}
	screen.Lines(6, 0, pt[0], pt[1], pt[2], pt[0])
	screen.Triangles(7, -5, pt[0], pt[1], pt[2], pt[0])

	m := screen.Mouse()
	p := fromScreen(m)
	cursor.Locate(2, 8)
	cursor.Color = 0
	cursor.Printf("A: %.3f, %.3f\n", points[0].X, points[0].Y)
	cursor.Printf("B: %.3f, %.3f\n", points[1].X, points[1].Y)
	cursor.Printf("C: %.3f, %.3f\n", points[2].X, points[2].Y)
	if p.X >= 0 {
		cursor.Printf("   %.3f, %.3f\n", p.X, p.Y)
	} else {
		cursor.Println(" ")
	}
	screen.Point(1, 1, m)

	cursor.Println()

	if plane.IsCCW(points[0], points[1], points[2]) {
		cursor.Color = 3
		cursor.Println("IsCCW: TRUE")
	} else {
		cursor.Color = 0
		cursor.Println("IsCCW: false")
	}

	if plane.InTriangle(points[0], points[1], points[2], p) {
		cursor.Color = 1
		cursor.Println("InTriangle: TRUE")
	} else {
		cursor.Color = 0
		cursor.Println("InTriangle: false")
	}

	a, b, c := 0, 1, 2
	if !plane.IsCCW(points[a], points[b], points[c]) {
		b, c = c, b
	}
	if plane.InTriangleCCW(points[a], points[b], points[c], p) {
		cursor.Color = 1
		cursor.Println("InTriangleCCW: TRUE")
	} else {
		cursor.Color = 0
		cursor.Println("InTriangleCCW: false")
	}

	if plane.InCircumcircle(points[a], points[b], points[c], p) {
		cursor.Color = 2
		cursor.Println("InCircumcircle: TRUE")
	} else {
		cursor.Color = 0
		cursor.Println("InCircumcircle: false")
	}

	cursor.Println(" ")

	d := plane.Circumcenter(points[0], points[1], points[2])
	cursor.Color = 0
	cursor.Printf("Circumcenter: %.3f, %.3f\n", d.X, d.Y)
	dd := toScreen(d)
	screen.Lines(5, -2, dd.Pluss(-2, -2), dd.Pluss(2, 2))
	screen.Lines(5, -2, dd.Pluss(2, -2), dd.Pluss(-2, 2))

	r := d.Minus(points[a]).Length()
	cir := []plane.Pixel{}
	for a := float32(0); a <= 2*math32.Pi+0.01; a += math32.Pi / 32 {
		cir = append(cir, toScreen(plane.Polar{r, a}.Cartesian().Plus(d)))
	}
	screen.Lines(5, -2, cir...)

	screen.Display()
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func toScreen(p plane.Coord) plane.Pixel {
	return plane.Pixel{
		X: int16(offset.X + p.X*ratio),
		Y: int16(offset.Y - p.Y*ratio),
	}
}

func fromScreen(p plane.Pixel) plane.Coord {
	return plane.Coord{
		X: (float32(p.X) - offset.X) / ratio,
		Y: (offset.Y - float32(p.Y)) / ratio,
	}
}

////////////////////////////////////////////////////////////////////////////////
