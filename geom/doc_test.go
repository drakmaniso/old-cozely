// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package geom_test

//------------------------------------------------------------------------------

import (
	"fmt"

	. "github.com/drakmaniso/glam/geom"
)

//------------------------------------------------------------------------------

func Example() {
	a := Vec2{1, 0}
	b := Vec2{0, 1}
	c := a.Minus(b).Normalized()
	fmt.Println(c)
	t := Mat3{
		{1, 0, 0},
		{0, 1, 0},
		{2, 3, 1},
	}
	r := Mat3{
		{0, 1, 0},
		{-1, 0, 0},
		{0, 0, 1},
	}
	m := r.Times(t)
	fmt.Println(m)
	// Output:
	// {0.70710677 -0.70710677}
	// [[0 1 0] [-1 0 0] [-3 2 1]]
}

//------------------------------------------------------------------------------
