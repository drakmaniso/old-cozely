// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
	err := glam.Setup()
	if err != nil {
		glam.ShowError("setting up glam", err)
		return
	}

	glam.Update = update
	glam.Draw = draw
	window.Handle = handler{}
	mouse.Handle = handler{}
	key.Handle = handler{}

	err = glam.LoopStable(1 / 60.0)
	if err != nil {
		glam.ShowError("running", err)
		return
	}
}

//------------------------------------------------------------------------------

type handler struct{}

func (h handler) KeyDown(l key.Label, p key.Position, ts uint32) {
	if l == key.LabelEscape {
		glam.Stop()
	}
	scroller.Print("%v: Key Down: %v %v\n", ts, l, p)
}

func (h handler) KeyUp(l key.Label, p key.Position, ts uint32) {
	scroller.Print("%v: Key Up: %v %v\n", ts, l, p)
}

func (h handler) MouseMotion(rel pixel.Coord, pos pixel.Coord, ts uint32) {
	scroller.Print("%v: mouse motion  %+d,%+d  %d,%d\n", ts, rel.X, rel.Y, pos.X, pos.Y)
}

func (h handler) MouseButtonDown(b mouse.Button, clicks int, ts uint32) {
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
	scroller.Print("%v: mouse button down  %s (%v), clicks=%v\n", ts, n, b, clicks)
}

func (h handler) MouseButtonUp(b mouse.Button, clicks int, ts uint32) {
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
	scroller.Print("%v: mouse button up: %s (%v), clicks=%v\n", ts, n, b, clicks)
}

func (h handler) MouseWheel(w pixel.Coord, ts uint32) {
	scroller.Print("%v: mouse wheel: %+d,%+d\n", ts, w.X, w.Y)
}

func (h handler) WindowShown(ts uint32) {
	scroller.Print("%v: window shown\n", ts)
}

func (h handler) WindowHidden(ts uint32) {
	scroller.Print("%v: window hidden\n", ts)
}

func (h handler) WindowResized(s pixel.Coord, ts uint32) {
	scroller.Print("%v: window resized %v\n", ts, s)
}

func (h handler) WindowMinimized(ts uint32) {
	scroller.Print("%v: window minimized\n", ts)
}

func (h handler) WindowMaximized(ts uint32) {
	scroller.Print("%v: window maximized\n", ts)
}

func (h handler) WindowRestored(ts uint32) {
	scroller.Print("%v: window restored\n", ts)
}

func (h handler) WindowMouseEnter(ts uint32) {
	scroller.Print("%v: window mouse enter\n", ts)
}

func (h handler) WindowMouseLeave(ts uint32) {
	scroller.Print("%v: window mouse leave\n", ts)
}

func (h handler) WindowFocusGained(ts uint32) {
	scroller.Print("%v: window focus gained\n", ts)
}

func (h handler) WindowFocusLost(ts uint32) {
	scroller.Print("%v: window focus lost\n", ts)
}

func (h handler) WindowQuit(ts uint32) {
	scroller.Print("%v: window quit\n", ts)
	glam.Stop()
}

var scroller = mtx.Clip{
	Left: 1, Top: 4,
	Right: -2, Bottom: -1,
	VScroll:   true,
	ClearChar: ' ',
}

//------------------------------------------------------------------------------
func update(_, _ float64) {
	topbar.Clear()

	d := mouse.Delta()
	topbar.Print("   mouse.Delta():%+6d,%+6d\v\n", d.X, d.Y)
	p := mouse.Position()
	topbar.Print("mouse.Position():%6d,%6d\v\n", p.X, p.Y)

	topbar.Print("   mouse buttons: ")
	if mouse.IsPressed(mouse.Left) {
		topbar.Print("\aleft ")
	} else {
		topbar.Print("left ")
	}
	if mouse.IsPressed(mouse.Middle) {
		topbar.Print("\amiddle ")
	} else {
		topbar.Print("middle ")
	}
	if mouse.IsPressed(mouse.Right) {
		topbar.Print("\aright ")
	} else {
		topbar.Print("right ")
	}
	if mouse.IsPressed(mouse.Extra1) {
		topbar.Print("\aextra1 ")
	} else {
		topbar.Print("extra1 ")
	}
	if mouse.IsPressed(mouse.Extra2) {
		topbar.Print("\aextra2\n")
	} else {
		topbar.Print("extra2\n")
	}
}

var topbar = mtx.Clip{
	Left: 1, Top: 0,
	Right: -2, Bottom: 2,
}

func draw() {
}

//------------------------------------------------------------------------------
