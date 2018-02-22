// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var cmdScreen = pixel.NewCanvas(pixel.TargetResolution(128, 128))

var points = []pixel.Coord{
	{4, 4},
	{4 + 1, 4 + 20},
	{4 + 1 + 20, 4 + 20 - 1},
	{16, 32},
}

//------------------------------------------------------------------------------

func TestPaint_commands(t *testing.T) {
	do(func() {
		err := glam.Run(cmdLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type cmdLoop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (cmdLoop) Enter() error {
	palette.Load("graphics/shape1")
	return nil
}

//------------------------------------------------------------------------------

func (cmdLoop) Draw() error {
	cmdScreen.Clear(0)
	m := cmdScreen.Mouse()
	cmdScreen.Triangles(2, -2, points...)
	if !key.IsPressed(key.PositionSpace) {
		cmdScreen.Lines(5, 0, points...)
		cmdScreen.Lines(13, -1, points[len(points)-1], m)
	}
	if !key.IsPressed(key.PositionLShift) {
		for _, p := range points {
			cmdScreen.Point(8, p.X, p.Y, 1)
			cmdScreen.Point(18, m.X, m.Y, 2)
		}
	}
	cmdScreen.Display()
	return nil
}

//------------------------------------------------------------------------------

func (cmdLoop) MouseButtonDown(b mouse.Button, _ int) {
	switch b {
	case mouse.Left:
		points = append(points, cmdScreen.Mouse())
	case mouse.Right:
		if len(points) > 0 {
			points = points[:len(points)-1]
		}
	}
}

//------------------------------------------------------------------------------
