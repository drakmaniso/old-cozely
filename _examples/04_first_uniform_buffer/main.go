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
	err := glam.Setup()
	if err != nil {
		glam.ShowError("setting up glam", err)
		return
	}

	err = setup()
	if err != nil {
		glam.ShowError("setting up the game", err)
		return
	}

	glam.Update = update
	glam.Draw = draw

	err = glam.Loop()
	if err != nil {
		glam.ShowError("running", err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline    *gfx.Pipeline
	perFrameUBO gfx.UniformBuffer
)

// Uniform buffer
var perFrame struct {
	transform plane.GPUMatrix
}

// Vertex buffer
type mesh []struct {
	position plane.Coord `layout:"0"`
	color    color.RGB   `layout:"1"`
}

// Animation state
var (
	angle float64
)

//------------------------------------------------------------------------------

func setup() error {
	var triangle mesh

	// Create and configure the pipeline
	vs, err := os.Open(glam.Path() + "/shader.vert")
	if err != nil {
		return glam.Error("opening vertex shader", err)
	}
	fs, err := os.Open(glam.Path() + "/shader.frag")
	if err != nil {
		return glam.Error("opening fragment shader", err)
	}
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(vs),
		gfx.FragmentShader(fs),
		gfx.VertexFormat(0, triangle),
		gfx.Topology(gfx.Triangles),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)

	// Fill and create the vertex buffer
	triangle = mesh{
		{plane.Coord{0, 0.75}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{plane.Coord{-0.65, -0.465}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{plane.Coord{0.65, -0.465}, color.RGB{R: 0, G: 0.6, B: 0.2}},
	}
	vbo := gfx.NewVertexBuffer(triangle, gfx.StaticStorage)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	return glam.Error("gfx", gfx.Err())
}

//------------------------------------------------------------------------------

func update(dt, _ float64) {
	angle -= 1.0 * dt
}

func draw() {
	pipeline.Bind()
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})

	perFrame.transform = plane.Rotation(float32(angle)).GPU()
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	gfx.Draw(0, 3)
	pipeline.Unbind()
}

//------------------------------------------------------------------------------
