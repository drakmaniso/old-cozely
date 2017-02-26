// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"os"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/plane"
)

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

// Uniform buffer
var perFrame struct {
	transform plane.GPUMatrix
}

// Vertex buffer
type vertex struct {
	position plane.Coord `layout:"0"`
	color    color.RGB   `layout:"1"`
}

// OpenGL objects
var (
	pipeline    gfx.Pipeline
	perFrameUBO gfx.UniformBuffer
	triangleVBO gfx.VertexBuffer
)

// Animation state
var (
	angle float32
)

//------------------------------------------------------------------------------

func setup() error {
	// Setup the pipeline
	vs, err := os.Open(glam.Path() + "/shader.vert")
	if err != nil {
		return err
	}
	fs, err := os.Open(glam.Path() + "/shader.frag")
	if err != nil {
		return err
	}
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(vs),
		gfx.FragmentShader(fs),
		gfx.VertexFormat(0, vertex{}),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)

	// Fill and create the vertex buffer
	var triangle = []vertex{
		{plane.Coord{0, 0.75}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{plane.Coord{-0.65, -0.465}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{plane.Coord{0.65, -0.465}, color.RGB{R: 0, G: 0.6, B: 0.2}},
	}
	triangleVBO = gfx.NewVertexBuffer(triangle, gfx.StaticStorage)

	return gfx.Err()
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update() {
	angle -= 0.01
}

func (l looper) Draw() {
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	pipeline.Bind()

	perFrame.transform = plane.Rotation(angle).GPU()
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	triangleVBO.Bind(0, 0)
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
