// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/x/vector"
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	err := cozely.Run(loop{})
	if err != nil {
		cozely.ShowError(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop) Enter() error {
	palette.Load("MSX2")
	return nil
}

func (loop) Leave() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() error {
	return nil
}

func (loop) React() error {
	return nil
}

func (loop) Render() error {
	vector.Line(color.SRGB{1, 0.5, 0}, 10, 10, 100, 100)
	w := cozely.WindowSize()
	m := input.Cursor.Position()
	vector.Line(color.SRGB{1, 1, 1}, w.X/2, w.Y/2, m.X, m.Y)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
