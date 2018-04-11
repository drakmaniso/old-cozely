// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"

	"github.com/drakmaniso/cozely"
	"github.com/drakmaniso/cozely/colour"
	"github.com/drakmaniso/cozely/input"
	"github.com/drakmaniso/cozely/palette"
	"github.com/drakmaniso/cozely/pixel"
	"github.com/drakmaniso/cozely/plane"
	"github.com/drakmaniso/cozely/space"
	"github.com/drakmaniso/cozely/x/gl"
	"github.com/drakmaniso/cozely/x/poly"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quit   = input.Bool("Quit")
	rotate = input.Bool("Rotate")
	move   = input.Bool("Move")
	zoom   = input.Bool("Zoom")
)

var context = input.Context("Default", quit, rotate, move, zoom)

var bindings = input.Bindings{
	"Default": {
		"Quit":   {"Escape"},
		"Rotate": {"Mouse Left"},
		"Move":   {"Mouse Right"},
		"Zoom":   {"Mouse Middle"},
	},
}

////////////////////////////////////////////////////////////////////////////////

var overlay = pixel.Canvas(pixel.Zoom(2))

var cursor = pixel.Cursor{Canvas: overlay}

var font = pixel.FontID(0)

var txtColor = palette.Index(1)

////////////////////////////////////////////////////////////////////////////////

var pipeline *gl.Pipeline

// Uniform buffer
var miscUBO gl.UniformBuffer
var misc struct {
	worldFromObject space.Matrix
	SunIlluminance  colour.LRGB
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

var gametime float64

////////////////////////////////////////////////////////////////////////////////

func main() {
	cozely.Configure(
		cozely.UpdateStep(1.0/50),
		cozely.Multisample(8),
	)
	cozely.Events.Resize = resize
	err := cozely.Run(loop{})
	if err != nil {
		cozely.ShowError(err)
		return
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop) Enter() error {
	input.Load(bindings)
	context.Activate(1)

	txtColor.SetColour(colour.SRGB8{0xFF, 0xFF, 0xFF})

	pipeline = gl.NewPipeline(
		poly.PipelineSetup(),
		poly.ToneMapACES(),
		gl.Shader(cozely.Path()+"shader.vert"),
		gl.Shader(cozely.Path()+"shader.frag"),
		gl.DepthTest(true),
		gl.DepthWrite(true),
	)

	// Create the uniform buffer
	miscUBO = gl.NewUniformBuffer(&misc, gl.DynamicStorage)

	//
	meshes = poly.Meshes{}
	// meshes.AddObj(cozely.Path() + "../../shared/cube.obj")
	// meshes.AddObj(cozely.Path() + "../../shared/teapot.obj")
	meshes.AddObj(cozely.Path() + "../../shared/suzanne.obj")
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

	return cozely.Error("gl", gl.Err())
}

func (loop) Leave() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (l loop) MouseMotion(_, _ int32, _, _ int32) {
	if cozely.GameTime() < gametime {
		fmt.Printf("***************ERROR************\n")
	}
	// fmt.Printf("  (%.4f: %.4f, %.4f)\n", cozely.GameTime(), cozely.RenderTime(), cozely.UpdateLag())
	gametime = cozely.GameTime()
}

func (loop) Update() error {
	if cozely.GameTime() < gametime {
		fmt.Printf("***************ERROR************\n")
	}
	// fmt.Printf(" - %.4f: %.4f, %.4f\n", cozely.GameTime(), cozely.RenderTime(), cozely.UpdateLag())
	gametime = cozely.GameTime()

	// prepare()

	// p := camera.Focus()
	// d := camera.Distance()
	// y, pt, r := camera.Orientation()

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() error {
	if cozely.GameTime() < gametime {
		fmt.Printf("***************ERROR************\n")
	}
	// fmt.Printf("## %.4f: %.4f, %.4f\n", cozely.GameTime(), cozely.RenderTime(), cozely.UpdateLag())
	gametime = cozely.GameTime()

	prepare()

	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	s := cozely.WindowSize()
	gl.Viewport(0, 0, int32(s.X), int32(s.Y))
	pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(colour.LRGBA{0.0, 0.0, 0.0, 1.0})
	// gl.ClearColorBuffer(colour.LRGBA{0.4, 0.45, 0.5, 1.0})
	gl.Disable(gl.Blend)
	gl.Enable(gl.FramebufferSRGB)

	camera.Bind()
	miscUBO.SubData(&misc, 0)
	miscUBO.Bind(1)

	poly.BindMeshBuffers()

	gl.Draw(0, int32(len(meshes.Faces)*6))

	pipeline.Unbind()

	overlay.Clear(0)
	cursor.Locate(2, 12)
	ft, or := cozely.RenderStats()
	cursor.Printf("% 3.2f", ft*1000)
	if or > 0 {
		cursor.Printf(" (%d)", or)
	}
	overlay.Display()

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func prepare() {
	/*
		dt := float32(cozely.RenderTime())

		camera.Move(forward*dt, lateral*dt, vertical*dt)

		// m := mouse.SmoothDelta()
		mx, my := mouse.Delta()
		m := plane.Coord{float32(mx), float32(my)}

		s := cozely.WindowSize().Cartesian()
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
	*/
}

////////////////////////////////////////////////////////////////////////////////

func resize() {
	s := cozely.WindowSize()
	gl.Viewport(0, 0, int32(s.X), int32(s.Y))
	if camera != nil {
		camera.WindowResized()
	}
}

////////////////////////////////////////////////////////////////////////////////
