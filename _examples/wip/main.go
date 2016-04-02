// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"log"
	"strings"
	"time"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	. "github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/geom/space"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

var pipeline gfx.Pipeline

type perVertex struct {
	position Vec3      `layout:"0"`
	color    color.RGB `layout:"1"`
}

type perObject struct {
	transform Mat4
}

var projection Mat4

var transform gfx.Buffer
var colorfulTriangle gfx.Buffer

//------------------------------------------------------------------------------

var vertexShader = strings.NewReader(`
#version 450 core

layout(location = 0) in vec3 Position;
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
	gl_Position = obj.Transform * vec4(Position, 1);
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
	window.Handler = g

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
		// Front Face
		{Vec3{0, 0, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 1, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{0, 1, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{0, 0, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 0, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 1, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		// Back Face
		{Vec3{0, 0, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{0, 1, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 1, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{0, 0, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 1, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 0, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		// Right Face
		{Vec3{1, 0, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{1, 1, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{1, 1, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{1, 0, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{1, 0, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{1, 1, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		// Left Face
		{Vec3{0, 0, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{0, 1, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{0, 1, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{0, 0, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{0, 1, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{0, 0, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		// Bottom Face
		{Vec3{0, 0, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{0, 0, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 0, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{0, 0, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 0, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 0, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		// Top Face
		{Vec3{0, 1, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 1, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{0, 1, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{0, 1, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 1, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 1, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
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

type game struct {
	window.DefaultHandler
}

//------------------------------------------------------------------------------

func (g *game) WindowResized(s IVec2, timestamp time.Duration) {
	r := float32(s.X) / float32(s.Y)
	projection = space.Perspective(0.535, r, 0.001, 1000.0)
}

//------------------------------------------------------------------------------

var angle float32

func (g *game) Update() {
	angle += 0.01
}

func (g *game) Draw() {
	pipeline.Bind()
	pipeline.UniformBuffer(0, transform)

	m := projection.Times(space.LookAt(Vec3{0, 0, 5}, Vec3{0, 0, 0}, Vec3{0, 1, 0}))
	m = m.Times(space.Rotation(angle, Vec3{1, -0.5, 0.25}.Normalized()))
	m = m.Times(space.Translation(Vec3{-0.5, -0.5, -0.5}))
	t := perObject{
		transform: m,
	}
	transform.Update(&t, 0)

	pipeline.VertexBuffer(0, colorfulTriangle, 0)
	gfx.Draw(gfx.Triangles, 0, 6*2*3)
}

//------------------------------------------------------------------------------
