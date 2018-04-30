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

var (
	canvas3  = pixel.Canvas(pixel.Zoom(2))
	palette3 = color.Palette()
	bg3      = palette3.Entry(color.SRGB8{0xFF, 0xFE, 0xFC})
	fg3      = palette3.Entry(color.SRGB8{0x07, 0x05, 0x00})
)

type loop3 struct{}

////////////////////////////////////////////////////////////////////////////////

func TestTest3(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		input.Load(bindings)
		err := cozely.Run(loop3{})
		if err != nil {
			t.Error(err)
		}
	})
}

func (loop3) Enter() {
	palette3.Activate()
	println(bg3, fg3)
	canvas3.Cursor().Color = fg3 - 1
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

func (loop3) Render() {
	canvas3.Clear(bg3)

	canvas3.Text(fg3, pixel.Monozela10)

	canvas3.Locate(coord.CR{2, 8})
	canvas3.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	canvas3.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	canvas3.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	canvas3.Println("12+34 56-7.8 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	canvas3.Println()

	canvas3.Locate(coord.CR{16, 100})
	canvas3.Write([]byte("Foo"))
	canvas3.Cursor().Position = canvas3.Cursor().Position.Plus(coord.CR{1, 3})
	canvas3.WriteRune('B')
	canvas3.Cursor().Position = canvas3.Cursor().Position.Plus(coord.CR{2, 2})
	canvas3.WriteRune('a')
	canvas3.Cursor().Position = canvas3.Cursor().Position.Plus(coord.CR{3, 1})
	canvas3.WriteRune('r')
	canvas3.Cursor().Position = coord.CR{32, 132}
	canvas3.Write([]byte("Boo\n"))
	canvas3.Write([]byte("Choo"))

	canvas3.Locate(coord.CR{16, 200})
	canvas3.Cursor().Font = tinela9
	canvas3.Print("Tinela")
	canvas3.Cursor().Font = simpela10
	canvas3.Print("Simpela10")
	canvas3.Cursor().Font = simpela12
	canvas3.Print("Simpela12")
	canvas3.Cursor().Font = cozela10
	canvas3.Print("Cozela10")
	canvas3.Cursor().Font = cozela12
	canvas3.Print("Cozela12")
	canvas3.Cursor().Font = chaotela12
	canvas3.Print("Chaotela12")

	canvas3.Locate(coord.CR{canvas3.Size().C - 200, 9})
	canvas3.Cursor().Font = pixel.FontID(0)
	m := canvas3.FromWindow(cursor.XY(0).CR())
	canvas3.Printf("Position x=%d, y=%d\n", m.C, m.R)

	canvas3.Display()
}
