// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var cursor *pixel.Cursor

//------------------------------------------------------------------------------

func main() {
	glam.Configure(
		pixel.Zoom(2),
	)
	err := glam.Run(setup, loop{})
	if err != nil {
		glam.ShowError(err)
	}
}

//------------------------------------------------------------------------------

func setup() error {
	palette.Change("MSX2")
	err := pixel.LoadFont("fonts/pixop9.font.png")
	if err != nil {
		return err
	}
	err = pixel.LoadFont("fonts/pixop11.font.png")
	if err != nil {
		return err
	}

	pixel.SetBackground(palette.Index(255))

	cursor = pixel.Screen().NewCursor()
	cursor.ColorShift(0x20-1)

	return nil
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw() error {
	s := pixel.Screen()

	cursor.Locate(16, 16)

	cursor.Font(0)
	cursor.Write([]byte("Foo"))
	cursor.Write([]byte("Bar"))
	cursor.Flush()
	cursor.Write([]byte("Boo"))
	cursor.Flush()

	cursor.Font(0)
	cursor.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	cursor.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	cursor.Println("0123456789!@#$%^&*()-+=_~[]{}|\\;:'\",.<>/?")
	cursor.Println("12+34 56-78 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	cursor.Println()

	cursor.Font(1)
	cursor.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	cursor.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	cursor.Println("0123456789!@#$%^&*()-+=_~[]{}|\\;:'\",.<>/?")
	cursor.Println("12+34 56-78 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	cursor.Println()

	cursor.Font(0)
	cursor.Println("My affection for my guest increases every day. He excites at once my")
	cursor.Println("admiration and my pity to an astonishing degree. How can I see so noble")
	cursor.Println("a creature destroyed by misery, without feeling the most poignant grief?")
	cursor.Println("He is so gentle, yet so wise; his mind is so cultivated; and when he")
	cursor.Println("speaks, although his words are culled with the choicest art, yet they")
	cursor.Println("flow with rapidity and unparalleled eloquence.")
	cursor.Println()

	cursor.Font(1)
	cursor.Println("My affection for my guest increases every day. He excites at once my")
	cursor.Println("admiration and my pity to an astonishing degree. How can I see so noble")
	cursor.Println("a creature destroyed by misery, without feeling the most poignant grief?")
	cursor.Println("He is so gentle, yet so wise; his mind is so cultivated; and when he")
	cursor.Println("speaks, although his words are culled with the choicest art, yet they")
	cursor.Println("flow with rapidity and unparalleled eloquence.")
	cursor.Println()

	cursor.Locate(s.Size().X-200, 2)
	cursor.Font(1)
	cursor.Printf("Position x=%d, y=%d\n", s.Mouse().X, s.Mouse().Y)

	s.Blit()
	return nil
}

//------------------------------------------------------------------------------
