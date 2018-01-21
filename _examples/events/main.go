// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"fmt"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
)

//------------------------------------------------------------------------------

func main() {
	err := glam.Run(nil, loop{})
	if err != nil {
		glam.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (loop) KeyDown(l key.Label, p key.Position) {
	if l == key.LabelEscape {
		glam.Stop()
	}
	fmt.Printf("%v: Key Down: %v %v\n", glam.GameTime(), l, p)
}

func (loop) KeyUp(l key.Label, p key.Position) {
	fmt.Printf("%v: Key Up: %v %v\n", glam.GameTime(), l, p)
}

//------------------------------------------------------------------------------

func (loop) MouseMotion(dx, dy int32, x, y int32) {
	fmt.Printf("%v: mouse motion  %+d,%+d  %d,%d\n", glam.GameTime(), dx, dy, x, y)
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
	fmt.Printf("%v: mouse button down  %s (%v), clicks=%v\n", glam.GameTime(), n, b, clicks)
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
	fmt.Printf("%v: mouse button up: %s (%v), clicks=%v\n", glam.GameTime(), n, b, clicks)
}

func (loop) MouseWheel(dx, dy int32) {
	fmt.Printf("%v: mouse wheel: %+d,%+d\n", glam.GameTime(), dx, dy)
}

//------------------------------------------------------------------------------

func (loop) WindowShown() {
	fmt.Printf("%v: window shown\n", glam.GameTime())
}

func (loop) WindowHidden() {
	fmt.Printf("%v: window hidden\n", glam.GameTime())
}

func (loop) WindowResized(w, h int32) {
	fmt.Printf("%v: window resized %dx%d\n", glam.GameTime(), w, h)
}

func (loop) WindowMinimized() {
	fmt.Printf("%v: window minimized\n", glam.GameTime())
}

func (loop) WindowMaximized() {
	fmt.Printf("%v: window maximized\n", glam.GameTime())
}

func (loop) WindowRestored() {
	fmt.Printf("%v: window restored\n", glam.GameTime())
}

func (loop) WindowMouseEnter() {
	fmt.Printf("%v: window mouse enter\n", glam.GameTime())
}

func (loop) WindowMouseLeave() {
	fmt.Printf("%v: window mouse leave\n", glam.GameTime())
}

func (loop) WindowFocusGained() {
	fmt.Printf("%v: window focus gained\n", glam.GameTime())
}

func (loop) WindowFocusLost() {
	fmt.Printf("%v: window focus lost\n", glam.GameTime())
}

func (loop) WindowQuit() {
	fmt.Printf("%v: window quit\n", glam.GameTime())
	glam.Stop()
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil //TODO

	dx, dy := mouse.Delta()
	fmt.Printf("   mouse.Delta():%+6d,%+6d\v\n", dx, dy)
	px, py := mouse.Position()
	fmt.Printf("mouse.Position():%6d,%6d\v\n", px, py)

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

	return nil
}

func (loop) Draw() error {
	return nil
}

//------------------------------------------------------------------------------
