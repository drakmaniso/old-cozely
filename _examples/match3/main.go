// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/carol/key"
	"math/rand"
	"time"

	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/colour"
	"github.com/drakmaniso/carol/mouse"
	"github.com/drakmaniso/carol/pixel"

	"github.com/drakmaniso/carol/_examples/match3/ecs"
	"github.com/drakmaniso/carol/_examples/match3/grid"
)

//------------------------------------------------------------------------------

var tilesPict [8]struct {
	normal, big *pixel.Picture
}

var current grid.Position

//------------------------------------------------------------------------------

func main() {
	err := carol.Run(loop{})
	if err != nil {
		carol.ShowError(err)
	}
}

//------------------------------------------------------------------------------

type loop struct {
	carol.Handlers
}

//------------------------------------------------------------------------------

func (loop) Setup() error {
	pixel.SetBackground(colour.SRGB8{0x5C, 0x82, 0x86})

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
		tilesPict[i].normal = pixel.GetPicture(n)
		tilesPict[i].big = pixel.GetPicture(n + "_big")
	}

	current = grid.Nowhere()

	grid.Setup(8, 8)
	grid.Fill(newTile)

	return nil
}

//------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (loop) MouseButtonDown(_ mouse.Button, _ int) {
	m := pixel.Mouse()
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

func testMatch(e1, e2 ecs.Entity) bool {
	if !e1.Has(ecs.Color) || !e2.Has(ecs.Color) {
		return false
	}
	c1 := colors[e1]
	c2 := colors[e2]
	return c1 == c2
}

func (loop) MouseButtonUp(_ mouse.Button, _ int) {
	current = grid.Nowhere()
}

//------------------------------------------------------------------------------

func (loop) KeyDown(l key.Label, _ key.Position) {
	switch l {
	case key.LabelSpace:
		f := func(e ecs.Entity) {
			if !e.Has(ecs.MatchFlag) {
				print(grid.PositionOf(e).String(), " ")
				e.Add(ecs.MatchFlag)
			}
		}
		grid.TestAndMark(testMatch, f)
		println()
	}
}

//------------------------------------------------------------------------------

func (loop) ScreenResized(width, height int16, _ int32) {
	grid.ScreenResized(width, height)
}

//------------------------------------------------------------------------------
