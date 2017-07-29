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
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/space"
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

	err = glam.Run()
	if err != nil {
		glam.ShowError("running", err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline    *gfx.Pipeline
	perFrameUBO gfx.UniformBuffer
	sampler     gfx.Sampler
	diffuse     gfx.Texture2D
)

// Uniform buffer
var perObject struct {
	screenFromObject space.Matrix
}

// Vertex buffer
type mesh []struct {
	position space.Coord `layout:"0"`
	uv       plane.Coord `layout:"1"`
}

// Transformation matrices
var (
	screenFromView  space.Matrix // projection matrix
	viewFromWorld   space.Matrix // view matrix
	worldFromObject space.Matrix // model matrix
)

// Cube state
var (
	position   space.Coord
	yaw, pitch float32
)

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
	pipeline = gfx.NewPipeline(
		gfx.Shader(glam.Path()+"shader.vert"),
		gfx.Shader(glam.Path()+"shader.frag"),
		gfx.VertexFormat(0, mesh{}),
		gfx.Topology(gfx.Triangles),
		gfx.CullFace(false, true),
		gfx.DepthTest(true),
	)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perObject, gfx.DynamicStorage)

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
		return glam.Error("opening texture", err)
	}
	defer r.Close()
	img, _, err := image.Decode(r)
	if err != nil {
		return glam.Error("decoding texture", err)
	}
	diffuse.Load(img, pixel.Coord{0, 0}, 0)
	diffuse.GenerateMipmap()

	// Initialize worldFromObject and viewFromWorld matrices
	position = space.Coord{0, 0, 0}
	yaw = -0.6
	pitch = 0.3
	computeWorldFromObject()
	computeViewFromWorld()

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	// MTX
	mtx.Color(color.RGB{0.0, 0.05, 0.1}, color.RGB{0.7, 0.6, 0.45})
	mtx.Opaque(false)
	mtx.ShowFrameTime(true, -1, 0, false)

	return glam.Error("gfx", gfx.Err())
}

//------------------------------------------------------------------------------

func computeWorldFromObject() {
	r := space.EulerZXY(pitch, yaw, 0)
	worldFromObject = space.Translation(position).Times(r)
}

func computeViewFromWorld() {
	viewFromWorld = space.LookAt(
		space.Coord{0, 0, 3},
		space.Coord{0, 0, 0},
		space.Coord{0, 1, 0},
	)
}

//------------------------------------------------------------------------------

type loop struct {
	glam.DefaultHandlers
}

func (loop) Update() {
}

func (loop) Draw(_, _ float64) {
	pipeline.Bind()
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	gfx.Enable(gfx.FramebufferSRGB)

	perObject.screenFromObject =
		screenFromView.
			Times(viewFromWorld).
			Times(worldFromObject)
	perFrameUBO.SubData(&perObject, 0)
	perFrameUBO.Bind(0)

	diffuse.Bind(0)
	sampler.Bind(0)
	gfx.Draw(0, 6*2*3)

	pipeline.Unbind()
}

//------------------------------------------------------------------------------
