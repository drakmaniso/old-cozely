// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package vector_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/palettes/msx2"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/x/vector"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop struct{}

// Initialization //////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		err := cozely.Run(loop{})
		if err != nil {
			cozely.ShowError(err)
		}
	})
}

func (loop) Enter() error {
	msx2.Palette.Activate()
	return nil
}

func (loop) Leave() error {
	return nil
}

// Game Loop ///////////////////////////////////////////////////////////////////

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
	vector.Line(color.SRGB{1, 1, 1}, w.C/2, w.R/2, m.C, m.R)
	return nil
}
