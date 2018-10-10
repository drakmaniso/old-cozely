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

		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (l *loop3) declare() {
	pixel.SetZoom(2)
	//TODO:
	l.bg = 7
	l.fg = 1

	l.tinela9 = pixel.Font("fonts/tinela9")
	l.monozela10 = pixel.Font("fonts/monozela10")
	l.simpela10 = pixel.Font("fonts/simpela10")
	l.simpela12 = pixel.Font("fonts/simpela12")
	l.cozela10 = pixel.Font("fonts/cozela10")
	l.cozela12 = pixel.Font("fonts/cozela12")
	l.chaotela12 = pixel.Font("fonts/chaotela12")
	l.font = l.monozela10
}

func (a *loop3) Enter() {
}

func (loop3) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop3) React() {
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
	}
}

func (loop3) Update() {
}

func (l *loop3) Render() {
	pixel.Clear(l.bg)

	cur := pixel.Cursor{}

	cur.Color = l.fg

	cur.Position = pixel.XY{2, 8}
	cur.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	cur.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	cur.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	cur.Println("12+34 56-7.8 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	cur.Println()

	cur.Position = pixel.XY{16, 100}
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

	cur.Position = pixel.XY{16, 200}
	cur.Font = l.tinela9
	cur.Print("Tinela")
	cur.Font = l.simpela10
	cur.Print("Simpela10")
	cur.Font = l.simpela12
	cur.Print("Simpela12")
	cur.Font = l.cozela10
	cur.Print("Cozela10")
	cur.Font = l.cozela12
	cur.Print("Cozela12")
	cur.Font = l.chaotela12
	cur.Print("Chaotela12")

	cur.Position = pixel.XY{pixel.Resolution().X - 200, 9}
	cur.Font = pixel.FontID(0)
	m := pixel.XYof(input.MenuPointer.XY())
	cur.Printf("Position x=%d, y=%d\n", m.X, m.Y)
}
