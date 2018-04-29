// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl_test

import (
	"math/rand"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/x/gl"
)

// Declarations ////////////////////////////////////////////////////////////////

// Input Bindings

var (
	randomize = input.Digital("Randomize")
)

var context06 = input.Context("Default06", quit, randomize)

var bindings06 = input.Bindings{
	"Default06": {
		"Quit":      {"Escape"},
		"Randomize": {"Space", "Mouse Left"},
	},
}

type loop06 struct {
	// OpenGL objects
	pipeline    *gl.Pipeline
	perFrameUBO gl.UniformBuffer
	rosesINBO   gl.VertexBuffer
	perFrame    perFrame06
}

// Instance Buffer
var roses [64]struct {
	position    coord.XY `layout:"0" divisor:"1"`
	size        float32  `layout:"1"`
	numerator   int32    `layout:"2"`
	denominator int32    `layout:"3"`
	offset      float32  `layout:"4"`
	speed       float32  `layout:"5"`
}

const nbPoints int32 = 512

// Uniform buffer
type perFrame06 struct {
	ratio float32
	time  float32
}

// Initialization //////////////////////////////////////////////////////////////

func Example_instancedDraw() {
	defer cozely.Recover()

	cozely.Configure(cozely.Multisample(8))
	l := loop06{}
	cozely.Events.Resize = func() {
		s := cozely.WindowSize()
		l.perFrame.ratio = float32(s.C) / float32(s.R)
		gl.Viewport(0, 0, int32(s.C), int32(s.R))
	}
	err := cozely.Run(&l)
	if err != nil {
		panic(err)
	}
	//Output:
}

func (l *loop06) Enter() {
	input.Load(bindings06)
	context06.Activate(1)

	// Setup the pipeline
	l.pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader06.vert"),
		gl.Shader(cozely.Path()+"shader06.frag"),
		gl.VertexFormat(1, roses[:]),
		gl.Topology(gl.LineStrip),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create the uniform buffer
	l.perFrameUBO = gl.NewUniformBuffer(&perFrame{}, gl.DynamicStorage)

	// Create the instance buffer
	l.randomizeRosesData()
	l.rosesINBO = gl.NewVertexBuffer(roses[:], gl.DynamicStorage)

	// Bind the instance buffer to the pipeline
	l.pipeline.Bind()
	l.rosesINBO.Bind(1, 0)
	l.pipeline.Unbind()
}

func (loop06) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (l *loop06) React() {
	if randomize.Started(1) {
		l.randomizeRosesData()
		l.rosesINBO.SubData(roses[:], 0)
	}

	if quit.Started(1) {
		cozely.Stop(nil)
	}
}

func (loop06) Update() {
}

func (l *loop06) Render() {
	l.perFrame.time += float32(cozely.RenderDelta())

	l.pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(color.LRGBA{0.9, 0.85, 0.80, 1.0})
	gl.Enable(gl.LineSmooth)
	gl.Enable(gl.Blend)
	gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)

	u := l.perFrame
	l.perFrameUBO.Bind(0)
	l.perFrameUBO.SubData(&u, 0)
	gl.DrawInstanced(0, nbPoints, int32(len(roses)))

	l.pipeline.Unbind()
}

func (l *loop06) randomizeRosesData() {
	for i := 0; i < len(roses); i++ {
		roses[i].position.X = rand.Float32()*2.0 - 1.0
		roses[i].position.Y = rand.Float32()*2.0 - 1.0
		roses[i].size = rand.Float32()*0.20 + 0.1
		roses[i].numerator = rand.Int31n(16) + 1
		roses[i].denominator = rand.Int31n(16) + 1
		roses[i].offset = rand.Float32()*2.8 + 0.2
		roses[i].speed = 0.5 + 1.5*rand.Float32()
		if rand.Int31n(2) > 0 {
			roses[i].speed = -roses[i].speed
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

// func rose(nbPoints int, num int, den int, offset float32) []perVertex {
// 	var m = []perVertex{}
// 	for i := den * nbPoints; i >= 0; i-- {
// 		var k = float32(num) / float32(den)
// 		var theta = float32(i) * 2 * math32.Pi / float32(nbPoints)
// 		var r = (math.Cos(k*theta) + offset) / (1.0 + offset)
// 		var p = plane.Polar{r, theta}
// 		m = append(m, perVertex{p.Coord()})
// 	}
// 	return m
// }
