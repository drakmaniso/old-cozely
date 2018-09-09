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
	quit     = input.Button("Quit")
	cursor   = input.Cursor("Cursor")
	next     = input.Button("Next")
	previous = input.Button("Previous")
	scene1   = input.Button("Scene1")
	scene2   = input.Button("Scene2")
	scene3   = input.Button("Scene3")
	scene4   = input.Button("Scene4")
	scene5   = input.Button("Scene5")
	scene6   = input.Button("Scene6")
	scene7   = input.Button("Scene7")
	scene8   = input.Button("Scene8")
	scene9   = input.Button("Scene9")
	scene10  = input.Button("Scene10")
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
