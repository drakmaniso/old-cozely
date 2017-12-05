// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"errors"

	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/palette"
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
	p, err := palette.New("Foo")
	if err != nil {
		return carol.Error("while creating palette", err)
	}
	p.New("dark grey", palette.RGBA{0.1, 0.1, 0.1, 1.0})
	p.New("orange", palette.RGBA{1.0, 0.5, 0.0, 1.0})
	p.New("violet", palette.RGBA{1.0, 0.0, 0.5, 1.0})
	p.New("turquoise", palette.RGBA{0.0, 1.0, 0.5, 1.0})

	return nil
}

func (loop) Update() error {
	return nil
}

var x = int16(1)

var (
	timer = 0.0
	count = 0
)

func (loop) Draw(delta, _ float64) error {
	timer += delta
	if timer > 0.25 {
		count++
		timer = 0.0
		if count%2 != 0 {
			palette.Palette(0).SetRGBA(2, palette.RGBA{1, 1, 1, 1})
		} else {
			palette.Palette(0).SetRGBA(2, palette.RGBA{1, 0, 0.5, 1})
		}
	}

	x++
	if x > 300 {
		x = 1
	}

	p, ok := picture.Get("logo")
	if !ok && false {
		return errors.New("picture not found")
	}
	p.Paint(x, 10)
	_ = p

	p2, ok := picture.Get("mire")
	if !ok && false {
		return errors.New("picture not found")
	}
	p2.Paint(10, 60)

	p.Paint(40, 30)

	return nil
}

//------------------------------------------------------------------------------
