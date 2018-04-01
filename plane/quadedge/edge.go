// Adapted by Laurent Moussault (2018) from C code:
// http://www.ic.unicamp.br/~stolfi/EXPORT/software/c/2000-05-04/libquad/
//
// See
//
//   "Primitives for the Manipulation of General Subdivisions
//   and the Computation of Voronoi Diagrams"
//
//   p. Guibas, J. Stolfi, ACM TOG, April 1985
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

package quadedge

//------------------------------------------------------------------------------

type edgeID uint32

type Edge struct {
	pool *Pool
	id   edgeID
}

//------------------------------------------------------------------------------

const (
	canonical edgeID = 0xFFFFFFFC
	quad      edgeID = 0x00000003
	noEdge    edgeID = 0xFFFFFFFF
)
//------------------------------------------------------------------------------

func (e Edge) Pool() *Pool {
	return e.pool
}

func (e Edge) ID() uint32 {
	return uint32(e.id)
}

//------------------------------------------------------------------------------

func (e edgeID) rot() edgeID {
	return (e & canonical) + ((e + 1) & quad)
}

func (e edgeID) sym() edgeID {
	return (e & canonical) + ((e + 2) & quad)
}

func (e edgeID) tor() edgeID {
	return (e & canonical) + ((e + 3) & quad)
}

func (e Edge) Canonical() Edge {
	e.id = e.id & canonical
	return e
}

func (e Edge) Rot() Edge {
	e.id = (e.id & canonical) + ((e.id + 1) & quad)
	return e
}

func (e Edge) Sym() Edge {
	e.id = (e.id & canonical) + ((e.id + 2) & quad)
	return e
}

func (e Edge) Tor() Edge {
	e.id = (e.id & canonical) + ((e.id + 3) & quad)
	return e
}

//------------------------------------------------------------------------------

func (e Edge) OrigNext() Edge {
	e.id = e.pool.next[e.id]
	return e
}

func (e Edge) RightNext() Edge {
	e.id = e.pool.next[e.id.rot()].tor()
	return e
}

func (e Edge) DestNext() Edge {
	e.id = e.pool.next[e.id.sym()].sym()
	return e
}

func (e Edge) LeftNext() Edge {
	e.id = e.pool.next[e.id.tor()].rot()
	return e
}

func (e Edge) OrigPrev() Edge {
	e.id = e.pool.next[e.id.rot()].rot()
	return e
}

func (e Edge) RightPrev() Edge {
	e.id = e.pool.next[e.id.sym()]
	return e
}

func (e Edge) DestPrev() Edge {
	e.id = e.pool.next[e.id.tor()].tor()
	return e
}

func (e Edge) LeftPrev() Edge {
	e.id = e.pool.next[e.id].sym()
	return e
}

//------------------------------------------------------------------------------

func (e Edge) Orig() uint32 {
	return e.pool.data[e.id]
}

func (e Edge) SetOrig(data uint32) {
	e.pool.data[e.id] = data
}

func (e Edge) Right() uint32 {
	return e.pool.data[e.id.rot()]
}

func (e Edge) SetRight(data uint32) {
	e.pool.data[e.id.rot()] = data
}

func (e Edge) Dest() uint32 {
	return e.pool.data[e.id.sym()]
}

func (e Edge) SetDest(data uint32) {
	e.pool.data[e.id.sym()] = data
}

func (e Edge) Left() uint32 {
	return e.pool.data[e.id.tor()]
}

func (e Edge) SetLeft(data uint32) {
	e.pool.data[e.id.tor()] = data
}

//------------------------------------------------------------------------------

func (e Edge) mark() uint32 {
	return e.pool.marks[e.id>>2]
}

func (e Edge) setMark(mark uint32) {
	e.pool.marks[e.id>>2] = mark
}

//------------------------------------------------------------------------------

func (e Edge) Walk(walkFn func(e Edge)) {
	m := e.pool.nextMark
	e.pool.nextMark++
	if e.pool.nextMark == 0 {
		e.pool.nextMark = 1
	}
	//TODO: non-recursive version?
	e.pool.walk(e.id, walkFn, m)
}

func (p *Pool) walk(eid edgeID, walkFn func(e Edge), m uint32) {
	for ; p.marks[eid>>2] != m; eid = p.next[eid] {
		walkFn(Edge{pool: p, id: eid})
		p.marks[eid>>2] = m
		p.walk(p.next[eid.sym()], walkFn, m)
	}
}

//------------------------------------------------------------------------------
