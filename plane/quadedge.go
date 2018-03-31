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

type QuadEdges struct {
	next []Edge
	data []VertexID
	free []Edge
}

type Edge uint32

//------------------------------------------------------------------------------

const (
	canonical Edge = 0xFFFFFFFC
	quad      Edge = 0x00000003
)

//------------------------------------------------------------------------------

func NewQuadEdges(capacity int) *QuadEdges {
	return &QuadEdges{
		next: make([]Edge, 0, capacity),
		data: make([]VertexID, 0, capacity),
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

func (q *QuadEdges) OrigNext(e Edge) Edge {
	return q.next[e]
}

func (q *QuadEdges) rotRightNext(e Edge) Edge {
	return q.next[e.Rot()]
}

func (q *QuadEdges) symDestNext(e Edge) Edge {
	return q.next[e.Sym()]
}

func (q *QuadEdges) torLeftNext(e Edge) Edge {
	return q.next[e.Tor()]
}

func (q *QuadEdges) RightNext(e Edge) Edge {
	return q.next[e.Rot()].Tor()
}

func (q *QuadEdges) DestNext(e Edge) Edge {
	return q.next[e.Sym()].Sym()
}

func (q *QuadEdges) LeftNext(e Edge) Edge {
	return q.next[e.Tor()].Rot()
}

func (q *QuadEdges) OrigPrev(e Edge) Edge {
	return q.next[e.Rot()].Rot()
}

func (q *QuadEdges) DestPrev(e Edge) Edge {
	return q.next[e.Tor()].Tor()
}

func (q *QuadEdges) RightPrev(e Edge) Edge {
	return q.next[e.Sym()]
}

func (q *QuadEdges) LeftPrev(e Edge) Edge {
	return q.next[e].Sym()
}

//------------------------------------------------------------------------------

func (q *QuadEdges) OrigData(e Edge) VertexID {
	return q.data[e]
}

func (q *QuadEdges) RightData(e Edge) VertexID {
	return q.data[e.Rot()]
}

func (q *QuadEdges) DestData(e Edge) VertexID {
	return q.data[e.Sym()]
}

func (q *QuadEdges) LeftData(e Edge) VertexID {
	return q.data[e.Tor()]
}

//------------------------------------------------------------------------------
