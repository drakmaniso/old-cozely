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
	next     []Edge
	data     []uint32
	mark     []uint32
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

func NewQuadEdges(capacity uint32) *QuadEdges {
	return &QuadEdges{
		next:     make([]Edge, 0, capacity),
		data:     make([]uint32, 0, capacity),
		mark:     make([]uint32, 0, capacity),
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

func (q *QuadEdges) RightPrev(e Edge) Edge {
	return q.next[e.Sym()]
}

func (q *QuadEdges) DestPrev(e Edge) Edge {
	return q.next[e.Tor()].Tor()
}

func (q *QuadEdges) LeftPrev(e Edge) Edge {
	return q.next[e].Sym()
}

//------------------------------------------------------------------------------

func (q *QuadEdges) OrigData(e Edge) uint32 {
	return q.data[e]
}

func (q *QuadEdges) RightData(e Edge) uint32 {
	return q.data[e.Rot()]
}

func (q *QuadEdges) DestData(e Edge) uint32 {
	return q.data[e.Sym()]
}

func (q *QuadEdges) LeftData(e Edge) uint32 {
	return q.data[e.Tor()]
}

//------------------------------------------------------------------------------

func (q *QuadEdges) MakeEdge(orig, dest uint32, leftRight uint32) Edge {

	//TODO: implement the free list

	s := uint32(len(q.next) + 4)
	if s > q.capacity {
		//TODO: grow the slices
		panic("growing QuadEdges not implemented")
	}

	// Allocate the new quad
	e := Edge(len(q.next))
	q.next = q.next[:s]
	q.data = q.data[:s]

	// Initialize the quad
	q.next[e] = e
	q.data[e] = orig
	q.next[e.Sym()] = e.Sym()
	q.data[e.Sym()] = dest
	q.next[e.Rot()] = e.Tor()
	q.data[e.Rot()] = leftRight
	q.next[e.Tor()] = e.Rot()
	q.data[e.Tor()] = leftRight

	return e
}

//------------------------------------------------------------------------------

func (q *QuadEdges) Splice(a, b Edge) {
	alpha := q.OrigNext(a).Rot()
	beta := q.OrigNext(b).Rot()

	q.next[a], q.data[a], q.next[b], q.data[b] =
		q.next[b], q.data[b], q.next[a], q.data[a]

	q.next[alpha], q.data[alpha], q.next[beta], q.data[beta] =
		q.next[beta], q.data[beta], q.next[alpha], q.data[alpha]
}

//------------------------------------------------------------------------------

func (q *QuadEdges) Destroy(e Edge) {
	f := e.Sym()
	if q.next[e] != e {
		q.Splice(e, q.OrigPrev(e))
	}
	if q.next[f] != f {
		q.Splice(f, q.OrigPrev(f))
	}

	q.next[e] = noEdge
	q.data[e] = 0xFFFFFFFF
	q.mark[e] = 0
	q.next[e.Rot()] = noEdge
	q.data[e.Rot()] = 0xFFFFFFFF
	q.mark[e.Rot()] = 0
	q.next[e.Sym()] = noEdge
	q.data[e.Sym()] = 0xFFFFFFFF
	q.mark[e.Sym()] = 0
	q.next[e.Tor()] = noEdge
	q.data[e.Tor()] = 0xFFFFFFFF
	q.mark[e.Tor()] = 0

	q.free = append(q.free, e)
}

//------------------------------------------------------------------------------

func (q *QuadEdges) Walk(e Edge, walkFn func(e Edge)) {
	m := q.nextMark
	q.nextMark++
	if q.nextMark == 0 {
		q.nextMark = 1
	}
	q.walkRec(e, walkFn, m)
}

func (q *QuadEdges) walkRec(e Edge, walkFn func(e Edge), m uint32) {
	for ; q.mark[e] != m; e = q.next[e] {
		q.walkRec(q.next[e.Sym()], walkFn, m)
	}
}

//------------------------------------------------------------------------------
