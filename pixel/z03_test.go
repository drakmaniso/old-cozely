// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var ()

type loop3 struct {
	bg, fg color.Index

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
	pixel.SetZoom(2)
	//TODO:
	a.bg = 7
	a.fg = 1

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
	pixel.Clear(a.bg)

	cur := pixel.Cursor{}

	cur.Style(a.fg, pixel.Monozela10)

	cur.Locate(0, pixel.XY{2, 8})
	cur.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	cur.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	cur.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	cur.Println("12+34 56-7.8 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	cur.Println()

	cur.Locate(0, pixel.XY{16, 100})
	cur.Write([]byte("Foo"))
	cur.Position = cur.Position.Plus(pixel.XY{1, 3})
	cur.WriteRune('B')
	cur.Position = cur.Position.Plus(pixel.XY{2, 2})
	cur.WriteRune('a')
	cur.Position = cur.Position.Plus(pixel.XY{3, 1})
	cur.WriteRune('r')
	cur.Position = pixel.XY{32, 132}
	cur.Write([]byte("Boo\n"))
	cur.Write([]byte("Choo"))

	cur.Locate(0, pixel.XY{16, 200})
	cur.Font = a.tinela9
	cur.Print("Tinela")
	cur.Font = a.simpela10
	cur.Print("Simpela10")
	cur.Font = a.simpela12
	cur.Print("Simpela12")
	cur.Font = a.cozela10
	cur.Print("Cozela10")
	cur.Font = a.cozela12
	cur.Print("Cozela12")
	cur.Font = a.chaotela12
	cur.Print("Chaotela12")

	cur.Locate(0, pixel.XY{pixel.Resolution().X - 200, 9})
	cur.Font = pixel.FontID(0)
	m := pixel.XYof(cursor.XY(0))
	cur.Printf("Position x=%d, y=%d\n", m.X, m.Y)
}
