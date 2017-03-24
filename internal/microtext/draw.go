// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package microtext

import (
	"strings"

	"unsafe"

	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/pixel"
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
	SetColor(color.RGB{1, 1, 1}, color.RGB{0, 0, 0})
	SetBgAlpha(true)
	text = make([]byte, screen.nbCols*screen.nbRows)
}

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
}

//------------------------------------------------------------------------------

const charWidth = 7
const charHeight = 11

// WindowResized is called each time resolution changes. It reallocates all GPU
// ressources accordingly.
func WindowResized(s pixel.Coord, ts uint32) {
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

	ShowFrameTime(ftenabled, ftx, fty, opaque)
}

//------------------------------------------------------------------------------

// Draw is called during the main loop, after the user's Draw.
func Draw() {
	pipeline.Bind()
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
		fgRed     float32
		fgGreen   float32
		fgBlue    float32
		pixelSize int32
		//
		bgRed   float32
		bgGreen float32
		bgBlue  float32
		bgAlpha float32
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

// SetColor changes the foreground and background colors.
func SetColor(fg, bg color.RGB) {
	screen.fgRed = fg.R
	screen.fgGreen = fg.G
	screen.fgBlue = fg.B

	screen.bgRed = bg.R
	screen.bgGreen = bg.G
	screen.bgBlue = bg.B

	screenUpdated = true
}

// SetBgAlpha sets wether the background of letters is dran or not.
func SetBgAlpha(o bool) {
	if o {
		screen.bgAlpha = 1.0
	} else {
		screen.bgAlpha = 0.0
	}
	screenUpdated = true
}

// GetBgAlpha returns true if the background of letters is currently drawn.
func GetBgAlpha() bool {
	return screen.bgAlpha != 0.0
}

// ToggleBgAlpha inverts the status of background transparency.
func ToggleBgAlpha() {
	if screen.bgAlpha != 0 {
		screen.bgAlpha = 0
	} else {
		screen.bgAlpha = 1
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

func ShowFrameTime(enable bool, x, y int, opaque bool) {
	ftenabled = enable
	opaque = opaque
	ftx, fty = x, y

	if x < 0 {
		x += int(screen.nbCols)
		if x < 0 {
			x = 0
		}
	}
	if x > int(screen.nbCols)-5 {
		x = int(screen.nbCols) - 5
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
	frametime = float64(int64(frametime*10.0+0.5)) / 10.0
	// Isolate digits
	v100 := uint(frametime / 100)
	v10 := uint(frametime/10) - v100*10
	v1 := uint(frametime) - v100*100 - v10*10
	v01 := uint(frametime*10) - v100*1000 - v10*100 - v1*10

	var c100, c10, c1, c01 byte

	switch {
	case v100 > 9:
		c100 = '~' + 1 | colour
	case v100 == 0:
		if opaque {
			c100 = ' ' | colour
		} else {
			c100 = '\x00'
		}
	default:
		c100 = '0' + byte(v100) | colour
	}

	switch {
	case v100 > 9 || v10 > 9:
		c10 = '~' + 1 | colour
	case v100 == 0 && v10 == 0:
		if opaque {
			c10 = ' ' | colour
		} else {
			c10 = '\x00'
		}
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

	text[ftloc+0] = c100
	text[ftloc+1] = c10
	text[ftloc+2] = c1
	text[ftloc+3] = '.' | colour
	text[ftloc+4] = c01

	textUpdated = true
}

var ftx, fty int
var ftenabled bool
var opaque bool
var ftloc int

//------------------------------------------------------------------------------
