// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/input"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
)

//------------------------------------------------------------------------------

var (
	newPoint   = input.Bool("NewPoint")
	hidePoints = input.Bool("HidePoints")
	hideLines  = input.Bool("HideLines")
)

var cmdContext = input.Context("TestCommands",
	quit, newPoint, previous, hidePoints, hideLines)

var cmdBindings = input.Bindings{
	"TestCommands": {
		"Quit":       {"Escape"},
		"NewPoint":   {"Mouse Left"},
		"Previous":   {"Mouse Right", "U"},
		"HidePoints": {"P"},
		"HideLines":  {"L"},
	},
}

//------------------------------------------------------------------------------

var cmdScreen = pixel.Canvas(pixel.TargetResolution(128, 128))

var points = []plane.Pixel{
	{4, 4},
	{4 + 1, 4 + 20},
	{4 + 1 + 20, 4 + 20 - 1},
	{16, 32},
}

var pointshidden, lineshidden bool

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

type cmdLoop struct{}

//------------------------------------------------------------------------------

func (cmdLoop) Enter() error {
	input.Load(cmdBindings)
	cmdContext.Activate(1)
	palette.Load("graphics/shape1")
	return nil
}

func (cmdLoop) Leave() error { return nil }

//------------------------------------------------------------------------------

func (cmdLoop) React() error {
	if quit.JustPressed(1) {
		glam.Stop()
	}

	if newPoint.JustPressed(1) {
		points = append(points, cmdScreen.Mouse())
	}

	if previous.JustPressed(1) {
		if len(points) > 0 {
			points = points[:len(points)-1]
		}
	}

	pointshidden = hidePoints.Pressed(1)
	lineshidden = hideLines.Pressed(1)

	return nil
}

//------------------------------------------------------------------------------

func (cmdLoop) Render() error {
	cmdScreen.Clear(0)
	m := cmdScreen.Mouse()
	cmdScreen.Triangles(2, -2, points...)
	if !lineshidden {
		cmdScreen.Lines(5, 0, points...)
		cmdScreen.Lines(13, -1, points[len(points)-1], m)
	}
	if !pointshidden {
		for _, p := range points {
			cmdScreen.Point(8, 1, p)
		}
		cmdScreen.Point(18, 2, m)
	}
	cmdScreen.Display()
	return nil
}

//------------------------------------------------------------------------------

func (cmdLoop) Update() error { return nil }

//------------------------------------------------------------------------------
