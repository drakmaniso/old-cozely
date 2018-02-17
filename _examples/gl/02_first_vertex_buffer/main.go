// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
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
	err := glam.Run(loop{})
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
	color    colour.LRGB `layout:"1"`
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (loop) Enter() error {
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
		{plane.Coord{0, 0.65}, colour.LRGB{R: 0.3, G: 0, B: 0.8}},
		{plane.Coord{-0.65, -0.475}, colour.LRGB{R: 0.8, G: 0.3, B: 0}},
		{plane.Coord{0.65, -0.475}, colour.LRGB{R: 0, G: 0.6, B: 0.2}},
	}
	vbo := gl.NewVertexBuffer(triangle, gl.StaticStorage)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	return glam.Error("gl", gl.Err())
}

//------------------------------------------------------------------------------

func (l loop) WindowResized(w, h int32) {
	gl.Viewport(0, 0, w, h)
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

func (loop) Draw() error {
	pipeline.Bind()
	gl.ClearColorBuffer(colour.LRGBA{0.9, 0.9, 0.9, 1.0})

	gl.Draw(0, 3)
	pipeline.Unbind()

	return gl.Err()
}

//------------------------------------------------------------------------------
