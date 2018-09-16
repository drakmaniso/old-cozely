// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"os"
	"testing"

	"github.com/cozely/cozely/input"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quit     = input.Button("Quit")
	next     = input.Button("Next")
	previous = input.Button("Previous")
	scenes   = []input.ButtonID{
		input.Button("Scene0"),
		input.Button("Scene1"),
		input.Button("Scene2"),
		input.Button("Scene3"),
		input.Button("Scene4"),
		input.Button("Scene5"),
		input.Button("Scene6"),
		input.Button("Scene7"),
		input.Button("Scene8"),
		input.Button("Scene9"),
	}
	scrollup   = input.Button("ScrollUp")
	scrolldown = input.Button("ScrollDown")
	cursor     = input.Cursor("Cursor")
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
