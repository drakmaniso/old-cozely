// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/plane"
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

	glam.Draw = draw

	err = glam.Loop()
	if err != nil {
		glam.ShowError("running", err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline *gfx.Pipeline
)

// Vertex buffer
type mesh []struct {
	position plane.Coord `layout:"0"`
	color    color.RGB   `layout:"1"`
}

//------------------------------------------------------------------------------

func setup() error {
	var triangle mesh

	// Create and configure the pipeline
	pipeline = gfx.NewPipeline(
		gfx.Shader(glam.Path()+"shader.vert"),
		gfx.Shader(glam.Path()+"shader.frag"),
		gfx.VertexFormat(0, triangle),
		gfx.Topology(gfx.Triangles),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create and fill the vertex buffer
	triangle = mesh{
		{plane.Coord{0, 0.65}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{plane.Coord{-0.65, -0.475}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{plane.Coord{0.65, -0.475}, color.RGB{R: 0, G: 0.6, B: 0.2}},
	}
	vbo := gfx.NewVertexBuffer(triangle, gfx.StaticStorage)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	return glam.Error("gfx", gfx.Err())
}

//------------------------------------------------------------------------------

func draw() {
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	pipeline.Bind()
	gfx.Draw(0, 3)
	pipeline.Unbind()
}

//------------------------------------------------------------------------------
