// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

/*
Package cozely is the starting point of the Cozely framework, a simple,
all-in-one set of packages for making games in Go.

It focuses on pixel art for 2D, and polygonal art (aka low-poly) for 3D,
supports windows and linux, and has only two dependencies (SDL2 and Opengl 4.6).

**THIS IS A WORK IN PROGRESS**, not usable yet: the framework is *very*
incomplete, and the API is subject to frequent changes.

Hello World example:

	package main

	import (
		"github.com/cozely/cozely"
		"github.com/cozely/cozely/pixel"
	)

	func main() {
		cozely.Run(loop{})
	}

	type loop struct{}

	func (loop) Enter()  {
		// Enter the game loop
	}

	func (loop) React()  {
		// React to user inputs
	}

	func (loop) Update() {
		// Update the game state
	}

	func (loop) Render() {
		// Render the game state
		pixel.Clear(1)
		cur := pixel.Cursor{
			Position: pixel.XY{8, 16},
			Color: 7,
		}
		cur.Print("Hello, World!")
	}

	func (loop) Leave()  {
		// Leave the game loop
	}

*/
package cozely
