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
	"github.com/cozely/cozely/window"
	"github.com/cozely/cozely/x/gl"
	"github.com/cozely/cozely/x/poly"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quit      = input.Button("Quit")
	rotate    = input.Button("Rotate")
	move      = input.Button("Move")
	onward    = input.Button("Onward")
	left      = input.Button("Left")
	back      = input.Button("Back")
	right     = input.Button("Right")
	up        = input.Button("Up")
	down      = input.Button("Down")
	rollleft  = input.Button("Roll Left")
	rollright = input.Button("Roll Right")
	resetview = input.Button("Reset View")
	resetobj  = input.Button("Reset Object")
	rotation  = input.Delta("Rotation")
	cursor    = input.Cursor("Cursor")
)

var context1 = input.Context("Default", quit, rotate, move, rotation, cursor,
	onward, back, left, right, up, down, rollleft, rollright, resetview, resetobj)

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
		window.Events.Resize = resize
		err := cozely.Run(loop{})
		if err != nil {
			panic(err)
		}
	})
}

func (loop) Enter() {
	input.ShowMouse(false)
	context1.ActivateOn(1)

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
	s := window.Size()
	gl.Viewport(0, 0, int32(s.X), int32(s.Y))
	if camera != nil {
		camera.WindowResized()
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop) React() {
	if move.Pressed() {
		dragStart = misc.worldFromObject
		current.dragDelta = coord.XY{0, 0}
		input.GrabMouse(true)
	}
	if rotate.Pressed() {
		input.GrabMouse(true)
	}

	const s = 2.0
	switch {
	case onward.Ongoing():
		forward = -s
	case back.Ongoing():
		forward = s
	default:
		forward = 0
	}
	switch {
	case left.Ongoing():
		lateral = -s
	case right.Ongoing():
		lateral = s
	default:
		lateral = 0
	}
	switch {
	case up.Ongoing():
		vertical = s
	case down.Ongoing():
		vertical = -s
	default:
		vertical = 0
	}
	switch {
	case rollleft.Ongoing():
		rolling = -s
	case rollright.Ongoing():
		rolling = s
	default:
		rolling = 0
	}

	if resetview.Pressed() {
		camera.SetFocus(coord.XYZ{0, 0, 0})
		camera.SetDistance(4)
		camera.SetOrientation(0, 0, 0)
	}
	if resetobj.Pressed() {
		misc.worldFromObject = space.Identity()
	}

	if move.Released() || rotate.Released() {
		input.GrabMouse(false)
	}

	if quit.Pressed() {
		cozely.Stop(nil)
	}
}

func (loop) Update() {
}

func (loop) Render() {
	prepare()

	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	s := window.Size()
	gl.Viewport(0, 0, int32(s.X), int32(s.Y))
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

	pixel.Clear(0)
	cur := pixel.Cursor{
		Color: 7,
	}
	cur.Locate(0, pixel.XY{2, 12})
	ft, or := cozely.RenderStats()
	cur.Printf("% 3.2f", ft*1000)
	if or > 0 {
		cur.Printf(" (%d)", or)
	}
	if window.HasMouseFocus() {
		pixel.MouseCursor.Paint(0, pixel.XYof(cursor.XY()))
	}
}

func prepare() {
	dt := float32(cozely.RenderDelta())

	camera.Move(forward*dt, lateral*dt, vertical*dt)

	m := rotation.XY()

	s := coord.XYof(window.Size())
	switch {
	case rollleft.Ongoing() || rollright.OngoingOn(1):
		camera.Rotate(0, 0, rolling*dt)
	case rotate.Ongoing():
		camera.Rotate(2*m.X/s.X, 2*m.Y/s.Y, rolling*dt)
	case move.Ongoing():
		current.dragDelta = current.dragDelta.Plus(coord.XY{2 * m.Y / s.Y, 2 * m.X / s.X})
		r := space.EulerXYZ(current.dragDelta.X, current.dragDelta.Y, 0)
		vr := camera.View().WithoutTranslation()
		r = vr.Transpose().Times(r.Times(vr))
		misc.worldFromObject = r.Times(dragStart)
	}
}
