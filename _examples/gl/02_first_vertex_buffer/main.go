// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

func main() {
	err := glam.Run(setup, loop{})
	if err != nil {
		glam.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline *gl.Pipeline
)

// Vertex buffer
type mesh []struct {
	position plane.Coord `layout:"0"`
	color    colour.RGB  `layout:"1"`
}

//------------------------------------------------------------------------------

func setup() error {
	var triangle mesh

	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(glam.Path()+"shader.vert"),
		gl.Shader(glam.Path()+"shader.frag"),
		gl.VertexFormat(0, triangle),
		gl.Topology(gl.Triangles),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create and fill the vertex buffer
	triangle = mesh{
		{plane.Coord{0, 0.65}, colour.RGB{R: 0.3, G: 0, B: 0.8}},
		{plane.Coord{-0.65, -0.475}, colour.RGB{R: 0.8, G: 0.3, B: 0}},
		{plane.Coord{0.65, -0.475}, colour.RGB{R: 0, G: 0.6, B: 0.2}},
	}
	vbo := gl.NewVertexBuffer(triangle, gl.StaticStorage)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

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

func (loop) Draw() error {
	pipeline.Bind()
	gl.ClearColorBuffer(colour.RGBA{0.9, 0.9, 0.9, 1.0})

	gl.Draw(0, 3)
	pipeline.Unbind()

	return gl.Err()
}

//------------------------------------------------------------------------------
