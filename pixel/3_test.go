package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var ()

type loop3 struct {
	pict1, pict2 pixel.PictureID
	position, size pixel.XY
}

////////////////////////////////////////////////////////////////////////////////

func TestTest3(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop3{}
		l.setup()

		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (l *loop3) setup() {
	pixel.SetZoom(4)
	l.pict1 = pixel.Picture("graphics/9tiles")
	l.pict2 = pixel.Picture("graphics/yellowbutton")
	l.position = pixel.XY{48, 48}
	l.size = pixel.XY{60, 40}
}

func (l *loop3) Enter() {
}

func (loop3) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (l *loop3) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}

	if input.Click.Ongoing() {
		l.size = pixel.XYof(input.Pointer.XY()).Minus(l.position)
	}
}

func (loop3) Update() {
}

func (l *loop3) Render() {
	pixel.Clear(7)
	l.pict2.Tile(pixel.XY{48,4}, l.pict2.Size(), -1)
	l.pict2.Tile(pixel.XY{64,4}, pixel.XY{32, 16}, -1)
	l.pict2.Tile(pixel.XY{102,4}, pixel.XY{16, 32}, -1)
	l.pict2.Tile(pixel.XY{128,4}, pixel.XY{32, 32}, -1)
	l.pict1.Tile(l.position, l.size, 0)
}

// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
