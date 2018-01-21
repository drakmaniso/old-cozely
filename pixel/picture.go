// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
)

//------------------------------------------------------------------------------

type Picture uint16

var pictures map[string]Picture

func init() {
	pictures = make(map[string]Picture, 128)
}

//------------------------------------------------------------------------------

type mapping struct {
	binFlip int16
	x, y    int16
	w, h    int16
}

var mappings = []mapping {
	mapping{},
}

//------------------------------------------------------------------------------

func newPicture(name string, w, h int16) Picture {
	mappings = append(mappings, mapping{w: w, h: h})
	p := Picture(len(mappings) - 1)
	pictures[name] = p
	return p
}

//------------------------------------------------------------------------------

// Size returns the width and height of the picture.
func (p Picture) Size() Coord {
	return Coord{mappings[p].w, mappings[p].h}
}

//------------------------------------------------------------------------------

func (p Picture) mapTo(binFlip int16, x, y int16) {
	mappings[p].binFlip = binFlip
	mappings[p].x, mappings[p].y = x, y
}

func (p Picture) getMap() (binFlip int16, x, y, w, h int16) {
	return mappings[p].binFlip, mappings[p].x, mappings[p].y, mappings[p].w, mappings[p].h
}

//------------------------------------------------------------------------------

// GetPicture returns the picture associated with a name. If there isn't any, an
// empty picture is returned, and a sticky error is set.
func GetPicture(name string) Picture {
	p, ok := pictures[name]
	if !ok {
		setErr("in GetPicture", errors.New("picture \""+name+"\" not found"))
	}
	return p
}

//------------------------------------------------------------------------------

func (p Picture) Paint(x, y int16) {
	appendCommand(cmdPicture, 4, 1)
	parameters = append(parameters, int16(p), x, y)
}

//------------------------------------------------------------------------------
