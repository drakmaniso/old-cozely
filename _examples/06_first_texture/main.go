// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"image"
	_ "image/png"
	"os"
	"time"
	"unsafe"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/color"
	. "github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/geom/space"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
	g := newGame()

	glam.Loop = g
	window.Handle = g
	mouse.Handle = g

	// Run the Game Loop
	err := glam.Run()
	check(err)
}

//------------------------------------------------------------------------------

type game struct {
	basic.WindowHandler
	basic.MouseHandler

	pipeline  gfx.Pipeline
	transform gfx.UniformBuffer
	cube      gfx.VertexBuffer
	diffuse   gfx.Texture2D

	distance                float32
	position                Vec3
	yaw, pitch              float32
	model, view, projection Mat4
}

type perVertex struct {
	position Vec3 `layout:"0"`
	uv       Vec2 `layout:"1"`
}

type perObject struct {
	transform Mat4
}

//------------------------------------------------------------------------------

func newGame() *game {
	g := &game{}

	// Setup the Pipeline
	vf, err := os.Open(glam.Path() + "shader.vert")
	check(err)
	vs, err := gfx.NewVertexShader(vf)
	check(err)
	ff, err := os.Open(glam.Path() + "shader.frag")
	check(err)
	fs, err := gfx.NewFragmentShader(ff)
	check(err)
	g.pipeline, err = gfx.NewPipeline(vs, fs)
	check(err)
	err = g.pipeline.VertexFormat(0, perVertex{})
	check(err)
	g.pipeline.ClearColor(Vec4{0.9, 0.9, 0.9, 1.0})

	// Create the Uniform Buffer
	g.transform, err = gfx.NewUniformBuffer(unsafe.Sizeof(perObject{}), gfx.DynamicStorage)
	check(err)

	// Create and fill the Vertex Buffer
	g.cube, err = gfx.NewVertexBuffer(cube(), gfx.StaticStorage)
	check(err)

	// Create and bind the sampler
	s := gfx.NewSampler()
	s.Filter(gfx.LinearMipmapLinear, gfx.Linear)
	s.Anisotropy(16.0)
	s.Wrap(gfx.Repeat, gfx.Repeat, gfx.Repeat)        // Default
	s.BorderColor(color.RGBA{R: 0, G: 0, B: 0, A: 0}) // Default
	g.pipeline.Sampler(0, s)

	// Create and load the textures
	g.diffuse = gfx.NewTexture2D(8, IVec2{512, 512}, gfx.SRGBA8)
	r, err := os.Open(glam.Path() + "../shared/testpattern.png")
	check(err)
	defer r.Close()
	img, _, err := image.Decode(r)
	check(err)
	g.diffuse.Data(img, IVec2{0, 0}, 0)
	g.diffuse.GenerateMipmap()

	// Initialize model and view matrices
	g.position = Vec3{0, 0, 0}
	g.yaw = -0.6
	g.pitch = 0.3
	g.updateModel()
	g.distance = 3
	g.updateView()

	return g
}

//------------------------------------------------------------------------------

func (g *game) WindowResized(s Vec2, timestamp time.Duration) {
	r := s.X / s.Y
	g.projection = space.Perspective(math.Pi/4, r, 0.001, 1000.0)
}

func (g *game) MouseWheel(motion Vec2, timestamp time.Duration) {
	g.distance -= motion.Y / 4
	g.updateView()
}

func (g *game) MouseButtonDown(b mouse.Button, clicks int, timestamp time.Duration) {
	mouse.SetRelativeMode(true)
}

func (g *game) MouseButtonUp(b mouse.Button, clicks int, timestamp time.Duration) {
	mouse.SetRelativeMode(false)
}

func (g *game) MouseMotion(motion Vec2, position Vec2, timestamp time.Duration) {
	s := window.Size()

	switch {
	case mouse.IsPressed(mouse.Left):
		g.yaw += 4 * motion.X / s.X
		g.pitch += 4 * motion.Y / s.Y
		switch {
		case g.pitch < -math.Pi/2:
			g.pitch = -math.Pi / 2
		case g.pitch > +math.Pi/2:
			g.pitch = +math.Pi / 2
		}
		g.updateModel()

	case mouse.IsPressed(mouse.Middle):
		g.position.X += 2 * motion.X / s.X
		g.position.Y -= 2 * motion.Y / s.Y
		g.updateModel()
	}
}

//------------------------------------------------------------------------------

func (g *game) updateModel() {
	g.model = space.Translation(g.position)
	g.model = g.model.Times(space.EulerZXY(g.pitch, g.yaw, 0))
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
	g.transform.Bind(0)

	mvp := g.projection.Times(g.view)
	mvp = mvp.Times(g.model)
	t := perObject{
		transform: mvp,
	}
	g.transform.Update(&t, 0)

	g.cube.Bind(0, 0)
	g.diffuse.Bind(0)
	gfx.Draw(gfx.Triangles, 0, 6*2*3)
}

//------------------------------------------------------------------------------

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//------------------------------------------------------------------------------
