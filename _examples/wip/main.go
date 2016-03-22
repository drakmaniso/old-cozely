// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"log"
	"strings"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

var pipeline *gfx.Pipeline

//------------------------------------------------------------------------------

func main() {
	g := &game{}
	glam.Handler = g

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

	if err := glam.Run(); err != nil {
		log.Print(err)
	}
}

//------------------------------------------------------------------------------

type game struct{}

func (g *game) Update() {
}

func (g *game) Draw() {
	pipeline.Use()
}

//------------------------------------------------------------------------------
