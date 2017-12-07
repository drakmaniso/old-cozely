// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/gfx"
)

//------------------------------------------------------------------------------

func main() {
	err := carol.Run(loop{})
	if err != nil {
		carol.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

type loop struct {
	carol.Handlers
}

var logo, mire gfx.Picture

func (loop) Setup() error {
	gfx.NewPaletteMSX2()

	logo = gfx.GetPicture("logo")
	mire = gfx.GetPicture("mire")

	return gfx.Err()
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
		// if count%2 != 0 {
		// 	gfx.Palette(0).SetRGBA(2, gfx.RGBA{1, 1, 1, 1})
		// } else {
		// 	gfx.Palette(0).SetRGBA(2, gfx.RGBA{1, 0, 0.5, 1})
		// }
	}

	x++
	if x > 300 {
		x = 1
	}
	logo.Paint(x, 10)

	mire.Paint(10, 40)
	logo.Paint(40, 30)
	p := gfx.GetPicture("msx2")
	p.Paint(8, 64)

	return gfx.Err()
}

//------------------------------------------------------------------------------
