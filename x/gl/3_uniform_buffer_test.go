// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/plane"
	"github.com/cozely/cozely/window"
	"github.com/cozely/cozely/x/gl"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop3 struct {
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
	defer cozely.Recover()

	window.Events.Resize = func() {
		s := window.Size()
		gl.Viewport(0, 0, int32(s.X), int32(s.Y))
	}
	l := loop3{}
	err := cozely.Run(&l)
	if err != nil {
		panic(err)
	}
	//Output:
}

func (l *loop3) Enter() {
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
}

func (loop3) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop3) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}
}

func (loop3) Update() {
}

func (l *loop3) Render() {
	l.angle -= 1.0 * cozely.RenderDelta()

	l.pipeline.Bind()
	gl.ClearColorBuffer(color.LRGBA{0.9, 0.9, 0.9, 1.0})

	u := perFrame{
		transform: plane.Rotation(float32(l.angle)).GPU(),
	}
	l.perFrameUBO.SubData(&u, 0)
	l.perFrameUBO.Bind(0)

	gl.Draw(0, 3)
	l.pipeline.Unbind()
}
