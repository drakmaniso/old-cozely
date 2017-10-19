// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/color"
	"github.com/drakmaniso/carol/gfx"
	"github.com/drakmaniso/carol/key"
	"github.com/drakmaniso/carol/mouse"
	"github.com/drakmaniso/carol/mtx"
	"github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

func main() {
	err := carol.Setup()
	if err != nil {
		carol.ShowError("setting up carol", err)
		return
	}

	carol.Loop(loop{})

	err = carol.Run()
	if err != nil {
		carol.ShowError("running", err)
		return
	}
}

//------------------------------------------------------------------------------

type loop struct {
	carol.DefaultHandlers
}

//------------------------------------------------------------------------------

func (loop) KeyDown(l key.Label, p key.Position) {
	if l == key.LabelEscape {
		carol.Stop()
	}
	scroller.Print("%v: Key Down: %v %v\n", carol.Now(), l, p)
}

func (loop) KeyUp(l key.Label, p key.Position) {
	scroller.Print("%v: Key Up: %v %v\n", carol.Now(), l, p)
}

//------------------------------------------------------------------------------

func (loop) MouseMotion(rel pixel.Coord, pos pixel.Coord) {
	scroller.Print("%v: mouse motion  %+d,%+d  %d,%d\n", carol.Now(), rel.X, rel.Y, pos.X, pos.Y)
}

func (loop) MouseButtonDown(b mouse.Button, clicks int) {
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
	scroller.Print("%v: mouse button down  %s (%v), clicks=%v\n", carol.Now(), n, b, clicks)
}

func (loop) MouseButtonUp(b mouse.Button, clicks int) {
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
	scroller.Print("%v: mouse button up: %s (%v), clicks=%v\n", carol.Now(), n, b, clicks)
}

func (loop) MouseWheel(w pixel.Coord) {
	scroller.Print("%v: mouse wheel: %+d,%+d\n", carol.Now(), w.X, w.Y)
}

//------------------------------------------------------------------------------

func (loop) WindowShown() {
	scroller.Print("%v: window shown\n", carol.Now())
}

func (loop) WindowHidden() {
	scroller.Print("%v: window hidden\n", carol.Now())
}

func (loop) WindowResized(s pixel.Coord) {
	scroller.Print("%v: window resized %v\n", carol.Now(), s)
}

func (loop) WindowMinimized() {
	scroller.Print("%v: window minimized\n", carol.Now())
}

func (loop) WindowMaximized() {
	scroller.Print("%v: window maximized\n", carol.Now())
}

func (loop) WindowRestored() {
	scroller.Print("%v: window restored\n", carol.Now())
}

func (loop) WindowMouseEnter() {
	scroller.Print("%v: window mouse enter\n", carol.Now())
}

func (loop) WindowMouseLeave() {
	scroller.Print("%v: window mouse leave\n", carol.Now())
}

func (loop) WindowFocusGained() {
	scroller.Print("%v: window focus gained\n", carol.Now())
}

func (loop) WindowFocusLost() {
	scroller.Print("%v: window focus lost\n", carol.Now())
}

func (loop) WindowQuit() {
	scroller.Print("%v: window quit\n", carol.Now())
	carol.Stop()
}

var scroller = mtx.Clip{
	Left: 1, Top: 4,
	Right: -2, Bottom: -1,
}

//------------------------------------------------------------------------------

func (loop) Update() {
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

func (loop) Draw(_, _ float64) {
	gfx.ClearColorBuffer(color.RGBA{})
}

//------------------------------------------------------------------------------
