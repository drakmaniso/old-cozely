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
	quit     = input.NewBool("Quit")
	next     = input.NewBool("Next")
	previous = input.NewBool("Previous")
	scene1   = input.NewBool("Scene1")
	scene2   = input.NewBool("Scene2")
	scene3   = input.NewBool("Scene3")
	scene4   = input.NewBool("Scene4")
	scene5   = input.NewBool("Scene5")
	scene6   = input.NewBool("Scene6")
	scene7   = input.NewBool("Scene7")
	scene8   = input.NewBool("Scene8")
	scene9   = input.NewBool("Scene9")
	scene10  = input.NewBool("Scene10")
)

var testContext = input.NewContext("Test", quit, next, previous,
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
