// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"math/rand"

	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/core/math32"
	"github.com/drakmaniso/carol/mouse"
	"github.com/drakmaniso/carol/plane"
)

//------------------------------------------------------------------------------

func main() {
	err := carol.Run(loop{})
	if err != nil {
		carol.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline    *gl.Pipeline
	perFrameUBO gl.UniformBuffer
	pointsVBO   gl.VertexBuffer
)

// Uniform buffer
var perFrame struct {
	ratio    plane.Coord
	Rotation float32
}

// Vertex buffer
var points [512]struct {
	Position plane.Coord `layout:"0"`
}

// Application State
var (
	bgColor  = colour.RGBA{0.9, 0.87, 0.85, 1.0}
	rotSpeed = float32(0.003)
	jitter   = float32(0.002)
	angles   []float32
	speeds   []float32
)

//------------------------------------------------------------------------------

type loop struct {
	carol.Handlers
}

//------------------------------------------------------------------------------

func (loop) Setup() error {
	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(carol.Path()+"shader.vert"),
		gl.Shader(carol.Path()+"shader.frag"),
		gl.Topology(gl.Points),
		gl.VertexFormat(0, points[:]),
	)
	gl.Enable(gl.FramebufferSRGB)
	gl.Enable(gl.Blend)
	gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)
	gl.PointSize(3.0)

	// Create the uniform buffer
	perFrameUBO = gl.NewUniformBuffer(&perFrame, gl.DynamicStorage)
	perFrame.Rotation = 0.0

	// Create and fill the vertex buffer
	// points = make(mesh, len(points))
	angles = make([]float32, len(points))
	speeds = make([]float32, len(points))
	setupPoints()
	pointsVBO = gl.NewVertexBuffer(points[:], gl.DynamicStorage)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	pointsVBO.Bind(0, 0)
	pipeline.Unbind()

	return carol.Error("gl", gl.Err())
}

//------------------------------------------------------------------------------

func (loop) Update() error {
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

	return nil
}

var updated bool

//------------------------------------------------------------------------------

var cleared bool

func (loop) Draw(_, _ float64) error {
	if !cleared {
		gl.ClearColorBuffer(bgColor)
		cleared = true
	}
	if updated {
		pipeline.Bind()

		perFrameUBO.Bind(0)
		perFrameUBO.SubData(&perFrame, 0)

		gl.Draw(0, int32(len(points)))

		pipeline.Unbind()
		updated = false
	}

	return nil
}

//------------------------------------------------------------------------------

func setupPoints() {
	n := float32(3 + rand.Intn(13))
	for i := range points {
		points[i].Position = plane.Coord{rand.Float32(), rand.Float32()}
		a := math32.Floor(rand.Float32() * n)
		a = a * (2.0 * math32.Pi) / n
		points[i].Position = plane.Coord{0.75 * math32.Cos(a), 0.75 * math32.Sin(a)}
		angles[i] = a + float32(i)*math32.Pi/float32(len(points)) + math32.Pi/2.0
		speeds[i] = 0.004 * rand.Float32()
	}
	rotSpeed = 0.01 * (rand.Float32() - 0.5)
	jitter = 0.014 * rand.Float32()
	cleared = false
}

//------------------------------------------------------------------------------

func (loop) WindowResized(w, h int32) {
	setupPoints()

	// Compute ratio
	if w > h {
		perFrame.ratio = plane.Coord{float32(h) / float32(w), 1.0}
	} else {
		perFrame.ratio = plane.Coord{1.0, float32(w) / float32(h)}
	}
	gl.Viewport(0, 0, w, h)
}

func (loop) MouseButtonDown(b mouse.Button, _ int) {
	setupPoints()
}

//------------------------------------------------------------------------------
