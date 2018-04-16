// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var (
	newPoint   = input.Bool("NewPoint")
	hidePoints = input.Bool("HidePoints")
	hideLines  = input.Bool("HideLines")
)

var context5 = input.Context("TestCommands",
	quit, newPoint, previous, hidePoints, hideLines)

var bindings5 = input.Bindings{
	"TestCommands": {
		"Quit":       {"Escape"},
		"NewPoint":   {"Mouse Left"},
		"Previous":   {"Mouse Right", "U"},
		"HidePoints": {"P"},
		"HideLines":  {"L"},
	},
}

////////////////////////////////////////////////////////////////////////////////

var canvas5 = pixel.Canvas(pixel.Resolution(128, 128))

var points = []coord.CR{
	{4, 4},
	{4 + 1, 4 + 20},
	{4 + 1 + 20, 4 + 20 - 1},
	{16, 32},
}

var pointshidden, lineshidden bool

type loop5 struct{}

////////////////////////////////////////////////////////////////////////////////

func TestTest5(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		err := cozely.Run(loop5{})
		if err != nil {
			t.Error(err)
		}
	})
}

func (loop5) Enter() {
	input.Bind(bindings5)
	context5.Activate(1)
	palette2.Activate()
}

func (loop5) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop5) React() {
	if quit.JustPressed(1) {
		cozely.Stop(nil)
	}

	if newPoint.JustPressed(1) {
		m := canvas5.FromWindow(input.Cursor.Position())
		points = append(points, m)
	}

	if previous.JustPressed(1) {
		if len(points) > 0 {
			points = points[:len(points)-1]
		}
	}

	pointshidden = hidePoints.Pressed(1)
	lineshidden = hideLines.Pressed(1)
}

func (loop5) Update() {
}

func (loop5) Render() {
	canvas5.Clear(1)
	m := canvas5.FromWindow(input.Cursor.Position())
	canvas5.Triangles(2, -2, points...)
	if !lineshidden {
		canvas5.Lines(5, 0, points...)
		canvas5.Lines(13, -1, points[len(points)-1], m)
	}
	if !pointshidden {
		for _, p := range points {
			canvas5.Point(8, 1, p)
		}
		canvas5.Point(18, 2, m)
	}
	canvas5.Display()
}
