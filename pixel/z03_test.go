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
	a.canvas.Cursor().Color = a.fg - 1
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

	a.canvas.Text(a.fg, pixel.Monozela10)

	a.canvas.Locate(coord.CR{2, 8})
	a.canvas.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	a.canvas.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	a.canvas.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	a.canvas.Println("12+34 56-7.8 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	a.canvas.Println()

	a.canvas.Locate(coord.CR{16, 100})
	a.canvas.Write([]byte("Foo"))
	a.canvas.Cursor().Position = a.canvas.Cursor().Position.Plus(coord.CR{1, 3})
	a.canvas.WriteRune('B')
	a.canvas.Cursor().Position = a.canvas.Cursor().Position.Plus(coord.CR{2, 2})
	a.canvas.WriteRune('a')
	a.canvas.Cursor().Position = a.canvas.Cursor().Position.Plus(coord.CR{3, 1})
	a.canvas.WriteRune('r')
	a.canvas.Cursor().Position = coord.CR{32, 132}
	a.canvas.Write([]byte("Boo\n"))
	a.canvas.Write([]byte("Choo"))

	a.canvas.Locate(coord.CR{16, 200})
	a.canvas.Cursor().Font = a.tinela9
	a.canvas.Print("Tinela")
	a.canvas.Cursor().Font = a.simpela10
	a.canvas.Print("Simpela10")
	a.canvas.Cursor().Font = a.simpela12
	a.canvas.Print("Simpela12")
	a.canvas.Cursor().Font = a.cozela10
	a.canvas.Print("Cozela10")
	a.canvas.Cursor().Font = a.cozela12
	a.canvas.Print("Cozela12")
	a.canvas.Cursor().Font = a.chaotela12
	a.canvas.Print("Chaotela12")

	a.canvas.Locate(coord.CR{a.canvas.Size().C - 200, 9})
	a.canvas.Cursor().Font = pixel.FontID(0)
	m := a.canvas.FromWindow(cursor.XY(0).CR())
	a.canvas.Printf("Position x=%d, y=%d\n", m.C, m.R)

	a.canvas.Display()
}
