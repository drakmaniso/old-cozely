// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var canvas5 = pixel.Canvas(pixel.Resolution(128, 128))

var points = []coord.CR{
	{4, 4},
	{4 + 1, 4 + 20},
	{4 + 1 + 20, 4 + 20 - 1},
	{16, 32},
}

var pointshidden, lineshidden bool

type loop5 struct{}

////////////////////////////////////////////////////////////////////////////////

func TestTest5(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		err := cozely.Run(loop5{})
		if err != nil {
			t.Error(err)
		}
	})
}

func (loop5) Enter() {
	palette2.Activate()
}

func (loop5) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop5) React() {
	if quit.Started(0) {
		cozely.Stop(nil)
	}

	if next.Started(0) {
		m := canvas5.FromWindow(cursor.XY(0).CR())
		points = append(points, m)
	}

	if previous.Started(0) {
		if len(points) > 0 {
			points = points[:len(points)-1]
		}
	}

	pointshidden = scene1.Ongoing(0)
	lineshidden = scene2.Ongoing(0)
}

func (loop5) Update() {
}

func (loop5) Render() {
	canvas5.Clear(1)
	m := canvas5.FromWindow(cursor.XY(0).CR())
	canvas5.Triangles(2, -2, points...)
	if !lineshidden {
		canvas5.Lines(5, 0, points...)
		canvas5.Lines(13, -1, points[len(points)-1], m)
	}
	if !pointshidden {
		for _, p := range points {
			canvas5.Point(8, 1, p)
		}
		canvas5.Point(18, 2, m)
	}
	canvas5.Display()
}
