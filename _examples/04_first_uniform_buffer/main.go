// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"strings"
	"unsafe"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/plane"
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
	err := setup()
	if err != nil {
		glam.ErrorDialog(err)
		return
	}

	// Run the main loop
	glam.Loop = looper{}
	err = glam.Run()
	if err != nil {
		glam.ErrorDialog(err)
	}
}

//------------------------------------------------------------------------------

// Vertex buffer layout
type perVertex struct {
	position plane.Coord `layout:"0"`
	color    color.RGB   `layout:"1"`
}

// Uniform buffer
type perObject struct {
	transform plane.GPUMatrix
}

// OpenGL objects
var (
	pipeline  gfx.Pipeline
	transform gfx.UniformBuffer
	triangle  gfx.VertexBuffer
)

// Animation state
var (
	angle float32
)

//------------------------------------------------------------------------------

func setup() error {
	// Setup the pipeline
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(vertexShader),
		gfx.FragmentShader(fragmentShader),
		gfx.VertexFormat(0, perVertex{}),
	)
	gfx.Enable(gfx.CullFace)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	transform = gfx.NewUniformBuffer(unsafe.Sizeof(perObject{}), gfx.DynamicStorage)

	// Create the vertex buffer
	data := []perVertex{
		{plane.Coord{0, 0.75}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{plane.Coord{-0.65, -0.465}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{plane.Coord{0.65, -0.465}, color.RGB{R: 0, G: 0.6, B: 0.2}},
	}
	triangle = gfx.NewVertexBuffer(data, 0)

	return gfx.Err()
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update() {
	angle -= 0.01
}

func (l looper) Draw() {
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	pipeline.Bind()
	transform.Bind(0)

	m := plane.Rotation(angle)
	t := perObject{
		transform: m.GPU(),
	}
	transform.Update(&t, 0)

	triangle.Bind(0, 0)
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
