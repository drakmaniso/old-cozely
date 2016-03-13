// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine_test

//------------------------------------------------------------------------------

import (
	"fmt"
	"os"
	"testing"

	"github.com/drakmaniso/glam/engine"
	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
)

//------------------------------------------------------------------------------

type game struct {
	engine.DefaultHandler
}

func (g *game) Update() {
	// fmt.Printf("--- Update delta=%v pos=%v rightBttn=%v\n",
	// 	mouse.Delta(), mouse.Position(), mouse.Buttons().IsPressed(mouse.Right))
}

func (g *game) Quit() {
	fmt.Println("*** Bye! ***")
}

func (g *game) KeyDown(l key.Label, p key.Position, time uint32) {
	fmt.Println("*** Key Down: ", l, p, time)
}

func (g *game) KeyUp(l key.Label, p key.Position, time uint32) {
	fmt.Println("*** Key Up: ", l, p, time)
}

func (g *game) MouseMotion(rel geom.IVec2, pos geom.IVec2, b mouse.ButtonState, time uint32) {
	fmt.Println("*** Mouse Motion: ", rel, pos, b, time)
	if b.IsPressed(mouse.Right) {
		fmt.Println("    (right button pressed)")
	}
}

func (g *game) MouseButtonDown(b mouse.Button, clicks int, pos geom.IVec2, time uint32) {
	fmt.Println("*** Mouse Button Down: ", b, clicks, pos, time)
	if b == mouse.Right {
		fmt.Println("    right button pressed...")
	}
}

func (g *game) MouseButtonUp(b mouse.Button, clicks int, pos geom.IVec2, time uint32) {
	fmt.Println("*** Mouse Button Up: ", b, clicks, pos, time)
	if b == mouse.Right {
		fmt.Println("    ...right button released.")
	}
}

func (g *game) MouseWheel(w geom.IVec2, time uint32) {
	fmt.Println("*** Mouse Wheel: ", w, time)
}

func TestMain(m *testing.M) {
	var g game
	engine.HandleLoop(&g)
	engine.HandleKey(&g)
	engine.HandleMouse(&g)
	err = engine.Run()
	os.Exit(m.Run())
}

//------------------------------------------------------------------------------

var err error

func TestEngine_Run(t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

//------------------------------------------------------------------------------
