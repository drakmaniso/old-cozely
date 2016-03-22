// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"log"
	"strings"
	"time"

	"github.com/drakmaniso/glam/engine"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

type game struct {
	engine.DefaultHandler
}

var pipeline *gfx.Pipeline

//------------------------------------------------------------------------------

func main() {
	g := &game{}
	engine.Handler = g
	key.Handler = g
	mouse.Handler = g
	window.Handler = g

	vs := strings.NewReader(`
		#version 450 core
		void main(void)
		{
			gl_Position = vec4(0.0, 0.0, 0.5, 1.0);
		}
	`)

	fs := strings.NewReader(`
		#version 450 core
		out vec4 color;
		void main(void)
		{
			color = vec4(0.0, 0.8, 1.0, 1.0);
		}	
	`)

	pipeline, _ = gfx.NewPipeline(vs, fs)

	err := engine.Run()
	if err != nil {
		log.Panic(err)
	}
}

//------------------------------------------------------------------------------

func (g *game) KeyDown(l key.Label, p key.Position, time time.Duration) {
	if l == key.LabelEscape {
		engine.Stop()
	}
}

func (g *game) Update() {
}

func (g *game) Draw() {
	pipeline.Use()
}

func (g *game) Quit() {
	pipeline.Close()
	engine.Stop()
}

//------------------------------------------------------------------------------
