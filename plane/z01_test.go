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
	col1 = pixel.DefaultPalette.ByName["White"]
	col2 = pixel.DefaultPalette.ByName["Red"]
	col3 = pixel.DefaultPalette.ByName["Green"]
	col4 = pixel.DefaultPalette.ByName["Blue"]
	col5 = pixel.DefaultPalette.ByName["Dark Gray"]
	col6 = pixel.DefaultPalette.ByName["Light Gray"]
	col7 = pixel.DefaultPalette.ByName["Black"]
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

		pixel.SetZoom(3)

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
}

func (loop1) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) React() {
	if next.Pushed() {
		newPoints()
	}

	if quit.Pushed() {
		cozely.Stop(nil)
	}
}

func (loop1) Update() {
}

func (loop1) Render() {
	pixel.Clear(1)
	cur := pixel.Cursor{}

	ratio = float32(pixel.Resolution().Y)
	offset = coord.XY{
		X: (float32(pixel.Resolution().X) - ratio),
		Y: float32(pixel.Resolution().Y),
	}
	pt := make([]pixel.XY, len(points))
	s := pixel.XY{5, 5}
	for i, sd := range points {
		pt[i] = toScreen(sd)
	}

	d := plane.Circumcenter(points[0], points[1], points[2])
	r := d.Minus(points[0]).Length()
	cir := []pixel.XY{}
	for a := float32(0); a <= 2*math32.Pi+0.01; a += math32.Pi / 32 {
		cir = append(cir, toScreen(coord.RA{r, a}.XY().Plus(d)))
	}
	pixel.Lines(col5, 0, cir...)
	pixel.Triangles(col7, 0, pt[0], pt[1], pt[2], pt[0])
	pixel.Lines(col6, 0, pt[0], pt[1], pt[2], pt[0])
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
		pixel.Box(c, c, 0, 2, pt[i].Minus(s), pt[i].Plus(s))
		cur.Locate(0, pixel.XY{pt[i].X - 2, pt[i].Y + 3})
		cur.Style(col1, 0)
		cur.Print([]string{"A", "B", "C"}[i])
	}

	m := pixel.XYof(cursor.XYon(0))
	p := fromScreen(m)
	cur.Locate(1, pixel.XY{2, 8})
	cur.Style(col1, 0)
	cur.Printf("A: %.3f, %.3f\n", points[0].X, points[0].Y)
	cur.Printf("B: %.3f, %.3f\n", points[1].X, points[1].Y)
	cur.Printf("C: %.3f, %.3f\n", points[2].X, points[2].Y)
	if p.X >= 0 {
		cur.Printf("   %.3f, %.3f\n", p.X, p.Y)
	} else {
		cur.Println(" ")
	}
	pixel.Point(col1, 0, m)

	cur.Println()

	if plane.IsCCW(points[0], points[1], points[2]) {
		cur.Style(col4, 0)
		cur.Println("IsCCW: TRUE")
	} else {
		cur.Style(col1, 0)
		cur.Println("IsCCW: false")
	}

	if plane.InTriangle(points[0], points[1], points[2], p) {
		cur.Style(col2, 0)
		cur.Println("InTriangle: TRUE")
	} else {
		cur.Style(col1, 0)
		cur.Println("InTriangle: false")
	}

	a, b, c := 0, 1, 2
	if !plane.IsCCW(points[a], points[b], points[c]) {
		b, c = c, b
	}
	if plane.InTriangleCCW(points[a], points[b], points[c], p) {
		cur.Style(col2, 0)
		cur.Println("InTriangleCCW: TRUE")
	} else {
		cur.Style(col1, 0)
		cur.Println("InTriangleCCW: false")
	}

	if plane.InCircumcircle(points[a], points[b], points[c], p) {
		cur.Style(col3, 0)
		cur.Println("InCircumcircle: TRUE")
	} else {
		cur.Style(col1, 0)
		cur.Println("InCircumcircle: false")
	}

	cur.Println(" ")

	cur.Style(col1, 0)
	cur.Printf("Circumcenter: %.3f, %.3f\n", d.X, d.Y)
	dd := toScreen(d)
	pixel.Lines(col5, 0, dd.MinusS(2), dd.PlusS(2))
	pixel.Lines(col5, 0, dd.Minus(pixel.XY{-2, 2}), dd.Plus(pixel.XY{-2, 2}))
}

func toScreen(p coord.XY) pixel.XY {
	return pixel.XY{
		int16(offset.X + p.X*ratio),
		int16(offset.Y - p.Y*ratio),
	}
}

func fromScreen(p pixel.XY) coord.XY {
	return coord.XYof(p.Coord().FlipY().Minus(offset.FlipY())).Slash(ratio)
}

func newPoints() {
	for i := range points {
		points[i] = coord.XY{X: rand.Float32(), Y: rand.Float32()}
	}
}
