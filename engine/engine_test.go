// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine_test

//------------------------------------------------------------------------------

import (
	"log"
	"os"
	"testing"

	"github.com/drakmaniso/glam/engine"
	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
)

//------------------------------------------------------------------------------

func TestMain(m *testing.M) {
	engine.HandleQuit(func() { log.Print("*** Bye! ***") })
	engine.HandleKeyDown(
		func(l key.Label, p key.Position, time uint32) {
			log.Print("Key Down: ", l, p, time)
		},
	)
	engine.HandleKeyUp(
		func(l key.Label, p key.Position, time uint32) {
			log.Print("Key Up: ", l, p, time)
		},
	)
	engine.HandleMouseMotion(
		func(rel geom.IVec2, pos geom.IVec2, b mouse.ButtonState, time uint32) {
			log.Print("Mouse Motion: ", rel, pos, b, time)
			if b.IsPressed(mouse.Right) {
				log.Println("(RightPressed)")
			}
		},
	)
	engine.HandleMouseButtonDown(
		func(b mouse.Button, clicks int, pos geom.IVec2, time uint32) {
			log.Print("Mouse Button Down: ", b, clicks, pos, time)
			if b == mouse.Right {
				log.Println(">>> Right")
			}
		},
	)
	engine.HandleMouseButtonUp(
		func(b mouse.Button, clicks int, pos geom.IVec2, time uint32) {
			log.Print("Mouse Button Up: ", b, clicks, pos, time)
			if b == mouse.Right {
				log.Println("<<< Right")
			}
		},
	)
	engine.HandleMouseWheel(
		func(w geom.IVec2, time uint32) {
			log.Print("Mouse Wheel: ", w, time)
		},
	)
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
