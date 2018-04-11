// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package machine_test

import (
	"fmt"

	"github.com/cozely/cozely/x/machine"
)

func Example() {
	counter := 0

	// Define the State Machine
	var s1, s2, s3 machine.State

	s1 = func() machine.State {
		fmt.Println("State 1")
		return s2
	}

	s2 = func() machine.State {
		fmt.Println("State 2")
		return s3
	}

	s3 = func() machine.State {
		fmt.Println("State 3")
		if counter > 6 {
			return nil
		}
		return s2
	}

	// Run the State Machine
	m := s1

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
