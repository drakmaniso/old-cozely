// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"math/rand"

	"github.com/drakmaniso/cozely/input"

	"github.com/drakmaniso/cozely"
	"github.com/drakmaniso/cozely/colour"
	"github.com/drakmaniso/cozely/plane"
	"github.com/drakmaniso/cozely/x/gl"
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
	rosesINBO   gl.VertexBuffer
)

// Uniform buffer
var perFrame struct {
	ratio float32
	time  float32
}

// Instance Buffer

var roses [64]struct {
	position    plane.Coord `layout:"0" divisor:"1"`
	size        float32     `layout:"1"`
	numerator   int32       `layout:"2"`
	denominator int32       `layout:"3"`
	offset      float32     `layout:"4"`
	speed       float32     `layout:"5"`
}

const nbPoints int32 = 512

////////////////////////////////////////////////////////////////////////////////

func main() {
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

	// Setup the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader.vert"),
		gl.Shader(cozely.Path()+"shader.frag"),
		gl.VertexFormat(1, roses[:]),
		gl.Topology(gl.LineStrip),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gl.NewUniformBuffer(&perFrame, gl.DynamicStorage)

	// Create the instance buffer
	randomizeRosesData()
	rosesINBO = gl.NewVertexBuffer(roses[:], gl.DynamicStorage)

	// Bind the instance buffer to the pipeline
	pipeline.Bind()
	rosesINBO.Bind(1, 0)
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
		randomizeRosesData()
		rosesINBO.SubData(roses[:], 0)
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() error {
	perFrame.time += float32(cozely.UpdateTime())

	pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(colour.LRGBA{0.9, 0.85, 0.80, 1.0})
	gl.Enable(gl.LineSmooth)
	gl.Enable(gl.Blend)
	gl.Blending(gl.SrcAlpha, gl.OneMinusSrcAlpha)

	perFrameUBO.Bind(0)
	perFrameUBO.SubData(&perFrame, 0)
	gl.DrawInstanced(0, nbPoints, int32(len(roses)))

	pipeline.Unbind()

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func resize() {
	s := cozely.WindowSize()
	perFrame.ratio = float32(s.X) / float32(s.Y)
	gl.Viewport(0, 0, int32(s.X), int32(s.Y))
}

////////////////////////////////////////////////////////////////////////////////

func randomizeRosesData() {
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
// 	// var m = []perVertex{{plane.Coord{0.0, 0.0}, colour.LRGB{0.9, 0.9, 0.9}}}
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

////////////////////////////////////////////////////////////////////////////////
