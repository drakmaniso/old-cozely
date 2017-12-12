// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/pixel"
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

var logo, mire pixel.Picture

func (loop) Setup() error {
	pixel.PaletteMSX2()

	logo = pixel.GetPicture("logo")
	mire = pixel.GetPicture("mire")

	return pixel.Err()
}

func (loop) Update() error {
	x++
	if x >= pixel.ScreenSize().X {
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
	// 		pixel.Color(2).SetRGBA(pixel.RGBA{1, 1, 1, 1})
	// 	} else {
	// 		pixel.Color(2).SetRGBA(pixel.RGBA{1, 0, 0.5, 1})
	// 	}
	// }

	logo.Paint(x, 10)

	s := pixel.ScreenSize()

	mire.Paint(0, 0)
	mire.Paint(s.X-32, 0)
	mire.Paint(0, s.Y-32)
	mire.Paint(s.X-32, s.Y-32)

	logo.Paint(s.X/2-32, 20)

	return pixel.Err()
}

//------------------------------------------------------------------------------
