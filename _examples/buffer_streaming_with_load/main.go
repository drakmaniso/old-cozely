// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"math/rand"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/math32"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
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

	glam.Loop(loop{})

	err = glam.Run()
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
	ratio    plane.Coord
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
	pipeline = gfx.NewPipeline(
		gfx.Shader(glam.Path()+"shader.vert"),
		gfx.Shader(glam.Path()+"shader.frag"),
		gfx.Topology(gfx.Points),
		gfx.VertexFormat(0, points[:]),
	)
	gfx.Enable(gfx.FramebufferSRGB)
	gfx.Enable(gfx.Blend)
	gfx.Blending(gfx.SrcAlpha, gfx.OneMinusSrcAlpha)
	gfx.ClearColorBuffer(bgColor)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)
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

	return glam.Error("gfx", gfx.Err())
}

//------------------------------------------------------------------------------
type loop struct {
	glam.DefaultHandlers
}

func (loop) Update() {
	for i, pt := range points {
		points[i].Position = plane.Coord{
			pt.Position.X + speeds[i]*math32.Cos(angles[i]) + jitter*(rand.Float32()-0.5),
			pt.Position.Y + speeds[i]*math32.Sin(angles[i]) + jitter*(rand.Float32()-0.5),
		}
		if points[i].Position.Length() > 0.75 {
			angles[i] += math32.Pi / 4.0
		}
	}
	pointsVBO.SubData(points[:], 0)

	perFrame.Rotation += rotSpeed

	updated = true
}

var updated bool

func (loop) Draw(_, _ float64) {
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
		a := math32.Floor(rand.Float32() * n)
		a = a * (2.0 * math32.Pi) / n
		points[i].Position = plane.Coord{0.75 * math32.Cos(a), 0.75 * math32.Sin(a)}
		angles[i] = a + float32(i)*math32.Pi/float32(len(points)) + math32.Pi/2.0
		speeds[i] = 0.004 * rand.Float32()
	}
	rotSpeed = 0.01 * (rand.Float32() - 0.5)
	jitter = 0.006*rand.Float32() - 0.001
	if jitter < 0.0 {
		jitter = 0.0
	}
}

//------------------------------------------------------------------------------

func (loop) WindowResized(sp pixel.Coord) {
	gfx.ClearColorBuffer(bgColor)

	setupPoints()

	// Compute ratio
	s := plane.CoordOf(sp)
	if s.X > s.Y {
		perFrame.ratio = plane.Coord{s.Y / s.X, 1.0}
	} else {
		perFrame.ratio = plane.Coord{1.0, s.X / s.Y}
	}
}

func (loop) MouseButtonDown(b mouse.Button, _ int) {
	gfx.ClearColorBuffer(bgColor)
	setupPoints()
}

//------------------------------------------------------------------------------
