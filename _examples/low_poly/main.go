// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
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

	glam.Loop(loop{})
	glam.SetTimeStep(1 / 2.0)

	err = glam.Run()
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
	worldFromObject space.Matrix
	SunIlluminance  color.RGB
	_               byte
}

// PlanarCamera

var camera *pbr.PlanarCamera

// State

var forward, lateral, vertical, rolling float32
var dragStart space.Matrix

var current struct {
	dragDelta plane.Coord
}

var previous struct {
	dragDelta plane.Coord
}

// worldFromObject

var meshes poly.Meshes

//------------------------------------------------------------------------------

func setup() error {
	pipeline = gfx.NewPipeline(
		poly.PipelineSetup(),
		pbr.ToneMapACES(),
		gfx.Shader(glam.Path()+"shader.vert"),
		gfx.Shader(glam.Path()+"shader.frag"),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	miscUBO = gfx.NewUniformBuffer(&misc, gfx.DynamicStorage)

	//
	meshes = poly.Meshes{}
	// meshes.AddObj(glam.Path() + "../shared/cube.obj")
	// meshes.AddObj(glam.Path() + "../shared/teapot.obj")
	meshes.AddObj(glam.Path() + "../shared/suzanne.obj")
	// meshes.AddObj("E:/objtestfiles/triceratops.obj")
	poly.SetupMeshBuffers(meshes)

	// Setup camera

	camera = pbr.NewPlanarCamera()
	camera.SetExposure(16.0, 1.0/125.0, 100.0)
	camera.SetFocus(space.Coord{0, 0, 0})
	camera.SetDistance(4)

	// Setup model
	misc.worldFromObject = space.Identity()

	// Setup light
	misc.SunIlluminance = pbr.DirectionalLightSpectralIlluminance(116400.0, 5400.0)

	// MTX
	mtx.ShowFrameTime(true, -1, 0)

	return glam.Error("gfx", gfx.Err())
}

//------------------------------------------------------------------------------

type loop struct {
	glam.DefaultHandlers
}

//------------------------------------------------------------------------------

func (loop) Update() {
	// prepare(glam.TimeStep())

	p := camera.Focus()
	d := camera.Distance()
	y, pt, r := camera.Orientation()
	mtx.Locate(0, 0)
	mtx.Print("cam: %6.2f,%6.2f,%6.2f\n", p.X, p.Y, p.Z)
	mtx.Print("~~~~~%6.2f\n", d)
	mtx.Print("~~~~~%6.2f,%6.2f,%6.2f\n", y, pt, r)
}

//------------------------------------------------------------------------------

func (loop) Draw(dt64, _ float64) {
	prepare(dt64)

	pipeline.Bind()
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.4, 0.45, 0.5, 1.0})
	gfx.Disable(gfx.Blend)

	camera.Bind()
	miscUBO.SubData(&misc, 0)
	miscUBO.Bind(1)

	poly.BindMeshBuffers()

	gfx.Draw(0, int32(len(meshes.Faces)*6))

	pipeline.Unbind()
}

//------------------------------------------------------------------------------

func prepare(dt64 float64) {
	dt := float32(dt64)

	mtx.Locate(1, 4)
	mtx.Print("%6.2f", glam.Now())

	camera.Move(forward*dt, lateral*dt, vertical*dt)

	m := mouse.SmoothDelta()
	s := plane.CoordOf(window.Size())
	switch {
	case mouse.IsPressed(mouse.Right):
		camera.Rotate(2*m.X/s.X, 2*m.Y/s.Y, rolling*dt)
	case mouse.IsPressed(mouse.Left):
		current.dragDelta = current.dragDelta.Plus(plane.Coord{2 * m.Y / s.Y, 2 * m.X / s.X})
		r := space.EulerXYZ(current.dragDelta.X, current.dragDelta.Y, 0)
		vr := camera.View().WithoutTranslation()
		r = vr.Transpose().Times(r.Times(vr))
		misc.worldFromObject = r.Times(dragStart)
	}
}

//------------------------------------------------------------------------------
