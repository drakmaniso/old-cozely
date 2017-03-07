// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

// Event handler
type handler struct {
	basic.WindowHandler
	basic.MouseHandler
}

//------------------------------------------------------------------------------

var writer = mtx.Clip{
	Left: 1, Top: 1,
	Right: -20, Bottom: -2,
	VScroll: true,
	HScroll: false,
}

func (h handler) WindowResized(s pixel.Coord, _ uint32) {
	sx, sy := window.Size().Cartesian()
	r := sx / sy
	projection = space.Perspective(math.Pi/4, r, 0.001, 1000.0)

	// MTX

	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			mtx.Poke(-16+x, -16+y, byte(x+15*y))
		}
	}
	// for y := 0; y < 60; y++ {
	// 	for x := 0; x < 200; x++ {
	// 		mtx.Poke(x, y, ('A'+byte(x+16*y)&0x1F)|0x80)
	// 	}
	// }

	// w.SetClearChar(' ')
	// writer.Clear()
	writer.Locate(0, 0)
	writer.Print("\tworld\n")
	// w.Print("Essai %x\n", 33333)
	// w.Print("123456\b\b\b\b\b\b\aBoo!\n\a")

	if true {
		writer.Print(`
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec libero
ligula, consectetur at congue et, ultricies placerat velit. Pellentesque
finibus tristique orci sit amet pharetra. Nullam arcu urna, tempus
malesuada aliquet quis, semper blandit ante. Proin vitae dignissim lacus.
Etiam in rutrum tortor. Nulla sed maximus dolor, quis venenatis ante.
Maecenas nec ante vel massa elementum varius nec vitae odio. Proin
tincidunt iaculis elit eu luctus. Donec dignissim ipsum in orci congue
rutrum a at turpis. Aliquam congue tristique dapibus. Pellentesque sed
aliquam ex, id blandit metus. Mauris egestas magna quis elit dignissim, a
laoreet sem facilisis. Duis tristique dapibus dictum. Lorem ipsum dolor
sit amet, consectetur adipiscing elit.
`)
	}

	// w.Locate(-2, 1)
	// w.Print("hello")
}

//------------------------------------------------------------------------------

func (h handler) MouseWheel(motion pixel.Coord, _ uint32) {
	distance -= float32(motion.Y) / 4
	updateView()
}

var ccc int

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ uint32) {
	// writer.SetClearChar('*')
	// writer.Scroll(0, -1)
	writer.Print("\n%v", ccc)
	ccc++
	mouse.SetRelativeMode(true)
}

func (h handler) MouseButtonUp(b mouse.Button, _ int, _ uint32) {
	mouse.SetRelativeMode(false)
}

func (h handler) MouseMotion(motion pixel.Coord, _ pixel.Coord, _ uint32) {
	mx, my := motion.Cartesian()
	sx, sy := window.Size().Cartesian()

	switch {
	case mouse.IsPressed(mouse.Left):
		position.X += 2 * mx / sx
		position.Y -= 2 * my / sy
		updateModel()

	case mouse.IsPressed(mouse.Right):
		yaw += 4 * mx / sx
		pitch += 4 * my / sy
		switch {
		case pitch < -math.Pi/2:
			pitch = -math.Pi / 2
		case pitch > +math.Pi/2:
			pitch = +math.Pi / 2
		}
		updateModel()
	}
}

//------------------------------------------------------------------------------
