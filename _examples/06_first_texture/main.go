// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"image"
	_ "image/png"
	"os"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
	glam.Setup()

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

// OpenGL objects
var (
	pipeline    gfx.Pipeline
	perFrameUBO gfx.UniformBuffer
	sampler     gfx.Sampler
	diffuse     gfx.Texture2D
)

// Uniform buffer
var perFrame struct {
	transform space.Matrix
}

// Vertex buffer
type mesh []struct {
	position space.Coord `layout:"0"`
	uv       plane.Coord `layout:"1"`
}

// Matrices
var (
	model      space.Matrix
	view       space.Matrix
	projection space.Matrix
)

// Cube state
var (
	distance   float32
	position   space.Coord
	yaw, pitch float32
)

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
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
		gfx.VertexFormat(0, mesh{}),
	)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)

	// Create and fill the vertex buffer
	vbo := gfx.NewVertexBuffer(cube(), gfx.StaticStorage)

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

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	// MTX
	mtx.Color(color.RGB{0.0, 0.05, 0.1}, color.RGB{0.7, 0.6, 0.45})
	mtx.Opaque(true)

	return gfx.Err()
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

func (l looper) Update(_, _ float64) {
	if glam.Overruns() > 0 {
		ftclip.Print("\f\a%5.1f\a", glam.AverageFrameTime()*1000.0)
	} else {
		ftclip.Print("\f%5.1f", glam.AverageFrameTime()*1000.0)
	}
}

var ftclip = mtx.Clip{
	Left: -5, Top: 0,
	Right: -1, Bottom: 0,
	TransparentSpace: true,
}

func (l looper) Draw(_ float64) {
	pipeline.Bind()
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	gfx.Enable(gfx.DepthTest)
	gfx.CullFace(false, true)
	gfx.Enable(gfx.FramebufferSRGB)

	perFrame.transform = projection.Times(view)
	perFrame.transform = perFrame.transform.Times(model)
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	diffuse.Bind(0)
	sampler.Bind(0)
	gfx.Draw(gfx.Triangles, 0, 6*2*3)

	pipeline.Unbind()
}

//------------------------------------------------------------------------------
