// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"log"
	"strings"

	"github.com/drakmaniso/glam"
	. "github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

var pipeline gfx.Pipeline

//------------------------------------------------------------------------------

var vertexShader = strings.NewReader(`
#version 450 core

out gl_PerVertex {
	vec4 gl_Position;
};

void main(void)
{
	const vec4 triangle[3] = vec4[3](
		vec4(0, 0.65, 0.5, 1),
		vec4(-0.65, -0.475, 0.5, 1),
		vec4(0.65, -0.475, 0.5, 1)
	);
	gl_Position = triangle[gl_VertexID];
}	
`)

var fragmentShader = strings.NewReader(`
#version 450 core

out vec4 color;

void main(void)
{
	color = vec4(0.3, 0.1, 0.6, 1.0);
}	
`)

//------------------------------------------------------------------------------

func main() {
	g := &game{}
	glam.Handler = g

	// Setup the Pipeline
	var vs, fs gfx.Shader
	vs.Create(gfx.VertexShader, vertexShader)
	fs.Create(gfx.FragmentShader, fragmentShader)
	if err := pipeline.Create(&vs, &fs); err != nil {
		log.Fatal(err)
	}
	pipeline.ClearColor(Vec4{0.9, 0.9, 0.9, 1.0})

	// Run the Game Loop
	if err := glam.Run(); err != nil {
		log.Fatal(err)
	}
}

//------------------------------------------------------------------------------

type game struct{}

func (g *game) Update() {
}

func (g *game) Draw() {
	pipeline.Use()
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
