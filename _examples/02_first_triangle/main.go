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
	g := newGame()

	glam.Loop = g

	// Run the Game Loop
	err := glam.Run()
	if err != nil {
		log.Fatal(err)
	}
}

//------------------------------------------------------------------------------

type game struct {
	pipeline gfx.Pipeline
}

//------------------------------------------------------------------------------

func newGame() *game {
	g := &game{}

	// Setup the Pipeline
	vs, err := gfx.NewVertexShader(vertexShader)
	if err != nil {
		log.Fatal(err)
	}
	fs, err := gfx.NewFragmentShader(fragmentShader)
	if err != nil {
		log.Fatal(err)
	}
	g.pipeline, err = gfx.NewPipeline(vs, fs)
	if err != nil {
		log.Fatal(err)
	}
	g.pipeline.ClearColor(Vec4{0.9, 0.9, 0.9, 1.0})

	return g
}

//------------------------------------------------------------------------------

func (g *game) Update() {
}

func (g *game) Draw() {
	g.pipeline.Bind()
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
