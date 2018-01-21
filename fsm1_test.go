// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam_test

import (
	"fmt"

	"github.com/drakmaniso/glam"
)

var counter int

func state1() glam.State {
	fmt.Println("State 1")
	return state2
}

func state2() glam.State {
	fmt.Println("State 2")
	return state3
}

func state3() glam.State {
	fmt.Println("State 3")
	if counter > 6 {
		return nil
	}
	return state2
}

func ExampleState_noAllocations() {
	m := glam.State(state1)

	for m != nil {
		counter++
		m.Update()
	}
	// Output:
	// State 1
	// State 2
	// State 3
	// State 2
	// State 3
	// State 2
	// State 3
}
