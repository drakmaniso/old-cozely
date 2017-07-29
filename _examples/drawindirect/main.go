// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/mtx"
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
)

// Uniform buffer
var perFrame struct {
	screenFromObject space.Matrix
}

// Indirect Command Buffer
var commands = []gfx.DrawIndirectCommand{
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   0,
		BaseInstance:  1,
	},
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   6,
		BaseInstance:  1,
	},
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   12,
		BaseInstance:  2,
	},
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   18,
		BaseInstance:  3,
	},
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   24,
		BaseInstance:  4,
	},
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   30,
		BaseInstance:  5,
	},
}

// Instance Buffer

var draws = []struct {
	color color.RGB `layout:"1" divisor:"1"`
}{
	{color.RGB{R: 0.2, G: 0, B: 0.6}},
	{color.RGB{R: 0.2, G: 0, B: 0.6}},
	{color.RGB{R: 0, G: 0.3, B: 0.1}},
	{color.RGB{R: 0, G: 0.3, B: 0.1}},
	{color.RGB{R: 0.8, G: 0.3, B: 0}},
	{color.RGB{R: 0.8, G: 0.3, B: 0}},
}

// Vertex buffer
type mesh []struct {
	position space.Coord `layout:"0"`
}

// Tranformation matrices
var (
	screenFromView  space.Matrix // projection matrix
	viewFromWorld   space.Matrix // view matrix
	worldFromObject space.Matrix // model matirx
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
		gfx.VertexFormat(1, draws),
		gfx.Topology(gfx.Triangles),
		gfx.CullFace(false, true),
		gfx.DepthTest(true),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)

	// Create the Indirect Command Buffer
	icbo := gfx.NewIndirectBuffer(commands, gfx.DynamicStorage)
	ibo := gfx.NewVertexBuffer(draws, gfx.DynamicStorage)

	// Create and fill the vertex buffer
	vbo := gfx.NewVertexBuffer(cube(), 0)

	// Initialize worldFromObject and viewFromWorld matrices
	position = space.Coord{0, 0, 0}
	yaw = -0.6
	pitch = 0.3
	computeWorldFromObject()
	computeViewFromWorld()

	// MTX
	mtx.Color(color.RGB{0.0, 0.05, 0.1}, color.RGB{0.7, 0.6, 0.45})
	mtx.Opaque(false)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	icbo.Bind()
	ibo.Bind(1, 0)
	pipeline.Unbind()

	return glam.Error("gfx", gfx.Err())
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

	perFrame.screenFromObject = screenFromView.Times(viewFromWorld)
	perFrame.screenFromObject = perFrame.screenFromObject.Times(worldFromObject)
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	gfx.DrawIndirect(0, 6)

	pipeline.Unbind()
}

//------------------------------------------------------------------------------

func computeWorldFromObject() {
	worldFromObject = space.Translation(position)
	worldFromObject = worldFromObject.Times(space.EulerZXY(pitch, yaw, 0))
}

func computeViewFromWorld() {
	viewFromWorld = space.LookAt(space.Coord{0, 0, 3}, space.Coord{0, 0, 0}, space.Coord{0, 1, 0})
}

//------------------------------------------------------------------------------
