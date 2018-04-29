// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package quadedge_test

import (
	"os"
	"testing"

	"github.com/cozely/cozely/input"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quit     = input.Digital("Quit")
	cursor   = input.Cursor("Cursor")
	next     = input.Digital("Next")
	previous = input.Digital("Previous")
	scene1   = input.Digital("Scene1")
	scene2   = input.Digital("Scene2")
	scene3   = input.Digital("Scene3")
	scene4   = input.Digital("Scene4")
	scene5   = input.Digital("Scene5")
	scene6   = input.Digital("Scene6")
	scene7   = input.Digital("Scene7")
	scene8   = input.Digital("Scene8")
	scene9   = input.Digital("Scene9")
	scene10  = input.Digital("Scene10")
)
var bindings = input.Bindings{
	"Default": {
		"Quit":     {"Escape"},
		"Cursor":   {"Mouse"},
		"Next":     {"Mouse Left", "Space"},
		"Previous": {"Mouse Right", "U"},
		"Scene1":   {"1"},
		"Scene2":   {"2"},
		"Scene3":   {"3"},
		"Scene4":   {"4"},
		"Scene5":   {"5"},
		"Scene6":   {"6"},
		"Scene7":   {"7"},
		"Scene8":   {"8"},
		"Scene9":   {"9"},
		"Scene10":  {"0"},
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
