// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package plane_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/plane"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

var (
	canvas  = pixel.Canvas(pixel.Zoom(2))
	scene   = pixel.Scene()
	palette = color.Palette()
	col1    = palette.Entry(color.LRGB{1, 1, 1})
	col2    = palette.Entry(color.LRGB{0.4, 0.05, 0.0})
	col3    = palette.Entry(color.LRGB{0.0, 0.4, 0.05})
	col4    = palette.Entry(color.LRGB{0.0, 0.05, 0.45})
	col5    = palette.Entry(color.LRGB{0.1, 0.0, 0.15})
	col6    = palette.Entry(color.LRGB{0.25, 0.25, 0.25})
	col7    = palette.Entry(color.LRGB{0.025, 0.025, 0.025})
)

var (
	points []coord.XY
)

var (
	ratio  float32
	offset coord.XY
)

type loop1 struct{}

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		input.Load(bindings)
		err := cozely.Run(loop1{})
		if err != nil {
			t.Error(err)
		}
	})
}

func (loop1) Enter() {
	input.ShowMouse(false)

	points = make([]coord.XY, 3)
	newPoints()

	palette.Activate()
}

func (loop1) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) React() {
	if next.Started(0) {
		newPoints()
	}

	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop1) Update() {
}

func (loop1) Render() {
	canvas.Clear(0)
	ratio = float32(canvas.Size().R)
	offset = coord.XY{
		X: (float32(canvas.Size().C) - ratio),
		Y: float32(canvas.Size().R),
	}
	pt := make([]coord.CR, len(points))
	s := coord.CR{5, 5}
	for i, sd := range points {
		pt[i] = toScreen(sd)
	}

	d := plane.Circumcenter(points[0], points[1], points[2])
	r := d.Minus(points[0]).Length()
	cir := []coord.CR{}
	for a := float32(0); a <= 2*math32.Pi+0.01; a += math32.Pi / 32 {
		cir = append(cir, toScreen(coord.RA{r, a}.XY().Plus(d)))
	}
	scene.Lines(col5, cir...)
	scene.Triangles(col7, pt[0], pt[1], pt[2], pt[0])
	scene.Lines(col6, pt[0], pt[1], pt[2], pt[0])
	for i := range points {
		var c color.Index
		switch i {
		case 0:
			c = col2
		case 1:
			c = col3
		case 2:
			c = col4
		}
		scene.Box(c, c, 2, pt[i].Minus(s), pt[i].Plus(s))
		scene.Locate(coord.CR{pt[i].C - 2, pt[i].R + 3})
		scene.Text(col1, 0)
		scene.Print([]string{"A", "B", "C"}[i])
	}

	m := canvas.FromWindow(cursor.XY(0).CR())
	p := fromScreen(m)
	scene.Locate(coord.CR{2, 8})
	scene.Text(col1, 0)
	scene.Printf("A: %.3f, %.3f\n", points[0].X, points[0].Y)
	scene.Printf("B: %.3f, %.3f\n", points[1].X, points[1].Y)
	scene.Printf("C: %.3f, %.3f\n", points[2].X, points[2].Y)
	if p.X >= 0 {
		scene.Printf("   %.3f, %.3f\n", p.X, p.Y)
	} else {
		scene.Println(" ")
	}
	scene.Point(col1, m)

	scene.Println()

	if plane.IsCCW(points[0], points[1], points[2]) {
		scene.Text(col4, 0)
		scene.Println("IsCCW: TRUE")
	} else {
		scene.Text(col1, 0)
		scene.Println("IsCCW: false")
	}

	if plane.InTriangle(points[0], points[1], points[2], p) {
		scene.Text(col2, 0)
		scene.Println("InTriangle: TRUE")
	} else {
		scene.Text(col1, 0)
		scene.Println("InTriangle: false")
	}

	a, b, c := 0, 1, 2
	if !plane.IsCCW(points[a], points[b], points[c]) {
		b, c = c, b
	}
	if plane.InTriangleCCW(points[a], points[b], points[c], p) {
		scene.Text(col2, 0)
		scene.Println("InTriangleCCW: TRUE")
	} else {
		scene.Text(col1, 0)
		scene.Println("InTriangleCCW: false")
	}

	if plane.InCircumcircle(points[a], points[b], points[c], p) {
		scene.Text(col3, 0)
		scene.Println("InCircumcircle: TRUE")
	} else {
		scene.Text(col1, 0)
		scene.Println("InCircumcircle: false")
	}

	scene.Println(" ")

	scene.Text(col1, 0)
	scene.Printf("Circumcenter: %.3f, %.3f\n", d.X, d.Y)
	dd := toScreen(d)
	scene.Lines(col5, dd.Minuss(2), dd.Pluss(2))
	scene.Lines(col5, dd.Minus(coord.CR{-2, 2}), dd.Plus(coord.CR{-2, 2}))

	canvas.Display(scene)
}

func toScreen(p coord.XY) coord.CR {
	return coord.CR{
		C: int16(offset.X + p.X*ratio),
		R: int16(offset.Y - p.Y*ratio),
	}
}

func fromScreen(p coord.CR) coord.XY {
	return (p.XY().FlipY().Minus(offset.FlipY())).Slash(ratio)
}

func newPoints() {
	for i := range points {
		points[i] = coord.XY{X: rand.Float32(), Y: rand.Float32()}
	}
}
