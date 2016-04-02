// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"log"
	"strings"
	"time"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	. "github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/geom/space"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

type perVertex struct {
	position Vec3      `layout:"0"`
	color    color.RGB `layout:"1"`
}

type perObject struct {
	transform Mat4
}

//------------------------------------------------------------------------------

var vertexShader = strings.NewReader(`
#version 450 core

layout(location = 0) in vec3 Position;
layout(location = 1) in vec3 Color;

layout(std140, binding = 0) uniform PerObject {
	mat4 Transform;
} obj;

out gl_PerVertex {
	vec4 gl_Position;
};

out PerVertex {
	layout(location = 0) out vec3 Color;
} vert;

void main(void) {
	gl_Position = obj.Transform * vec4(Position, 1);
	vert.Color = Color;
}
`)

var fragmentShader = strings.NewReader(`
#version 450 core

in PerVertex {
	layout(location = 0) in vec3 Color;
} vert;

out vec4 Color;

void main(void) {
	Color = vec4(vert.Color, 1);
}
`)

//------------------------------------------------------------------------------

func main() {
	log.SetFlags(log.Lshortfile)

	g := &game{distance: 3, yaw: -0.6, pitch: 0.3}
	glam.Handler = g
	window.Handler = g
	mouse.Handler = g

	var err error

	// Setup the Pipeline
	vs, err := gfx.NewVertexShader(vertexShader)
	if err != nil {
		log.Fatal(err)
	}
	fs, err := gfx.NewFragmentShader(fragmentShader)
	if err != nil {
		log.Fatal(err)
	}
	g.pipeline, err = gfx.NewPipeline(vs, fs)
	if err != nil {
		log.Fatal(err)
	}
	err = g.pipeline.VertexFormat(0, perVertex{})
	if err != nil {
		log.Fatal(err)
	}
	g.pipeline.ClearColor(Vec4{0.9, 0.9, 0.9, 1.0})

	// Create the Uniform Buffer
	g.transform, err = gfx.NewBuffer(uintptr(64), gfx.DynamicStorage)
	if err != nil {
		log.Fatal(err)
	}

	// Create the Vertex Buffer
	data := []perVertex{
		// Front Face
		{Vec3{0, 0, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 1, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{0, 1, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{0, 0, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 0, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 1, 1}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		// Back Face
		{Vec3{0, 0, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{0, 1, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 1, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{0, 0, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 1, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		{Vec3{1, 0, 0}, color.RGB{R: 0.3, G: 0, B: 0.8}},
		// Right Face
		{Vec3{1, 0, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{1, 1, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{1, 1, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{1, 0, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{1, 0, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{1, 1, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		// Left Face
		{Vec3{0, 0, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{0, 1, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{0, 1, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{0, 0, 1}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{0, 1, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		{Vec3{0, 0, 0}, color.RGB{R: 0.8, G: 0.3, B: 0}},
		// Bottom Face
		{Vec3{0, 0, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{0, 0, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 0, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{0, 0, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 0, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 0, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		// Top Face
		{Vec3{0, 1, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 1, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{0, 1, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{0, 1, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 1, 1}, color.RGB{R: 0, G: 0.6, B: 0.2}},
		{Vec3{1, 1, 0}, color.RGB{R: 0, G: 0.6, B: 0.2}},
	}
	g.colorfulTriangle, err = gfx.NewBuffer(data, 0)
	if err != nil {
		log.Fatal(err)
	}

	g.updateModel()
	g.updateView()

	// Run the Game Loop
	err = glam.Run()
	if err != nil {
		log.Fatal(err)
	}
}

//------------------------------------------------------------------------------

type game struct {
	window.DefaultWindowHandler
	mouse.DefaultMouseHandler

	distance                float32
	yaw, pitch              float32
	model, view, projection Mat4

	pipeline         gfx.Pipeline
	transform        gfx.Buffer
	colorfulTriangle gfx.Buffer
}

//------------------------------------------------------------------------------

func (g *game) WindowResized(s IVec2, timestamp time.Duration) {
	r := float32(s.X) / float32(s.Y)
	g.projection = space.Perspective(math.Pi/4, r, 0.001, 1000.0)
}

func (g *game) MouseWheel(motion IVec2, timestamp time.Duration) {
	g.distance -= float32(motion.Y) / 4
	g.updateView()
}

func (g *game) MouseButtonDown(b mouse.Button, clicks int, timestamp time.Duration) {
	mouse.SetRelativeMode(true)
}

func (g *game) MouseButtonUp(b mouse.Button, clicks int, timestamp time.Duration) {
	mouse.SetRelativeMode(false)
}

func (g *game) MouseMotion(motion IVec2, position IVec2, timestamp time.Duration) {
	switch {
	case mouse.IsPressed(mouse.Left):
		g.yaw += 4 * float32(motion.X) / 1280
		g.pitch += 4 * float32(motion.Y) / 720
		switch {
		case g.pitch < -math.Pi/2+0.01:
			g.pitch = -math.Pi/2 + 0.01
		case g.pitch > math.Pi/2-0.01:
			g.pitch = math.Pi/2 - 0.01
		}
		g.updateModel()
	case mouse.IsPressed(mouse.Middle):
		g.distance += 4 * float32(motion.Y) / 720
		g.updateView()
	}
}

//------------------------------------------------------------------------------

func (g *game) updateModel() {

	g.model = space.EulerZXY(g.pitch, g.yaw, 0)
	g.model = g.model.Times(space.Translation(Vec3{-0.5, -0.5, -0.5}))
}

func (g *game) updateView() {
	if g.distance < 1 {
		g.distance = 1
	}
	g.view = space.LookAt(Vec3{0, 0, g.distance}, Vec3{0, 0, 0}, Vec3{0, 1, 0})
}

//------------------------------------------------------------------------------

func (g *game) Update() {
}

func (g *game) Draw() {
	g.pipeline.Bind()
	g.pipeline.UniformBuffer(0, g.transform)

	mvp := g.projection.Times(g.view)
	mvp = mvp.Times(g.model)
	t := perObject{
		transform: mvp,
	}
	g.transform.Update(&t, 0)

	g.pipeline.VertexBuffer(0, g.colorfulTriangle, 0)
	gfx.Draw(gfx.Triangles, 0, 6*2*3)
}

//------------------------------------------------------------------------------
