// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"os"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////

var tests = make(chan func())

////////////////////////////////////////////////////////////////////////////////

func do(f func()) {
	done := make(chan bool, 1)
	tests <- func() {
		f()
		done <- true
	}
	<-done
}

////////////////////////////////////////////////////////////////////////////////

func TestMain(m *testing.M) {
	result := make(chan int, 1)

	go func() {
		result <- m.Run()
	}()

	go func() {
		os.Exit(<-result)
	}()

	for f := range tests {
		f()
	}
}
