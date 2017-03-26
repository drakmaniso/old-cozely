// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"math/rand"
	"os"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
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

	err = glam.LoopStable(1.0 / 60.0)
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
	pointsVBO   gfx.VertexBuffer
)

// Uniform buffer
var perFrame struct {
	Scale    plane.Coord
	Rotation float32
}

// Vertex buffer
var points [4096]struct {
	Position plane.Coord `layout:"0"`
}

// Application State
var (
	bgColor  = color.RGBA{0.9, 0.87, 0.85, 1.0}
	rotSpeed = float32(0.003)
	jitter   = float32(0.002)
	angles   []float32
	speeds   []float32
)

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
	v, err := os.Open(glam.Path() + "shader.vert")
	if err != nil {
		return glam.Error("opening vertex shader", err)
	}
	f, err := os.Open(glam.Path() + "shader.frag")
	if err != nil {
		return glam.Error("opening fragment shader", err)
	}
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(v),
		gfx.FragmentShader(f),
		gfx.Topology(gfx.Points),
		gfx.VertexFormat(0, points[:]),
	)
	gfx.Enable(gfx.FramebufferSRGB)
	gfx.Enable(gfx.Blend)
	gfx.Blending(gfx.SrcAlpha, gfx.OneMinusSrcAlpha)
	gfx.ClearColorBuffer(bgColor)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)
	updateView()
	perFrame.Rotation = 0.0

	// Create and fill the vertex buffer
	// points = make(mesh, len(points))
	angles = make([]float32, len(points))
	speeds = make([]float32, len(points))
	setupPoints()
	pointsVBO = gfx.NewVertexBuffer(points[:], gfx.DynamicStorage)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	pointsVBO.Bind(0, 0)
	pipeline.Unbind()

	// MTX
	mtx.Color(color.RGB{0.1, 0.05, 0.0}, color.RGB{1, 1, 1})

	return glam.Error("gfx", gfx.Err())
}

//------------------------------------------------------------------------------

var updated bool

func update(_, _ float64) {
	for i, pt := range points {
		points[i].Position = plane.Coord{
			pt.Position.X + speeds[i]*math.Cos(angles[i]) + jitter*(rand.Float32()-0.5),
			pt.Position.Y + speeds[i]*math.Sin(angles[i]) + jitter*(rand.Float32()-0.5),
		}
		if points[i].Position.Length() > 0.75 {
			angles[i] += math.Pi / 4.0
		}
	}
	pointsVBO.SubData(points[:], 0)

	perFrame.Rotation += rotSpeed

	updated = true
}

func draw() {
	if updated {
		pipeline.Bind()

		perFrameUBO.Bind(0)
		perFrameUBO.SubData(&perFrame, 0)

		gfx.Draw(0, int32(len(points)))

		pipeline.Unbind()
		updated = false
	}
}

//------------------------------------------------------------------------------

func setupPoints() {
	n := float32(6 + rand.Intn(60))
	for i := range points {
		points[i].Position = plane.Coord{rand.Float32(), rand.Float32()}
		a := math.Floor(rand.Float32() * n)
		a = a * (2.0 * math.Pi) / n
		points[i].Position = plane.Coord{0.75 * math.Cos(a), 0.75 * math.Sin(a)}
		angles[i] = a + float32(i)*math.Pi/float32(len(points)) + math.Pi/2.0
		speeds[i] = 0.004 * rand.Float32()
	}
	rotSpeed = 0.01 * (rand.Float32() - 0.5)
	jitter = 0.006*rand.Float32() - 0.001
	if jitter < 0.0 {
		jitter = 0.0
	}
}

//------------------------------------------------------------------------------

// Event handler
type handler struct {
	basic.WindowHandler
	basic.MouseHandler
}

func (h handler) WindowResized(s pixel.Coord, _ uint32) {
	gfx.ClearColorBuffer(bgColor)
	setupPoints()
	updateView()
}

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ uint32) {
	gfx.ClearColorBuffer(bgColor)
	setupPoints()
}

//------------------------------------------------------------------------------

func updateView() {
	sx, sy := window.Size().Cartesian()
	if sx > sy {
		perFrame.Scale = plane.Coord{sy / sx, 1.0}
	} else {
		perFrame.Scale = plane.Coord{1.0, sx / sy}
	}
	pipeline.Bind()
}

//------------------------------------------------------------------------------
