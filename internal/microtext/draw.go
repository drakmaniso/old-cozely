// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package microtext

import (
	"strings"

	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/pixel"
	"time"
	"unsafe"
)

//------------------------------------------------------------------------------

var (
	pipeline   gfx.Pipeline
	fontSSBO   gfx.StorageBuffer
	screenSSBO gfx.StorageBuffer
)

//------------------------------------------------------------------------------

func Setup() {
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(strings.NewReader(vertexShader)),
		gfx.FragmentShader(strings.NewReader(fragmentShader)),
	)

	screen.nbCols = 80
	screen.nbRows = 30
	screen.pixelSize = 2

	Text = make([]byte, screen.nbCols*screen.nbRows)

	nc, nr := int(screen.nbCols), int(screen.nbRows)
	for i := range Text {
		Text[i] = byte(i & 0xFF) // 0x20
		if i%nc == 0 {
			Text[i] = 152
		}
		if i%nc == nc-1 {
			Text[i] = 153
		}
		if i/nc <= 2 {
			Text[i] = 160
		}
		if i/nc == nr-1 {
			Text[i] = 155
		}
	}
	Text[nc+1] = byte('D') | 0x80
	Text[nc+2] = byte('i') | 0x80
	Text[nc+3] = byte('a') | 0x80
	Text[nc+4] = byte('l') | 0x80
	Text[nc+5] = byte('o') | 0x80
	Text[nc+6] = byte('g') | 0x80
	Text[nc+nc-2] = 159
	Text[(nr-1)*nc] = 154
	Text[(nr-1)*nc+nc-1] = 156
	fontSSBO = gfx.NewStorageBuffer(&Font, gfx.StaticStorage)
	SetColor(color.RGB{0, 0, 0}, color.RGB{1, 1, 1})
	SetOpacity(false)
	TextUpdated = true
	screenSSBO = gfx.NewStorageBuffer(
		unsafe.Sizeof(screen)+uintptr(screen.nbCols*screen.nbRows),
		gfx.DynamicStorage,
	)
}

//------------------------------------------------------------------------------

func Draw() {
	pipeline.Bind()
	gfx.Disable(gfx.DepthTest)
	gfx.CullFace(false, false)
	if screenUpdated {
		screenSSBO.SubData(&screen, 0)
		screenUpdated = false
	}
	if TextUpdated {
		screenSSBO.SubData(Text, unsafe.Sizeof(screen))
		TextUpdated = false
	}
	fontSSBO.Bind(0)
	screenSSBO.Bind(1)
	gfx.Draw(gfx.TriangleStrip, 0, 4)
	pipeline.Unbind()
}

//------------------------------------------------------------------------------

// Data for the SSBO
var (
	screen struct {
		left   uint32
		top    uint32
		nbCols uint32
		nbRows uint32
		//
		pixelSize uint32
		fgRed     float32
		fgGreen   float32
		fgBlue    float32
		//
		opacity uint32
		bgRed   float32
		bgGreen float32
		bgBlue  float32
	}

	Text []byte
)

var (
	screenUpdated bool
	TextUpdated   bool
)

//------------------------------------------------------------------------------

func SetColor(fg, bg color.RGB) {
	screen.fgRed = fg.R
	screen.fgGreen = fg.G
	screen.fgBlue = fg.B
	screen.bgRed = bg.R
	screen.bgGreen = bg.G
	screen.bgBlue = bg.B
	screenUpdated = true
}

func SetOpacity(o bool) {
	if o {
		screen.opacity = 1
	} else {
		screen.opacity = 0
	}
	screenUpdated = true
}

func ToggleOpacity() {
	if screen.opacity != 0 {
		screen.opacity = 0
	} else {
		screen.opacity = 1
	}
	screenUpdated = true
}

func Size() (cols, rows int) {
	return int(screen.nbCols), int(screen.nbRows)
}

//------------------------------------------------------------------------------

const charWidth = 8
const charHeight = 12

func WindowResized(s pixel.Coord, ts time.Duration) {
	// First, calculate the pixel size needed to display the whole screen
	px := s.X / (charWidth * int32(screen.nbCols))
	py := s.Y / (charHeight * int32(screen.nbRows))
	if py < px {
		px = py
	}
	if px < 1 {
		px = 1
	}
	screen.pixelSize = uint32(px)

	// Calculate the margins
	l := (s.X - (charWidth * int32(screen.pixelSize) * int32(screen.nbCols))) / 2
	if l > 0 {
		screen.left = uint32(l)
	} else {
		screen.left = 0
	}
	if true {
		t := (s.Y - (charHeight * int32(screen.pixelSize) * int32(screen.nbRows))) / 2
		if t > 0 {
			screen.top = uint32(t)
		} else {
			screen.top = 0
		}
	} else {
		screen.top = 0
	}

	screenUpdated = true
}

//------------------------------------------------------------------------------
