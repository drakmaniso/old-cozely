// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

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

	c := s.NewCursor()
	c.ColorShift(254)
	c.Locate(20, 2)

	c.Font(0)
	c.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	c.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	c.Println("0123456789!@#$%^&*()-+=_~[]{}|\\;:'\",.<>/?")
	c.Println("12+34 56-78 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	c.Println()

	c.Font(1)
	c.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	c.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	c.Println("0123456789!@#$%^&*()-+=_~[]{}|\\;:'\",.<>/?")
	c.Println("12+34 56-78 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	c.Println()

	c.Font(0)
	c.Println("My affection for my guest increases every day. He excites at once my")
	c.Println("admiration and my pity to an astonishing degree. How can I see so noble")
	c.Println("a creature destroyed by misery, without feeling the most poignant grief?")
	c.Println("He is so gentle, yet so wise; his mind is so cultivated; and when he")
	c.Println("speaks, although his words are culled with the choicest art, yet they")
	c.Println("flow with rapidity and unparalleled eloquence.")
	c.Println()

	c.Font(1)
	c.Println("My affection for my guest increases every day. He excites at once my")
	c.Println("admiration and my pity to an astonishing degree. How can I see so noble")
	c.Println("a creature destroyed by misery, without feeling the most poignant grief?")
	c.Println("He is so gentle, yet so wise; his mind is so cultivated; and when he")
	c.Println("speaks, although his words are culled with the choicest art, yet they")
	c.Println("flow with rapidity and unparalleled eloquence.")
	c.Println()


	c.Locate(s.Size().X-200, 2)
	c.Font(1)
	c.Printf("Position x=%d, y=%d\n", s.Mouse().X, s.Mouse().Y)

	s.Blit()
	return nil
}

//------------------------------------------------------------------------------
