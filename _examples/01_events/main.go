// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"fmt"
	"time"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
	g := &game{}

	glam.Loop = g
	window.Handle = g
	mouse.Handle = g
	key.Handle = g

	err := glam.Run()
	check(err)
}

//------------------------------------------------------------------------------

type game struct{}

func (g *game) KeyDown(l key.Label, p key.Position, ts time.Duration) {
	if l == key.LabelEscape {
		glam.Stop()
	}
	fmt.Println("*** Key Down: ", l, p, ts)
}

func (g *game) KeyUp(l key.Label, p key.Position, ts time.Duration) {
	fmt.Println("*** Key Up: ", l, p, ts)
}

func (g *game) MouseMotion(rel geom.Vec2, pos geom.Vec2, ts time.Duration) {
	fmt.Println("*** Mouse Motion: ", rel, pos, ts)
}

func (g *game) MouseButtonDown(b mouse.Button, clicks int, ts time.Duration) {
	fmt.Println("*** Mouse Button Down: ", b, clicks, ts)
	if b == mouse.Left {
		fmt.Println("    (Click!)")
	}
}

func (g *game) MouseButtonUp(b mouse.Button, clicks int, ts time.Duration) {
	fmt.Println("*** Mouse Button Up: ", b, clicks, ts)
}

func (g *game) MouseWheel(w geom.Vec2, ts time.Duration) {
	fmt.Println("*** Mouse Wheel: ", w, ts)
}

func (g *game) WindowShown(ts time.Duration) {
	fmt.Println("*** Window Shown: ", ts)
}

func (g *game) WindowHidden(ts time.Duration) {
	fmt.Println("*** Window Hidden: ", ts)
}

func (g *game) WindowResized(s geom.Vec2, ts time.Duration) {
	fmt.Println("*** Window Resized: ", s, ts)
}

func (g *game) WindowMinimized(ts time.Duration) {
	fmt.Println("*** Window Minimized: ", ts)
}

func (g *game) WindowMaximized(ts time.Duration) {
	fmt.Println("*** Window Maximized: ", ts)
}

func (g *game) WindowRestored(ts time.Duration) {
	fmt.Println("*** Window Restored: ", ts)
}

func (g *game) WindowMouseEnter(ts time.Duration) {
	fmt.Println("*** Window Mouse Enter: ", ts)
}

func (g *game) WindowMouseLeave(ts time.Duration) {
	fmt.Println("*** Window Mouse Leave: ", ts)
}

func (g *game) WindowFocusGained(ts time.Duration) {
	fmt.Println("*** Window Focus Gained: ", ts)
}

func (g *game) WindowFocusLost(ts time.Duration) {
	fmt.Println("*** Window Focus Lost: ", ts)
}

func (g *game) WindowQuit(ts time.Duration) {
	fmt.Println("*** Window Quit ***")
	glam.Stop()
}

func (g *game) Draw() {}

func (g *game) Update() {
	// fmt.Printf("--- Update delta=%v pos=%v rightBttn=%v\n",
	// 	mouse.Delta(), mouse.Position(), mouse.IsPressed(mouse.Right))
	// fmt.Printf("--- w = %v\n", key.IsPressed(key.PositionW))
	// fmt.Printf("--- window.HasFocus = %v\n", window.HasFocus())
	// fmt.Printf("--- window.HasMouseFocus = %v\n", window.HasMouseFocus())
}

//------------------------------------------------------------------------------

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//------------------------------------------------------------------------------
