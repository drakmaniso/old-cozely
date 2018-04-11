// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"math/rand"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/colour"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/plane"
	"github.com/cozely/cozely/x/gl"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quit      = input.Bool("Quit")
	randomize = input.Bool("Randomize")
)

var context = input.Context("Default", quit, randomize)

var bindings = input.Bindings{
	"Default": {
		"Quit":      {"Escape"},
		"Randomize": {"Space", "Mouse Left"},
	},
}

////////////////////////////////////////////////////////////////////////////////

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
	bgColor  = colour.LRGBA{0.9, 0.87, 0.85, 1.0}
	rotSpeed = float32(0.003)
	jitter   = float32(0.002)
	angles   []float32
	speeds   []float32
)

var cleared bool

////////////////////////////////////////////////////////////////////////////////

func main() {
	angles = make([]float32, len(points))
	speeds = make([]float32, len(points))

	cozely.Configure(cozely.Multisample(8))
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
	input.Load(bindings)
	context.Activate(1)

	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader.vert"),
		gl.Shader(cozely.Path()+"shader.frag"),
		gl.Topology(gl.Points),
		gl.VertexFormat(0, points[:]),
		gl.DepthTest(false),
		gl.DepthWrite(false),
	)

	// Create the uniform buffer
	perFrameUBO = gl.NewUniformBuffer(&perFrame, gl.DynamicStorage)
	perFrame.Rotation = 0.0

	// Create and fill the vertex buffer
	// points = make(mesh, len(points))
	setupPoints()
	pointsVBO = gl.NewVertexBuffer(points[:], gl.DynamicStorage)

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	pointsVBO.Bind(0, 0)
	pipeline.Unbind()

	return cozely.Error("gl", gl.Err())
}

func (loop) Leave() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}

	if randomize.JustPressed(1) {
		setupPoints()
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() error {
	if !cleared {
		gl.ClearColorBuffer(bgColor)
		cleared = true
	}
	if updated {
		pipeline.Bind()
		gl.Enable(gl.FramebufferSRGB)
		gl.Enable(gl.Blend)
		gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)
		gl.PointSize(3.0)

		perFrameUBO.Bind(0)
		perFrameUBO.SubData(&perFrame, 0)

		gl.Draw(0, int32(len(points)))

		pipeline.Unbind()
		updated = false
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////

func resize() {
	setupPoints()

	s := cozely.WindowSize().Cartesian()

	// Compute ratio
	if s.X > s.Y {
		perFrame.ratio = plane.Coord{s.Y / s.X, 1.0}
	} else {
		perFrame.ratio = plane.Coord{1.0, s.X / s.Y}
	}
	gl.Viewport(0, 0, int32(s.X), int32(s.Y))
}

////////////////////////////////////////////////////////////////////////////////
