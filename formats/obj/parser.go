// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Package obj implements a simple parser for Wavefront ".obj" file format.
package obj

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

//------------------------------------------------------------------------------

// Parse reads r as a ".obj" file and calls b for each statement recognized.
func Parse(r io.Reader, b Builder) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		w := strings.Fields(s.Text())
		if len(w) < 1 {
			continue
		}
		switch w[0] {

		case "#":
			b.Comment(s.Text())

		case "v":
			v := parseFloats(w)
			b.V(v...)

		case "vt":
			v := parseFloats(w)
			b.VT(v...)

		case "vn":
			v := parseFloats(w)
			b.VN(v...)

		case "f":
			verts := parseIndices(w)
			b.F(verts...)

		case "b":
			verts := parseIndices(w)
			b.P(verts...)

		case "l":
			verts := parseIndices(w)
			b.L(verts...)

		case "mtllib":
			b.MtlLib(w[1:]...)

		case "usemtl":
			if len(w) >= 2 {
				b.UseMtl(w[1])
			}

		case "maplib":
			b.MapLib(w[1:]...)

		case "usemap":
			if len(w) >= 2 {
				b.UseMap(w[1])
			}

		case "g":
			b.G(w[1:]...)

		case "s":
			if len(w) >= 2 {
				val, _ := strconv.Atoi(w[1])
				b.S(val)
			}

		case "lod":
			if len(w) >= 2 {
				val, _ := strconv.Atoi(w[1])
				b.LOD(val)
			}

		case "o":
			if len(w) >= 2 {
				b.O(w[1])
			}
		}
	}
}

//------------------------------------------------------------------------------

// A Builder is used when parsing a ".obj" file. The methods are called each
// time a statement is encountered, in the order they appear in the file.
type Builder interface {
	Comment(txt string) error
	O(name string) error
	G(names ...string) error
	V(coords ...float32) error
	VT(coords ...float32) error
	VN(coords ...float32) error
	P(verts ...Indices) error
	L(verts ...Indices) error
	F(verts ...Indices) error
	S(group int) error
	LOD(level int) error
	MtlLib(names ...string) error
	UseMtl(name string) error
	MapLib(names ...string) error
	UseMap(name string) error
	NotSupported(txt string) error
}

// Indices regroups the indices of the vertex position, texture coordinates and
// normal. They follow the convention of ".obj" files: counting start at 1, and
// negative numbers reference backward from the current position (i.e. -1 is the
// last defined vertex). When TexCoord or Normal weren't specified in the file,
// thay are set to 0 in the struct.
type Indices struct {
	Vertex   int
	TexCoord int
	Normal   int
}

//------------------------------------------------------------------------------

func parseFloats(words []string) []float32 {
	var v []float32
	v = make([]float32, len(words)-1, len(words)-1)
	for i := range v {
		val, _ := strconv.ParseFloat(words[1+i], 32)
		v[i] = float32(val)
	}
	return v
}

//------------------------------------------------------------------------------

func parseIndices(words []string) []Indices {
	var verts []Indices
	verts = make([]Indices, len(words)-1, len(words)-1)
	for i := range verts {
		ss := strings.Split(words[1+i], "/")
		if len(ss) >= 1 {
			val, _ := strconv.Atoi(ss[0])
			verts[i].Vertex = val
		}
		if len(ss) >= 2 {
			val, _ := strconv.Atoi(ss[1])
			verts[i].TexCoord = val
		}
		if len(ss) >= 3 {
			val, _ := strconv.Atoi(ss[2])
			verts[i].Normal = val
		}
	}
	return verts
}

//------------------------------------------------------------------------------

type DefaultBuilder struct{}

func (DefaultBuilder) Comment(txt string) error {
	return nil
}

func (DefaultBuilder) O(name string) error {
	return nil
}

func (DefaultBuilder) G(names ...string) error {
	return nil
}

func (DefaultBuilder) S(group int) error {
	return nil
}

func (DefaultBuilder) LOD(level int) error {
	return nil
}

func (DefaultBuilder) V(coords ...float32) error {
	return nil
}

func (DefaultBuilder) VT(coords ...float32) error {
	return nil
}

func (DefaultBuilder) VN(coords ...float32) error {
	return nil
}

func (DefaultBuilder) F(verts ...Indices) error {
	return nil
}

func (DefaultBuilder) P(verts ...Indices) error {
	return nil
}

func (DefaultBuilder) L(verts ...Indices) error {
	return nil
}

func (DefaultBuilder) MtlLib(name ...string) error {
	return nil
}

func (DefaultBuilder) UseMtl(name string) error {
	return nil
}

func (DefaultBuilder) MapLib(name ...string) error {
	return nil
}

func (DefaultBuilder) UseMap(name string) error {
	return nil
}

func (DefaultBuilder) NotSupported(txt string) error {
	return nil
}

//------------------------------------------------------------------------------
