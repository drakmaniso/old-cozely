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
	gfx.PaletteMSX2()

	logo = gfx.GetPicture("logo")
	mire = gfx.GetPicture("mire")

	return gfx.Err()
}

func (loop) Update() error {
	x++
	if x >= 320 {
		x = -64
	}
	return nil
}

var x = int16(0)

var (
	timer = 0.0
	count = 0
)

func (loop) Draw(delta, _ float64) error {
	timer += delta
	// if timer > 0.25 {
	// 	count++
	// 	timer = 0.0
	// 	if count%2 != 0 {
	// 		gfx.Color(1).SetRGBA(gfx.RGBA{1, 1, 1, 1})
	// 	} else {
	// 		gfx.Color(1).SetRGBA(gfx.RGBA{1, 0, 0.5, 1})
	// 	}
	// }

	logo.Paint(x, 10)

	mire.Paint(0, 0)
	mire.Paint(320-32, 0)
	mire.Paint(0, 180-32)
	mire.Paint(320-32, 180-32)

	logo.Paint(320/2-32, 20)

	return gfx.Err()
}

//------------------------------------------------------------------------------
