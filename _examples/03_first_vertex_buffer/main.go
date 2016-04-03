// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"log"
	"strings"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	. "github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

var vertexShader = strings.NewReader(`
#version 450 core

layout(location = 0) in vec2 Position;
layout(location = 1) in vec3 Color;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) out vec3 Color;
} vert;

void main(void) {
	gl_Position = vec4(Position, 0.5, 1);
	vert.Color = Color;
}
`)

var fragmentShader = strings.NewReader(`
#version 450 core

in PerVertex {
	layout(location = 0) in vec3 Color;
} vert;

out vec4 Color;

void main(void) {
	Color = vec4(vert.Color, 1);
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

type perVertex struct {
	position Vec2      `layout:"0"`
	color    color.RGB `layout:"1"`
}

type game struct {
	pipeline gfx.Pipeline
	triangle gfx.Buffer
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
	err = g.pipeline.VertexFormat(0, perVertex{})
	if err != nil {
		log.Fatal(err)
	}
	g.pipeline.ClearColor(Vec4{0.9, 0.9, 0.9, 1.0})

	// Create the Vertex Buffer
	data := []perVertex{
		{Vec2{0, 0.65}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec2{-0.65, -0.475}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec2{0.65, -0.475}, color.RGB{R: 0, G: 0.6, B: 0.2}},
	}
	g.triangle, err = gfx.NewBuffer(data, 0)
	if err != nil {
		log.Fatal(err)
	}

	return g
}

//------------------------------------------------------------------------------

func (g *game) Update() {
}

func (g *game) Draw() {
	g.pipeline.Bind()
	g.pipeline.VertexBuffer(0, g.triangle, 0)
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
