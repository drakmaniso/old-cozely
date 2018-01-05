// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/mouse"
	"github.com/drakmaniso/carol/pixel"

	"github.com/drakmaniso/carol/_examples/match3/grid"
)

//------------------------------------------------------------------------------

var tilesPict [8]struct {
	normal, big *pixel.Picture
}

var current grid.Position

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

	current = grid.Nowhere()

	grid.Setup(8, 8)
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
	current = grid.PositionAt(m)
}

func (loop) MouseButtonUp(_ mouse.Button, _ int) {
	current = grid.Nowhere()
}

//------------------------------------------------------------------------------

func (loop) ScreenResized(width, height int16, _ int32) {
	grid.ScreenResized(width, height)
}

//------------------------------------------------------------------------------
