// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine_test

//------------------------------------------------------------------------------

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/drakmaniso/glam/engine"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

type game struct {
	engine.DefaultHandler
}

func (g *game) Update() {
	// fmt.Printf("--- Update delta=%v pos=%v rightBttn=%v\n",
	// 	mouse.Delta(), mouse.Position(), mouse.IsPressed(mouse.Right))
	// fmt.Printf("--- w = %v\n", key.IsPressed(key.PositionW))
	// fmt.Printf("--- window.HasFocus = %v\n", window.HasFocus())
	// fmt.Printf("--- window.HasMouseFocus = %v\n", window.HasMouseFocus())
}

func (g *game) Quit() {
	fmt.Println("*** Bye! ***")
	engine.Stop()
}

func (g *game) KeyDown(l key.Label, p key.Position, time time.Duration) {
	if l == key.LabelEscape {
		engine.Stop()
	}
	fmt.Println("*** Key Down: ", l, p, time)
}

// func (g *game) KeyUp(l key.Label, p key.Position, time time.Duration) {
// 	fmt.Println("*** Key Up: ", l, p, time)
// }

// func (g *game) MouseMotion(rel geom.IVec2, pos geom.IVec2, time time.Duration) {
// 	fmt.Println("*** Mouse Motion: ", rel, pos, time)
// 	if mouse.IsPressed(mouse.Right) {
// 		fmt.Println("    (right button pressed)")
// 	}
// }

// func (g *game) MouseButtonDown(b mouse.Button, clicks int, time time.Duration) {
// 	fmt.Println("*** Mouse Button Down: ", b, clicks, time)
// 	if b == mouse.Left {
// 		mouse.SetRelativeMode(true)
// 		fmt.Println("    Relative Mode ON")
// 	}
// 	if b == mouse.Right {
// 		mouse.SetRelativeMode(false)
// 		fmt.Println("    Relative Mode OFF")
// 	}
// }

// func (g *game) MouseButtonUp(b mouse.Button, clicks int, time time.Duration) {
// 	fmt.Println("*** Mouse Button Up: ", b, clicks, time)
// }

// func (g *game) MouseWheel(w geom.IVec2, time time.Duration) {
// 	fmt.Println("*** Mouse Wheel: ", w, time)
// }

func (g *game) WindowShown(ts time.Duration) {
	fmt.Println("*** Window Shown: ", ts)
}

func (g *game) WindowHidden(ts time.Duration) {
	fmt.Println("*** Window Hidden: ", ts)
}

// func (g *game) WindowResized(s geom.IVec2, ts time.Duration) {
// 	fmt.Println("*** Window Resized: ", s, ts)
// }

// func (g *game) WindowMinimized(ts time.Duration) {
// 	fmt.Println("*** Window Minimized: ", ts)
// }

// func (g *game) WindowMaximized(ts time.Duration) {
// 	fmt.Println("*** Window Maximized: ", ts)
// }

// func (g *game) WindowRestored(ts time.Duration) {
// 	fmt.Println("*** Window Restored: ", ts)
// }

// func (g *game) WindowMouseEnter(ts time.Duration) {
// 	fmt.Println("*** Window Mouse Enter: ", ts)
// }

// func (g *game) WindowMouseLeave(ts time.Duration) {
// 	fmt.Println("*** Window Mouse Leave: ", ts)
// }

// func (g *game) WindowFocusGained(ts time.Duration) {
// 	fmt.Println("*** Window Focus Gained: ", ts)
// }

// func (g *game) WindowFocusLost(ts time.Duration) {
// 	fmt.Println("*** Window Focus Lost: ", ts)
// }

func TestMain(m *testing.M) {
	var g game
	engine.Handler = &g
	key.Handler = &g
	mouse.Handler = &g
	window.Handler = &g
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
