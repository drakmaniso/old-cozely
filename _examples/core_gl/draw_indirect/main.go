// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/space"
)

//------------------------------------------------------------------------------

func main() {
	err := carol.Run(loop{})
	if err != nil {
		carol.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline    *gl.Pipeline
	perFrameUBO gl.UniformBuffer
)

// Uniform buffer
var perFrame struct {
	screenFromObject space.Matrix
}

// Indirect Command Buffer
var commands = []gl.DrawIndirectCommand{
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
	colour colour.RGB `layout:"1" divisor:"1"`
}{
	{colour.RGB{R: 0.2, G: 0, B: 0.6}},
	{colour.RGB{R: 0.2, G: 0, B: 0.6}},
	{colour.RGB{R: 0, G: 0.3, B: 0.1}},
	{colour.RGB{R: 0, G: 0.3, B: 0.1}},
	{colour.RGB{R: 0.8, G: 0.3, B: 0}},
	{colour.RGB{R: 0.8, G: 0.3, B: 0}},
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

type loop struct {
	carol.Handlers
}

//------------------------------------------------------------------------------

func (loop) Setup() error {
	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(carol.Path()+"shader.vert"),
		gl.Shader(carol.Path()+"shader.frag"),
		gl.VertexFormat(0, mesh{}),
		gl.VertexFormat(1, draws),
		gl.Topology(gl.Triangles),
		gl.CullFace(false, true),
		gl.DepthTest(true),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gl.NewUniformBuffer(&perFrame, gl.DynamicStorage)

	// Create the Indirect Command Buffer
	icbo := gl.NewIndirectBuffer(commands, gl.DynamicStorage)
	ibo := gl.NewVertexBuffer(draws, gl.DynamicStorage)

	// Create and fill the vertex buffer
	vbo := gl.NewVertexBuffer(cube(), 0)

	// Initialize worldFromObject and viewFromWorld matrices
	position = space.Coord{0, 0, 0}
	yaw = -0.6
	pitch = 0.3
	computeWorldFromObject()
	computeViewFromWorld()

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	icbo.Bind()
	ibo.Bind(1, 0)
	pipeline.Unbind()

	return carol.Error("gl", gl.Err())
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw(_, _ float64) error {
	pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(colour.RGBA{0.9, 0.9, 0.9, 1.0})

	perFrame.screenFromObject = screenFromView.Times(viewFromWorld)
	perFrame.screenFromObject = perFrame.screenFromObject.Times(worldFromObject)
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	gl.DrawIndirect(0, 6)

	pipeline.Unbind()

	return gl.Err()
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
