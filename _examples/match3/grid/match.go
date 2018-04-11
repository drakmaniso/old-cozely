// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package grid

import (
	"github.com/cozely/cozely/_examples/match3/ecs"
)

////////////////////////////////////////////////////////////////////////////////

func (p Position) TestAndMark(test func(e1, e2 ecs.Entity) bool, mark func(e ecs.Entity)) {
	e := get(p.x, p.y)

	// Count rightward matches
	r := int8(0)
	for i := int8(1); p.x+i < width && test(e, get(p.x+i, p.y)); i++ {
		r++
	}

	// Count leftward matches
	l := int8(0)
	for i := int8(1); p.x-i >= 0 && test(e, get(p.x-i, p.y)); i++ {
		l++
	}

	// Count upward matches
	u := int8(0)
	for i := int8(1); p.y+i < height && test(e, get(p.x, p.y+i)); i++ {
		u++
	}

	// Count downward matches
	d := int8(0)
	for i := int8(1); p.y-i >= 0 && test(e, get(p.x, p.y-i)); i++ {
		d++
	}

	if r+l >= 2 || u+d >= 2 {
		mark(e)
		if r+l >= 2 {
			for i := int8(1); i <= r; i++ {
				mark(get(p.x+i, p.y))
			}
			for i := int8(1); i <= l; i++ {
				mark(get(p.x-i, p.y))
			}
		}
		if u+d >= 2 {
			for i := int8(1); i <= u; i++ {
				mark(get(p.x, p.y+i))
			}
			for i := int8(1); i <= d; i++ {
				mark(get(p.x, p.y-i))
			}
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestAndMark(test func(e1, e2 ecs.Entity) bool, mark func(e ecs.Entity)) {
	for y := int8(0); y < height; y++ {
		for x := int8(0); x < width; x++ {
			p := Position{x: x, y: y}
			p.TestAndMark(test, mark)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
