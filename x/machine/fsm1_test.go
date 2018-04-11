// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package machine_test

import (
	"fmt"

	"github.com/drakmaniso/glam/x/machine"
)

var counter int

func state1() machine.State {
	fmt.Println("State 1")
	return state2
}

func state2() machine.State {
	fmt.Println("State 2")
	return state3
}

func state3() machine.State {
	fmt.Println("State 3")
	if counter > 6 {
		return nil
	}
	return state2
}

func Example_noAllocations() {
	m := machine.State(state1)

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
