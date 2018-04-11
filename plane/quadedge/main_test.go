// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package quadedge_test

import (
	"os"
	"testing"

	"github.com/drakmaniso/glam/input"
)

//------------------------------------------------------------------------------

var (
	quit     = input.Bool("Quit")
	next     = input.Bool("Next")
	previous = input.Bool("Previous")
	scene1   = input.Bool("Scene1")
	scene2   = input.Bool("Scene2")
	scene3   = input.Bool("Scene3")
	scene4   = input.Bool("Scene4")
	scene5   = input.Bool("Scene5")
	scene6   = input.Bool("Scene6")
	scene7   = input.Bool("Scene7")
	scene8   = input.Bool("Scene8")
	scene9   = input.Bool("Scene9")
	scene10  = input.Bool("Scene10")
)

var testContext = input.Context("Test", quit, next, previous,
	scene1, scene2, scene3, scene4, scene5, scene6, scene7, scene8, scene9)

var testBindings = map[string]map[string][]string{
	"Test": {
		"Quit":     {"Escape"},
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
		"Scene10":  {"10"},
	},
}

//------------------------------------------------------------------------------

var tests = make(chan func())

//------------------------------------------------------------------------------

func do(f func()) {
	done := make(chan bool, 1)
	tests <- func() {
		f()
		done <- true
	}
	<-done
}

//------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------
