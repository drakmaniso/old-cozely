// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mtx

import (
	"fmt"
	micro "github.com/drakmaniso/glam/internal/microtext"
)

//------------------------------------------------------------------------------

func Size() (x, y int) {
	return micro.Size()
}

func Clamp(x, y int) (int, int) {
	sx, sy := micro.Size()

	if x < 0 {
		x += sx
		if x < 0 {
			x = 0
		}
	}
	if x >= sx {
		x = sx - 1
	}

	if y < 0 {
		y += sy
		if y < 0 {
			y = 0
		}
	}
	if y >= sy {
		y = sy - 1
	}

	return x, y
}

//------------------------------------------------------------------------------

func Clear() {
	for i := range micro.Text {
		micro.Text[i] = '\x00'
	}
	micro.TextUpdated = true
}

//------------------------------------------------------------------------------

func Peek(x, y int) byte {
	sx, _ := micro.Size()
	x, y = Clamp(x, y)
	return micro.Text[x+y*sx]
}

func Poke(x, y int, value byte) {
	sx, _ := micro.Size()
	x, y = Clamp(x, y)
	ov := micro.Text[x+y*sx]
	if value != ov {
		micro.Text[x+y*sx] = value
		micro.TextUpdated = true
	}
}

//------------------------------------------------------------------------------

func Printf(x, y int, format string, a ...interface{}) (int, int) {
	var w micro.Writer
	w.Left, w.Top = Clamp(x, y)
	w.Right, w.Bottom = micro.Size()
	w.X, w.Y = w.Left, w.Top

	fmt.Fprintf(&w, format, a...)
	return x, y //TODO
}

//------------------------------------------------------------------------------
