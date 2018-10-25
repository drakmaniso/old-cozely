// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"math/rand"
	"time"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/resource"
	"github.com/cozely/cozely/window"

	"github.com/cozely/cozely/examples/match3/ecs"
	"github.com/cozely/cozely/examples/match3/grid"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quit   = input.Button("Quit")
	click  = input.Button("Clicko")
	test   = input.Button("Test")
	cursor = input.Cursor("Cursor")
)

////////////////////////////////////////////////////////////////////////////////

var tilesPict [8]struct {
	normal, big pixel.PictureID
}

var current grid.Position

type loop struct{}

////////////////////////////////////////////////////////////////////////////////

func main() {
	var err error
	defer cozely.Recover()

	cozely.Configure(
		cozely.Title("Match 3"),
	)
	window.Events.Resize = resize
	pixel.SetResolution(pixel.XY{180, 180})

	err = resource.Path(cozely.Path())
	if err != nil {
		panic(err)
	}

	err = cozely.Run(loop{})
	if err != nil {
		panic(err)
	}
}

func (loop) Enter() {
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
	resize()
}

func (loop) Leave() {
}

func resize() {
	grid.ScreenResized(pixel.Resolution().X, pixel.Resolution().Y)
}

func newTile() ecs.Entity {
	e := ecs.New(ecs.Color)
	c := colour(rand.Int31n(7))
	// if rand.Int31n(16) == 0 {
	// 	c = 7
	// }
	colours[e] = c

	return e
}

func init() {
	rand.Seed(int64(time.Now().Unix()))
}

////////////////////////////////////////////////////////////////////////////////

func (loop) React() {
	if click.Pressed() {
		m := pixel.XYof(cursor.XY())
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

	if click.Released() {
		current = grid.Nowhere()
	}

	if test.Pressed() {
		f := func(e ecs.Entity) {
			if !e.Has(ecs.MatchFlag) {
				print(grid.PositionOf(e).String(), " ")
				e.Add(ecs.MatchFlag)
			}
		}
		grid.TestAndMark(testMatch, f)
		println()
	}

	if quit.Pressed() {
		cozely.Stop(nil)
	}
}

func testMatch(e1, e2 ecs.Entity) bool {
	if !e1.Has(ecs.Color) || !e2.Has(ecs.Color) {
		return false
	}
	c1 := colours[e1]
	c2 := colours[e2]
	return c1 == c2
}

func (loop) Update() {
}
