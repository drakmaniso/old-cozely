// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"math/rand"
	"time"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/pixel"

	"github.com/cozely/cozely/_examples/match3/ecs"
	"github.com/cozely/cozely/_examples/match3/grid"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quit  = input.Bool("Quit")
	selct = input.Bool("Select")
	test  = input.Bool("Test")
)

var context = input.Context("Default", quit, selct, test)

var bindings = input.Bindings{
	"Default": {
		"Quit":   {"Escape"},
		"Select": {"Mouse Left"},
		"Test":   {"Space"},
	},
}

////////////////////////////////////////////////////////////////////////////////

var tilesPict [8]struct {
	normal, big pixel.PictureID
}

var current grid.Position

var screen = pixel.Canvas(pixel.TargetResolution(180, 180))

////////////////////////////////////////////////////////////////////////////////

func main() {
	setup()

	cozely.Configure(
		cozely.Title("Match 3"),
	)

	err := cozely.Run(loop{})
	if err != nil {
		cozely.ShowError(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

func setup() error {
	cozely.Events.Resize = resize

	for i, n := range []string{
		"red",
		"yellow",
		"green",
		"blue",
		"violet",
		"pink",
		"dark",
		"multi",
	} {
		tilesPict[i].normal = pixel.Picture("graphics/" + n)
		tilesPict[i].big = pixel.Picture("graphics/" + n + "_big")
	}

	current = grid.Nowhere()

	grid.Setup(8, 8)
	grid.Fill(newTile)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop) Enter() error {
	input.Load(bindings)
	context.Activate(1)

	err := palette.Load("graphics/blue")
	if err != nil {
		return err
	}

	return nil
}

func (loop) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func newTile() ecs.Entity {
	e := ecs.New(ecs.Color)
	c := color(rand.Int31n(7))
	// if rand.Int31n(16) == 0 {
	// 	c = 7
	// }
	colors[e] = c

	return e
}

func init() {
	rand.Seed(int64(time.Now().Unix()))
}

////////////////////////////////////////////////////////////////////////////////

func (loop) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}

	if selct.JustPressed(1) {
		m := screen.Mouse()
		current = grid.PositionAt(m)
		if current != grid.Nowhere() {
			e := grid.At(current)
			n := 0
			f := func(e ecs.Entity) {
				print(grid.PositionOf(e).String(), " ")
				n++
			}
			grid.PositionOf(e).TestAndMark(testMatch, f)
			println("-> ", n)
		}
	}

	if selct.JustReleased(1) {
		current = grid.Nowhere()
	}

	if test.JustPressed(1) {
		f := func(e ecs.Entity) {
			if !e.Has(ecs.MatchFlag) {
				print(grid.PositionOf(e).String(), " ")
				e.Add(ecs.MatchFlag)
			}
		}
		grid.TestAndMark(testMatch, f)
		println()
	}

	return nil
}

func testMatch(e1, e2 ecs.Entity) bool {
	if !e1.Has(ecs.Color) || !e2.Has(ecs.Color) {
		return false
	}
	c1 := colors[e1]
	c2 := colors[e2]
	return c1 == c2
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func resize() {
	w, h := screen.Size().X, screen.Size().Y
	grid.ScreenResized(w, h)
}

////////////////////////////////////////////////////////////////////////////////
