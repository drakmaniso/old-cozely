// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/x/gl"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop02 struct {
	pipeline *gl.Pipeline
}

// Vertex buffer
type mesh2d []struct {
	position coord.XY   `layout:"0"`
	color    color.LRGB `layout:"1"`
}

// Initialization //////////////////////////////////////////////////////////////

func Example_vertexBuffer() {
	defer cozely.Recover()

	cozely.Events.Resize = func() {
		s := cozely.WindowSize()
		gl.Viewport(0, 0, int32(s.X), int32(s.Y))
	}
	l := loop02{}
	err := cozely.Run(&l)
	if err != nil {
		panic(err)
	}
	//Output:
}

func (l *loop02) Enter() {
	var triangle mesh2d

	// Create and configure the pipeline
	l.pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader02.vert"),
		gl.Shader(cozely.Path()+"shader02.frag"),
		gl.VertexFormat(0, triangle),
		gl.Topology(gl.Triangles),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create and fill the vertex buffer
	triangle = mesh2d{
		{coord.XY{0, 0.65}, color.LRGB{R: 0.3, G: 0, B: 0.8}},
		{coord.XY{-0.65, -0.475}, color.LRGB{R: 0.8, G: 0.3, B: 0}},
		{coord.XY{0.65, -0.475}, color.LRGB{R: 0, G: 0.6, B: 0.2}},
	}
	vbo := gl.NewVertexBuffer(triangle, gl.StaticStorage)

	// Bind the vertex buffer to the pipeline
	l.pipeline.Bind()
	vbo.Bind(0, 0)
	l.pipeline.Unbind()
}

func (loop02) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop02) React() {
}

func (loop02) Update() {
}

func (l *loop02) Render() {
	l.pipeline.Bind()
	gl.ClearColorBuffer(color.LRGBA{0.9, 0.9, 0.9, 1.0})

	gl.Draw(0, 3)
	l.pipeline.Unbind()
}
