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

// An Edge is an "edge reference", i.e. it identifies a specific edge of a
// quad-edge. To create a new Edge, use Pool.New().
//
// Note that Edge contains a pointer to the pool used to create it.
type Edge struct {
	pool *Pool
	id   edgeID
}

// edgeID is an "edge reference", i.e. it identifies a specific edge of a
// quad-edge.
type edgeID uint32

//------------------------------------------------------------------------------

const (
	// canonical is a mask to isolate the quad ID in an edge reference.
	canonical edgeID = 0xFFFFFFFC

	// quad is a mask to isolate the rotation part of an edge reference.
	quad edgeID = 0x00000003

	// noEdge is a dummy edge reference, for uninitialized quads.
	noEdge edgeID = 0xFFFFFFFF
)

//------------------------------------------------------------------------------

// Pool returns the allocator that was used to create e.
func (e Edge) Pool() *Pool {
	return e.pool
}

//------------------------------------------------------------------------------

// rot returns the rotated version of e (counter-clockwise), i.e. the edge that
// belongs to the same quad-edge, but is the dual of e, directed from right to
// left.
func (e edgeID) rot() edgeID {
	return (e & canonical) + ((e + 1) & quad)
}

// sym returns the symmetric of e, i.e. the edge that belongs to the same
// quad-edge but with opposite direction.
func (e edgeID) sym() edgeID {
	return (e & canonical) + ((e + 2) & quad)
}

// tor returns the inverse rotated version of e (so, clockwise), i.e. the edge
// that belongs to the same quad-edge, but is the dual of e, directed from left
// to right. (It is notated Rot⁻¹ in the article)
func (e edgeID) tor() edgeID {
	return (e & canonical) + ((e + 3) & quad)
}

//------------------------------------------------------------------------------

// Canonical returns the canonical representation of e. All four edges of a
// quad-edge return the same value.
func (e Edge) Canonical() Edge {
	e.id = e.id & canonical
	return e
}

// Rot returns the rotated version of e (counter-clockwise), i.e. the edge that
// belongs to the same quad-edge, but is the dual of e, directed from right to
// left.
func (e Edge) Rot() Edge {
	e.id = (e.id & canonical) + ((e.id + 1) & quad)
	return e
}

// Sym returns the symmetric of e, i.e. the edge that belongs to the same
// quad-edge but with opposite direction.
func (e Edge) Sym() Edge {
	e.id = (e.id & canonical) + ((e.id + 2) & quad)
	return e
}

// Tor returns the inverse rotated version of e (so, clockwise), i.e. the edge
// that belongs to the same quad-edge, but is the dual of e, directed from left
// to right. (It is notated Rot⁻¹ in the article)
func (e Edge) Tor() Edge {
	e.id = (e.id & canonical) + ((e.id + 3) & quad)
	return e
}

//------------------------------------------------------------------------------

// OrigNext returns the next counter-clockwise edge with the same origin vertex.
func (e Edge) OrigNext() Edge {
	e.id = e.pool.next[e.id]
	return e
}

// RightNext returns the next counter-clockwise edge with the same right face.
func (e Edge) RightNext() Edge {
	e.id = e.pool.next[e.id.rot()].tor()
	return e
}

// DestNext returns the next counter-clockwise edge with the same destination
// vertex.
func (e Edge) DestNext() Edge {
	e.id = e.pool.next[e.id.sym()].sym()
	return e
}

// LeftNext returns the next counter-clockwise edge with the same left face.
func (e Edge) LeftNext() Edge {
	e.id = e.pool.next[e.id.tor()].rot()
	return e
}

// OrigPrev returns the previous counter-clockwise edge with the same origin vertex.
func (e Edge) OrigPrev() Edge {
	e.id = e.pool.next[e.id.rot()].rot()
	return e
}

// RightPrev returns the previous counter-clockwise edge with the same right face.
func (e Edge) RightPrev() Edge {
	e.id = e.pool.next[e.id.sym()]
	return e
}

// DestPrev returns the previous counter-clockwise edge with the same destination
// vertex.
func (e Edge) DestPrev() Edge {
	e.id = e.pool.next[e.id.tor()].tor()
	return e
}

// LeftPrev returns the previous counter-clockwise edge with the same left face.
func (e Edge) LeftPrev() Edge {
	e.id = e.pool.next[e.id].sym()
	return e
}

//------------------------------------------------------------------------------

// Orig returns the ID of the origin vertex of e.
func (e Edge) Orig() uint32 {
	return e.pool.data[e.id]
}

// SetOrig changes the ID of the origin vertex of e. If there is other edges in
// the same origin edge ring, they will not be updated.
func (e Edge) SetOrig(data uint32) {
	e.pool.data[e.id] = data
}

// Right returns the ID of the right face of e.
func (e Edge) Right() uint32 {
	return e.pool.data[e.id.rot()]
}

// SetRight changes the ID of the right face of e. If there is other edges in
// the same right edge ring, they will not be updated.
func (e Edge) SetRight(data uint32) {
	e.pool.data[e.id.rot()] = data
}

// Dest returns the ID of the destination vertex of e.
func (e Edge) Dest() uint32 {
	return e.pool.data[e.id.sym()]
}

// SetDest changes the ID of the destination vertex of e. If there is other
// edges in the same destination edge ring, they will not be updated.
func (e Edge) SetDest(data uint32) {
	e.pool.data[e.id.sym()] = data
}

// Left returns the ID of the left face of e.
func (e Edge) Left() uint32 {
	return e.pool.data[e.id.tor()]
}

// SetLeft changes the ID of the left face of e. If there is other edges in
// the same left edge ring, they will not be updated.
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

// Walk calls walkFn for every undirected primal edge reachable from e.
//
// It does so by a chain of Sym() and OrigNext() calls, but ensures that in each
// quad-edge, only one edge is visited (i.e. the symetric of an already
// encountered edge is never visited).
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
