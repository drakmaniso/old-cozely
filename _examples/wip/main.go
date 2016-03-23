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

var pipeline gfx.Pipeline

//------------------------------------------------------------------------------

func main() {
	g := &game{}
	glam.Handler = g

	vs := strings.NewReader(`
		#version 420 core
		void main(void)
		{
			const float Pi = 3.14;
			const float r = 0.75;
			const vec4 v[3] = vec4[3](
				vec4(r*sin(0),       r*cos(0),       0.5, 1.0),
				vec4(r*sin(-Pi*2/3), r*cos(-Pi*2/3), 0.5, 1.0),
				vec4(r*sin(-Pi*4/3), r*cos(-Pi*4/3), 0.5, 1.0)
			);
			gl_Position = v[gl_VertexID];
		}	
	`)

	fs := strings.NewReader(`
		#version 420 core
		out vec4 color;
		void main(void)
		{
			color = vec4(0.84, 0.00, 0.44, 1.0);
		}	
	`)

	_ = pipeline.CompileShaders(vs, fs)

	if err := glam.Run(); err != nil {
		log.Print(err)
	}
}

//------------------------------------------------------------------------------

type game struct{}

func (g *game) Update() {
}

func (g *game) Draw() {
	pipeline.Bind()
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
