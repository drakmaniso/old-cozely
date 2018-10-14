package gl_test

import (
	"image"
	_ "image/png"
	"os"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/space"
	"github.com/cozely/cozely/window"
	"github.com/cozely/cozely/x/gl"
	"github.com/cozely/cozely/x/math32"
)

// Declarations ////////////////////////////////////////////////////////////////

// Input Bindings
// (same as in FirstCube example)

// OpenGL Objects

type loop5 struct {
	// OpenGL objects
	pipeline    *gl.Pipeline
	perFrameUBO gl.UniformBuffer
	sampler     gl.Sampler
	diffuse     gl.Texture2D

	// Transformation matrices
	screenFromView  space.Matrix // projection matrix
	viewFromWorld   space.Matrix // view matrix
	worldFromObject space.Matrix // model matrix

	// Cube state
	position   coord.XYZ
	yaw, pitch float32
}

// Uniform buffer
// (same as in example 04)
// type perObject struct {
// 	screenFromObject space.Matrix
// }

// Vertex buffer
type uvmesh []struct {
	position coord.XYZ `layout:"0"`
	uv       coord.XY  `layout:"1"`
}

// Initialization //////////////////////////////////////////////////////////////

func Example_texture() {
	defer cozely.Recover()

	cozely.Configure(cozely.Multisample(8))
	l := loop5{}
	window.Events.Resize = func() {
		s := window.Size()
		gl.Viewport(0, 0, int32(s.X), int32(s.Y))
		r := float32(s.X) / float32(s.Y)
		l.screenFromView = space.Perspective(math32.Pi/4, r, 0.001, 1000.0)
	}
	err := cozely.Run(&l)
	if err != nil {
		panic(err)
	}
	//Output:
}

func (l *loop5) Enter() {
	// Create and configure the pipeline
	l.pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader05.vert"),
		gl.Shader(cozely.Path()+"shader05.frag"),
		gl.VertexFormat(0, uvmesh{}),
		gl.Topology(gl.Triangles),
		gl.CullFace(false, true),
		gl.DepthTest(true),
		gl.DepthWrite(true),
		gl.DepthComparison(gl.LessOrEqual),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create the uniform buffer
	l.perFrameUBO = gl.NewUniformBuffer(&perObject{}, gl.DynamicStorage)

	// Create and fill the vertex buffer
	vbo := gl.NewVertexBuffer(uvcube(), gl.StaticStorage)

	// Create and bind the sampler
	l.sampler = gl.NewSampler(
		gl.Minification(gl.LinearMipmapLinear),
		gl.Anisotropy(16.0),
	)

	// Create and load the textures
	l.diffuse = gl.NewTexture2D(8, gl.SRGBA8, 512, 512)
	r, err := os.Open(cozely.Path() + "testpattern.png")
	if err != nil {
		panic(cozely.Wrap("opening texture", err))
	}
	defer r.Close()
	img, _, err := image.Decode(r)
	if err != nil {
		panic(cozely.Wrap("decoding texture", err))
	}
	l.diffuse.SubImage(0, 0, 0, img)
	l.diffuse.GenerateMipmap()

	// Initialize worldFromObject and viewFromWorld matrices
	l.position = coord.XYZ{0, 0, 0}
	l.yaw = -0.6
	l.pitch = 0.3
	l.computeWorldFromObject()
	l.computeViewFromWorld()

	// Bind the vertex buffer to the pipeline
	l.pipeline.Bind()
	vbo.Bind(0, 0)
	l.pipeline.Unbind()
}

func (loop5) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (l *loop5) React() {
	m := delta.XY()
	s := window.Size().Coord()

	if rotate.Pressed() || move.Pressed() || zoom.Pressed() {
		input.GrabMouse(true)
	}
	if rotate.Released() || move.Released() || zoom.Released() {
		input.GrabMouse(false)
	}

	if rotate.Ongoing() {
		l.yaw += 4 * m.X / s.X
		l.pitch += 4 * m.Y / s.Y
		switch {
		case l.pitch < -math32.Pi/2:
			l.pitch = -math32.Pi / 2
		case l.pitch > +math32.Pi/2:
			l.pitch = +math32.Pi / 2
		}
		l.computeWorldFromObject()
	}

	if move.Ongoing() {
		d := m.Times(2).Slashxy(s)
		l.position.X += d.X
		l.position.Y -= d.Y
		l.computeWorldFromObject()
	}

	if zoom.Ongoing() {
		d := m.Times(2).Slashxy(s)
		l.position.X += d.X
		l.position.Z += d.Y
		l.computeWorldFromObject()
	}

	if quit.Pressed() {
		cozely.Stop(nil)
	}
}

func (l *loop5) computeWorldFromObject() {
	rot := space.EulerZXY(l.pitch, l.yaw, 0)
	l.worldFromObject = space.Translation(l.position).Times(rot)
}

func (l *loop5) computeViewFromWorld() {
	l.viewFromWorld = space.LookAt(
		coord.XYZ{0, 0, 3},
		coord.XYZ{0, 0, 0},
		coord.XYZ{0, 1, 0},
	)
}

func (loop5) Update() {
}

func (l *loop5) Render() {
	l.pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(color.LRGBA{0.9, 0.9, 0.9, 1.0})

	u := perObject{
		screenFromObject: l.screenFromView.
			Times(l.viewFromWorld).
			Times(l.worldFromObject),
	}
	l.perFrameUBO.SubData(&u, 0)
	l.perFrameUBO.Bind(0)

	l.diffuse.Bind(0)
	l.sampler.Bind(0)
	gl.Draw(0, 6*2*3)

	l.pipeline.Unbind()
}

////////////////////////////////////////////////////////////////////////////////

func uvcube() uvmesh {
	return uvmesh{
		// Front Face
		{coord.XYZ{-0.5, -0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, -0.5, +0.5}, coord.XY{1, 1}},
		{coord.XYZ{+0.5, +0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, -0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, +0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, +0.5, +0.5}, coord.XY{0, 0}},
		// Back Face
		{coord.XYZ{+0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{-0.5, -0.5, -0.5}, coord.XY{1, 1}},
		{coord.XYZ{-0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{+0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{-0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{+0.5, +0.5, -0.5}, coord.XY{0, 0}},
		// Right Face
		{coord.XYZ{+0.5, -0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, -0.5, -0.5}, coord.XY{1, 1}},
		{coord.XYZ{+0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{+0.5, -0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{+0.5, +0.5, +0.5}, coord.XY{0, 0}},
		// Left Face
		{coord.XYZ{-0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{-0.5, -0.5, +0.5}, coord.XY{1, 1}},
		{coord.XYZ{-0.5, +0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{-0.5, +0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, +0.5, -0.5}, coord.XY{0, 0}},
		// Bottom Face
		{coord.XYZ{-0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, -0.5, -0.5}, coord.XY{1, 1}},
		{coord.XYZ{+0.5, -0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, -0.5, -0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, -0.5, +0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, -0.5, +0.5}, coord.XY{0, 0}},
		// Top Face
		{coord.XYZ{-0.5, +0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, +0.5, +0.5}, coord.XY{1, 1}},
		{coord.XYZ{+0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, +0.5, +0.5}, coord.XY{0, 1}},
		{coord.XYZ{+0.5, +0.5, -0.5}, coord.XY{1, 0}},
		{coord.XYZ{-0.5, +0.5, -0.5}, coord.XY{0, 0}},
	}
}

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
