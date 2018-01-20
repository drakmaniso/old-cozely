// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/colour"
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

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(glam.Path()+"shader.vert"),
		gl.Shader(glam.Path()+"shader.frag"),
		gl.Topology(gl.Triangles),
	)

	return glam.Error("gfx", gl.Err())
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (l loop) WindowResized(w, h int32) {
	gl.Viewport(0, 0, w, h)
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw() error {
	pipeline.Bind()
	gl.ClearColorBuffer(colour.LRGBA{0.9, 0.9, 0.9, 1.0})

	gl.Draw(0, 3)
	pipeline.Unbind()

	return gl.Err()
}

//------------------------------------------------------------------------------
