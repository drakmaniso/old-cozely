// Adapted by Laurent Moussault (2018) from C code: see ORIGINAL_LICENSE

package quadedge

//------------------------------------------------------------------------------

// Splice implements the second operator from Guibas and Stolfi Quad-Edge
// article.
//
// It affects the two origin edge rings of a and b, as well as their two left
// edge rings: if the two rings are distinct, they are combined into one; if the
// two rings are the same, they are broken in separate pieces.
//
// In the origin edge rings, the change happens immediately after a.Orig() and
// b.Orig() (in counterclockwise order); and in the left edge rings, the change
// happens immediately before a.Left() and b.Left().
func Splice(a, b Edge) {
	p := a.pool
	alpha := p.next[a.id].rot()
	beta := p.next[b.id].rot()

	p.next[a.id], p.next[b.id] = p.next[b.id], p.next[a.id]

	p.next[alpha], p.next[beta] = p.next[beta], p.next[alpha]
}

//------------------------------------------------------------------------------

// Connect creates and returns a new edge connecting the destination of a to the
// origin of b, in such a way that a, e and b share the same left face. For
// convenience, it also sets the origin and destination vertex ID of the new
// edge (but not the right and left face ID).
func Connect(a, b Edge) Edge {
	e := a.pool.New()
	e.SetOrig(a.Dest())
	e.SetDest(b.Orig())
	Splice(e, a.LeftNext())
	Splice(e.Sym(), b)
	return e
}

//------------------------------------------------------------------------------

// SwapTriangles disconnects e, which must have two triangles as left and right
// faces, and reconnect it to the other two vertices of the quadrilateral formed
// by the first step. For example, if v1, v2, v3 and v4 form a quadrilateral,
// and e is a diagonal between v1, and v3, after the swap e will connect v2 to
// v4.
func SwapTriangles(e Edge) {
	a := e.OrigPrev()
	b := e.Sym().OrigPrev()
	Splice(e, a)
	Splice(e.Sym(), b)
	Splice(e, a.LeftNext())
	Splice(e.Sym(), b.LeftNext())
	e.SetOrig(a.Dest())
	e.SetDest(b.Dest())
}

//------------------------------------------------------------------------------
