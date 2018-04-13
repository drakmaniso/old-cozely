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

////////////////////////////////////////////////////////////////////////////////

func TestTest3(t *testing.T) {
	do(func() {
		err := cozely.Run(loop3{})
		if err != nil {
			t.Error(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type loop3 struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop3) Enter() error {
	input.Load(bindings)
	context.Activate(1)
	palette3.Activate()
	println(bg3, fg3)
	canvas3.Cursor().Color = fg3 - 1
	return nil
}

func (loop3) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop3) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop3) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop3) Render() error {
	canvas3.Clear(bg3)

	canvas3.Text(fg3-1, pixel.Monozela10)

	canvas3.Locate(2, 8, 0)
	canvas3.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	canvas3.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	canvas3.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	canvas3.Println("12+34 56-7.8 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	canvas3.Println()

	canvas3.Locate(16, 100, 0)
	canvas3.Write([]byte("Foo"))
	canvas3.Cursor().Position = canvas3.Cursor().Position.Pluss(1, 3, 0)
	canvas3.WriteRune('B')
	canvas3.Cursor().Position = canvas3.Cursor().Position.Pluss(2, 2, 0)
	canvas3.WriteRune('a')
	canvas3.Cursor().Position = canvas3.Cursor().Position.Pluss(3, 1, 0)
	canvas3.WriteRune('r')
	canvas3.Cursor().Position = coord.CRD{32, 132, 0}
	canvas3.Write([]byte("Boo\n"))
	canvas3.Write([]byte("Choo"))

	canvas3.Locate(16, 200, 0)
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

	canvas3.Locate(canvas3.Size().C-200, 9, 0)
	canvas3.Cursor().Font = pixel.FontID(0)
	canvas3.Printf("Position x=%d, y=%d\n", canvas3.Mouse().C, canvas3.Mouse().R)

	canvas3.Display()
	return nil
}

////////////////////////////////////////////////////////////////////////////////
