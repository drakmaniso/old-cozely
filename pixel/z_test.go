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
	quit     = input.Digital("Quit")
	next     = input.Digital("Next")
	previous = input.Digital("Previous")
	scenes   = []input.DigitalID{
		input.Digital("Scene1"),
		input.Digital("Scene2"),
		input.Digital("Scene3"),
		input.Digital("Scene4"),
		input.Digital("Scene5"),
		input.Digital("Scene6"),
		input.Digital("Scene7"),
		input.Digital("Scene8"),
		input.Digital("Scene9"),
		input.Digital("Scene10"),
	}
	scrollup   = input.Digital("ScrollUp")
	scrolldown = input.Digital("ScrollDown")
	cursor     = input.Cursor("Cursor")
)

var bindings = input.Bindings{
	"Default": {
		"Quit":       {"Escape"},
		"Next":       {"Mouse Left", "Space"},
		"Previous":   {"Mouse Right", "U"},
		"Scene1":     {"1"},
		"Scene2":     {"2"},
		"Scene3":     {"3"},
		"Scene4":     {"4"},
		"Scene5":     {"5"},
		"Scene6":     {"6"},
		"Scene7":     {"7"},
		"Scene8":     {"8"},
		"Scene9":     {"9"},
		"Scene10":    {"0"},
		"ScrollUp":   {"Mouse Scroll Up"},
		"ScrollDown": {"Mouse Scroll Down"},
		"Cursor":     {"Mouse"},
	},
}

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

////////////////////////////////////////////////////////////////////////////////
