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

////////////////////////////////////////////////////////////////////////////////

func TestTest5(t *testing.T) {
	do(func() {
		err := cozely.Run(loop5{})
		if err != nil {
			t.Error(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type loop5 struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop5) Enter() error {
	input.Load(bindings5)
	context5.Activate(1)
	palette2.Activate()
	return nil
}

func (loop5) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop5) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}

	if newPoint.JustPressed(1) {
		points = append(points, canvas5.Mouse())
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

////////////////////////////////////////////////////////////////////////////////

func (loop5) Render() error {
	canvas5.Clear(0)
	m := canvas5.Mouse()
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
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop5) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////
