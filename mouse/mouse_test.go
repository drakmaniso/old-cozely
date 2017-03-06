// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mouse_test

//------------------------------------------------------------------------------

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

func TestMain(m *testing.M) {
	glam.Setup()

	mouse.Handle = handler{}

	// Run the main loop
	glam.Loop = looper{}
	err := glam.Run()
	if err != nil {
		fmt.Printf("Glam: Error: %v\n", err)
		os.Exit(-1)
	}

	os.Exit(m.Run())
}

func Test(t *testing.T) {
}

//------------------------------------------------------------------------------

type handler struct {
}

func (h handler) MouseMotion(rel pixel.Coord, pos pixel.Coord, ts time.Duration) {
	scroller.Print("%s:  Motion: relative=%v, position=%v\n", ts, rel, pos)
}

func (h handler) MouseButtonDown(b mouse.Button, clicks int, ts time.Duration) {
	var n string
	switch b {
	case mouse.Left:
		n = "Left"
	case mouse.Middle:
		n = "Middle"
	case mouse.Right:
		n = "Right"
	case mouse.Extra1:
		n = "Extra1"
	case mouse.Extra2:
		n = "Extra2"
	default:
		n = "UNKOWN!"
	}
	scroller.Print("%v:  Button Down: %s (%v), clicks=%v\n", ts, n, b, clicks)
}

func (h handler) MouseButtonUp(b mouse.Button, clicks int, ts time.Duration) {
	var n string
	switch b {
	case mouse.Left:
		n = "Left"
	case mouse.Middle:
		n = "Middle"
	case mouse.Right:
		n = "Right"
	case mouse.Extra1:
		n = "Extra1"
	case mouse.Extra2:
		n = "Extra2"
	default:
		n = "UNKOWN!"
	}
	scroller.Print("%v:  Button Up: %s (%v), clicks=%v\n", ts, n, b, clicks)
}

func (h handler) MouseWheel(w pixel.Coord, ts time.Duration) {
	scroller.Print("%v:  Wheel: %v\n", ts, w)
}

var scroller = mtx.Clip{
	Left: 1, Top: 4,
	Right: -2, Bottom: -1,
	VScroll:   true,
	ClearChar: ' ',
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Draw() {}

func (l looper) Update() {
	timer += glam.TimeStep
	if timer > time.Second/10 {
		timer = 0

		topbar.Clear()

		d := mouse.Delta()
		topbar.Print("   mouse.Delta():%+6d,%+6d\n", d.X, d.Y)
		p := mouse.Position()
		topbar.Print("mouse.Position():%+6d,%+6d\n", p.X, p.Y)

		topbar.Print("   mouse buttons: ")
		if mouse.IsPressed(mouse.Left) {
			topbar.Print("\aleft\a ")
		} else {
			topbar.Print("left ")
		}
		if mouse.IsPressed(mouse.Middle) {
			topbar.Print("\amiddle\a ")
		} else {
			topbar.Print("middle ")
		}
		if mouse.IsPressed(mouse.Right) {
			topbar.Print("\aright\a ")
		} else {
			topbar.Print("right ")
		}
		if mouse.IsPressed(mouse.Extra1) {
			topbar.Print("\aextra1\a ")
		} else {
			topbar.Print("extra1 ")
		}
		if mouse.IsPressed(mouse.Extra2) {
			topbar.Print("\aextra2\a\n")
		} else {
			topbar.Print("extra2\n")
		}
	}
}

var timer time.Duration

var topbar = mtx.Clip{
	Left: 1, Top: 0,
	Right: -2, Bottom: 2,
}

//------------------------------------------------------------------------------
