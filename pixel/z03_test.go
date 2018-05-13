// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var ()

type loop3 struct {
	canvas  pixel.CanvasID
	scene pixel.SceneID
	palette color.PaletteID
	bg, fg  color.Index

	tinela9, monozela10, simpela10, simpela12,
	cozela10, cozela12, chaotela12, font pixel.FontID
}

////////////////////////////////////////////////////////////////////////////////

func TestTest3(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop3{}
		l.declare()

		input.Load(bindings)
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (a *loop3) declare() {
	a.canvas = pixel.Canvas(pixel.Zoom(2))
	a.scene = pixel.Scene()
	a.palette = color.Palette()
	a.bg = a.palette.Entry(color.SRGB8{0xFF, 0xFE, 0xFC})
	a.fg = a.palette.Entry(color.SRGB8{0x07, 0x05, 0x00})

	a.tinela9 = pixel.Font("fonts/tinela9")
	a.monozela10 = pixel.Font("fonts/monozela10")
	a.simpela10 = pixel.Font("fonts/simpela10")
	a.simpela12 = pixel.Font("fonts/simpela12")
	a.cozela10 = pixel.Font("fonts/cozela10")
	a.cozela12 = pixel.Font("fonts/cozela12")
	a.chaotela12 = pixel.Font("fonts/chaotela12")
	a.font = a.monozela10
}

func (a *loop3) Enter() {
	a.palette.Activate()
	println(a.bg, a.fg)
	a.scene.Cursor().Color = a.fg - 1
}

func (loop3) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop3) React() {
	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop3) Update() {
}

func (a *loop3) Render() {
	a.canvas.Clear(a.bg)
	a.scene.Clear()

	a.scene.Text(a.fg, pixel.Monozela10)

	a.scene.Locate(coord.CR{2, 8})
	a.scene.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	a.scene.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	a.scene.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	a.scene.Println("12+34 56-7.8 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	a.scene.Println()

	a.scene.Locate(coord.CR{16, 100})
	a.scene.Write([]byte("Foo"))
	a.scene.Cursor().Position = a.scene.Cursor().Position.Plus(coord.CR{1, 3})
	a.scene.WriteRune('B')
	a.scene.Cursor().Position = a.scene.Cursor().Position.Plus(coord.CR{2, 2})
	a.scene.WriteRune('a')
	a.scene.Cursor().Position = a.scene.Cursor().Position.Plus(coord.CR{3, 1})
	a.scene.WriteRune('r')
	a.scene.Cursor().Position = coord.CR{32, 132}
	a.scene.Write([]byte("Boo\n"))
	a.scene.Write([]byte("Choo"))

	a.scene.Locate(coord.CR{16, 200})
	a.scene.Cursor().Font = a.tinela9
	a.scene.Print("Tinela")
	a.scene.Cursor().Font = a.simpela10
	a.scene.Print("Simpela10")
	a.scene.Cursor().Font = a.simpela12
	a.scene.Print("Simpela12")
	a.scene.Cursor().Font = a.cozela10
	a.scene.Print("Cozela10")
	a.scene.Cursor().Font = a.cozela12
	a.scene.Print("Cozela12")
	a.scene.Cursor().Font = a.chaotela12
	a.scene.Print("Chaotela12")

	a.scene.Locate(coord.CR{a.canvas.Size().C - 200, 9})
	a.scene.Cursor().Font = pixel.FontID(0)
	m := a.canvas.FromWindow(cursor.XY(0).CR())
	a.scene.Printf("Position x=%d, y=%d\n", m.C, m.R)

	a.canvas.Display(a.scene)
}
