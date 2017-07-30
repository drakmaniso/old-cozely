// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package overl

import (
	"strings"
	"unsafe"

	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var (
	pipeline *gfx.Pipeline
	fontSSBO gfx.StorageBuffer
)

var overlays []*Overlay

// Used as temporary for SSBO uploading
var header struct {
	left, top     float32
	right, bottom float32
	//
	x, y          int32
	columns, rows int32
	//
	pixelSize uint32
	flags     int32
	_         int32
	_         int32
}

const FontWidth = 7
const FontHeight = 11
const pixelSize = 1

//------------------------------------------------------------------------------

// Setup is called during glam setup.
func Setup() {
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(strings.NewReader(vertexShader)),
		gfx.FragmentShader(strings.NewReader(fragmentShader)),
		gfx.CullFace(false, false),
		gfx.DepthTest(false),
		gfx.Topology(gfx.TriangleStrip),
	)

	fontSSBO = gfx.NewStorageBuffer(&Font, gfx.StaticStorage)

	if err := gfx.Err(); err != nil {
		internal.Log(err.Error())
	}
}

//------------------------------------------------------------------------------

func Draw() {
	for _, o := range overlays {
		o.Draw()
	}
}

//------------------------------------------------------------------------------

func WindowResized(s pixel.Coord) {
	for _, o := range overlays {
		o.WindowResized(s)
	}
}

//------------------------------------------------------------------------------

type Overlay struct {
	ssbo gfx.StorageBuffer

	// Data for the SSBO

	header struct {
		left, top     float32
		right, bottom float32
		//
		x, y          int32
		columns, rows int32
		//
		pixelSize uint32
		flags     int32
		_         int32
		_         int32
	}

	text []byte

	// Flags

	headerUpdated bool
	textUpdated   bool
}

//------------------------------------------------------------------------------

func FontSize() pixel.Coord {
	ps := int32(1 << pixelSize) //TODO
	return pixel.Coord{FontWidth, FontHeight}.Times(ps)
}

//------------------------------------------------------------------------------

func Create(position pixel.Coord, columns, rows int) *Overlay {
	o := Overlay{}
	o.header.x = position.X
	o.header.y = position.Y
	o.header.columns = int32(columns)
	o.header.rows = int32(rows)
	o.header.pixelSize = pixelSize

	o.header.top = 0.0
	o.header.left = 0.0
	o.header.bottom = 0.0
	o.header.right = 0.0
	// o.header.flags = 1

	o.text = make([]byte, o.header.columns*o.header.rows)

	o.headerUpdated = true
	o.textUpdated = true

	overlays = append(overlays, &o)
	return &o
}

//------------------------------------------------------------------------------

func (o *Overlay) WindowResized(s pixel.Coord) {
	r := o.header.left == 0 &&
		o.header.right == 0 &&
		o.header.top == 0 &&
		o.header.bottom == 0
	_ = r

	ps := int32(1 << o.header.pixelSize)
	//TODO: handle negative positions
	o.header.left = 2*float32(o.header.x)/float32(s.X) - 1
	o.header.top = 1 - 2*float32(o.header.y)/float32(s.Y)
	o.header.right = o.header.left + 2*float32(o.header.columns*FontWidth*ps)/float32(s.X)
	o.header.bottom = o.header.top - 2*float32(o.header.rows*FontHeight*ps)/float32(s.Y)

	if r {
		o.ssbo.Delete()
		o.ssbo = gfx.NewStorageBuffer(
			unsafe.Sizeof(o.header)+uintptr(o.header.columns*o.header.rows),
			gfx.DynamicStorage,
		)
	}

	o.headerUpdated = true
}

//------------------------------------------------------------------------------

func (o *Overlay) Bounds() (topLeft, bottomRight pixel.Coord) {
	tl := pixel.Coord{o.header.x, o.header.y}
	br := tl.Plus(pixel.Coord{FontWidth * o.header.columns, FontHeight * o.header.rows})
	return tl, br
}

//------------------------------------------------------------------------------

// Clear erases the overlay.
func (o *Overlay) Clear() {
	for i := range o.text {
		o.text[i] = '\x00'
	}
	o.textUpdated = true
}

//------------------------------------------------------------------------------

// Peek returns the character at given coordinates.
func (o *Overlay) Peek(x, y int) byte {
	return o.text[x+y*int(o.header.columns)]
}

// Poke sets the character at given coordinates.
func (o *Overlay) Poke(x, y int, c byte) {
	i := x + y*int(o.header.columns)
	if o.text[i] != c {
		o.text[i] = c
		o.textUpdated = true
	}
}

//------------------------------------------------------------------------------

// Draw is
func (o *Overlay) Draw() {
	pipeline.Bind()
	gfx.Blending(gfx.SrcAlpha, gfx.OneMinusSrcAlpha)
	gfx.Enable(gfx.Blend)
	if o.headerUpdated {
		header = o.header
		o.ssbo.SubData(&header, 0)
		o.headerUpdated = false
	}
	if o.textUpdated {
		o.ssbo.SubData(o.text, unsafe.Sizeof(o.header))
		o.textUpdated = false
	}
	fontSSBO.Bind(0)
	o.ssbo.Bind(1)
	gfx.Draw(0, 4)
	pipeline.Unbind()
}

//------------------------------------------------------------------------------

func (o *Overlay) Size() (columns, rows int) {
	return int(o.header.columns), int(o.header.rows)
}

//------------------------------------------------------------------------------
