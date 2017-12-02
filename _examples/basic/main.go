// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"errors"

	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/picture"
)

//------------------------------------------------------------------------------

func main() {
	err := carol.Run(loop{})
	if err != nil {
		carol.ShowError("in game loop", err)
		return
	}
}

//------------------------------------------------------------------------------

type loop struct {
	carol.Handlers
}

func (loop) Setup() error {
	p, err := colour.NewPalette("Foo")
	if err != nil {
		return carol.Error("while creating palette", err)
	}
	p.NewColour("dark grey", colour.RGBA{0.1, 0.1, 0.1, 1.0})
	p.NewColour("orange", colour.RGBA{1.0, 0.5, 0.0, 1.0})
	p.NewColour("violet", colour.RGBA{1.0, 0.0, 0.5, 1.0})
	p.NewColour("turquoise", colour.RGBA{0.0, 1.0, 0.5, 1.0})

	return nil
}

func (loop) Update() error {
	return nil
}

var x = int16(1)

func (loop) Draw(_, _ float64) error {
	x++
	if x > 300 {
		x = 1
	}

	p, ok := picture.Named("logo")
	if !ok && false {
		return errors.New("picture not found")
	}
	p.Paint(x, 10)
	_ = p

	p2, ok := picture.Named("mire")
	if !ok && false {
		return errors.New("picture not found")
	}
	p2.Paint(10, 60)

	p.Paint(40, 30)

	return nil
}

//------------------------------------------------------------------------------
