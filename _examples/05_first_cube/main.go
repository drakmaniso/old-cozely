// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"os"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
	glam.Setup()

	err := setup()
	if err != nil {
		glam.Log("ERROR during setup: %s\n", err)
		return
	}

	window.Handle = handler{}
	mouse.Handle = handler{}

	// Run the Game Loop
	glam.Loop = looper{}
	err = glam.Run()
	if err != nil {
		glam.Log("ERROR: %s\n", err)
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
	transform space.Matrix
}

// Vertex buffer
type mesh []struct {
	position space.Coord `layout:"0"`
	color    color.RGB   `layout:"1"`
}

// Matrices
var (
	model      space.Matrix
	view       space.Matrix
	projection space.Matrix
)

// Cube state
var (
	distance   float32
	position   space.Coord
	yaw, pitch float32
)

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
	v, err := os.Open(glam.Path() + "shader.vert")
	if err != nil {
		return glam.Error("unable to open vertex shader", err)
	}
	f, err := os.Open(glam.Path() + "shader.frag")
	if err != nil {
		return glam.Error("unable to fragment shader", err)
	}
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(v),
		gfx.FragmentShader(f),
		gfx.VertexFormat(0, mesh{}),
		gfx.CullFace(false, false),
		gfx.DepthTest(true),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)

	// Create and fill the vertex buffer
	vbo := gfx.NewVertexBuffer(cube(), 0)

	// Initialize model and view matrices
	position = space.Coord{0, 0, 0}
	yaw = -0.6
	pitch = 0.3
	updateModel()
	distance = 3
	updateView()

	// MTX
	mtx.Color(color.RGB{0.0, 0.05, 0.1}, color.RGB{0.7, 0.6, 0.45})
	mtx.Opaque(false)
	printState()

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	return glam.Error("setup", gfx.Err())
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update(_, _ float64) {
}

func (l looper) Draw(_ float64) {
	pipeline.Bind()
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})

	perFrame.transform = projection.Times(view)
	perFrame.transform = perFrame.transform.Times(model)
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	gfx.Draw(gfx.Triangles, 0, 6*2*3)

	pipeline.Unbind()
}

//------------------------------------------------------------------------------

func updateModel() {
	model = space.Translation(position)
	model = model.Times(space.EulerZXY(pitch, yaw, 0))
}

func updateView() {
	if distance < 1 {
		distance = 1
	}
	view = space.LookAt(space.Coord{0, 0, distance}, space.Coord{0, 0, 0}, space.Coord{0, 1, 0})
}

//------------------------------------------------------------------------------
