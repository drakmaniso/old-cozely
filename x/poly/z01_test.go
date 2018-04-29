// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/space"
	"github.com/cozely/cozely/x/gl"
	"github.com/cozely/cozely/x/poly"
)

////////////////////////////////////////////////////////////////////////////////

var (
	rotate    = input.Digital("Rotate")
	move      = input.Digital("Move")
	onward    = input.Digital("Onward")
	left      = input.Digital("Left")
	back      = input.Digital("Back")
	right     = input.Digital("Right")
	up        = input.Digital("Up")
	down      = input.Digital("Down")
	rollleft  = input.Digital("Roll Left")
	rollright = input.Digital("Roll Right")
	resetview = input.Digital("Reset View")
	resetobj  = input.Digital("Reset Object")
	rotation  = input.Delta("Rotation")
	cursor    = input.Cursor("Cursor")
)

var context1 = input.Context("Default", quit, rotate, move, rotation, cursor,
	onward, back, left, right, up, down, rollleft, rollright, resetview, resetobj)

var bindings1 = input.Bindings{
	"Default": {
		"Quit":         {"Escape"},
		"Rotate":       {"Mouse Right"},
		"Move":         {"Mouse Left"},
		"Rotation":     {"Mouse", "Right Stick"},
		"Onward":       {"W", "Up"},
		"Left":         {"A", "Left"},
		"Back":         {"S", "Down"},
		"Right":        {"D", "Right"},
		"Up":           {"Space"},
		"Down":         {"Left Shift"},
		"Roll Left":    {"Q"},
		"Roll Right":   {"E"},
		"Reset View":   {"Mouse Back"},
		"Reset Object": {"Mouse Forward"},
		"Cursor":       {"Mouse"},
	},
}

////////////////////////////////////////////////////////////////////////////////

var (
	overlay = pixel.Canvas(pixel.Zoom(2))
	palette = color.Palette(color.SRGB8{0xFF, 0xFF, 0xFF})
)

////////////////////////////////////////////////////////////////////////////////

var pipeline *gl.Pipeline

// Uniform buffer
var miscUBO gl.UniformBuffer
var misc struct {
	worldFromObject space.Matrix
	SunIlluminance  color.LRGB
	_               byte
}

// PlanarCamera

var camera *poly.PlanarCamera

// State

var forward, lateral, vertical, rolling float32
var dragStart space.Matrix

var current struct {
	dragDelta coord.XY
}

var previous struct {
	dragDelta coord.XY
}

// worldFromObject

var meshes poly.Meshes

type loop struct{}

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		cozely.Configure(
			cozely.UpdateStep(1.0/50),
			cozely.Multisample(8),
		)
		cozely.Events.Resize = resize
		err := cozely.Run(loop{})
		if err != nil {
			panic(err)
		}
	})
}

func (loop) Enter() {
	input.ShowMouse(false)
	input.Load(bindings1)
	context1.Activate(1)
	palette.Activate()

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
	// meshes.AddObj(cozely.Path() + "cube.obj")
	// meshes.AddObj(cozely.Path() + "teapot.obj")
	meshes.AddObj(cozely.Path() + "suzanne.obj")
	poly.SetupMeshBuffers(meshes)

	// Setup camera

	camera = poly.NewPlanarCamera()
	camera.SetExposure(16.0, 1.0/125.0, 100.0)
	camera.SetFocus(coord.XYZ{0, 0, 0})
	camera.SetDistance(4)

	// Setup model
	misc.worldFromObject = space.Identity()

	// Setup light
	misc.SunIlluminance = poly.DirectionalLightSpectralIlluminance(116400.0, 5400.0)
}

func (loop) Leave() {
}

func resize() {
	s := cozely.WindowSize()
	gl.Viewport(0, 0, int32(s.C), int32(s.R))
	if camera != nil {
		camera.WindowResized()
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop) React() {
	if move.Started(0) {
		dragStart = misc.worldFromObject
		current.dragDelta = coord.XY{0, 0}
		input.GrabMouse(true)
	}
	if rotate.Started(0) {
		input.GrabMouse(true)
	}

	const s = 2.0
	switch {
	case onward.Ongoing(0):
		forward = -s
	case back.Ongoing(0):
		forward = s
	default:
		forward = 0
	}
	switch {
	case left.Ongoing(0):
		lateral = -s
	case right.Ongoing(0):
		lateral = s
	default:
		lateral = 0
	}
	switch {
	case up.Ongoing(0):
		vertical = s
	case down.Ongoing(0):
		vertical = -s
	default:
		vertical = 0
	}
	switch {
	case rollleft.Ongoing(0):
		rolling = -s
	case rollright.Ongoing(0):
		rolling = s
	default:
		rolling = 0
	}

	if resetview.Started(0) {
		camera.SetFocus(coord.XYZ{0, 0, 0})
		camera.SetDistance(4)
		camera.SetOrientation(0, 0, 0)
	}
	if resetobj.Started(0) {
		misc.worldFromObject = space.Identity()
	}

	if move.Stopped(0) || rotate.Stopped(0) {
		input.GrabMouse(false)
	}

	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop) Update() {
}

func (loop) Render() {
	prepare()

	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	s := cozely.WindowSize()
	gl.Viewport(0, 0, int32(s.C), int32(s.R))
	pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(color.LRGBA{0.025, 0.025, 0.025, 1.0})
	// gl.ClearColorBuffer(color.LRGBA{0.4, 0.45, 0.5, 1.0})
	gl.Disable(gl.Blend)
	gl.Enable(gl.FramebufferSRGB)

	camera.Bind()
	miscUBO.SubData(&misc, 0)
	miscUBO.Bind(1)

	poly.BindMeshBuffers()

	gl.Draw(0, int32(len(meshes.Faces)*6))

	pipeline.Unbind()

	overlay.Clear(0)
	overlay.Locate(0, coord.CR{2, 12})
	ft, or := cozely.RenderStats()
	overlay.Printf("% 3.2f", ft*1000)
	if or > 0 {
		overlay.Printf(" (%d)", or)
	}
	if cozely.HasMouseFocus() {
		overlay.Picture(pixel.MouseCursor, 10, overlay.FromWindow(cursor.XY(0).CR()))
	}
	overlay.Display()
}

func prepare() {
	dt := float32(cozely.RenderDelta())

	camera.Move(forward*dt, lateral*dt, vertical*dt)

	m := rotation.XY(0)

	s := cozely.WindowSize().XY()
	switch {
	case rollleft.Ongoing(1) || rollright.Ongoing(1):
		camera.Rotate(0, 0, rolling*dt)
	case rotate.Ongoing(1):
		camera.Rotate(2*m.X/s.X, 2*m.Y/s.Y, rolling*dt)
	case move.Ongoing(1):
		current.dragDelta = current.dragDelta.Plus(coord.XY{2 * m.Y / s.Y, 2 * m.X / s.X})
		r := space.EulerXYZ(current.dragDelta.X, current.dragDelta.Y, 0)
		vr := camera.View().WithoutTranslation()
		r = vr.Transpose().Times(r.Times(vr))
		misc.worldFromObject = r.Times(dragStart)
	}
}
