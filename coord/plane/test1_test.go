// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/coord/plane"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

var (
	canvas1  = pixel.Canvas(pixel.Zoom(2))
	palette1 = color.Palette()
	col1     = palette1.Entry(color.LRGB{1, 1, 1})
	col2     = palette1.Entry(color.LRGB{0.4, 0.05, 0.0})
	col3     = palette1.Entry(color.LRGB{0.0, 0.4, 0.05})
	col4     = palette1.Entry(color.LRGB{0.0, 0.05, 0.45})
	col5     = palette1.Entry(color.LRGB{0.1, 0.0, 0.15})
	col6     = palette1.Entry(color.LRGB{0.25, 0.25, 0.25})
	col7     = palette1.Entry(color.LRGB{0.025, 0.025, 0.025})
)

var (
	points []coord.XY
)

var (
	ratio  float32
	offset coord.XY
)

const aboveall = int16(0x7FFF)

////////////////////////////////////////////////////////////////////////////////

func newPoints() {
	for i := range points {
		points[i] = coord.XY{X: rand.Float32(), Y: rand.Float32()}
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		err := cozely.Run(loop1{})
		if err != nil {
			t.Error(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type loop1 struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Enter() error {
	bindings.Load()
	context.Activate(1)

	points = make([]coord.XY, 3)
	newPoints()

	palette1.Activate()
	return nil
}

func (loop1) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop1) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}

	if next.JustPressed(1) {
		newPoints()
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop1) Render() error {
	canvas1.Clear(0)
	ratio = float32(canvas1.Size().R)
	offset = coord.XY{
		X: (float32(canvas1.Size().C) - ratio),
		Y: float32(canvas1.Size().R),
	}
	pt := make([]coord.CR, len(points))
	s := coord.CR{5, 5}
	for i, sd := range points {
		pt[i] = toScreen(sd)
		var c color.Index
		switch i {
		case 0:
			c = col2
		case 1:
			c = col3
		case 2:
			c = col4
		}
		canvas1.Box(c, c, 2, 2, pt[i].Minus(s), pt[i].Plus(s))
		canvas1.Locate(pt[i].C-2, pt[i].R+3, aboveall)
		canvas1.Text(col1-1, 0)
		canvas1.Print([]string{"A", "B", "C"}[i])
	}
	canvas1.Lines(col6, 0, pt[0], pt[1], pt[2], pt[0])
	canvas1.Triangles(col7, -5, pt[0], pt[1], pt[2], pt[0])

	m := canvas1.Mouse()
	p := fromScreen(m)
	canvas1.Locate(2, 8, aboveall)
	canvas1.Text(col1-1, 0)
	canvas1.Printf("A: %.3f, %.3f\n", points[0].X, points[0].Y)
	canvas1.Printf("B: %.3f, %.3f\n", points[1].X, points[1].Y)
	canvas1.Printf("C: %.3f, %.3f\n", points[2].X, points[2].Y)
	if p.X >= 0 {
		canvas1.Printf("   %.3f, %.3f\n", p.X, p.Y)
	} else {
		canvas1.Println(" ")
	}
	canvas1.Point(col1, 1, m)

	canvas1.Println()

	if plane.IsCCW(points[0], points[1], points[2]) {
		canvas1.Text(col4-1, 0)
		canvas1.Println("IsCCW: TRUE")
	} else {
		canvas1.Text(col1-1, 0)
		canvas1.Println("IsCCW: false")
	}

	if plane.InTriangle(points[0], points[1], points[2], p) {
		canvas1.Text(col2-1, 0)
		canvas1.Println("InTriangle: TRUE")
	} else {
		canvas1.Text(col1-1, 0)
		canvas1.Println("InTriangle: false")
	}

	a, b, c := 0, 1, 2
	if !plane.IsCCW(points[a], points[b], points[c]) {
		b, c = c, b
	}
	if plane.InTriangleCCW(points[a], points[b], points[c], p) {
		canvas1.Text(col2-1, 0)
		canvas1.Println("InTriangleCCW: TRUE")
	} else {
		canvas1.Text(col1-1, 0)
		canvas1.Println("InTriangleCCW: false")
	}

	if plane.InCircumcircle(points[a], points[b], points[c], p) {
		canvas1.Text(col3-1, 0)
		canvas1.Println("InCircumcircle: TRUE")
	} else {
		canvas1.Text(col1-1, 0)
		canvas1.Println("InCircumcircle: false")
	}

	canvas1.Println(" ")

	d := plane.Circumcenter(points[0], points[1], points[2])
	canvas1.Text(col1-1, 0)
	canvas1.Printf("Circumcenter: %.3f, %.3f\n", d.X, d.Y)
	dd := toScreen(d)
	canvas1.Lines(col5, -2, dd.Pluss(-2, -2), dd.Pluss(2, 2))
	canvas1.Lines(col5, -2, dd.Pluss(2, -2), dd.Pluss(-2, 2))

	r := d.Minus(points[a]).Length()
	cir := []coord.CR{}
	for a := float32(0); a <= 2*math32.Pi+0.01; a += math32.Pi / 32 {
		cir = append(cir, toScreen(coord.RA{r, a}.XY().Plus(d)))
	}
	canvas1.Lines(col5, -2, cir...)

	canvas1.Display()
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func toScreen(p coord.XY) coord.CR {
	return coord.CR{
		C: int16(offset.X + p.X*ratio),
		R: int16(offset.Y - p.Y*ratio),
	}
}

func fromScreen(p coord.CR) coord.XY {
	return (p.XY().FlipY().Minus(offset.FlipY())).Slash(ratio)
}

////////////////////////////////////////////////////////////////////////////////
