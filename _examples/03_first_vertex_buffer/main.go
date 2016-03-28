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

type vertex struct {
	position Vec2 `layout:"0"`
	color    Vec3 `layout:"1"`
}

var colorfulTriangle gfx.Buffer

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
	g := &game{}
	glam.Handler = g

	// Setup the Pipeline
	var vs, fs gfx.Shader
	vs.Create(gfx.VertexShader, vertexShader)
	fs.Create(gfx.FragmentShader, fragmentShader)
	if err := pipeline.Create(&vs, &fs); err != nil {
		log.Fatal(err)
	}
	if err := pipeline.VertexFormat(0, vertex{}); err != nil {
		log.Fatal(err)
	}
	pipeline.ClearColor(Vec4{0.9, 0.9, 0.9, 1.0})

	// Create the Vertex Buffer
	data := []vertex{
		{Vec2{0, 0.65}, Vec3{0.3, 0, 0.8}},
		{Vec2{-0.65, -0.475}, Vec3{0.8, 0.3, 0}},
		{Vec2{0.65, -0.475}, Vec3{0, 0.6, 0.2}},
	}
	if err := colorfulTriangle.Create(data, 0); err != nil {
		log.Fatal(err)
	}

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
	pipeline.VertexBuffer(0, &colorfulTriangle, 0)
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
