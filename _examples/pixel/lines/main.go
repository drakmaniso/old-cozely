// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var screen = pixel.NewCanvas(pixel.TargetResolution(640, 360))

var (
	background = palette.Index(1)
)

//------------------------------------------------------------------------------

type line struct {
	ax, ay int16
	bx, by int16
}

var lines []line

//------------------------------------------------------------------------------

func main() {
	glam.Configure(
		glam.TimeStep(1 / 60.0),
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

func (loop) Enter() error {
	palette.Load("MSX2")
	return pixel.Err()
}

//------------------------------------------------------------------------------

var step int16

func (loop) Update() error {
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
		screen.Line(palette.Index(i+int(step)), ox+l.ax, oy+l.ay, ox+l.bx, oy+l.by)
	}
	screen.Display()
	return pixel.Err()
}

//------------------------------------------------------------------------------
