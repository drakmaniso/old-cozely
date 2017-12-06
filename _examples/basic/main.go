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

func (loop) Setup() error {
	p, err := gfx.NewPalette("MSX2 Palette")
	if err != nil {
		return carol.Error("while creating palette", err)
	}
	for i := 0; i < 256; i++ {
		p.New("", gfx.RGBA{
			float32(i>>5) / 7.0,
			float32((i&0x1C)>>2) / 7.0,
			float32(i&0x3) / 3.0,
			1.0,
		})
	}

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
			gfx.Palette(0).SetRGBA(2, gfx.RGBA{1, 1, 1, 1})
		} else {
			gfx.Palette(0).SetRGBA(2, gfx.RGBA{1, 0, 0.5, 1})
		}
	}

	x++
	if x > 300 {
		x = 1
	}

	p := gfx.GetPicture("logo")
	p.Paint(x, 10)
	_ = p

	p2 := gfx.GetPicture("mire")
	p2.Paint(10, 40)

	p.Paint(40, 30)

	gfx.GetPicture("msx2").Paint(8, 64)

	return gfx.Err()
}

//------------------------------------------------------------------------------
