// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"os"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
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

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
	vs, err := os.Open(glam.Path() + "shader.vert")
	if err != nil {
		return glam.Error("opemning vertex shader", err)
	}
	fs, err := os.Open(glam.Path() + "shader.frag")
	if err != nil {
		return glam.Error("opening fragment shader", err)
	}
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(vs),
		gfx.FragmentShader(fs),
		gfx.Topology(gfx.Triangles),
	)

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
