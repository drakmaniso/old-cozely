// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"log"
	"strings"

	"github.com/drakmaniso/glam"
	. "github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/geom/space"
	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

var pipeline gfx.Pipeline

type perVertex struct {
	position Vec2 `layout:"0"`
	color    Vec3 `layout:"1"`
}

type perObject struct {
	transform Mat4
}

var transform gfx.Buffer
var colorfulTriangle gfx.Buffer

//------------------------------------------------------------------------------

var vertexShader = strings.NewReader(`
#version 450 core

layout(location = 0) in vec2 Position;
layout(location = 1) in vec3 Color;

layout(std140, binding = 0) uniform PerObject {
	mat4 Transform;
} obj;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) out vec3 Color;
} vert;

void main(void) {
	gl_Position = obj.Transform * vec4(Position, 0.5, 1);
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
	log.SetFlags(log.Lshortfile)

	g := &game{}
	glam.Handler = g

	var err error

	// Setup the Pipeline
	vs, err := gfx.NewVertexShader(vertexShader)
	if err != nil {
		log.Fatal(err)
	}
	fs, err := gfx.NewFragmentShader(fragmentShader)
	if err != nil {
		log.Fatal(err)
	}
	pipeline, err = gfx.NewPipeline(vs, fs)
	if err != nil {
		log.Fatal(err)
	}
	err = pipeline.VertexFormat(0, perVertex{})
	if err != nil {
		log.Fatal(err)
	}
	pipeline.ClearColor(Vec4{0.9, 0.9, 0.9, 1.0})

	// Create the Uniform Buffer
	transform, err = gfx.NewBuffer(uintptr(64), gfx.DynamicStorage)
	if err != nil {
		log.Fatal(err)
	}

	// Create the Vertex Buffer
	data := []perVertex{
		{Vec2{0, 0.75}, Vec3{0.3, 0, 0.8}},
		{Vec2{-0.65, -0.465}, Vec3{0.8, 0.3, 0}},
		{Vec2{0.65, -0.465}, Vec3{0, 0.6, 0.2}},
	}
	colorfulTriangle, err = gfx.NewBuffer(data, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Run the Game Loop
	err = glam.Run()
	if err != nil {
		log.Fatal(err)
	}
}

//------------------------------------------------------------------------------

type game struct{}

var angle float32

func (g *game) Update() {
	angle += 0.01
}

func (g *game) Draw() {
	pipeline.Bind()
	pipeline.UniformBuffer(0, transform)

	t := perObject{
		transform: space.Rotation(angle, Vec3{0, 0, 1}),
	}
	transform.Update(&t, 0)

	pipeline.VertexBuffer(0, colorfulTriangle, 0)
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
