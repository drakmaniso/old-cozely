// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"os"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/plane"
)

//------------------------------------------------------------------------------

func main() {
	glam.Setup()

	err := setup()
	if err != nil {
		glam.ErrorDialog(err)
		return
	}

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
	pipeline gfx.Pipeline
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
	vs, err := os.Open(glam.Path() + "/shader.vert")
	if err != nil {
		return err
	}
	fs, err := os.Open(glam.Path() + "/shader.frag")
	if err != nil {
		return err
	}
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(vs),
		gfx.FragmentShader(fs),
		gfx.VertexFormat(0, triangle),
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

	return gfx.Err()
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update() {
}

func (l looper) Draw() {
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	pipeline.Bind()
	gfx.Draw(gfx.Triangles, 0, 3)
	pipeline.Unbind()
}

//------------------------------------------------------------------------------
