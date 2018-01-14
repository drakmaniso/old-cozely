// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/mouse"
	"github.com/drakmaniso/carol/pbr"
	"github.com/drakmaniso/carol/plane"
	"github.com/drakmaniso/carol/poly"
	"github.com/drakmaniso/carol/space"
)

//------------------------------------------------------------------------------

func main() {
	carol.SetTimeStep(1 / 2.0)

	err := carol.Run(loop{})
	if err != nil {
		carol.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

var pipeline *gl.Pipeline

// Uniform buffer
var miscUBO gl.UniformBuffer
var misc struct {
	worldFromObject space.Matrix
	SunIlluminance  colour.RGB
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

type loop struct {
	carol.Handlers
}

//------------------------------------------------------------------------------

func (loop) Setup() error {
	pipeline = gl.NewPipeline(
		poly.PipelineSetup(),
		pbr.ToneMapACES(),
		gl.Shader(carol.Path()+"shader.vert"),
		gl.Shader(carol.Path()+"shader.frag"),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create the uniform buffer
	miscUBO = gl.NewUniformBuffer(&misc, gl.DynamicStorage)

	//
	meshes = poly.Meshes{}
	// meshes.AddObj(carol.Path() + "../shared/cube.obj")
	// meshes.AddObj(carol.Path() + "../shared/teapot.obj")
	meshes.AddObj(carol.Path() + "../shared/suzanne.obj")
	// meshes.AddObj("E:/objtestfiles/pony.obj")
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

	return carol.Error("gl", gl.Err())
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	// prepare(carol.TimeStep())

	// p := camera.Focus()
	// d := camera.Distance()
	// y, pt, r := camera.Orientation()

	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw(dt64, _ float64) error {
	prepare(dt64)

	pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(colour.RGBA{0.0, 0.0, 0.0, 1.0})
	// gl.ClearColorBuffer(colour.RGBA{0.4, 0.45, 0.5, 1.0})
	gl.Disable(gl.Blend)

	camera.Bind()
	miscUBO.SubData(&misc, 0)
	miscUBO.Bind(1)

	poly.BindMeshBuffers()

	gl.Draw(0, int32(len(meshes.Faces)*6))

	pipeline.Unbind()

	return gl.Err()
}

//------------------------------------------------------------------------------

func prepare(dt64 float64) {
	dt := float32(dt64)

	camera.Move(forward*dt, lateral*dt, vertical*dt)

	// m := mouse.SmoothDelta()
	mx, my := mouse.Delta()
	m := plane.Coord{float32(mx), float32(my)}

	w, h := carol.WindowSize()
	s := plane.Coord{float32(w), float32(h)}
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
