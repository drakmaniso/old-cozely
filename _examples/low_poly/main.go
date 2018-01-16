// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"fmt"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/poly"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

func main() {
	glam.Configure(
		glam.TimeStep(1.0 / 50),
	)

	err := glam.Run(setup, loop{})
	if err != nil {
		glam.ShowError(err)
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

var camera *poly.PlanarCamera

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
	pipeline = gl.NewPipeline(
		poly.PipelineSetup(),
		poly.ToneMapACES(),
		gl.Shader(glam.Path()+"shader.vert"),
		gl.Shader(glam.Path()+"shader.frag"),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create the uniform buffer
	miscUBO = gl.NewUniformBuffer(&misc, gl.DynamicStorage)

	//
	meshes = poly.Meshes{}
	// meshes.AddObj(glam.Path() + "../shared/cube.obj")
	// meshes.AddObj(glam.Path() + "../shared/teapot.obj")
	meshes.AddObj(glam.Path() + "../shared/suzanne.obj")
	// meshes.AddObj("E:/objtestfiles/pony.obj")
	poly.SetupMeshBuffers(meshes)

	// Setup camera

	camera = poly.NewPlanarCamera()
	camera.SetExposure(16.0, 1.0/125.0, 100.0)
	camera.SetFocus(space.Coord{0, 0, 0})
	camera.SetDistance(4)

	// Setup model
	misc.worldFromObject = space.Identity()

	// Setup light
	misc.SunIlluminance = poly.DirectionalLightSpectralIlluminance(116400.0, 5400.0)

	return glam.Error("gl", gl.Err())
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------
var gametime float64

func (l loop) MouseMotion(_, _ int32, _, _ int32) {
	if glam.GameTime() < gametime {
		fmt.Printf("***************ERROR************\n")
	}
	// fmt.Printf("  (%.4f: %.4f, %.4f)\n", glam.GameTime(), glam.FrameTime(), glam.UpdateLag())
	gametime = glam.GameTime()
}
func (loop) Update() error {
	if glam.GameTime() < gametime {
		fmt.Printf("***************ERROR************\n")
	}
	// fmt.Printf(" - %.4f: %.4f, %.4f\n", glam.GameTime(), glam.FrameTime(), glam.UpdateLag())
	gametime = glam.GameTime()

	// prepare()

	// p := camera.Focus()
	// d := camera.Distance()
	// y, pt, r := camera.Orientation()

	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw() error {
	if glam.GameTime() < gametime {
		fmt.Printf("***************ERROR************\n")
	}
	// fmt.Printf("## %.4f: %.4f, %.4f\n", glam.GameTime(), glam.FrameTime(), glam.UpdateLag())
	gametime = glam.GameTime()

	prepare()

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

func prepare() {
	dt := float32(glam.FrameTime())

	camera.Move(forward*dt, lateral*dt, vertical*dt)

	// m := mouse.SmoothDelta()
	mx, my := mouse.Delta()
	m := plane.Coord{float32(mx), float32(my)}

	w, h := glam.WindowSize()
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
