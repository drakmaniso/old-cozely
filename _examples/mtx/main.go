// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"bufio"
	"os"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/window"
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

	glam.Update = update
	glam.Draw = draw
	window.Handle = handler{}
	mouse.Handle = handler{}

	err = glam.Loop()
	if err != nil {
		glam.ShowError("running", err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline    *gfx.Pipeline
	perFrameUBO gfx.UniformBuffer
)

// Uniform buffer
var perFrame struct {
	viewProjection space.Matrix
	time           float32
}

// Vertex buffer
type mesh []struct {
	position space.Coord `layout:"0"`
	color    color.RGB   `layout:"1"`
}

// Matrices
var (
	view       space.Matrix
	projection space.Matrix
)

// State
var (
	file    *os.File
	scanner *bufio.Scanner
)

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
	pipeline = gfx.NewPipeline(
		gfx.Shader(glam.Path()+"shader.vert"),
		gfx.Shader(glam.Path()+"shader.frag"),
		gfx.VertexFormat(0, mesh{}),
		gfx.Topology(gfx.Triangles),
		gfx.CullFace(false, true),
		gfx.DepthTest(true),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)

	// Create and fill the vertex buffer
	vbo := gfx.NewVertexBuffer(cube(), 0)

	// Initialize model and view matrices
	updateView()

	// MTX
	mtx.Color(color.RGB{1.00, 0.98, 0.89}, color.RGB{0.0, 0.0, 0.0})
	mtx.Opaque(false)
	mtx.ShowFrameTime(true, -1, 0, false)

	// File
	file, err := os.Open(glam.Path() + "main.go")
	if err != nil {
		return glam.Error("opening text file", err)
	}
	scanner = bufio.NewScanner(file)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	return glam.Error("gfx", gfx.Err())
}

//------------------------------------------------------------------------------

func update(dt, _ float64) {
	perFrame.time += float32(dt)

	timer += dt
	if timer < 0.1 {
		return
	}

	timer = 0

	if !scanner.Scan() {
		file.Close()
		file, err := os.Open(glam.Path() + "main.go")
		if err == nil {
			scanner = bufio.NewScanner(file)
		}
	}
	clip1.Print("\n%s", scanner.Text())

	clip2.Print("%c", ' '+incr%('~'-' '))
	incr++
}

var clip1 = mtx.Clip{
	Left: 1, Top: 2,
	Right: -20, Bottom: -1,
	VScroll: true,
}

var clip2 = mtx.Clip{
	Left: 0, Top: 0,
	Right: -7, Bottom: 0,
	HScroll: true,
}

var timer float64
var incr int

func draw() {
	pipeline.Bind()
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.05, 0.10, 0.11, 1.0})

	perFrame.viewProjection = projection.Times(view)
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	gfx.DrawInstanced(0, 6*2*3, 28*1)

	pipeline.Unbind()
}

//------------------------------------------------------------------------------

func updateView() {
	view = space.LookAt(
		space.Coord{0, 0, 10},
		space.Coord{0, 0, 0}, space.Coord{0, 1, 0},
	)
}

//------------------------------------------------------------------------------
