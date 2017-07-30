// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/overlay"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var hud, hud2, debug, txt *overlay.Overlay

//------------------------------------------------------------------------------

func main() {
	err := glam.Setup()
	if err != nil {
		glam.ShowError("setting up glam", err)
		return
	}

	err = setup()
	if err != nil {
		glam.ShowError("running", err)
		return
	}

	glam.Loop(loop{})

	err = glam.Run()
	if err != nil {
		glam.ShowError("running", err)
		return
	}
}

//------------------------------------------------------------------------------

func setup() error {
	hud = overlay.Create(pixel.Coord{100, 100}, 16, 16)
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			hud.Poke(x, y, byte(x+16*y))
		}
	}

	hud2 = overlay.Create(pixel.Coord{150, 150}, 16, 16)
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			hud2.Poke(x, y, byte(x+16*y))
		}
	}

	debug = overlay.Create(pixel.Coord{0, 0}, 100, 1)
	for x := 0; x < 10; x++ {
		debug.Poke(x, 0, 'a'+byte(x))
	}

	txt := overlay.Create(pixel.Coord{400, 100}, 60, 20)
	txt.Clear()
	// txt.Locate(1, 1)
	// txt.Print("0\n1\n2\n3\n4\n5\n6\n7\n8\n9\nABCDEF")
	txt.Print("Package overlay implements a \"text mode\" overlay, useful for development and debugging.\n")
	txt.Print("\nSpecial characters:\n")
	txt.Print("\t- '\\a': toggle \ahighlight\a\n")
	txt.Print("\t- '\\b': move cursor two\b\b\bone character left\n")
	txt.Print("\t- '\\f': blank\fspace (i.e. fully transparent)\n")
	txt.Print("\t- '\\n': newline\n")
	txt.Print("\t- '\\r': move cursor to beginning of line\n")
	txt.Print("\t- '\\t': tabulation\n")
	txt.Print("\t- '\\v': clear until end of line\n")
	txt.Print("INVISIBLE\r\v")
	txt.Locate(0, 0)
	txt.Print("PLOP\nPLIP\nPLUP\n")

	return nil
}

//------------------------------------------------------------------------------

type loop struct {
	glam.DefaultHandlers
}

func (loop) Update() {
}

func (loop) Draw(_, _ float64) {
	gfx.ClearColorBuffer(color.RGBA{0.1, 0.3, 0.4, 1.0})
}

//------------------------------------------------------------------------------

func (loop) KeyDown(l key.Label, p key.Position) {
	switch p {
	case key.PositionSpace:
		hud.Clear()
	case key.PositionUp:
		hud2.Scroll(0, -1)
	case key.PositionDown:
		hud2.Scroll(0, 1)
	case key.PositionLeft:
		hud2.Scroll(-1, 0)
	case key.PositionRight:
		hud2.Scroll(1, 0)
	default:
		glam.DefaultHandlers{}.KeyDown(l, p)
	}
}

//------------------------------------------------------------------------------
