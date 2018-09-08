package quadedge_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/plane/quadedge"
	"github.com/cozely/cozely/window"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

var (
	col1 = pixel.DefaultPalette.Names["Indigo"]
	col2 = pixel.DefaultPalette.Names["Orange"]
	col3 = pixel.DefaultPalette.Names["White"]
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

		pixel.SetZoom(2)

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
		m := pixel.ToCanvas(window.XYof(cursor.XY(0)))
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
	pixel.Clear(1)
	ratio = float32(pixel.Resolution().Y)
	orig = coord.XY{
		X: (float32(pixel.Resolution().X) - ratio) / 2,
		Y: float32(pixel.Resolution().Y),
	}

	m := pixel.ToCanvas(window.XYof(cursor.XY(0)))
	p := fromScreen(m)
	pixel.Locate(pixel.XY{2, 8})
	pixel.Text(col3, 0)
	fsr, fso := cozely.RenderStats()
	pixel.Printf("Framerate: %.2f (%d)\n", 1000*fsr, fso)
	if p.X >= 0 && p.X <= 1.0 {
		pixel.Printf("Position: %.3f, %.3f\n", p.X, p.Y)
	} else {
		pixel.Println(" ")
	}

	triangulation.Walk(func(e quadedge.Edge) {
		pixel.Lines(col1, 0, toScreen(points[e.Orig()]), toScreen(points[e.Dest()]))
	})

	pt := make([]pixel.XY, len(points))
	for i, sd := range points {
		pt[i] = toScreen(sd)
		pixel.Box(col2, col2, 0, 1, pt[i].MinusS(2), pt[i].PlusS(2))
	}

	if window.HasMouseFocus() {
		pixel.MouseCursor. Paint(0, m)
	}
}

func toScreen(p coord.XY) pixel.XY {
	return pixel.XYof(orig.Plus(p.FlipY().Times(ratio)))
}

func fromScreen(p pixel.XY) coord.XY {
	return (p.Coord().FlipY().Minus(orig.FlipY())).Slash(ratio)
}

func newPoints() {
	for i := range points {
		points[i] = coord.XY{rand.Float32(), rand.Float32()}
	}
	triangulation = quadedge.Delaunay(points)
}
