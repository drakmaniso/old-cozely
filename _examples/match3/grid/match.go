// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package grid

//------------------------------------------------------------------------------

func (p Position) TestAndMark(test func(e1, e2 uint32) bool, mark func(e uint32)) {
	e := get(p.x, p.y)

	// Count rightward matches
	r := 0
	for i := int8(1); p.x+i < width && test(e, get(p.x+i, p.y)); i++ {
		r++
	}

	// Count leftward matches
	l := 0
	for i := int8(1); p.x-i >= 0 && test(e, get(p.x-i, p.y)); i++ {
		l++
	}

	// Count upward matches
	u := 0
	for i := int8(1); p.y+i < height && test(e, get(p.x, p.y+i)); i++ {
		u++
	}

	// Count downward matches
	d := 0
	for i := int8(1); p.y-i >= 0 && test(e, get(p.x, p.y-i)); i++ {
		d++
	}

	if r+l >= 2 || u+d >= 2 {
		mark(e)
	}
}

//------------------------------------------------------------------------------
