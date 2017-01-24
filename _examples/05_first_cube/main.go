// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"os"
	"time"
	"unsafe"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
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

// Vertex buffer layout
type perVertex struct {
	position space.Coord `layout:"0"`
	color    color.RGB   `layout:"1"`
}

// Uniform buffer
type perObject struct {
	transform space.Matrix
}

// OpenGL objects
var (
	pipeline  gfx.Pipeline
	transform gfx.UniformBuffer
	mesh      gfx.VertexBuffer
)

// Cube state
var (
	distance   float32
	position   space.Coord
	yaw, pitch float32
)

// Matrices
var (
	model      space.Matrix
	view       space.Matrix
	projection space.Matrix
)

//------------------------------------------------------------------------------

func setup() error {
	// Setup the pipeline
	v, err := os.Open(glam.Path() + "shader.vert")
	if err != nil {
		return err
	}
	f, err := os.Open(glam.Path() + "shader.frag")
	if err != nil {
		return err
	}
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(v),
		gfx.FragmentShader(f),
		gfx.VertexFormat(0, perVertex{}),
	)
	gfx.Enable(gfx.DepthTest)
	gfx.CullFace(false, true)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	transform = gfx.NewUniformBuffer(unsafe.Sizeof(perObject{}), gfx.DynamicStorage)

	// Create and fill the vertex buffer
	mesh = gfx.NewVertexBuffer(cube(), 0)

	// Initialize model and view matrices
	position = space.Coord{0, 0, 0}
	yaw = -0.6
	pitch = 0.3
	updateModel()
	distance = 3
	updateView()

	return gfx.Err()
}

//------------------------------------------------------------------------------

type handler struct {
	basic.WindowHandler
	basic.MouseHandler
}

func (h handler) WindowResized(s pixel.Coord, _ time.Duration) {
	sx, sy := window.Size().Cartesian()
	r := sx / sy
	projection = space.Perspective(math.Pi/4, r, 0.001, 1000.0)
}

func (h handler) MouseWheel(motion pixel.Coord, _ time.Duration) {
	distance -= float32(motion.Y) / 4
	updateView()
}

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ time.Duration) {
	mouse.SetRelativeMode(true)
}

func (h handler) MouseButtonUp(b mouse.Button, _ int, _ time.Duration) {
	mouse.SetRelativeMode(false)
}

func (h handler) MouseMotion(motion pixel.Coord, _ pixel.Coord, _ time.Duration) {
	mx, my := motion.Cartesian()
	sx, sy := window.Size().Cartesian()

	switch {
	case mouse.IsPressed(mouse.Left):
		yaw += 4 * mx / sx
		pitch += 4 * my / sy
		switch {
		case pitch < -math.Pi/2:
			pitch = -math.Pi / 2
		case pitch > +math.Pi/2:
			pitch = +math.Pi / 2
		}
		updateModel()

	case mouse.IsPressed(mouse.Middle):
		position.X += 2 * mx / sx
		position.Y -= 2 * my / sy
		updateModel()
	}
}

//------------------------------------------------------------------------------

func updateModel() {
	model = space.Translation(position)
	model = model.Times(space.EulerZXY(pitch, yaw, 0))
}

func updateView() {
	if distance < 1 {
		distance = 1
	}
	view = space.LookAt(space.Coord{0, 0, distance}, space.Coord{0, 0, 0}, space.Coord{0, 1, 0})
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update() {
}

func (l looper) Draw() {
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	pipeline.Bind()
	transform.Bind(0)

	mvp := projection.Times(view)
	mvp = mvp.Times(model)
	t := perObject{
		transform: mvp,
	}
	transform.Load(&t, 0)

	mesh.Bind(0, 0)
	gfx.Draw(gfx.Triangles, 0, 6*2*3)
}

//------------------------------------------------------------------------------

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//------------------------------------------------------------------------------
