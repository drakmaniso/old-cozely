// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import "github.com/drakmaniso/glam"

//------------------------------------------------------------------------------

func main() {
	g := &game{}

	glam.Loop = g

	err := glam.Run()
	if err != nil {
		glam.ErrorDialog(err)
	}
}

//------------------------------------------------------------------------------

type game struct{}

func (g *game) Update() {
}

func (g *game) Draw() {
}

//------------------------------------------------------------------------------
