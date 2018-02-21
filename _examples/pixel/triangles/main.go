// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var screen = pixel.NewCanvas(pixel.TargetResolution(128, 128))

var (
	foreground    = palette.Entry(1, "fg", colour.SRGB{1, 1, 1})
	background    = palette.Entry(2, "bg", colour.SRGB{0, 0, 0})
	pointColor    = palette.Entry(3, "pointColor", colour.SRGB{1, 0, 0})
	lineColor = palette.Entry(6, "lineColor", colour.SRGB{0, 1, 0})
	triangleColor = palette.Entry(5, "triangleColor", colour.SRGB{0, 0, 1})
)

//------------------------------------------------------------------------------

var points = []pixel.Coord{}

//------------------------------------------------------------------------------

func main() {
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

func (loop) Draw() error {
	screen.Clear(background)
	screen.Triangles(triangleColor, 0, points...)
	if !key.IsPressed(key.PositionSpace) {
		screen.Lines(lineColor, 0, points...)
	}
	if !key.IsPressed(key.PositionLShift) {
		for _, p := range points {
			screen.Point(pointColor, p.X, p.Y, 0)
		}
	}
	screen.Display()
	return nil
}

//------------------------------------------------------------------------------

func (loop) MouseButtonDown(b mouse.Button, _ int) {
	switch b {
	case mouse.Left:
		points = append(points, screen.Mouse())
	case mouse.Right:
		if len(points) > 0 {
			points = points[:len(points)-1]
		}
	}
}

//------------------------------------------------------------------------------
