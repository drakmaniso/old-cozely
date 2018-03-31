// Adapted by Laurent Moussault (2018) from C code:
// http://www.ic.unicamp.br/~stolfi/EXPORT/software/c/2000-05-04/libquad/
//
// See
//
//   "Primitives for the Manipulation of General Subdivisions
//   and the Computation of Voronoi Diagrams"
//
//   L. Guibas, J. Stolfi, ACM TOG, April 1985
//
// Originally written by Jim Roth (DEC CADM Advanced Group) on May 1986.
// Modified by J. Stolfi on April 1993.
//
// Original Copyright notice:
//
// Copyright 1996 Institute of Computing, Unicamp.
//
// Permission to use this software for any purpose is hereby granted,
// provided that any substantial copy or mechanically derived version
// of this file that is made available to other parties is accompanied
// by this copyright notice in full, and is distributed under these same
// terms.
//
// DISCLAIMER: This software is provided "as is" with no explicit or
// implicit warranty of any kind.  Neither the authors nor their
// employers can be held responsible for any losses or damages
// that might be attributed to its use.
//
// End of copyright notice.

package plane

//------------------------------------------------------------------------------

type Subdivision struct {
	next     []Edge
	data     []uint32
	marks    []uint32
	free     []Edge
	capacity uint32
	nextMark uint32
}

type Edge uint32

//------------------------------------------------------------------------------

const (
	canonical Edge = 0xFFFFFFFC
	quad      Edge = 0x00000003
	noEdge    Edge = 0xFFFFFFFF
)

//------------------------------------------------------------------------------

func NewSubdivision(capacity uint32) *Subdivision {
	return &Subdivision{
		next:     make([]Edge, 0, 4*capacity),
		data:     make([]uint32, 0, 4*capacity),
		marks:    make([]uint32, 0, capacity),
		capacity: capacity,
		nextMark: 1,
	}
}

//------------------------------------------------------------------------------

func (e Edge) Canonical() Edge {
	return e & canonical
}

func (e Edge) Rot() Edge {
	return (e & canonical) + ((e + 1) & quad)
}

func (e Edge) Sym() Edge {
	return (e & canonical) + ((e + 2) & quad)
}

func (e Edge) Tor() Edge {
	return (e & canonical) + ((e + 3) & quad)
}

//------------------------------------------------------------------------------

func (s *Subdivision) OrigNext(e Edge) Edge {
	return s.next[e]
}

func (s *Subdivision) RightNext(e Edge) Edge {
	return s.next[e.Rot()].Tor()
}

func (s *Subdivision) DestNext(e Edge) Edge {
	return s.next[e.Sym()].Sym()
}

func (s *Subdivision) LeftNext(e Edge) Edge {
	return s.next[e.Tor()].Rot()
}

func (s *Subdivision) OrigPrev(e Edge) Edge {
	return s.next[e.Rot()].Rot()
}

func (s *Subdivision) RightPrev(e Edge) Edge {
	return s.next[e.Sym()]
}

func (s *Subdivision) DestPrev(e Edge) Edge {
	return s.next[e.Tor()].Tor()
}

func (s *Subdivision) LeftPrev(e Edge) Edge {
	return s.next[e].Sym()
}

//------------------------------------------------------------------------------

func (s *Subdivision) Orig(e Edge) uint32 {
	return s.data[e]
}

func (s *Subdivision) SetOrig(e Edge, data uint32) {
	s.data[e] = data
}

func (s *Subdivision) Right(e Edge) uint32 {
	return s.data[e.Rot()]
}

func (s *Subdivision) SetRight(e Edge, data uint32) {
	s.data[e.Rot()] = data
}

func (s *Subdivision) Dest(e Edge) uint32 {
	return s.data[e.Sym()]
}

func (s *Subdivision) SetDest(e Edge, data uint32) {
	s.data[e.Sym()] = data
}

func (s *Subdivision) Left(e Edge) uint32 {
	return s.data[e.Tor()]
}

func (s *Subdivision) SetLeft(e Edge, data uint32) {
	s.data[e.Tor()] = data
}

//------------------------------------------------------------------------------

func (s *Subdivision) mark(e Edge) uint32 {
	return s.marks[e>>2]
}

func (s *Subdivision) setMark(e Edge, mark uint32) {
	s.marks[e>>2] = mark
}

//------------------------------------------------------------------------------

func (s *Subdivision) MakeEdge() Edge {

	//TODO: implement the free list

	sz := uint32(len(s.next) + 4)
	if sz > s.capacity {
		//TODO: grow the slices
		panic("growing QuadEdges not implemented")
	}

	// Allocate the new quad
	e := Edge(len(s.next))
	s.next = s.next[:sz]
	s.data = s.data[:sz]

	// Initialize the quad
	s.next[e] = e
	s.data[e] = 0xFFFFFFFF
	s.next[e.Sym()] = e.Sym()
	s.data[e.Sym()] = 0xFFFFFFFF
	s.next[e.Rot()] = e.Tor()
	s.data[e.Rot()] = 0xFFFFFFFF
	s.next[e.Tor()] = e.Rot()
	s.data[e.Tor()] = 0xFFFFFFFF

	return e
}

//------------------------------------------------------------------------------

func (s *Subdivision) Splice(a, b Edge) {
	alpha := s.OrigNext(a).Rot()
	beta := s.OrigNext(b).Rot()

	s.next[a], s.next[b] = s.next[b], s.next[a]

	s.next[alpha], s.next[beta] = s.next[beta], s.next[alpha]
}

//------------------------------------------------------------------------------

func (s *Subdivision) Destroy(e Edge) {
	f := e.Sym()
	if s.next[e] != e {
		s.Splice(e, s.OrigPrev(e))
	}
	if s.next[f] != f {
		s.Splice(f, s.OrigPrev(f))
	}

	s.next[e] = noEdge
	s.data[e] = 0xFFFFFFFF
	s.next[e.Rot()] = noEdge
	s.data[e.Rot()] = 0xFFFFFFFF
	s.next[e.Sym()] = noEdge
	s.data[e.Sym()] = 0xFFFFFFFF
	s.next[e.Tor()] = noEdge
	s.data[e.Tor()] = 0xFFFFFFFF

	s.marks[e>>2] = 0

	s.free = append(s.free, e)
}

//------------------------------------------------------------------------------

func (s *Subdivision) Walk(e Edge, walkFn func(s *Subdivision, e Edge)) {
	m := s.nextMark
	s.nextMark++
	if s.nextMark == 0 {
		s.nextMark = 1
	}
	//TODO: non-recursive version?
	s.walkRec(e, walkFn, m)
}

func (s *Subdivision) walkRec(e Edge, walkFn func(s *Subdivision, e Edge), m uint32) {
	for ; s.mark(e) != m; e = s.next[e] {
		walkFn(s, e)
		s.setMark(e, m)
		s.walkRec(s.next[e.Sym()], walkFn, m)
	}
}

//------------------------------------------------------------------------------
