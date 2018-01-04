// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/mouse"
	"github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

var tilesPict [8]struct {
	normal, big *pixel.Picture
}

var currentX, currentY int8

const tileSize = int16(20)

var gridOrigin pixel.Coord

//------------------------------------------------------------------------------

func main() {
	err := carol.Run(loop{})
	if err != nil {
		carol.ShowError(err)
	}
}

//------------------------------------------------------------------------------

type loop struct {
	carol.Handlers
}

//------------------------------------------------------------------------------

func (loop) Setup() error {
	pixel.SetBackground(colour.SRGB8{0x5C, 0x82, 0x86})

	for i, n := range []string{
		"red",
		"yellow",
		"green",
		"blue",
		"violet",
		"pink",
		"dark",
		"multi",
	} {
		tilesPict[i].normal = pixel.GetPicture(n)
		tilesPict[i].big = pixel.GetPicture(n + "_big")
	}

	currentX, currentY = -1, -1

	fillGrid()

	return nil
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw(_, _ float64) error {
	sysDraw()
	return nil
}

//------------------------------------------------------------------------------

func (loop) MouseButtonDown(_ mouse.Button, _ int) {
	m := pixel.Mouse()
	x := (int16(m.X) - gridOrigin.X) / tileSize
	y := (int16(m.Y) - gridOrigin.Y) / tileSize
	if 0 <= x && x < 8 && 0 <= y && y < 8 {
		currentX, currentY = int8(x), int8(y)
	}
}

func (loop) MouseButtonUp(_ mouse.Button, _ int) {
	currentX, currentY = -1, -1
}

//------------------------------------------------------------------------------

func (loop) ScreenResized(width, height int16, _ int32) {
	gridOrigin.X = (width - (8 * tileSize)) / 2
	gridOrigin.Y = (height - (8 * tileSize)) / 2
}

//------------------------------------------------------------------------------
