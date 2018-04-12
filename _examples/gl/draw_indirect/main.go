// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/colour"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/space"
	"github.com/cozely/cozely/x/gl"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quit   = input.Bool("Quit")
	rotate = input.Bool("Rotate")
	move   = input.Bool("Move")
	zoom   = input.Bool("Zoom")
)

var context = input.Context("Default", quit, rotate, move, zoom)

var bindings = input.Bindings{
	"Default": {
		"Quit":   {"Escape"},
		"Rotate": {"Mouse Left"},
		"Move":   {"Mouse Right"},
		"Zoom":   {"Mouse Middle"},
	},
}

////////////////////////////////////////////////////////////////////////////////

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
	colour colour.LRGB `layout:"1" divisor:"1"`
}{
	{colour.LRGB{R: 0.2, G: 0, B: 0.6}},
	{colour.LRGB{R: 0.2, G: 0, B: 0.6}},
	{colour.LRGB{R: 0, G: 0.3, B: 0.1}},
	{colour.LRGB{R: 0, G: 0.3, B: 0.1}},
	{colour.LRGB{R: 0.8, G: 0.3, B: 0}},
	{colour.LRGB{R: 0.8, G: 0.3, B: 0}},
}

// Vertex buffer
type mesh []struct {
	position coord.XYZ `layout:"0"`
}

// Tranformation matrices
var (
	screenFromView  space.Matrix // projection matrix
	viewFromWorld   space.Matrix // view matrix
	worldFromObject space.Matrix // model matirx
)

// Cube state
var (
	position   coord.XYZ
	yaw, pitch float32
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	cozely.Configure(cozely.Multisample(8))
	cozely.Events.Resize = resize
	err := cozely.Run(loop{})
	if err != nil {
		cozely.ShowError(err)
		return
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop) Enter() error {
	input.Load(bindings)
	context.Activate(1)

	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader.vert"),
		gl.Shader(cozely.Path()+"shader.frag"),
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
	position = coord.XYZ{0, 0, 0}
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

	return cozely.Error("gl", gl.Err())
}

func (loop) Leave() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() error {
	pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(colour.LRGBA{0.9, 0.9, 0.9, 1.0})

	perFrame.screenFromObject = screenFromView.Times(viewFromWorld)
	perFrame.screenFromObject = perFrame.screenFromObject.Times(worldFromObject)
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	gl.DrawIndirect(0, 6)

	pipeline.Unbind()

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func resize() {
	s := cozely.WindowSize()
	gl.Viewport(0, 0, int32(s.C), int32(s.R))
	r := float32(s.C) / float32(s.R)
	screenFromView = space.Perspective(math32.Pi/4, r, 0.001, 1000.0)
}

////////////////////////////////////////////////////////////////////////////////
