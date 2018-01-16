// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

func main() {
	glam.Configure(
		glam.Multisample(8),
	)

	err := glam.Run(setup, loop{})
	if err != nil {
		glam.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	facePipeline *gl.Pipeline
	edgePipeline *gl.Pipeline
	perFrameUBO  gl.UniformBuffer
)

// Uniform buffer
var perObject struct {
	screenFromObject space.Matrix
}

// Vertex buffer
type mesh []struct {
	position space.Coord `layout:"0"`
	color    colour.RGB  `layout:"1"`
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
	// Create and configure the pipelines
	facePipeline = gl.NewPipeline(
		gl.Shader(glam.Path()+"shader.vert"),
		gl.Shader(glam.Path()+"shader.frag"),
		gl.VertexFormat(0, mesh{}),
		gl.Topology(gl.Triangles),
		gl.CullFace(false, true),
		gl.DepthTest(true),
	)
	edgePipeline = gl.NewPipeline(
		gl.Shader(glam.Path()+"shader.vert"),
		gl.Shader(glam.Path()+"shader.frag"),
		gl.VertexFormat(0, mesh{}),
		gl.Topology(gl.Lines),
		gl.CullFace(false, false),
		gl.DepthTest(true),
		gl.DepthComparison(gl.LessOrEqual),
	)

	gl.Enable(gl.FramebufferSRGB)
	// gl.DepthRange(0.1, 10.0)

	// Create the uniform buffer
	perFrameUBO = gl.NewUniformBuffer(&perObject, gl.DynamicStorage)

	// Create and fill the vertex buffer
	fvbo := gl.NewVertexBuffer(cubeFaces(), 0)
	evbo := gl.NewVertexBuffer(cubeEdges(), 0)

	// Initialize worldFromObject and viewFromWorld matrices
	position = space.Coord{0, 0, 0}
	yaw = -0.6
	pitch = 0.3
	computeWorldFromObject()
	computeViewFromWorld()

	// Bind the vertex buffer to the pipeline
	facePipeline.Bind()
	fvbo.Bind(0, 0)

	edgePipeline.Bind()
	evbo.Bind(0, 0)

	edgePipeline.Unbind()

	return glam.Error("gl", gl.Err())
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw() error {
	gl.ClearColorBuffer(colour.RGBA{0.9, 0.9, 0.9, 1.0})
	gl.ClearDepthBuffer(1.0)

	perObject.screenFromObject =
		screenFromView.
			Times(viewFromWorld).
			Times(worldFromObject)
	perFrameUBO.SubData(&perObject, 0)
	perFrameUBO.Bind(0)

	facePipeline.Bind()
	gl.Disable(gl.Blend)
	gl.Draw(0, 6*2*3)

	edgePipeline.Bind()
	gl.Disable(gl.LineSmooth)
	gl.Disable(gl.Blend)
	gl.Draw(0, 6*2*3)

	// edgePipeline.Bind()
	// gl.Enable(gl.LineSmooth)
	// gl.Enable(gl.Blend)
	// gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)
	// gl.Draw(0, 6*2*3)

	edgePipeline.Unbind()

	return gl.Err()
}

//------------------------------------------------------------------------------

func computeWorldFromObject() {
	rot := space.EulerZXY(pitch, yaw, 0)
	worldFromObject = space.Translation(position).Times(rot)
}

func computeViewFromWorld() {
	viewFromWorld = space.LookAt(
		space.Coord{0, 0, 3},
		space.Coord{0, 0, 0},
		space.Coord{0, 1, 0},
	)
}

//------------------------------------------------------------------------------
