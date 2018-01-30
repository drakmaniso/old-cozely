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
	c.Println("I am by birth a Genevese, and my family is one of the most distinguished of that republic. My ancestors had been for many years counsellors and syndics, and my")
	c.Println("father had filled several public situations with honour and reputation. He was respected by all who knew him for his integrity and indefatigable attention to public")
	c.Println("business. He passed his younger days perpetually occupied by the affairs of his country; a variety of circumstances had prevented his marrying early, nor was it until")
	c.Println("the decline of life that he became a husband and the father of a family.")
	c.Println("Any question? Or answers! No matter what, if it's pertinent... Anyway; no more (important) punctuation. The flow often.")
	c.Println()

	c.Font(1)
	c.Println("I am by birth a Genevese, and my family is one of the most distinguished of that republic. My ancestors had been for many years counsellors and syndics, and my")
	c.Println("father had filled several public situations with honour and reputation. He was respected by all who knew him for his integrity and indefatigable attention to public")
	c.Println("business. He passed his younger days perpetually occupied by the affairs of his country; a variety of circumstances had prevented his marrying early, nor was it until")
	c.Println("the decline of life that he became a husband and the father of a family.")
	c.Println("Any question? Or answers! No matter what, if it's pertinent... Anyway; no more (important) punctuation. The flow often.")
	c.Println()


	c.Locate(s.Size().X-200, 2)
	c.Font(0)
	c.Printf("Position x=%d, y=%d\n", s.Mouse().X, s.Mouse().Y)

	s.Blit()
	return nil
}

//------------------------------------------------------------------------------
