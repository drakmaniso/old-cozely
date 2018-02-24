// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var curScreen = pixel.NewCanvas(pixel.Zoom(2))

var cursor = pixel.NewCursor()

func init() {
	cursor.Canvas(curScreen)
}

var (
	pixop9  = pixel.NewFont("fonts/pixop9")
	pixop11 = pixel.NewFont("fonts/pixop11")
)

//------------------------------------------------------------------------------

func TestCursor_print(t *testing.T) {
	do(func() {
		err := glam.Run(curLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type curLoop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (curLoop) Enter() error {
	palette.Load("graphics/shape1")
	return nil
}

//------------------------------------------------------------------------------

func (curLoop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (curLoop) Draw() error {
	curScreen.Clear(21)

	cursor.Locate(16, 8, 0)

	cursor.Font(pixop9)
	cursor.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	cursor.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	cursor.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	cursor.Println("12+34 56-78 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	cursor.Println()

	cursor.Font(pixop11)
	cursor.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	cursor.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	cursor.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	cursor.Println("12+34 56-78 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	cursor.Println()

	cursor.Font(pixop9)
	cursor.Println("My affection for my guest increases every day. He excites at once my")
	cursor.Println("admiration and my pity to an astonishing degree. How can I see so noble")
	cursor.Println("a creature destroyed by misery, without feeling the most poignant grief?")
	cursor.Println("He is so gentle, yet so wise; his mind is so cultivated; and when he")
	cursor.Println("speaks, although his words are culled with the choicest art, yet they")
	cursor.Println("flow with rapidity and unparalleled eloquence.")
	cursor.Println()

	cursor.Font(pixop11)
	cursor.Println("My affection for my guest increases every day. He excites at once my")
	cursor.Println("admiration and my pity to an astonishing degree. How can I see so noble")
	cursor.Println("a creature destroyed by misery, without feeling the most poignant grief?")
	cursor.Println("He is so gentle, yet so wise; his mind is so cultivated; and when he")
	cursor.Println("speaks, although his words are culled with the choicest art, yet they")
	cursor.Println("flow with rapidity and unparalleled eloquence.")
	cursor.Println()

	cursor.Font(pixop9)
	cursor.Locate(500, 100, 0)
	cursor.Write([]byte("Foo"))
	cursor.Move(1, 3, 0)
	cursor.WriteRune('B')
	cursor.Move(2, 2, 0)
	cursor.WriteRune('a')
	cursor.Move(3, 1, 0)
	cursor.WriteRune('r')
	cursor.MoveTo(532, 132, 0)
	cursor.Flush()
	cursor.Write([]byte("Boo\n"))
	cursor.Write([]byte("Choo"))
	cursor.Flush()

	cursor.Locate(curScreen.Size().X-200, 2, 0)
	cursor.Font(pixop11)
	cursor.Printf("Position x=%d, y=%d\n", curScreen.Mouse().X, curScreen.Mouse().Y)

	curScreen.Display()
	return nil
}

//------------------------------------------------------------------------------
