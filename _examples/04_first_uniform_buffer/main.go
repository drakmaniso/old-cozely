// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"strings"
	"unsafe"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	. "github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/geom/plane"
	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

var vertexShader = strings.NewReader(`
#version 450 core

layout(location = 0) in vec2 Position;
layout(location = 1) in vec3 Color;

layout(std140, binding = 0) uniform PerObject {
	mat3 Transform;
} obj;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) out vec3 Color;
} vert;

void main(void) {
	vec3 p = obj.Transform * vec3(Position, 1);
	gl_Position = vec4(p.xy, 0.5, 1);
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
	g, err := newGame()
	if err != nil {
		glam.ErrorDialog(err)
		return
	}

	glam.Loop = g

	// Run the Game Loop
	err = glam.Run()
	if err != nil {
		glam.ErrorDialog(err)
	}
}

//------------------------------------------------------------------------------

type game struct {
	pipeline  gfx.Pipeline
	transform gfx.UniformBuffer
	triangle  gfx.VertexBuffer

	angle float32
}

type perVertex struct {
	position Vec2      `layout:"0"`
	color    color.RGB `layout:"1"`
}

type perObject struct {
	transform Mat3x4
}

//------------------------------------------------------------------------------

func newGame() (*game, error) {
	g := &game{}

	// Setup the Pipeline
	g.pipeline = gfx.NewPipeline(
		gfx.VertexShader(vertexShader),
		gfx.FragmentShader(fragmentShader),
		gfx.VertexFormat(0, perVertex{}),
	)
	gfx.Enable(gfx.CullFace)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the Uniform Buffer
	g.transform = gfx.NewUniformBuffer(unsafe.Sizeof(perObject{}), gfx.DynamicStorage)

	// Create the Vertex Buffer
	data := []perVertex{
		{Vec2{0, 0.75}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec2{-0.65, -0.465}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec2{0.65, -0.465}, color.RGB{R: 0, G: 0.6, B: 0.2}},
	}
	g.triangle = gfx.NewVertexBuffer(data, 0)

	return g, gfx.Err()
}

//------------------------------------------------------------------------------

func (g *game) Update() {
	g.angle -= 0.01
}

func (g *game) Draw() {
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	g.pipeline.Bind()
	g.transform.Bind(0)

	m := plane.Rotation(g.angle)
	t := perObject{
		transform: m.Mat3x4(),
	}
	g.transform.Update(&t, 0)

	g.triangle.Bind(0, 0)
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
