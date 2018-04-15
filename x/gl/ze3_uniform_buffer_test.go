// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/plane"
	"github.com/cozely/cozely/x/gl"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop03 struct {
	// OpenGL objects
	pipeline    *gl.Pipeline
	perFrameUBO gl.UniformBuffer
	// Animation state
	angle float64
}

// Uniform buffer
type perFrame struct {
	transform plane.GPUMatrix
}

// Vertex buffer
type coloredmesh2d []struct {
	position coord.XY   `layout:"0"`
	color    color.LRGB `layout:"1"`
}

// Initialization //////////////////////////////////////////////////////////////

func Example_uniformBuffer() {
	cozely.Events.Resize = func() {
		s := cozely.WindowSize()
		gl.Viewport(0, 0, int32(s.C), int32(s.R))
	}
	l := loop03{}
	err := cozely.Run(&l)
	if err != nil {
		cozely.ShowError(err)
		return
	}
	//Output:
}

func (l *loop03) Enter() error {
	var triangle coloredmesh2d

	// Create and configure the pipeline
	l.pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader03.vert"),
		gl.Shader(cozely.Path()+"shader03.frag"),
		gl.VertexFormat(0, triangle),
		gl.Topology(gl.Triangles),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create the uniform buffer
	l.perFrameUBO = gl.NewUniformBuffer(&perFrame{}, gl.DynamicStorage)

	// Fill and create the vertex buffer
	triangle = coloredmesh2d{
		{coord.XY{0, 0.75}, color.LRGB{0.3, 0, 0.8}},
		{coord.XY{-0.65, -0.465}, color.LRGB{0.8, 0.3, 0}},
		{coord.XY{0.65, -0.465}, color.LRGB{0, 0.6, 0.2}},
	}
	vbo := gl.NewVertexBuffer(triangle, gl.StaticStorage)

	// Bind the vertex buffer to the pipeline
	l.pipeline.Bind()
	vbo.Bind(0, 0)
	l.pipeline.Unbind()

	return cozely.Error("gl", gl.Err())
}

func (loop03) Leave() error {
	return nil
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop03) React() error {
	return nil
}

func (loop03) Update() error {
	return nil
}

func (l *loop03) Render() error {
	l.angle -= 1.0 * cozely.RenderTime()

	l.pipeline.Bind()
	gl.ClearColorBuffer(color.LRGBA{0.9, 0.9, 0.9, 1.0})

	u := perFrame{
		transform: plane.Rotation(float32(l.angle)).GPU(),
	}
	l.perFrameUBO.SubData(&u, 0)
	l.perFrameUBO.Bind(0)

	gl.Draw(0, 3)
	l.pipeline.Unbind()

	return gl.Err()
}
