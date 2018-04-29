// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package vector_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/palettes/msx2"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/x/vector"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop struct{}

// Initialization //////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		err := cozely.Run(loop{})
		if err != nil {
			panic(err)
		}
	})
}

func (loop) Enter() {
	msx2.Palette.Activate()
}

func (loop) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop) Update() {
}

func (loop) React() {
}

func (loop) Render() {
	vector.Line(color.SRGB{1, 0.5, 0}, 10, 10, 100, 100)
	w := cozely.WindowSize()
	m := coord.CR{} //input.Mouse.CR()
	vector.Line(color.SRGB{1, 1, 1}, w.C/2, w.R/2, m.C, m.R)
}
