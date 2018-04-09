// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

var Bools struct {
	Name    []string
	Active  []bool
	Just    []bool
	Pressed []bool
}

var Floats struct {
	Name   []string
	Active []bool
	Value  []float32
}

var Coords struct {
	Name   []string
	Active []bool
	X      []float32
	Y      []float32
}

var Deltas struct {
	Name   []string
	Active []bool
	X      []float32
	Y      []float32
}
