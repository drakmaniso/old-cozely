// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"math/rand"
	"os"
	"time"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
	err := setup()
	if err != nil {
		glam.ErrorDialog(err)
		return
	}

	window.Handle = handler{}
	mouse.Handle = handler{}

	// Run the main loop
	glam.Loop = looper{}
	err = glam.Run()
	if err != nil {
		glam.ErrorDialog(err)
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline    gfx.Pipeline
	perFrameUBO gfx.UniformBuffer
	rosesINBO   gfx.VertexBuffer
)

// Uniform buffer
var perFrame struct {
	ratio float32
	time  float32
}

// Instance Buffer

type perInstance struct {
	position    plane.Coord `layout:"0" divisor:"1"`
	size        float32     `layout:"1"`
	numerator   int32       `layout:"2"`
	denominator int32       `layout:"3"`
	offset      float32     `layout:"4"`
	speed       float32     `layout:"5"`
}

//------------------------------------------------------------------------------

func setup() error {
	// Setup the pipeline
	var v, err = os.Open(glam.Path() + "shader.vert")
	if err != nil {
		return err
	}
	f, err := os.Open(glam.Path() + "shader.frag")
	if err != nil {
		return err
	}
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(v),
		gfx.FragmentShader(f),
		gfx.VertexFormat(1, perInstance{}),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)

	// Create the instance buffer
	rosesINBO = gfx.NewVertexBuffer(generateInstances(), gfx.DynamicStorage)

	// Bind the instance buffer to the pipeline
	pipeline.Bind()
	rosesINBO.Bind(1, 0)
	pipeline.Unbind()

	return gfx.Err()
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update() {
	perFrame.time += float32(glam.TimeStep) / float32(time.Second)
}

func (l looper) Draw() {
	pipeline.Bind()
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.85, 0.80, 1.0})

	perFrameUBO.Bind(0)
	perFrameUBO.SubData(&perFrame, 0)
	gfx.DrawInstanced(gfx.LineStrip, 0, nbPoints, int32(nbInstances))

	pipeline.Unbind()
}

//------------------------------------------------------------------------------

const nbPoints int32 = 512
const nbInstances int = 64

func generateInstances() []perInstance {
	var data = []perInstance{}
	for i := 0; i < nbInstances; i++ {
		var inst = perInstance{}
		inst.position.X = rand.Float32()*2.0 - 1.0
		inst.position.Y = rand.Float32()*2.0 - 1.0
		inst.size = rand.Float32()*0.20 + 0.1
		inst.numerator = rand.Int31n(16) + 1
		inst.denominator = rand.Int31n(16) + 1
		inst.offset = rand.Float32()*2.8 + 0.2
		inst.speed = 0.5 + 1.5*rand.Float32()
		if rand.Int31n(2) > 0 {
			inst.speed = -inst.speed
		}
		data = append(data, inst)
	}
	return data
}

//------------------------------------------------------------------------------

// func rose(nbPoints int, num int, den int, offset float32) []perVertex {
// 	// var m = []perVertex{{plane.Coord{0.0, 0.0}, color.RGB{0.9, 0.9, 0.9}}}
// 	var m = []perVertex{}
// 	for i := den * nbPoints; i >= 0; i-- {
// 		var k = float32(num) / float32(den)
// 		var theta = float32(i) * 2 * math.Pi / float32(nbPoints)
// 		var r = (math.Cos(k*theta) + offset) / (1.0 + offset)
// 		var p = plane.Polar{r, theta}
// 		m = append(m, perVertex{p.Coord()})
// 	}
// 	return m
// }

//------------------------------------------------------------------------------

type handler struct {
	basic.WindowHandler
	basic.MouseHandler
}

func (h handler) WindowResized(s pixel.Coord, _ time.Duration) {
	var sx, sy = window.Size().Cartesian()
	perFrame.ratio = sy / sx
}

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ time.Duration) {
	rosesINBO.SubData(generateInstances(), 0)
}

//------------------------------------------------------------------------------
