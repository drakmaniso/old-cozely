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

var pipeline gfx.Pipeline

type perVertex struct {
	position Vec3      `layout:"0"`
	color    color.RGB `layout:"1"`
}

type perObject struct {
	transform Mat4
}

var model, view, projection Mat4

var transform gfx.Buffer
var colorfulTriangle gfx.Buffer

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

	g := &game{}
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
	pipeline, err = gfx.NewPipeline(vs, fs)
	if err != nil {
		log.Fatal(err)
	}
	err = pipeline.VertexFormat(0, perVertex{})
	if err != nil {
		log.Fatal(err)
	}
	pipeline.ClearColor(Vec4{0.9, 0.9, 0.9, 1.0})

	// Create the Uniform Buffer
	transform, err = gfx.NewBuffer(uintptr(64), gfx.DynamicStorage)
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
	colorfulTriangle, err = gfx.NewBuffer(data, 0)
	if err != nil {
		log.Fatal(err)
	}

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
}

//------------------------------------------------------------------------------

func (g *game) WindowResized(s IVec2, timestamp time.Duration) {
	r := float32(s.X) / float32(s.Y)
	projection = space.Perspective(math.Pi/4, r, 0.001, 1000.0)
}

var distance = float32(3)
var yaw = float32(-0.6)
var pitch = float32(0.4)

func (g *game) MouseWheel(motion IVec2, timestamp time.Duration) {
	distance -= float32(motion.Y) / 4
	g.updateView()
}

func (g *game) MouseButtonDown(b mouse.Button, clicks int, timestamp time.Duration) {
	if b == mouse.Left {
		mouse.SetRelativeMode(true)
	}
}

func (g *game) MouseButtonUp(b mouse.Button, clicks int, timestamp time.Duration) {
	if b == mouse.Left {
		mouse.SetRelativeMode(false)
	}
}

func (g *game) MouseMotion(motion IVec2, position IVec2, timestamp time.Duration) {
	if mouse.IsPressed(mouse.Left) {
		yaw -= 4 * float32(motion.X) / 1280
		pitch += 4 * float32(motion.Y) / 720
		switch {
		case pitch < -math.Pi/2+0.01:
			pitch = -math.Pi/2 + 0.01
		case pitch > math.Pi/2-0.01:
			pitch = math.Pi/2 - 0.01
		}
		g.updateView()
	}
}

func (g *game) updateView() {
	p := Vec3{
		math.Cos(pitch) * math.Sin(yaw),
		math.Sin(pitch),
		math.Cos(pitch) * math.Cos(yaw),
	}.Times(distance)
	view = space.LookAt(p, Vec3{0, 0, 0}, Vec3{0, 1, 0})
}

//------------------------------------------------------------------------------

var angle float32

func (g *game) Update() {
	// angle += 0.01
	model = space.Rotation(angle, Vec3{0, -1, 0}.Normalized())
	model = model.Times(space.Translation(Vec3{-0.5, -0.5, -0.5}))
}

func (g *game) Draw() {
	pipeline.Bind()
	pipeline.UniformBuffer(0, transform)

	mvp := projection.Times(view)
	mvp = mvp.Times(model)
	t := perObject{
		transform: mvp,
	}
	transform.Update(&t, 0)

	pipeline.VertexBuffer(0, colorfulTriangle, 0)
	gfx.Draw(gfx.Triangles, 0, 6*2*3)
}

//------------------------------------------------------------------------------
