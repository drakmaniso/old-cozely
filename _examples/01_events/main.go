// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"fmt"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
	glam.Setup()

	window.Handle = handler{}
	mouse.Handle = handler{}
	key.Handle = handler{}

	// Run the main loop
	glam.Loop = looper{}
	err := glam.Run()
	if err != nil {
		glam.Log("ERROR: %s\n", err)
	}
}

//------------------------------------------------------------------------------

type handler struct{}

func (h handler) KeyDown(l key.Label, p key.Position, ts uint32) {
	if l == key.LabelEscape {
		glam.Stop()
	}
	fmt.Println("*** Key Down: ", l, p, ts)
}

func (h handler) KeyUp(l key.Label, p key.Position, ts uint32) {
	fmt.Println("*** Key Up: ", l, p, ts)
}

func (h handler) MouseMotion(rel pixel.Coord, pos pixel.Coord, ts uint32) {
	fmt.Println("*** Mouse Motion: ", rel, pos, ts)
}

func (h handler) MouseButtonDown(b mouse.Button, clicks int, ts uint32) {
	fmt.Println("*** Mouse Button Down: ", b, clicks, ts)
	if b == mouse.Left {
		fmt.Println("    (Click!)")
	}
}

func (h handler) MouseButtonUp(b mouse.Button, clicks int, ts uint32) {
	fmt.Println("*** Mouse Button Up: ", b, clicks, ts)
}

func (h handler) MouseWheel(w pixel.Coord, ts uint32) {
	fmt.Println("*** Mouse Wheel: ", w, ts)
}

func (h handler) WindowShown(ts uint32) {
	fmt.Println("*** Window Shown: ", ts)
}

func (h handler) WindowHidden(ts uint32) {
	fmt.Println("*** Window Hidden: ", ts)
}

func (h handler) WindowResized(s pixel.Coord, ts uint32) {
	fmt.Println("*** Window Resized: ", s, ts)
}

func (h handler) WindowMinimized(ts uint32) {
	fmt.Println("*** Window Minimized: ", ts)
}

func (h handler) WindowMaximized(ts uint32) {
	fmt.Println("*** Window Maximized: ", ts)
}

func (h handler) WindowRestored(ts uint32) {
	fmt.Println("*** Window Restored: ", ts)
}

func (h handler) WindowMouseEnter(ts uint32) {
	fmt.Println("*** Window Mouse Enter: ", ts)
}

func (h handler) WindowMouseLeave(ts uint32) {
	fmt.Println("*** Window Mouse Leave: ", ts)
}

func (h handler) WindowFocusGained(ts uint32) {
	fmt.Println("*** Window Focus Gained: ", ts)
}

func (h handler) WindowFocusLost(ts uint32) {
	fmt.Println("*** Window Focus Lost: ", ts)
}

func (h handler) WindowQuit(ts uint32) {
	fmt.Println("*** Window Quit ***")
	glam.Stop()
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update(_, _ float64) {
	// fmt.Printf("--- Update delta=%v pos=%v rightBttn=%v\n",
	// 	mouse.Delta(), mouse.Position(), mouse.IsPressed(mouse.Right))
	// fmt.Printf("--- w = %v\n", key.IsPressed(key.PositionW))
	// fmt.Printf("--- window.HasFocus = %v\n", window.HasFocus())
	// fmt.Printf("--- window.HasMouseFocus = %v\n", window.HasMouseFocus())
}

func (l looper) Draw(_ float64) {
	mtx.Print(1, 0, "   mouse.Delta()=%+6d, %+6d", mouse.Delta().X, mouse.Delta().Y)
	mtx.Print(1, 1, "mouse.Position()=%v", mouse.Position())
}

//------------------------------------------------------------------------------
