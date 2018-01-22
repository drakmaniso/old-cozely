// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

func main() {
	err := glam.Run(setup, loop{})
	if err != nil {
		glam.ShowError(err)
	}
}

//------------------------------------------------------------------------------

func setup() error {
	pixel.LoadFont("fonts/pixop_mono_7x11.png")
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
	s.Print(0, 32, 20, 20, "Ceci est un essai")
	s.Blit()
	return nil
}

//------------------------------------------------------------------------------
