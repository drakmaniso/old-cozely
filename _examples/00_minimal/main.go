// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"log"
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
}

func (g *game) Quit() {
	engine.Stop()
}

func (g *game) KeyDown(l key.Label, p key.Position, time time.Duration) {
	if l == key.LabelEscape {
		engine.Stop()
	}
}

//------------------------------------------------------------------------------

func main() {
	g := &game{}
	engine.Handler = g
	key.Handler = g
	mouse.Handler = g
	window.Handler = g
	err := engine.Run()
	if err != nil {
		log.Panic(err)
	}
}

//------------------------------------------------------------------------------
