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
	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/geom/space"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
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

// Vertex buffer layout
type perVertex struct {
	position geom.Vec3 `layout:"0"`
	uv       geom.Vec2 `layout:"1"`
}

// Uniform buffer
type perObject struct {
	transform geom.Mat4
}

// OpenGL objects
var (
	pipeline  gfx.Pipeline
	transform gfx.UniformBuffer
	mesh      gfx.VertexBuffer
	diffuse   gfx.Texture2D
)

// Cube state
var (
	distance   float32
	position   geom.Vec3
	yaw, pitch float32
)

// Matrices
var (
	model      geom.Mat4
	view       geom.Mat4
	projection geom.Mat4
)

//------------------------------------------------------------------------------

func setup() error {
	// Setup the pipeline
	v, err := os.Open(glam.Path() + "shader.vert")
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
		gfx.VertexFormat(0, perVertex{}),
	)
	gfx.Enable(gfx.DepthTest)
	gfx.Enable(gfx.CullFace)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	transform = gfx.NewUniformBuffer(unsafe.Sizeof(perObject{}), gfx.DynamicStorage)

	// Create and fill the vertex buffer
	mesh = gfx.NewVertexBuffer(cube(), gfx.StaticStorage)

	// Create and bind the sampler
	s := gfx.NewSampler()
	s.Filter(gfx.LinearMipmapLinear, gfx.Linear)
	s.Anisotropy(16.0)
	s.Wrap(gfx.Repeat, gfx.Repeat, gfx.Repeat)        // Default
	s.BorderColor(color.RGBA{R: 0, G: 0, B: 0, A: 0}) // Default
	s.Bind(0)

	// Create and load the textures
	diffuse = gfx.NewTexture2D(8, pixel.XY{512, 512}, gfx.SRGBA8)
	r, err := os.Open(glam.Path() + "../shared/testpattern.png")
	if err != nil {
		return err
	}
	defer r.Close()
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	diffuse.Data(img, pixel.XY{0, 0}, 0)
	diffuse.GenerateMipmap()

	// Initialize model and view matrices
	position = geom.Vec3{0, 0, 0}
	yaw = -0.6
	pitch = 0.3
	updateModel()
	distance = 3
	updateView()

	return gfx.Err()
}

//------------------------------------------------------------------------------

// Event handler
type handler struct {
	basic.WindowHandler
	basic.MouseHandler
}

func (h handler) WindowResized(s pixel.XY, _ time.Duration) {
	r := float32(s.X) / float32(s.Y)
	projection = space.Perspective(math.Pi/4, r, 0.001, 1000.0)
}

func (h handler) MouseWheel(motion pixel.XY, _ time.Duration) {
	distance -= float32(motion.Y) / 4
	updateView()
}

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ time.Duration) {
	mouse.SetRelativeMode(true)
}

func (h handler) MouseButtonUp(b mouse.Button, _ int, _ time.Duration) {
	mouse.SetRelativeMode(false)
}

func (h handler) MouseMotion(motion pixel.XY, _ pixel.XY, _ time.Duration) {
	s := window.Size()

	switch {
	case mouse.IsPressed(mouse.Left):
		yaw += 4 * float32(motion.X) / float32(s.X)
		pitch += 4 * float32(motion.Y) / float32(s.Y)
		switch {
		case pitch < -math.Pi/2:
			pitch = -math.Pi / 2
		case pitch > +math.Pi/2:
			pitch = +math.Pi / 2
		}
		updateModel()

	case mouse.IsPressed(mouse.Middle):
		position.X += 2 * float32(motion.X) / float32(s.X)
		position.Y -= 2 * float32(motion.Y) / float32(s.Y)
		updateModel()
	}
}

//------------------------------------------------------------------------------

func updateModel() {
	model = space.Translation(position)
	model = model.Times(space.EulerZXY(pitch, yaw, 0))
}

func updateView() {
	if distance < 1 {
		distance = 1
	}
	view = space.LookAt(geom.Vec3{0, 0, distance}, geom.Vec3{0, 0, 0}, geom.Vec3{0, 1, 0})
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update() {
}

func (l looper) Draw() {
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	pipeline.Bind()
	transform.Bind(0)

	mvp := projection.Times(view)
	mvp = mvp.Times(model)
	t := perObject{
		transform: mvp,
	}
	transform.Update(&t, 0)

	mesh.Bind(0, 0)
	diffuse.Bind(0)
	gfx.Draw(gfx.Triangles, 0, 6*2*3)
}

//------------------------------------------------------------------------------
