// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package grid

import (
	"fmt"

	"github.com/drakmaniso/glam/_examples/match3/ecs"
)

//------------------------------------------------------------------------------

// A Position on the grid
type Position struct {
	x, y int8
}

var grid []ecs.Entity
var positions [ecs.Size]Position

var width, height int8

//------------------------------------------------------------------------------

// Setup prepare a new grid of width w and height h.
func Setup(w, h int8) {
	width, height = w, h
	grid = make([]ecs.Entity, w*h, w*h)
}

//------------------------------------------------------------------------------

// Fill the grid with new tiles.
func Fill(newTile func() ecs.Entity) {
	for y := int8(0); y < height; y++ {
		for x := int8(0); x < width; x++ {
			e := newTile()
			put(e, x, y)
		}
	}
}

//------------------------------------------------------------------------------

// put entity e at grid coordinates x, y.
func put(e ecs.Entity, x, y int8) {
	grid[x+y*width] = e
	positions[e] = Position{x: x, y: y}
	e.Add(ecs.GridPosition)
}

// get returns the entity at grid coordinates x, y.
func get(x, y int8) ecs.Entity {
	return grid[x+y*width]
}

//------------------------------------------------------------------------------

// PositionOf returns the grid position of entity e.
func PositionOf(e ecs.Entity) Position {
	if !e.Has(ecs.GridPosition) {
		return Nowhere()
	}
	return positions[e]
}

//------------------------------------------------------------------------------

// At returns the entity currently at the grid position.
func At(p Position) ecs.Entity {
	return grid[p.x+p.y*width]
}

//------------------------------------------------------------------------------

// String returns a string representation of p.
func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

//------------------------------------------------------------------------------

// Nowhere returns the nil value for grid positions.
func Nowhere() Position {
	return Position{x: -1, y: -1}
}

//------------------------------------------------------------------------------
