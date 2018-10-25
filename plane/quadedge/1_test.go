package quadedge_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/pico8"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/plane/quadedge"
	"github.com/cozely/cozely/resource"
	"github.com/cozely/cozely/window"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

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

		color.Load(&pico8.Palette)
		pixel.SetZoom(2)
		err := resource.Path("testdata/")
		if err != nil {
			t.Error(err)
			return
		}

		err = cozely.Run(loop1{})
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
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}
	if next.Pressed() {
		newPoints()
	}

	if previous.Pressed() {
		m := pixel.XYof(input.Pointer.XYon(0))
		p := fromScreen(m)
		points = append(points, p)
		triangulation = quadedge.Delaunay(points)
	}

	if scene1.Pressed() {
		points = points[:0]
		const st = 1.0 / 8
		for x := float32(st); x < 1.0; x += st {
			for y := float32(st); y < 1.0; y += st {
				points = append(points, coord.XY{x, y})
			}
		}
		triangulation = quadedge.Delaunay(points)
	}
	if scene2.Pressed() {
		points = points[:0]
		for a := float32(0); a < 2*math32.Pi; a += math32.Pi / 8 {
			points = append(points, coord.XY{
				X: .5 - math32.Cos(a)*.5,
				Y: .5 + math32.Sin(a)*.5,
			})
		}
		triangulation = quadedge.Delaunay(points)
	}
	if scene3.Pressed() {
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
	if scene4.Pressed() {
		points = points[:0]
		const st = 1.0 / 6
		for x := float32(st); x < 1.0; x += st {
			points = append(points, coord.XY{x, 0.5 + 0.17*x})
		}
		triangulation = quadedge.Delaunay(points)
	}
	if scene5.Pressed() {
		points = make([]coord.XY, 25000)
		newPoints()
	}
	if scene6.Pressed() {
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
}

func (loop1) Update() {
}

func (loop1) Render() {
	pixel.Clear(pico8.DarkBlue)
	cur := pixel.Cursor{}

	ratio = float32(pixel.Resolution().Y)
	orig = coord.XY{
		X: (float32(pixel.Resolution().X) - ratio) / 2,
		Y: float32(pixel.Resolution().Y),
	}

	m := pixel.XYof(input.Pointer.XYon(0))
	p := fromScreen(m)
	cur.Position = pixel.XY{2, 8}
	cur.Color = pico8.White
	fsr, fso := cozely.RenderStats()
	cur.Printf("Framerate: %.2f (%d)\n", 1000*fsr, fso)
	if p.X >= 0 && p.X <= 1.0 {
		cur.Printf("Position: %.3f, %.3f\n", p.X, p.Y)
	} else {
		cur.Println(" ")
	}

	triangulation.Walk(func(e quadedge.Edge) {
		pixel.Line(toScreen(points[e.Orig()]), toScreen(points[e.Dest()]), 0, pico8.Indigo)
	})

	pt := make([]pixel.XY, len(points))
	for i, sd := range points {
		pt[i] = toScreen(sd)
		pixel.Fill.TileMod(pt[i].MinusS(1), pixel.XY{3, 3}, 0, pico8.Orange)
	}

	if window.HasMouseFocus() {
		pixel.MouseCursor.Paint(m, 0)
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
