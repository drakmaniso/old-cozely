package quadedge_test

import (
	"math/rand"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/colour"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/coord/plane/quadedge"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

var (
	screen = pixel.Canvas(pixel.Zoom(2))
	cursor = pixel.Cursor{Canvas: screen}
)

var (
	points        []coord.XY
	triangulation quadedge.Edge
)

var (
	ratio float32
	orig  coord.XY
)

////////////////////////////////////////////////////////////////////////////////

func TestDelaunay(t *testing.T) {
	do(func() {
		err := cozely.Run(delLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type delLoop struct{}

////////////////////////////////////////////////////////////////////////////////

func (delLoop) Enter() error {
	input.Load(testBindings)
	testContext.Activate(1)

	points = make([]coord.XY, 64)
	newPoints()

	palette.Clear()
	palette.Index(1).SetColour(colour.LRGB{0.1, 0.2, 0.5})
	palette.Index(2).SetColour(colour.LRGB{0.5, 0.1, 0.0})
	return nil
}

func (delLoop) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (delLoop) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}

	if next.JustPressed(1) {
		newPoints()
	}

	if previous.JustPressed(1) {
		p := fromScreen(screen.Mouse())
		points = append(points, p)
		triangulation = quadedge.Delaunay(points)
	}

	if scene1.JustPressed(1) {
		points = points[:0]
		const st = 1.0 / 16
		for x := float32(st); x < 1.0; x += st {
			for y := float32(st); y < 1.0; y += st {
				points = append(points, coord.XY{x, y})
			}
		}
		triangulation = quadedge.Delaunay(points)
	}
	if scene2.JustPressed(1) {
		points = points[:0]
		for a := float32(0); a < 2*math32.Pi; a += math32.Pi / 8 {
			points = append(points, coord.XY{
				X: .5 - math32.Cos(a)*.5,
				Y: .5 + math32.Sin(a)*.5,
			})
		}
		triangulation = quadedge.Delaunay(points)
	}
	if scene3.JustPressed(1) {
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
	if scene4.JustPressed(1) {
		points = points[:0]
		const st = 1.0 / 6
		for x := float32(st); x < 1.0; x += st {
			points = append(points, coord.XY{x, 0.5 + 0.17*x})
		}
		triangulation = quadedge.Delaunay(points)
	}
	if scene5.JustPressed(1) {
		points = make([]coord.XY, 25000)
		newPoints()
	}
	if scene6.JustPressed(1) {
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

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (delLoop) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (delLoop) Render() error {
	screen.Clear(0)
	ratio = float32(screen.Size().R)
	orig = coord.XY{
		X: (float32(screen.Size().C) - ratio) / 2,
		Y: float32(screen.Size().R),
	}

	m := screen.Mouse()
	p := fromScreen(m)
	cursor.Locate(2, 8)
	cursor.Color = 0
	fsr, fso := cozely.RenderStats()
	cursor.Printf("Framerate: %.2f (%d)\n", 1000*fsr, fso)
	if p.X >= 0 && p.X <= 1.0 {
		cursor.Printf("Position: %.3f, %.3f\n", p.X, p.Y)
	} else {
		cursor.Println(" ")
	}

	pt := make([]coord.CR, len(points))
	for i, sd := range points {
		pt[i] = toScreen(sd)
		screen.Box(2, 2, 1, 0, pt[i].Minuss(2, 2), pt[i].Pluss(2, 2))
	}

	triangulation.Walk(func(e quadedge.Edge) {
		screen.Lines(1, -1, toScreen(points[e.Orig()]), toScreen(points[e.Dest()]))
	})

	screen.Display()
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func toScreen(p coord.XY) coord.CR {
	// return pixel.Coord(p.Times(ratio).Pixel()).Plus(orig)
	// return pixel.From(p.Times(ratio).XY()).Plus(orig)
	// return pixel.From(p.Times(ratio)).Plus(orig)
	// return p.Times(ratio).Pixel().Plus(orig)
	return coord.CR{
		C: int16(orig.X + p.X*ratio),
		R: int16(orig.Y - p.Y*ratio),
	}
}

func fromScreen(p coord.CR) coord.XY {
	// return plane.From(p.Minus(orig).XY()).Slash(ratio)
	// return plane.From(p.Minus(orig)).Slash(ratio)
	// return p.Minus(orig).Cartesian().Slash(ratio)
	return coord.XY{
		X: (float32(p.C) - orig.X) / ratio,
		Y: (orig.Y - float32(p.R)) / ratio,
	}
}

////////////////////////////////////////////////////////////////////////////////

func newPoints() {
	for i := range points {
		points[i] = coord.XY{rand.Float32(), rand.Float32()}
	}
	triangulation = quadedge.Delaunay(points)
}

////////////////////////////////////////////////////////////////////////////////
