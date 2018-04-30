package quadedge_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/plane/quadedge"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

var (
	canvas1  = pixel.Canvas(pixel.Zoom(2))
	palette1 = color.Palette(color.LRGB{1, 1, 1})
	col1     = palette1.Entry(color.LRGB{0.1, 0.2, 0.5})
	col2     = palette1.Entry(color.LRGB{0.5, 0.1, 0.0})
)

var (
	points        []coord.XY
	triangulation quadedge.Edge
)

var (
	ratio float32
	orig  coord.XY
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

	points = make([]coord.XY, 64)
	newPoints()

	palette1.Activate()
}

func (loop1) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) React() {
	if quit.Started(0) {
		cozely.Stop(nil)
	}
	if next.Started(0) {
		newPoints()
	}

	if previous.Started(0) {
		m := canvas1.FromWindow(cursor.XY(0).CR())
		p := fromScreen(m)
		points = append(points, p)
		triangulation = quadedge.Delaunay(points)
	}

	if scene1.Started(0) {
		points = points[:0]
		const st = 1.0 / 8
		for x := float32(st); x < 1.0; x += st {
			for y := float32(st); y < 1.0; y += st {
				points = append(points, coord.XY{x, y})
			}
		}
		triangulation = quadedge.Delaunay(points)
	}
	if scene2.Started(0) {
		points = points[:0]
		for a := float32(0); a < 2*math32.Pi; a += math32.Pi / 8 {
			points = append(points, coord.XY{
				X: .5 - math32.Cos(a)*.5,
				Y: .5 + math32.Sin(a)*.5,
			})
		}
		triangulation = quadedge.Delaunay(points)
	}
	if scene3.Started(0) {
		points = points[:0]
		const n = 26
		for a := float32(0); a < n*2*math32.Pi; a += math32.Pi / 26 {
			points = append(points, coord.XY{
				X: .5 + math32.Cos(a)*.5*a/(n*2*math32.Pi),
				Y: .5 + math32.Sin(a)*.5*a/(n*2*math32.Pi),
			})
		}
		triangulation = quadedge.Delaunay(points)
	}
	if scene4.Started(0) {
		points = points[:0]
		const st = 1.0 / 6
		for x := float32(st); x < 1.0; x += st {
			points = append(points, coord.XY{x, 0.5 + 0.17*x})
		}
		triangulation = quadedge.Delaunay(points)
	}
	if scene5.Started(0) {
		points = make([]coord.XY, 25000)
		newPoints()
	}
	if scene6.Started(0) {
		points = points[:0]
		const st = 1.0 / 27
		const h = 0.5 * 1.732050807568877 * st
		for x := float32(st); x < 1.0; x += st {
			for y := float32(st); y < 1.0; y += h {
				points = append(points, coord.XY{x + 0.5*y - 0.25, y})
			}
		}
		triangulation = quadedge.Delaunay(points)
	}

	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop1) Update() {
}

func (loop1) Render() {
	canvas1.Clear(0)
	ratio = float32(canvas1.Size().R)
	orig = coord.XY{
		X: (float32(canvas1.Size().C) - ratio) / 2,
		Y: float32(canvas1.Size().R),
	}

	m := canvas1.FromWindow(cursor.XY(0).CR())
	p := fromScreen(m)
	canvas1.Locate(0, coord.CR{2, 8})
	canvas1.Text(col1, 0)
	fsr, fso := cozely.RenderStats()
	canvas1.Printf("Framerate: %.2f (%d)\n", 1000*fsr, fso)
	if p.X >= 0 && p.X <= 1.0 {
		canvas1.Printf("Position: %.3f, %.3f\n", p.X, p.Y)
	} else {
		canvas1.Println(" ")
	}

	triangulation.Walk(func(e quadedge.Edge) {
		canvas1.Lines(col1, -1, toScreen(points[e.Orig()]), toScreen(points[e.Dest()]))
	})

	pt := make([]coord.CR, len(points))
	for i, sd := range points {
		pt[i] = toScreen(sd)
		canvas1.Box(col2, col2, 1, 0, pt[i].Minuss(2), pt[i].Pluss(2))
	}

	if cozely.HasMouseFocus() {
		canvas1.Picture(pixel.MouseCursor, 100, m)
	}

	canvas1.Display()
}

func toScreen(p coord.XY) coord.CR {
	return orig.Plus(p.FlipY().Times(ratio)).CR()
}

func fromScreen(p coord.CR) coord.XY {
	return (p.XY().FlipY().Minus(orig.FlipY())).Slash(ratio)
}

func newPoints() {
	for i := range points {
		points[i] = coord.XY{rand.Float32(), rand.Float32()}
	}
	triangulation = quadedge.Delaunay(points)
}
