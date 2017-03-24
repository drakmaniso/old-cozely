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
	glam.Setup()

	err := setup()
	if err != nil {
		glam.Log("ERROR during setup: \n", err)
		return
	}

	// Run the main Loop
	glam.Loop = looper{}
	err = glam.Run()
	if err != nil {
		glam.Log("ERROR: %s\n", err)
	}
}

//------------------------------------------------------------------------------

// OpenGL pipeline object
var (
	pipeline gfx.Pipeline
)

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
	vs, err := os.Open(glam.Path() + "shader.vert")
	if err != nil {
		return glam.Error("unable to open vertex shader", err)
	}
	fs, err := os.Open(glam.Path() + "shader.frag")
	if err != nil {
		return glam.Error("unable to fragment shader", err)
	}
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(vs),
		gfx.FragmentShader(fs),
	)

	return glam.Error("setup", gfx.Err())
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update(_, _ float64) {
}

func (l looper) Draw(_ float64) {
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	pipeline.Bind()
	gfx.Draw(gfx.Triangles, 0, 3)
	pipeline.Unbind()
}

//------------------------------------------------------------------------------
