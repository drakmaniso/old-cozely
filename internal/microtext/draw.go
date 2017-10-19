// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package microtext

import (
	"strings"
	"unsafe"

	"github.com/drakmaniso/carol/gfx"
	"github.com/drakmaniso/carol/internal"
	"github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

var (
	pipeline   *gfx.Pipeline
	fontSSBO   gfx.StorageBuffer
	screenSSBO gfx.StorageBuffer
)

//------------------------------------------------------------------------------

func init() {
	screen.nbCols = 1
	screen.nbRows = 1
	screen.pixelSize = 2
	text = make([]byte, screen.nbCols*screen.nbRows)
}

//------------------------------------------------------------------------------

// Setup is called during carol setup.
func Setup() {
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(strings.NewReader(vertexShader)),
		gfx.FragmentShader(strings.NewReader(fragmentShader)),
		gfx.CullFace(false, false),
		gfx.DepthTest(false),
		gfx.Topology(gfx.TriangleStrip),
	)

	fontSSBO = gfx.NewStorageBuffer(&Font, gfx.StaticStorage)

	s := pixel.Coord{internal.Window.Width, internal.Window.Height}
	WindowResized(s)
}

//------------------------------------------------------------------------------

const charWidth = 7
const charHeight = 11

// WindowResized is called each time resolution changes. It reallocates all GPU
// ressources accordingly.
func WindowResized(s pixel.Coord) {
	screen.nbCols = s.X / (charWidth * screen.pixelSize)
	screen.nbRows = s.Y / (charHeight * screen.pixelSize)
	screen.top = 0
	screen.left = 0

	// Reallocate the SSBO
	text = make([]byte, screen.nbCols*screen.nbRows)
	screenSSBO.Delete()
	screenSSBO = gfx.NewStorageBuffer(
		unsafe.Sizeof(screen)+uintptr(screen.nbCols*screen.nbRows),
		gfx.DynamicStorage,
	)

	textUpdated = true

	// Calculate the margins
	l := (s.X - (charWidth * int32(screen.pixelSize) * int32(screen.nbCols))) / 2
	if l > 0 {
		screen.left = int32(l)
	} else {
		screen.left = 0
	}
	if true {
		t := (s.Y - (charHeight * int32(screen.pixelSize) * int32(screen.nbRows))) / 2
		if t > 0 {
			screen.top = int32(t)
		} else {
			screen.top = 0
		}
	} else {
		screen.top = 0
	}

	screenUpdated = true

	ShowFrameTime(ftenabled, ftx, fty)
}

//------------------------------------------------------------------------------

// Draw is called during the main loop, after the user's Draw.
func Draw() {
	pipeline.Bind()
	gfx.Blending(gfx.SrcAlpha, gfx.OneMinusSrcAlpha)
	gfx.Enable(gfx.Blend)
	if screenUpdated {
		screenSSBO.SubData(&screen, 0)
		screenUpdated = false
	}
	if textUpdated {
		screenSSBO.SubData(text, unsafe.Sizeof(screen))
		textUpdated = false
	}
	fontSSBO.Bind(0)
	screenSSBO.Bind(1)
	gfx.Draw(0, 4)
	pipeline.Unbind()
}

//------------------------------------------------------------------------------

// Data for the SSBO
var (
	screen struct {
		left   int32
		top    int32
		nbCols int32
		nbRows int32
		//
		pixelSize    int32
		reverseVideo int32
		_            int32
		_            int32
	}

	text []byte
)

var (
	screenUpdated bool
	textUpdated   bool
)

//------------------------------------------------------------------------------

// Size returns the number of column and rows in the MTX screen.
func Size() (cols, rows int) {
	return int(screen.nbCols), int(screen.nbRows)
}

//------------------------------------------------------------------------------

// SetReverseVideo activates or deactivates reverse video.
func SetReverseVideo(o bool) {
	if o {
		screen.reverseVideo = 1
	} else {
		screen.reverseVideo = 0
	}
	screenUpdated = true
}

// GetReverseVideo returns true if microtext is in reverse video.
func GetReverseVideo() bool {
	return screen.reverseVideo != 0
}

// ToggleReverseVideo toggles reverse video mode.
func ToggleReverseVideo() {
	if screen.reverseVideo != 0 {
		screen.reverseVideo = 0
	} else {
		screen.reverseVideo = 1
	}
	screenUpdated = true
}

//------------------------------------------------------------------------------

// Clear erases the MTX screen.
func Clear() {
	for i := range text {
		text[i] = '\x00'
	}
	textUpdated = true
}

//------------------------------------------------------------------------------

// Peek returns the character at given coordinates.
func Peek(x, y int) byte {
	return text[x+y*int(screen.nbCols)]
}

// Poke sets the character at given coordinates.
func Poke(x, y int, c byte) {
	i := x + y*int(screen.nbCols)
	if text[i] != c {
		text[i] = c
		textUpdated = true
	}
}

//------------------------------------------------------------------------------

func ShowFrameTime(enable bool, x, y int) {
	ftenabled = enable
	ftx, fty = x, y

	if x < 0 {
		x += int(screen.nbCols)
		if x < 0 {
			x = 0
		}
	}
	if x > int(screen.nbCols)-6 {
		x = int(screen.nbCols) - 6
	}

	switch {
	case y < 0:
		y += int(screen.nbRows)
		if y < 0 {
			y = 0
		}
	case y > int(screen.nbRows)-1:
		y = int(screen.nbRows) - 1
	}

	ftloc = x + y*int(screen.nbCols)
}

func PrintFrameTime(frametime float64, xruns int) {
	if !ftenabled {
		return
	}

	colour := byte(0x00)
	if xruns > 0 {
		colour = 0x80
	}

	// Convert to milliseconds
	frametime *= 1000.0
	// Round to the first decimal
	frametime = float64(int64(frametime*100.0+0.5)) / 100.0
	// Isolate digits
	v100 := uint(frametime / 100)
	v10 := uint(frametime/10) - v100*10
	v1 := uint(frametime) - v100*100 - v10*10
	v01 := uint(frametime*10) - v100*1000 - v10*100 - v1*10
	v001 := uint(frametime*100) - v100*10000 - v10*1000 - v1*100 - v01*10

	var c100, c10, c1, c01, c001 byte

	switch {
	case v100 > 9:
		c100 = '~' + 1 | colour
	case v100 == 0:
		c100 = '\x00'
	default:
		c100 = '0' + byte(v100) | colour
	}

	switch {
	case v100 > 9 || v10 > 9:
		c10 = '~' + 1 | colour
	case v100 == 0 && v10 == 0:
		c10 = '\x00'
	default:
		c10 = '0' + byte(v10) | colour
	}

	switch {
	case v100 > 9 || v1 > 9:
		c1 = '~' + 1 | colour
	default:
		c1 = '0' + byte(v1) | colour
	}

	switch {
	case v100 > 9 || v01 > 9:
		c01 = '~' + 1 | colour
	default:
		c01 = '0' + byte(v01) | colour
	}

	switch {
	case v100 > 9 || v01 > 9:
		c001 = '~' + 1 | colour
	default:
		c001 = '0' + byte(v001) | colour
	}

	text[ftloc+0] = c100
	text[ftloc+1] = c10
	text[ftloc+2] = c1
	text[ftloc+3] = '.' | colour
	text[ftloc+4] = c01
	text[ftloc+5] = c001

	textUpdated = true
}

var ftx, fty int
var ftenabled bool
var ftloc int

//------------------------------------------------------------------------------
