// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"time"

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

func (h handler) WindowResized(s pixel.Coord, _ time.Duration) {
	sx, sy := window.Size().Cartesian()
	r := sx / sy
	projection = space.Perspective(math.Pi/4, r, 0.001, 1000.0)

	// MTX

	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			mtx.Poke(-16+x, -16+y, byte(x+16*y))
		}
	}

	w := mtx.Writer{}
	// w.Clip(2, 2, -2, -2)
	// w.SetClearChar('*')
	// w.Clear()
	w.SetOverwrite(true)
	w.Print("Hello world!\n")
	w.Print("Essai %x\n", 33333)
	w.Print("123456\b\b\b\b\b\b\aBoo!\n\a")

	if false {
		mtx.Print(2, 5,
			`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec libero ligula,
consectetur at congue et, ultricies placerat velit. Pellentesque finibus
tristique orci sit amet pharetra. Nullam arcu urna, tempus malesuada aliquet
quis, semper blandit ante. Proin vitae dignissim lacus. Etiam in rutrum
tortor. Nulla sed maximus dolor, quis venenatis ante. Maecenas nec ante vel
massa elementum varius nec vitae odio. Proin tincidunt iaculis elit eu luctus.
Donec dignissim ipsum in orci congue rutrum a at turpis. Aliquam congue
tristique dapibus. Pellentesque sed aliquam ex, id blandit metus. Mauris
egestas magna quis elit dignissim, a laoreet sem facilisis. Duis tristique
dapibus dictum. Lorem ipsum dolor sit amet, consectetur adipiscing elit.`)
	}

	w.Locate(0, -1)
	w.Print("Abc \a%f,\a %10.3f", 1.2345678901234567890, 1.2345678901234567890)

}

//------------------------------------------------------------------------------

func (h handler) MouseWheel(motion pixel.Coord, _ time.Duration) {
	distance -= float32(motion.Y) / 4
	updateView()
}

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ time.Duration) {
	mouse.SetRelativeMode(true)
}

func (h handler) MouseButtonUp(b mouse.Button, _ int, _ time.Duration) {
	mouse.SetRelativeMode(false)
}

func (h handler) MouseMotion(motion pixel.Coord, _ pixel.Coord, _ time.Duration) {
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
