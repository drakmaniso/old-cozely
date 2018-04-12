// Adapted by Laurent Moussault (2018) from C code: see ORIGINAL_LICENSE

package quadedge

import (
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////

// An Edge identifies one of the four directed edges of a specific quad-edge. It
// corresponds to what Guibas and Stolfi call an edge reference. To obtain a new
// Edge, create a new quad-edge with Pool.New().
//
// Note that for convenience, Edge objects contain a pointer to the pool used to
// create it.
type Edge struct {
	pool *Pool
	id   edgeID
}

// edgeID is a pure "edge reference", without the pool pointer.
type edgeID uint32

////////////////////////////////////////////////////////////////////////////////

const (
	// Nil is the value used to initialize the origin, destination, left and
	// right fields of new Edge objects.
	Nil = 0xFFFFFFFF

	// canonical is a mask to isolate the quad ID in an edge reference.
	canonical edgeID = 0xFFFFFFFC

	// quad is a mask to isolate the rotation part of an edge reference.
	quad edgeID = 0x00000003
)

////////////////////////////////////////////////////////////////////////////////

func (e Edge) String() string {
	if e.id == Nil {
		return "[no edge]"
	}
	on := e.pool.next[e.id]
	o := "orig:" +
		datastring(e.pool, e.id) +
		"->" + on.String()
	rn := e.pool.next[e.id.rot()]
	r := " right:" +
		datastring(e.pool, e.id.rot()) +
		"->" + rn.String()
	dn := e.pool.next[e.id.sym()]
	d := " dest:" +
		datastring(e.pool, e.id.sym()) +
		"->" + dn.String()
	ln := e.pool.next[e.id.tor()]
	l := " left:" +
		datastring(e.pool, e.id.tor()) +
		"->" + ln.String()
	return e.id.String() + "=[" + o + r + d + l + "]"
}

func (e edgeID) String() string {
	if e == Nil {
		return "no_edge"
	}
	var s string
	switch e & quad {
	case 0:
		s = "o"
	case 1:
		s = "r"
	case 2:
		s = "d"
	case 3:
		s = "l"
	default:
		s = "error"
	}
	return strconv.Itoa(int(e>>2)) + "" + s
}

func datastring(p *Pool, e edgeID) string {
	d := p.data[e]
	if d == Nil {
		return ""
	}
	return strconv.Itoa(int(d))
}

////////////////////////////////////////////////////////////////////////////////

// Pool returns the allocator that was used to create e.
func (e Edge) Pool() *Pool {
	return e.pool
}

////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////

// OrigLoop calls visit on every edges in the origin edge-ring of e, in
// conter-clockwise order (starting with e).
func (e Edge) OrigLoop(visit func(e Edge)) {
	f := e.id
	for e.id != Nil {
		visit(e)
		e.id = e.pool.next[e.id]
		if e.id == f {
			return
		}
	}
}

// RightLoop calls visit on every edges in the right edge-ring of e, in
// conter-clockwise order (starting with e).
func (e Edge) RightLoop(visit func(e Edge)) {
	f := e.id
	for e.id != Nil {
		visit(e)
		e.id = e.pool.next[e.id.rot()].tor()
		if e.id == f {
			return
		}
	}
}

// DestLoop calls visit on every edges in the destination edge-ring of e, in
// conter-clockwise order (starting with e).
func (e Edge) DestLoop(visit func(e Edge)) {
	f := e.id
	for e.id != Nil {
		visit(e)
		e.id = e.pool.next[e.id.sym()].sym()
		if e.id == f {
			return
		}
	}
}

// LeftLoop calls visit on every edges in the left edge-ring of e, in
// conter-clockwise order (starting with e).
func (e Edge) LeftLoop(visit func(e Edge)) {
	f := e.id
	for e.id != Nil {
		visit(e)
		e.id = e.pool.next[e.id.tor()].rot()
		if e.id == f {
			return
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

// SameRing returns true if o is in the origin edge-ring of e (i.e., if the two
// directed edges share the same origin).
func (e Edge) SameRing(o Edge) bool {
	f := e.id
	for e.id != Nil {
		if e.id == o.id {
			return true
		}
		e.id = e.pool.next[e.id]
		if e.id == f {
			return false
		}
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////

// Orig returns the vertex ID of the origin of e.
func (e Edge) Orig() uint32 {
	return e.pool.data[e.id]
}

// SetOrig changes the vertex ID of the origin of e. Note that if there is other
// edges with the same origin, they will not be updated.
func (e Edge) SetOrig(data uint32) {
	e.pool.data[e.id] = data
}

// Right returns the face ID at the right of e.
func (e Edge) Right() uint32 {
	return e.pool.data[e.id.rot()]
}

// SetRight changes the face ID at the right of e. Note that if there is other
// edges with the same right face, they will not be updated.
func (e Edge) SetRight(data uint32) {
	e.pool.data[e.id.rot()] = data
}

// Dest returns the vertex ID of the destination of e.
func (e Edge) Dest() uint32 {
	return e.pool.data[e.id.sym()]
}

// SetDest changes the vertex ID of the destination of e. If there is other
// edges in the same destination edge ring, they will not be updated.
func (e Edge) SetDest(data uint32) {
	e.pool.data[e.id.sym()] = data
}

// Left returns the face ID at the left of e.
func (e Edge) Left() uint32 {
	return e.pool.data[e.id.tor()]
}

// SetLeft changes the face ID at the left of e. Note that if there is other
// edges with the same left face, they will not be updated.
func (e Edge) SetLeft(data uint32) {
	e.pool.data[e.id.tor()] = data
}

////////////////////////////////////////////////////////////////////////////////

func (e Edge) mark() uint32 {
	return e.pool.marks[e.id>>2]
}

func (e Edge) setMark(mark uint32) {
	e.pool.marks[e.id>>2] = mark
}

////////////////////////////////////////////////////////////////////////////////

// Walk calls visit for every undirected primal edge reachable from e.
//
// It does so by a chain of Sym() and OrigNext() calls, but ensures that in each
// quad-edge, only one edge is visited (i.e. the symetric of an already
// encountered edge is never visited).
func (e Edge) Walk(visit func(e Edge)) {
	if e.id == Nil {
		return
	}
	m := e.pool.nextMark
	e.pool.nextMark++
	if e.pool.nextMark == 0 {
		e.pool.nextMark = 1
	}
	//TODO: non-recursive version?
	e.pool.walk(e.id, visit, m)
}

func (p *Pool) walk(eid edgeID, visit func(e Edge), m uint32) {
	for p.marks[eid>>2] != m {
		visit(Edge{pool: p, id: eid})
		p.marks[eid>>2] = m
		p.walk(p.next[eid.sym()], visit, m)
		eid = p.next[eid]
	}
}

////////////////////////////////////////////////////////////////////////////////
