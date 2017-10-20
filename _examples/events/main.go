// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"fmt"

	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/key"
	"github.com/drakmaniso/carol/mouse"
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
	fmt.Printf("%v: Key Down: %v %v\n", carol.Now(), l, p)
}

func (loop) KeyUp(l key.Label, p key.Position) {
	fmt.Printf("%v: Key Up: %v %v\n", carol.Now(), l, p)
}

//------------------------------------------------------------------------------

func (loop) MouseMotion(rel pixel.Coord, pos pixel.Coord) {
	fmt.Printf("%v: mouse motion  %+d,%+d  %d,%d\n", carol.Now(), rel.X, rel.Y, pos.X, pos.Y)
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
	fmt.Printf("%v: mouse button down  %s (%v), clicks=%v\n", carol.Now(), n, b, clicks)
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
	fmt.Printf("%v: mouse button up: %s (%v), clicks=%v\n", carol.Now(), n, b, clicks)
}

func (loop) MouseWheel(w pixel.Coord) {
	fmt.Printf("%v: mouse wheel: %+d,%+d\n", carol.Now(), w.X, w.Y)
}

//------------------------------------------------------------------------------

func (loop) WindowShown() {
	fmt.Printf("%v: window shown\n", carol.Now())
}

func (loop) WindowHidden() {
	fmt.Printf("%v: window hidden\n", carol.Now())
}

func (loop) WindowResized(s pixel.Coord) {
	fmt.Printf("%v: window resized %v\n", carol.Now(), s)
}

func (loop) WindowMinimized() {
	fmt.Printf("%v: window minimized\n", carol.Now())
}

func (loop) WindowMaximized() {
	fmt.Printf("%v: window maximized\n", carol.Now())
}

func (loop) WindowRestored() {
	fmt.Printf("%v: window restored\n", carol.Now())
}

func (loop) WindowMouseEnter() {
	fmt.Printf("%v: window mouse enter\n", carol.Now())
}

func (loop) WindowMouseLeave() {
	fmt.Printf("%v: window mouse leave\n", carol.Now())
}

func (loop) WindowFocusGained() {
	fmt.Printf("%v: window focus gained\n", carol.Now())
}

func (loop) WindowFocusLost() {
	fmt.Printf("%v: window focus lost\n", carol.Now())
}

func (loop) WindowQuit() {
	fmt.Printf("%v: window quit\n", carol.Now())
	carol.Stop()
}

//------------------------------------------------------------------------------

func (loop) Update() {
	return //TODO

	d := mouse.Delta()
	fmt.Printf("   mouse.Delta():%+6d,%+6d\v\n", d.X, d.Y)
	p := mouse.Position()
	fmt.Printf("mouse.Position():%6d,%6d\v\n", p.X, p.Y)

	fmt.Print("   mouse buttons: ")
	if mouse.IsPressed(mouse.Left) {
		fmt.Print("\aleft ")
	} else {
		fmt.Print("left ")
	}
	if mouse.IsPressed(mouse.Middle) {
		fmt.Print("\amiddle ")
	} else {
		fmt.Print("middle ")
	}
	if mouse.IsPressed(mouse.Right) {
		fmt.Print("\aright ")
	} else {
		fmt.Print("right ")
	}
	if mouse.IsPressed(mouse.Extra1) {
		fmt.Print("\aextra1 ")
	} else {
		fmt.Print("extra1 ")
	}
	if mouse.IsPressed(mouse.Extra2) {
		fmt.Print("\aextra2\n")
	} else {
		fmt.Print("extra2\n")
	}
}

func (loop) Draw(_, _ float64) {
}

//------------------------------------------------------------------------------
