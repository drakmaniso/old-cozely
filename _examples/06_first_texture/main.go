// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
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

	// Run the main loop
	glam.Loop = looper{}
	err = glam.Run()
	if err != nil {
		glam.ErrorDialog(err)
	}
}

//------------------------------------------------------------------------------

// Vertex buffer layout
type vertex struct {
	position space.Coord `layout:"0"`
	uv       plane.Coord `layout:"1"`
}

// Uniform buffer
var perFrame struct {
	transform space.Matrix
}

// OpenGL objects
var (
	pipeline    gfx.Pipeline
	perFrameUBO gfx.UniformBuffer
	cubeVBO     gfx.VertexBuffer
	sampler     gfx.Sampler
	diffuse     gfx.Texture2D
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
		gfx.VertexFormat(0, vertex{}),
	)
	gfx.Enable(gfx.DepthTest)
	gfx.CullFace(false, true)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)

	// Create and fill the vertex buffer
	cubeVBO = gfx.NewVertexBuffer(cube(), gfx.StaticStorage)

	// Create and bind the sampler
	sampler = gfx.NewSampler(
		gfx.Minification(gfx.LinearMipmapLinear),
		gfx.Anisotropy(16.0),
	)

	// Create and load the textures
	diffuse = gfx.NewTexture2D(8, pixel.Coord{512, 512}, gfx.SRGBA8)
	r, err := os.Open(glam.Path() + "../shared/testpattern.png")
	if err != nil {
		return err
	}
	defer r.Close()
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	diffuse.Load(img, pixel.Coord{0, 0}, 0)
	diffuse.GenerateMipmap()

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

// Event handler
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
		position.X += 2 * mx / sx
		position.Y -= 2 * my / sy
		updateModel()

	case mouse.IsPressed(mouse.Right):
		yaw += 4 * mx / sx
		pitch += 4 * my / sy
		switch {
		case pitch < -math.Pi/2:
			pitch = -math.Pi / 2
		case pitch > +math.Pi/2:
			pitch = +math.Pi / 2
		}
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
	sampler.Bind(0)

	perFrame.transform = projection.Times(view)
	perFrame.transform = perFrame.transform.Times(model)
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	cubeVBO.Bind(0, 0)
	diffuse.Bind(0)
	gfx.Draw(gfx.Triangles, 0, 6*2*3)
}

//------------------------------------------------------------------------------
