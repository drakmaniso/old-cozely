// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package microtext

import (
	"strings"

	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/pixel"
	"time"
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
	nc, nr := int(screen.nbCols), int(screen.nbRows)
	for i := range screen.chars {
		screen.chars[i] = byte(i & 0x7F)
		if (i/4)%13 == 0 {
			screen.chars[i] |= 0x80
		}
		if i%nc == 0 || i%nc == nc-1 {
			screen.chars[i] = 5
		}
		if i/nc == 0 || i/nc == nr-1 {
			screen.chars[i] = 10
		}
	}
	screen.chars[0] = 6
	screen.chars[nc-1] = 12
	screen.chars[(nr-1)*nc] = 3
	screen.chars[(nr-1)*nc+nc-1] = 9
	fontSSBO = gfx.NewStorageBuffer(&Font, gfx.StaticStorage)
	SetColor(color.RGB{0, 0, 0}, color.RGB{1, 1, 1})
	SetOpaque(false)
	updated = false
	screenSSBO = gfx.NewStorageBuffer(&screen, gfx.DynamicStorage)
}

//------------------------------------------------------------------------------

var screen struct {
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
	opaque  uint32
	bgRed   float32
	bgGreen float32
	bgBlue  float32
	//
	chars [120 * 45]byte
}

var updated bool

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
	if false {
		t := (s.Y - (charHeight * int32(screen.pixelSize) * int32(screen.nbRows))) / 2
		if t > 0 {
			screen.top = uint32(t)
		} else {
			screen.top = 0
		}
	} else {
		screen.top = 0
	}

	updated = true
}

func SetColor(fg, bg color.RGB) {
	screen.fgRed = fg.R
	screen.fgGreen = fg.G
	screen.fgBlue = fg.B
	screen.bgRed = bg.R
	screen.bgGreen = bg.G
	screen.bgBlue = bg.B
	updated = true
}

func SetOpaque(o bool) {
	if o {
		screen.opaque = 1
	} else {
		screen.opaque = 0
	}
	updated = true
}

//------------------------------------------------------------------------------

func Draw() {
	pipeline.Bind()
	gfx.Disable(gfx.DepthTest)
	gfx.CullFace(false, false)
	if updated {
		screenSSBO.SubData(&screen, 0)
		updated = false
	}
	fontSSBO.Bind(0)
	screenSSBO.Bind(1)
	gfx.Draw(gfx.TriangleStrip, 0, 4)
	pipeline.Unbind()
}

//------------------------------------------------------------------------------
