package quadedge_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/plane/quadedge"
)

//------------------------------------------------------------------------------

var screen = pixel.NewCanvas(pixel.Zoom(2))

var cursor = pixel.NewCursor()

func init() {
	cursor.Canvas(screen)
}

var (
	points        []plane.Coord
	triangulation quadedge.Edge
)

//------------------------------------------------------------------------------

var c0 = []plane.Coord{
	{1, 2},
	{1, 2},
	{1, 2},
	{1, 2},
	{1, 2},
}

var four = []plane.Coord{
	{-10, -4},
	{10, -4},
	{0, 10},
	{0, 7.5},
}

func TestDelaunay(t *testing.T) {
	e := quadedge.Delaunay(c0)
	e.Walk(func(f quadedge.Edge) {
		fmt.Println(f.Orig(), "->", f.Dest())
	})
}

//------------------------------------------------------------------------------

func TestDelaunay_graphic(t *testing.T) {
	do(func() {
		err := glam.Run(delLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type delLoop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (delLoop) Enter() error {
	points = make([]plane.Coord, 15)
	newPoints()

	palette.Clear()
	palette.Index(1).SetColour(colour.LRGB{0.1, 0.3, 0.6})
	palette.Index(2).SetColour(colour.LRGB{0.8, 0.1, 0.0})
	return nil
}

//------------------------------------------------------------------------------

var (
	ratio  float32
	offset plane.Coord
)

func (delLoop) Draw() error {
	screen.Clear(0)
	ratio = float32(screen.Size().Y)
	offset = plane.Coord{
		X: (float32(screen.Size().X) - ratio) / 2,
		Y: float32(screen.Size().Y),
	}
	pt := make([]pixel.Coord, len(points))
	l2 := pixel.Coord{2, 2}
	l1 := pixel.Coord{1, 1}
	for i, sd := range points {
		pt[i] = toScreen(sd)
		screen.Lines(2, 0, pt[i].Minus(l2), pt[i].Plus(l2.Perp()),
			pt[i].Plus(l2), pt[i].Minus(l2.Perp()), pt[i].Minus(l2))
		screen.Lines(2, 0, pt[i].Minus(l1), pt[i].Plus(l1.Perp()),
			pt[i].Plus(l1), pt[i].Minus(l1.Perp()), pt[i].Minus(l1))
	}

	triangulation.Walk(func(e quadedge.Edge) {
		screen.Lines(1, 1, toScreen(points[e.Orig()]), toScreen(points[e.Dest()]))
	})

	screen.Display()
	return nil
}

//------------------------------------------------------------------------------

func toScreen(p plane.Coord) pixel.Coord {
	return pixel.Coord{
		X: int16(offset.X + p.X*ratio),
		Y: int16(offset.Y - p.Y*ratio),
	}
}

func fromScreen(p pixel.Coord) plane.Coord {
	return plane.Coord{
		X: (float32(p.X) - offset.X) / ratio,
		Y: (offset.Y - float32(p.Y)) / ratio,
	}
}

//------------------------------------------------------------------------------

func (delLoop) MouseButtonDown(b mouse.Button, _ int) {
	switch b {
	case mouse.Left:
		newPoints()
	case mouse.Right:
		p := fromScreen(screen.Mouse())
		points = append(points, p)
		triangulation = quadedge.Delaunay(points)
	}
}

//------------------------------------------------------------------------------

func newPoints() {
	for i := range points {
		points[i] = plane.Coord{X: rand.Float32(), Y: rand.Float32()}
	}
	triangulation = quadedge.Delaunay(points)
}

//------------------------------------------------------------------------------
