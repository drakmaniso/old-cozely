// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"image"
	_ "image/png"
	"os"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/space"
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
	sampler     gfx.Sampler
	diffuse     gfx.Texture2D
)

// Uniform buffer
var perFrame struct {
	transform space.Matrix
}

// Vertex buffer
type mesh []struct {
	position space.Coord `layout:"0"`
	uv       plane.Coord `layout:"1"`
}

// Matrices
var (
	model      space.Matrix
	view       space.Matrix
	projection space.Matrix
)

// Cube state
var (
	distance   float32
	position   space.Coord
	yaw, pitch float32
)

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
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
		gfx.VertexFormat(0, mesh{}),
	)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)

	// Create and fill the vertex buffer
	vbo := gfx.NewVertexBuffer(cube(), gfx.StaticStorage)

	// Create and bind the sampler
	sampler = gfx.NewSampler(
		gfx.Minification(gfx.LinearMipmapLinear),
		gfx.Anisotropy(16.0),
	)

	// Create and load the textures
	diffuse = gfx.NewTexture2D(8, pixel.Coord{512, 512}, gfx.SRGBA8)
	r, err := os.Open(glam.Path() + "../shared/testpattern.png")
	if err != nil {
		return err
	}
	defer r.Close()
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	diffuse.Load(img, pixel.Coord{0, 0}, 0)
	diffuse.GenerateMipmap()

	// Initialize model and view matrices
	position = space.Coord{0, 0, 0}
	yaw = -0.6
	pitch = 0.3
	updateModel()
	distance = 3
	updateView()

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	// MTX

	return gfx.Err()
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
	view = space.LookAt(space.Coord{0, 0, distance}, space.Coord{0, 0, 0}, space.Coord{0, 1, 0})
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update() {
	x, y := 1, 1
	x, y = mtx.Print(x, y, "One two Three ", "FOUR ", "FIVE", "\n")
	x, y = mtx.Print(x, y, "Six\n")
	x, y = mtx.Print(x, y, "Seven\n")
	mtx.Print(0, 0, "TOP LEFT")
	mtx.Print(-12, -1, "BOTTOM RIGHT")

	_, y = mtx.Print(0, y, "Un es", "sai ", 2, " print ", 1.1, "\n")
	mtx.ReverseVideo(true)
	_, y = mtx.Print(0, y, "Essai\n")
	mtx.ReverseVideo(false)
	_, y = mtx.Print(0, y, "int=", 33)
	_, y = mtx.Print(0, y, "\nbool=", true)
	mtx.SetPrecision(-1)
	_, y = mtx.Print(0, y, "\nZoom=", 33.0/27.0)
	_, y = mtx.Print(0, y, "\nPosition: x=", 114.0/23.0)
	_, y = mtx.Print(0, y, "\nOther:    y=", 237.0/31.0)
	mtx.SetPrecision(6)
	_, y = mtx.Print(0, y, "\nfloat32=", 12.123456789123456789123456789*1.65487)
	_, y = mtx.Print(0, y, "\nfloat64=", 4.56789123456789123456789*1.1543)
	mtx.Print(5, y+1, `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec libero ligula,
consectetur at congue et, ultricies placerat velit. Pellentesque finibus
tristique orci sit amet pharetra. Nullam arcu urna, tempus malesuada aliquet
quis, semper blandit ante. Proin vitae dignissim lacus. Etiam in rutrum
tortor. Nulla sed maximus dolor, quis venenatis ante. Maecenas nec ante vel
massa elementum varius nec vitae odio. Proin tincidunt iaculis elit eu luctus.
Donec dignissim ipsum in orci congue rutrum a at turpis. Aliquam congue
tristique dapibus. Pellentesque sed aliquam ex, id blandit metus. Mauris
egestas magna quis elit dignissim, a laoreet sem facilisis. Duis tristique
dapibus dictum. Lorem ipsum dolor sit amet, consectetur adipiscing elit.`)

	// mtx.Clear('\000')
	// c := byte(0)
	// sx, sy := mtx.Size()
	// for y := 0; y < sy; y++ {
	// 	for x := 0; x < sx && x < 128; x++ {
	// 		mtx.Poke(x, y, c)
	// 		c++
	// 	}
	// }

}

func (l looper) Draw() {
	pipeline.Bind()
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	gfx.Enable(gfx.DepthTest)
	gfx.CullFace(false, true)
	gfx.Enable(gfx.FramebufferSRGB)

	perFrame.transform = projection.Times(view)
	perFrame.transform = perFrame.transform.Times(model)
	perFrameUBO.SubData(&perFrame, 0)
	perFrameUBO.Bind(0)

	diffuse.Bind(0)
	sampler.Bind(0)
	gfx.Draw(gfx.Triangles, 0, 6*2*3)

	pipeline.Unbind()
}

//------------------------------------------------------------------------------
