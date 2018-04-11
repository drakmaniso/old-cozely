// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/cozely"
	"github.com/drakmaniso/cozely/colour"
	"github.com/drakmaniso/cozely/mouse"
	"github.com/drakmaniso/cozely/palette"
	"github.com/drakmaniso/cozely/x/vector"
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	err := cozely.Run(loop{})
	if err != nil {
		cozely.ShowError(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop struct {
	cozely.EmptyLoop
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Enter() error {
	palette.Load("MSX2")
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() error {
	return nil
}

func (loop) Draw() error {
	vector.Line(colour.SRGB{1, 0.5, 0}, 10, 10, 100, 100)
	wx, wy := cozely.WindowSize()
	mx, my := mouse.Position()
	vector.Line(colour.SRGB{1, 1, 1}, int16(wx/2), int16(wy/2), int16(mx), int16(my))
	return nil
}

////////////////////////////////////////////////////////////////////////////////
