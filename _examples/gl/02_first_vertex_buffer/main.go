// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

// OpenGL objects
var (
	pipeline *gl.Pipeline
)

// Vertex buffer
type mesh []struct {
	position coord.XY   `layout:"0"`
	color    color.LRGB `layout:"1"`
}

////////////////////////////////////////////////////////////////////////////////

func main() {
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
	var triangle mesh

	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader.vert"),
		gl.Shader(cozely.Path()+"shader.frag"),
		gl.VertexFormat(0, triangle),
		gl.Topology(gl.Triangles),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create and fill the vertex buffer
	triangle = mesh{
		{coord.XY{0, 0.65}, color.LRGB{R: 0.3, G: 0, B: 0.8}},
		{coord.XY{-0.65, -0.475}, color.LRGB{R: 0.8, G: 0.3, B: 0}},
		{coord.XY{0.65, -0.475}, color.LRGB{R: 0, G: 0.6, B: 0.2}},
	}
	vbo := gl.NewVertexBuffer(triangle, gl.StaticStorage)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	return cozely.Error("gl", gl.Err())
}

func (loop) Leave() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) React() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() error {
	pipeline.Bind()
	gl.ClearColorBuffer(color.LRGBA{0.9, 0.9, 0.9, 1.0})

	gl.Draw(0, 3)
	pipeline.Unbind()

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func resize() {
	s := cozely.WindowSize()
	gl.Viewport(0, 0, int32(s.C), int32(s.R))
}

////////////////////////////////////////////////////////////////////////////////
