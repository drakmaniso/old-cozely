// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"math/rand"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var screen = pixel.NewCanvas(pixel.TargetResolution(300, 300))

var (
	background = palette.Index(1)
)

//------------------------------------------------------------------------------

type line struct {
	ax, ay int16
	bx, by int16
}

var lines []line

type point struct {
	pos   pixel.Coord
	color palette.Index
}

var points [2048]point

//------------------------------------------------------------------------------

func main() {
	glam.Configure(
		glam.TimeStep(1 / 30.0),
	)
	err := glam.Run(loop{})
	if err != nil {
		glam.ShowError(err)
	}
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (loop) WindowResized(_, _ int32) {
	for i := range points {
		points[i].pos.X = int16(rand.Intn(int(screen.Size().X)))
		points[i].pos.Y = int16(rand.Intn(int(screen.Size().Y)))
		points[i].color = palette.Index(1 + rand.Intn(255))
	}
}

//------------------------------------------------------------------------------

func (loop) Enter() error {
	palette.Load("MSX2")
	return pixel.Err()
}

//------------------------------------------------------------------------------

var step int16

func (loop) Update() error {
	for i := range points {
		points[i].color += 11
	}
	switch {
	case step < 100:
		s := step
		lines = append(
			lines,
			line{
				ax: 0, ay: 3 * s,
				bx: 3 * s, by: 300,
			},
		)
	case step < 200:
		s := step - 100
		lines = append(
			lines,
			line{
				ax: 3 * s, ay: 300,
				bx: 300, by: 300 - 3*s,
			},
		)
	case step < 300:
		s := step - 200
		lines = append(
			lines,
			line{
				ax: 300, ay: 300 - 3*s,
				bx: 300 - 3*s, by: 0,
			},
		)
	case step < 400:
		s := step - 300
		lines = append(
			lines,
			line{
				ax: 300 - 3*s, ay: 0,
				bx: 0, by: 3 * s,
			},
		)
	default:
	}
	step++
	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw() error {
	screen.Clear(0)
	s := screen.Size()
	ox, oy := (s.X-300)/2, (s.Y-300)/2
	for i, l := range lines {
		screen.Lines(
			palette.Index(i+int(step)),
			int16(-i),
			pixel.Coord{ox+l.ax, oy+l.ay},
			pixel.Coord{ox+l.bx, oy+l.by},
		)
	}
	for _, p := range points {
		screen.Point(p.color, p.pos.X, p.pos.Y, 1)
	}
	screen.Display()
	return pixel.Err()
}

//------------------------------------------------------------------------------
