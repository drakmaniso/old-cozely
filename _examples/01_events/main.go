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
	window.Handle = handler{}
	mouse.Handle = handler{}
	key.Handle = handler{}

	// Run the main loop
	glam.Loop = looper{}
	err := glam.Run()
	if err != nil {
		glam.ErrorDialog(err)
	}
}

//------------------------------------------------------------------------------

type handler struct{}

func (h handler) KeyDown(l key.Label, p key.Position, ts time.Duration) {
	if l == key.LabelEscape {
		glam.Stop()
	}
	fmt.Println("*** Key Down: ", l, p, ts)
}

func (h handler) KeyUp(l key.Label, p key.Position, ts time.Duration) {
	fmt.Println("*** Key Up: ", l, p, ts)
}

func (h handler) MouseMotion(rel geom.Vec2, pos geom.Vec2, ts time.Duration) {
	fmt.Println("*** Mouse Motion: ", rel, pos, ts)
}

func (h handler) MouseButtonDown(b mouse.Button, clicks int, ts time.Duration) {
	fmt.Println("*** Mouse Button Down: ", b, clicks, ts)
	if b == mouse.Left {
		fmt.Println("    (Click!)")
	}
}

func (h handler) MouseButtonUp(b mouse.Button, clicks int, ts time.Duration) {
	fmt.Println("*** Mouse Button Up: ", b, clicks, ts)
}

func (h handler) MouseWheel(w geom.Vec2, ts time.Duration) {
	fmt.Println("*** Mouse Wheel: ", w, ts)
}

func (h handler) WindowShown(ts time.Duration) {
	fmt.Println("*** Window Shown: ", ts)
}

func (h handler) WindowHidden(ts time.Duration) {
	fmt.Println("*** Window Hidden: ", ts)
}

func (h handler) WindowResized(s geom.Vec2, ts time.Duration) {
	fmt.Println("*** Window Resized: ", s, ts)
}

func (h handler) WindowMinimized(ts time.Duration) {
	fmt.Println("*** Window Minimized: ", ts)
}

func (h handler) WindowMaximized(ts time.Duration) {
	fmt.Println("*** Window Maximized: ", ts)
}

func (h handler) WindowRestored(ts time.Duration) {
	fmt.Println("*** Window Restored: ", ts)
}

func (h handler) WindowMouseEnter(ts time.Duration) {
	fmt.Println("*** Window Mouse Enter: ", ts)
}

func (h handler) WindowMouseLeave(ts time.Duration) {
	fmt.Println("*** Window Mouse Leave: ", ts)
}

func (h handler) WindowFocusGained(ts time.Duration) {
	fmt.Println("*** Window Focus Gained: ", ts)
}

func (h handler) WindowFocusLost(ts time.Duration) {
	fmt.Println("*** Window Focus Lost: ", ts)
}

func (h handler) WindowQuit(ts time.Duration) {
	fmt.Println("*** Window Quit ***")
	glam.Stop()
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Draw() {}

func (l looper) Update() {
	// fmt.Printf("--- Update delta=%v pos=%v rightBttn=%v\n",
	// 	mouse.Delta(), mouse.Position(), mouse.IsPressed(mouse.Right))
	// fmt.Printf("--- w = %v\n", key.IsPressed(key.PositionW))
	// fmt.Printf("--- window.HasFocus = %v\n", window.HasFocus())
	// fmt.Printf("--- window.HasMouseFocus = %v\n", window.HasMouseFocus())
}

//------------------------------------------------------------------------------
