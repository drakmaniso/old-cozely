// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pbr"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/poly"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/window"
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
	window.Handle = handler{}
	mouse.Handle = handler{}
	key.Handle = handler{}

	err = glam.LoopStable(1 / 60.0)
	if err != nil {
		glam.ShowError("running", err)
		return
	}
}

//------------------------------------------------------------------------------

var pipeline *gfx.Pipeline

// Uniform buffer
var miscUBO gfx.UniformBuffer
var misc struct {
	Model          space.Matrix
	SunIlluminance color.RGB
	_              byte
}

// Camera

var camera *pbr.Camera
var forward, lateral, vertical, rolling float32

// Model

var meshes poly.Meshes

var object struct {
	position         space.Coord
	yaw, pitch, roll float32
	scale            float32
}

//------------------------------------------------------------------------------

func setup() error {
	pipeline = gfx.NewPipeline(
		poly.PipelineSetup(),
		gfx.Shader(glam.Path()+"shader.vert"),
		gfx.Shader(glam.Path()+"shader.frag"),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	miscUBO = gfx.NewUniformBuffer(&misc, gfx.DynamicStorage)

	//
	meshes = poly.Meshes{}
	meshes.AddObj(glam.Path() + "../shared/suzanne.obj")
	// meshes.AddObj("E:/objtestfiles/elephant_quads.obj")
	poly.SetupMeshBuffers(meshes)

	// Setup camera

	camera = pbr.NewCamera()
	camera.SetExposure(16.0, 1.0/125.0, 100.0)
	camera.SetPosition(space.Coord{0, 0, 0})

	// Setup model

	object.position = space.Coord{0, 0, -4}
	object.scale = 1.0
	updateModel()

	// Setup light
	misc.SunIlluminance = pbr.DirectionalLightSpectralIlluminance(116400.0, 5400.0)

	// MTX
	mtx.Color(color.RGB{0.0, 0.05, 0.1}, color.RGB{0.7, 0.6, 0.45})
	mtx.Opaque(false)
	mtx.ShowFrameTime(true, -1, 0, false)

	// Bind the vertex buffer to the pipeline
	// pipeline.Bind()

	// pipeline.Bind()
	// vbo.Bind(0, 0)
	// pipeline.Unbind()

	return glam.Error("gfx", gfx.Err())
}

//------------------------------------------------------------------------------

func update(dt64, _ float64) {
	dt := float32(dt64)

	camera.NextState()

	camera.Move(forward*dt, lateral*dt, vertical*dt)

	if firstPerson {
		m := mouse.SmoothDelta()
		s := plane.CoordOf(window.Size())
		camera.Rotate(2*m.X/s.X, 2*m.Y/s.Y, rolling*dt)
	}

	p := camera.Position()
	y, pt, r := camera.Orientation()
	mtx.Print(1, 0, "cam: %6.2f,%6.2f,%6.2f", p.X, p.Y, p.Z)
	mtx.Print(1, 1, "     %6.2f,%6.2f,%6.2f", y, pt, r)
}

//------------------------------------------------------------------------------

func draw() {
	pipeline.Bind()
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.4, 0.45, 0.5, 1.0})

	camera.Bind()
	miscUBO.SubData(&misc, 0)
	miscUBO.Bind(1)

	poly.BindMeshBuffers()

	// poly.Draw()
	gfx.Draw(0, int32(len(meshes.Faces)*6))

	pipeline.Unbind()
}

func updateModel() {
	misc.Model = space.Translation(object.position)
	misc.Model = misc.Model.Times(space.EulerXZY(object.pitch, object.yaw, object.roll))
	misc.Model = misc.Model.Times(space.Scaling(space.Coord{object.scale, object.scale, object.scale}))
	mtx.Print(1, 3, "obj: %6.2f,%6.2f,%6.2f", object.position.X, object.position.Y, object.position.Z)
	mtx.Print(1, 4, "     %6.2f,%6.2f", object.yaw, object.pitch)
}

//------------------------------------------------------------------------------
