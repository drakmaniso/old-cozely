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
	projection space.Matrix
	view       space.Matrix
	model      space.Matrix
}

// Camera
var (
	distance   float32
	position   space.Coord
	yaw, pitch float32
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
	poly.SetupMeshBuffers(meshes)

	// Initialize view matrix
	position = space.Coord{0, 0, 0}
	// yaw = -0.6
	pitch = 0.3
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

var meshes = poly.Meshes{
	Faces: []poly.Face{
		// poly.Face{0, [3]uint16{0, 1, 2}},
		// poly.Face{1, [3]uint16{3, 4, 5}},
		// Front
		poly.Face{0, [3]uint16{5, 0, 4}},
		poly.Face{0, [3]uint16{5, 1, 0}},
		// Top
		poly.Face{1, [3]uint16{3, 5, 7}},
		poly.Face{1, [3]uint16{3, 1, 5}},
		// Bottom
		poly.Face{2, [3]uint16{6, 0, 2}},
		poly.Face{2, [3]uint16{6, 4, 0}},
		// Back
		poly.Face{3, [3]uint16{3, 6, 2}},
		poly.Face{3, [3]uint16{3, 7, 6}},
		// Left
		poly.Face{4, [3]uint16{7, 4, 6}},
		poly.Face{4, [3]uint16{7, 5, 4}},
		// Right
		poly.Face{5, [3]uint16{1, 2, 0}},
		poly.Face{5, [3]uint16{1, 3, 2}},
	},
	Vertices: []space.Coord{
		// space.Coord{0, 0.65, 0.5},
		// space.Coord{-0.65, -0.475, 0.5},
		// space.Coord{0.65, -0.475, 0.5},
		// space.Coord{0, -0.65, 0.6},
		// space.Coord{0.65, +0.475, 0.6},
		// space.Coord{-0.65, +0.475, 0.6},

		space.Coord{1.000000, 1.000000, -1.000000},
		space.Coord{1.000000, 1.000000, 1.000000},
		space.Coord{1.000000, -1.000000, -1.000000},
		space.Coord{1.000000, -1.000000, 1.000000},
		space.Coord{-1.000000, 1.000000, -1.000000},
		space.Coord{-1.000000, 1.000000, 1.000000},
		space.Coord{-1.000000, -1.000000, -1.000000},
		space.Coord{-1.000000, -1.000000, 1.000000},
	},
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update(_, _ float64) {
}

func (l looper) Draw(_ float64) {
	poly.BindPipeline()
	gfx.ClearColorBuffer(color.RGBA{0.8, 0.8, 0.8, 1.0})

	frameUBO.SubData(&frame, 0)
	frameUBO.Bind(0)

	poly.BindMeshBuffers()

	// poly.Draw()
	gfx.Draw(gfx.Triangles, 0, 6*2*3)

	poly.UnbindPipeline()
}

func updateModel() {
	frame.model = space.Translation(position)
	frame.model = frame.model.Times(space.EulerZXY(pitch, yaw, 0))
}

func updateView() {
	if distance < 1 {
		distance = 1
	}
	frame.view = space.LookAt(
		space.Coord{0, 0, distance},
		space.Coord{0, 0, 0},
		space.Coord{0, 1, 0},
	)
}

//------------------------------------------------------------------------------
