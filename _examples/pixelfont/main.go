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
	err := pixel.LoadFont("fonts/pixop11.font.png")
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
	h := pixel.Font(0).Height() + 2

	s.Print(0, 253, 20, 20+0*h, "A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	s.Print(0, 253, 20, 20+1*h, "a quick brown fox \"jumps\" over the (lazy) dog.")
	s.Print(0, 253, 20, 20+2*h, "0123456789!@#$%^&*()-+=_~[]{}|\\;:'\",.<>/?")
	s.Print(0, 253, 20, 20+3*h, "12+34 56-78 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")

	s.Print(0, 253, 20, 100+0*h, "I am by birth a Genevese, and my family is one of the most distinguished of that republic. My ancestors had been for many years counsellors and syndics, and my")
	s.Print(0, 253, 20, 100+1*h, "father had filled several public situations with honour and reputation. He was respected by all who knew him for his integrity and indefatigable attention to public")
	s.Print(0, 253, 20, 100+2*h, "business. He passed his younger days perpetually occupied by the affairs of his country; a variety of circumstances had prevented his marrying early, nor was it until")
	s.Print(0, 253, 20, 100+3*h, "the decline of life that he became a husband and the father of a family.")

	s.Print(0, 253, 20, 200+0*h, "Any question? Or answers! No matter what, if it's pertinent... Anyway; no more (important) punctuation. The flow often.")
	s.Blit()
	return nil
}

//------------------------------------------------------------------------------
