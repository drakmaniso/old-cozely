// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

import (
	"errors"

	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

type Picture struct {
	mode    Mode
	mapping uint16
}

var pictures map[string]*Picture

func init() {
	pictures = make(map[string]*Picture, 128)
}

//------------------------------------------------------------------------------

type mapping struct {
	binFlip int16
	x, y    int16
	w, h    int16
}

var mappings []mapping

//------------------------------------------------------------------------------

// Mode describes the way a picture is stored in memory.
type Mode uint8

// The three modes supported for pictures.
const (
	Indexed   Mode = 1
	FullColor Mode = 2
	GrayScale Mode = 3
)

//------------------------------------------------------------------------------

func newPicture(name string, mode Mode, w, h int16) *Picture {
	var p Picture
	p.mode = mode
	mappings = append(mappings, mapping{w: w, h: h})
	p.mapping = uint16(len(mappings) - 1)
	pictures[name] = &p
	return &p
}

//------------------------------------------------------------------------------

// Size returns the width and height of the picture.
func (p *Picture) Size() Coord {
	m := p.mapping
	return Coord{mappings[m].w, mappings[m].h}
}

//------------------------------------------------------------------------------

func (p *Picture) mapTo(binFlip int16, x, y int16) {
	m := p.mapping
	mappings[m].binFlip = binFlip
	mappings[m].x, mappings[m].y = x, y
}

func (p *Picture) getMap() (binFlip int16, x, y, w, h int16) {
	m := p.mapping
	return mappings[m].binFlip, mappings[m].x, mappings[m].y, mappings[m].w, mappings[m].h
}

//------------------------------------------------------------------------------

// GetPicture returns the picture associated with a name. If there isn't any, an
// empty picture is returned, and a sticky error is set.
func GetPicture(name string) *Picture {
	p, ok := pictures[name]
	if !ok {
		setErr("in GetPicture", errors.New("picture \""+name+"\" not found"))
	}
	return p
}

//------------------------------------------------------------------------------

func (p *Picture) Paint(x, y int16) {
	commands = append(commands, gl.DrawIndirectCommand{
		VertexCount:   4,
		InstanceCount: 1,
		FirstVertex:   uint32(4 * p.mode),
		BaseInstance:  uint32(len(parameters)),
	})
	parameters = append(parameters, int16(p.mapping), x, y)
}

//------------------------------------------------------------------------------
