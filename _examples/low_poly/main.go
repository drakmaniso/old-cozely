// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"fmt"
	"os"

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
	glam.Setup()

	err := setup()
	if err != nil {
		fmt.Println(err)
		return
	}

	window.Handle = handler{}
	mouse.Handle = handler{}
	key.Handle = handler{}

	glam.TimeStep = 1.0 / 60.0

	// Run the Game Loop
	glam.Loop = looper{}
	err = glam.Run()
	if err != nil {
		fmt.Println(err)
	}
}

//------------------------------------------------------------------------------

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
	// Create and configure the pipeline
	vs, err := os.Open(glam.Path() + "shader.vert")
	if err != nil {
		return err
	}
	fs, err := os.Open(glam.Path() + "shader.frag")
	if err != nil {
		return err
	}
	poly.SetupPipeline(gfx.VertexShader(vs),
		gfx.FragmentShader(fs),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	miscUBO = gfx.NewUniformBuffer(&misc, gfx.DynamicStorage)

	//
	meshes = poly.Meshes{}
	meshes.AddObj(glam.Path() + "../shared/suzanne.obj")
	// meshes.AddObj("E:/objtestfiles/halfsphere.obj")
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
	poly.BindPipeline()

	// pipeline.Bind()
	// vbo.Bind(0, 0)
	// pipeline.Unbind()

	return gfx.Err()
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update(_, dt float64) {
	camera.NextState()

	camera.Move(forward*float32(dt), lateral*float32(dt), vertical*float32(dt))

	if firstPerson {
		mx, my := mouse.Delta().Cartesian()

		const d = 0.4
		smoothedMouse.X += (mx - smoothedMouse.X) * d
		smoothedMouse.Y += (my - smoothedMouse.Y) * d

		sx, sy := window.Size().Cartesian()

		camera.Rotate(2*smoothedMouse.X/sx, 2*smoothedMouse.Y/sy, rolling*float32(dt))
	}

	p := camera.Position()
	y, pt, r := camera.Orientation()
	mtx.Print(1, 0, "cam: %6.2f,%6.2f,%6.2f", p.X, p.Y, p.Z)
	mtx.Print(1, 1, "     %6.2f,%6.2f,%6.2f", y, pt, r)
}

var smoothedMouse plane.Coord

//------------------------------------------------------------------------------

func (l looper) Draw(interpolation float64) {
	poly.BindPipeline()
	gfx.ClearColorBuffer(color.RGBA{0.4, 0.45, 0.5, 1.0})

	camera.Bind()
	miscUBO.SubData(&misc, 0)
	miscUBO.Bind(1)

	poly.BindMeshBuffers()

	// poly.Draw()
	gfx.Draw(gfx.Triangles, 0, int32(len(meshes.Faces)*6))

	poly.UnbindPipeline()
}

func updateModel() {
	misc.Model = space.Translation(object.position)
	misc.Model = misc.Model.Times(space.EulerXZY(object.pitch, object.yaw, object.roll))
	misc.Model = misc.Model.Times(space.Scaling(space.Coord{object.scale, object.scale, object.scale}))
	mtx.Print(1, 3, "obj: %6.2f,%6.2f,%6.2f", object.position.X, object.position.Y, object.position.Z)
	mtx.Print(1, 4, "     %6.2f,%6.2f", object.yaw, object.pitch)
}

//------------------------------------------------------------------------------
