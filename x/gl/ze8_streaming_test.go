// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl_test

import (
	"math/rand"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/x/gl"
	"github.com/cozely/cozely/x/math32"
)

// Declarations ////////////////////////////////////////////////////////////////

// Input Bindings
// (same as in InstancedDraw example)

type loop08 struct {
	// OpenGL objects
	pipeline    *gl.Pipeline
	perFrameUBO gl.UniformBuffer
	pointsVBO   gl.VertexBuffer
	perFrame    perFrame08

	// Application State
	bgColor  color.LRGBA
	rotSpeed float32
	jitter   float32
	angles   []float32
	speeds   []float32

	cleared, updated bool
}

// Uniform buffer
type perFrame08 struct {
	ratio    coord.XY
	Rotation float32
}

// Vertex buffer
var points [512]struct {
	Position coord.XY `layout:"0"`
}

// Initialization //////////////////////////////////////////////////////////////

func Example_streaming() {
	defer cozely.Recover()

	l := loop08{
		bgColor:  color.LRGBA{0.9, 0.87, 0.85, 1.0},
		rotSpeed: float32(0.003),
		jitter:   float32(0.002),
		angles:   make([]float32, len(points)),
		speeds:   make([]float32, len(points)),
	}

	cozely.Configure(cozely.Multisample(8), gl.NoClear())
	cozely.Events.Resize = l.resize
	err := cozely.Run(&l)
	if err != nil {
		panic(err)
	}
	//Output:
}

func (l *loop08) Enter() {
	input.Load(bindings06)
	context06.Activate(1)

	// Create and configure the pipeline
	l.pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader08.vert"),
		gl.Shader(cozely.Path()+"shader08.frag"),
		gl.Topology(gl.Points),
		gl.VertexFormat(0, points[:]),
		gl.DepthTest(false),
		gl.DepthWrite(false),
	)

	// Create the uniform buffer
	l.perFrameUBO = gl.NewUniformBuffer(&perFrame08{}, gl.DynamicStorage)
	l.perFrame.Rotation = 0.0

	// Create and fill the vertex buffer
	// points = make(mesh, len(points))
	l.setupPoints()
	l.pointsVBO = gl.NewVertexBuffer(points[:], gl.DynamicStorage)

	// Bind the vertex buffer to the pipeline
	l.pipeline.Bind()
	l.pointsVBO.Bind(0, 0)
	l.pipeline.Unbind()
}

func (loop08) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (l *loop08) React() {
	if randomize.Started(1) {
		l.setupPoints()
	}

	if quit.Started(1) {
		cozely.Stop(nil)
	}
}

func (l *loop08) Update() {
	for i, pt := range points {
		points[i].Position = coord.XY{
			pt.Position.X +
				l.speeds[i]*math32.Cos(l.angles[i]) +
				l.jitter*(rand.Float32()-0.5),
			pt.Position.Y +
				l.speeds[i]*math32.Sin(l.angles[i]) +
				l.jitter*(rand.Float32()-0.5),
		}
		if points[i].Position.Length() > 0.75 {
			l.angles[i] += math32.Pi / 4.0
		}
	}
	l.pointsVBO.SubData(points[:], 0)

	l.perFrame.Rotation += l.rotSpeed

	l.updated = true
}

func (l *loop08) Render() {
	if !l.cleared {
		gl.ClearColorBuffer(l.bgColor)
		l.cleared = true
	}
	if l.updated {
		l.pipeline.Bind()
		gl.Enable(gl.FramebufferSRGB)
		gl.Enable(gl.Blend)
		gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)
		gl.PointSize(3.0)

		u := l.perFrame
		l.perFrameUBO.Bind(0)
		l.perFrameUBO.SubData(&u, 0)

		gl.Draw(0, int32(len(points)))

		l.pipeline.Unbind()
		l.updated = false
	}
}

func (l *loop08) setupPoints() {
	n := float32(3 + rand.Intn(13))
	for i := range points {
		points[i].Position = coord.XY{rand.Float32(), rand.Float32()}
		a := math32.Floor(rand.Float32() * n)
		a = a * (2.0 * math32.Pi) / n
		points[i].Position = coord.XY{0.75 * math32.Cos(a), 0.75 * math32.Sin(a)}
		l.angles[i] = a + float32(i)*math32.Pi/float32(len(points)) + math32.Pi/2.0
		l.speeds[i] = 0.004 * rand.Float32()
	}
	l.rotSpeed = 0.01 * (rand.Float32() - 0.5)
	l.jitter = 0.014 * rand.Float32()
	l.cleared = false
}

func (l *loop08) resize() {
	l.setupPoints()

	s := cozely.WindowSize().XY()

	// Compute ratio
	if s.X > s.Y {
		l.perFrame.ratio = coord.XY{s.Y / s.X, 1.0}
	} else {
		l.perFrame.ratio = coord.XY{1.0, s.X / s.Y}
	}
	gl.Viewport(0, 0, int32(s.X), int32(s.Y))
}
