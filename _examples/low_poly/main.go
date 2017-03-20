// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
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
	"github.com/drakmaniso/glam/poly"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
	glam.Setup()

	err := setup()
	if err != nil {
		glam.ErrorDialog(err)
		return
	}

	window.Handle = handler{}
	mouse.Handle = handler{}

	// Run the Game Loop
	glam.Loop = looper{}
	err = glam.Run()
	if err != nil {
		glam.ErrorDialog(err)
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	frameUBO gfx.UniformBuffer
)

// Uniform buffer
var frame struct {
	ProjectionView space.Matrix
	Model          space.Matrix
}

var meshes poly.Meshes

// Camera
var (
	distance   float32
	position   space.Coord
	yaw, pitch float32
)

var (
	projection space.Matrix
	view       space.Matrix
)

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
	f, err := os.Open(glam.Path() + "shader.frag")
	if err != nil {
		return err
	}
	poly.SetupPipeline(
		gfx.FragmentShader(f),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	frameUBO = gfx.NewUniformBuffer(&frame, gfx.DynamicStorage)

	//
	meshes = poly.Meshes{}
	meshes.AddObj(glam.Path() + "../shared/suzanne.obj")
	// meshes.AddObj("E:/objtestfiles/scifigirl.obj")
	poly.SetupMeshBuffers(meshes)

	// Initialize view matrix
	position = space.Coord{0, 0, 0}
	yaw = 0.3
	pitch = 0.2
	updateModel()
	distance = 6
	updateView()

	// MTX
	mtx.Color(color.RGB{0.0, 0.05, 0.1}, color.RGB{0.7, 0.6, 0.45})
	mtx.Opaque(false)

	// Bind the vertex buffer to the pipeline
	poly.BindPipeline()

	// pipeline.Bind()
	// vbo.Bind(0, 0)
	// pipeline.Unbind()

	return gfx.Err()
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update(_, _ float64) {
}

func (l looper) Draw(_ float64) {
	poly.BindPipeline()
	gfx.ClearColorBuffer(color.RGBA{0.8, 0.8, 0.8, 1.0})

	frame.ProjectionView = projection.Times(view)
	frameUBO.SubData(&frame, 0)
	frameUBO.Bind(0)

	poly.BindMeshBuffers()

	// poly.Draw()
	gfx.Draw(gfx.Triangles, 0, int32(len(meshes.Faces)*6))

	poly.UnbindPipeline()
}

func updateModel() {
	frame.Model = space.Translation(position)
	frame.Model = frame.Model.Times(space.EulerZXY(pitch, yaw, 0))
	// frame.Model = frame.Model.Times(space.Scaling(space.Coord{2, 2, 2}))
}

func updateView() {
	if distance < 1 {
		distance = 1
	}
	view = space.LookAt(
		space.Coord{0, 0, distance},
		space.Coord{0, 0, 0},
		space.Coord{0, 1, 0},
	)
}

//------------------------------------------------------------------------------
