// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"log"
	"time"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

type game struct {
	glam.DefaultHandler
}

//------------------------------------------------------------------------------

func main() {
	g := &game{}
	glam.Handler = g
	key.Handler = g
	mouse.Handler = g
	window.Handler = g
	err := glam.Run()
	if err != nil {
		log.Panic(err)
	}
}

//------------------------------------------------------------------------------

func (g *game) Update() {
}

func (g *game) WindowQuit(ts time.Duration) {
	glam.Stop()
}

func (g *game) KeyDown(l key.Label, p key.Position, ts time.Duration) {
	if l == key.LabelEscape {
		glam.Stop()
	}
}

//------------------------------------------------------------------------------
