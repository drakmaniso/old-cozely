// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/drakmaniso/cozely"
	"github.com/drakmaniso/cozely/input"
	"github.com/drakmaniso/cozely/palette"
	"github.com/drakmaniso/cozely/pixel"
	"github.com/drakmaniso/cozely/plane"
)

//------------------------------------------------------------------------------

var (
	curScreen = pixel.Canvas(pixel.Zoom(2))
	cursor    = pixel.Cursor{Canvas: curScreen}
)

var curBg, curFg palette.Index

//------------------------------------------------------------------------------

func TestCursor_print(t *testing.T) {
	do(func() {
		err := cozely.Run(curLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type curLoop struct{}

//------------------------------------------------------------------------------

func (curLoop) Enter() error {
	input.Load(testBindings)
	testContext.Activate(1)
	palette.Load("C64")
	curBg = palette.Find("white")
	curFg = palette.Find("black")
	cursor.Color = curFg - 1
	return nil
}

func (curLoop) Leave() error { return nil }

//------------------------------------------------------------------------------

func (curLoop) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}
	return nil
}

//------------------------------------------------------------------------------

func (curLoop) Update() error { return nil }

//------------------------------------------------------------------------------

func (curLoop) Render() error {
	curScreen.Clear(curBg)

	cursor.Locate(2, 8)

	cursor.Font = pixel.FontID(0)
	cursor.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	cursor.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	cursor.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	cursor.Println("12+34 56-7.8 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	cursor.Println()

	cursor.Font = pixel.FontID(0)
	cursor.Locate(16, 100)
	cursor.Write([]byte("Foo"))
	cursor.Position = cursor.Position.Pluss(1, 3)
	cursor.WriteRune('B')
	cursor.Position = cursor.Position.Pluss(2, 2)
	cursor.WriteRune('a')
	cursor.Position = cursor.Position.Pluss(3, 1)
	cursor.WriteRune('r')
	cursor.Position = plane.Pixel{32, 132}
	cursor.Write([]byte("Boo\n"))
	cursor.Write([]byte("Choo"))

	cursor.Locate(16, 200)
	cursor.Font = tinela9
	cursor.Print("Tinela")
	cursor.Font = simpela10
	cursor.Print("Simpela10")
	cursor.Font = simpela12
	cursor.Print("Simpela12")
	cursor.Font = cozela10
	cursor.Print("Cozela10")
	cursor.Font = cozela12
	cursor.Print("Cozela12")
	cursor.Font = chaotela12
	cursor.Print("Chaotela12")

	cursor.Locate(curScreen.Size().X-200, 9)
	cursor.Font = pixel.FontID(0)
	cursor.Printf("Position x=%d, y=%d\n", curScreen.Mouse().X, curScreen.Mouse().Y)

	curScreen.Display()
	return nil
}

//------------------------------------------------------------------------------
