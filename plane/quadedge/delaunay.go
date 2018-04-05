// Adapted by Laurent Moussault (2018) from C code: see ORIGINAL_LICENSE

package quadedge

import (
	"fmt"
	"sort"

	"github.com/drakmaniso/glam/plane"
)

//------------------------------------------------------------------------------

func Delaunay(sites []plane.Coord) Edge {
	fmt.Println(sites)
	// Construct indices and remove duplicates
	v := make([]uint32, 0, len(sites))
	for i := range sites {
		n := true
		for j := range sites[:i] {
			if sites[j] == sites[i] {
				n = false
				break
			}
		}
		if n {
			v = append(v, uint32(i))
		}
	}
	if len(v) < 2 {
		return Edge{id: Nil}
	}

	// Sort
	sort.Slice(v, func(i int, j int) bool {
		switch {
		case sites[v[i]].X < sites[v[j]].X:
			return true
		case sites[v[i]].X > sites[v[j]].X:
			return false
		default:
			return sites[v[i]].Y < sites[v[j]].Y
		}
	})

	// Divide and conquer algorithm
	p := NewPool(uint32(len(v) * 10)) //TODO: correct size
	l, _ := delaunay(sites, p, v)

	return l
}

//------------------------------------------------------------------------------

func delaunay(coords []plane.Coord, p *Pool, sub []uint32) (l, r Edge) {
	if len(sub) == 2 {
		// Create an edge connecting sub[0] to sub[1]
		a := New(p)
		a.SetOrig(sub[0])
		a.SetDest(sub[1])
		return a, a.Sym()
	}

	if len(sub) == 3 {
		// Create edges connecting sub[0] to sub[1] and sub[1] to sub[2]
		a := New(p)
		b := New(p)
		Splice(a.Sym(), b)
		a.SetOrig(sub[0])
		a.SetDest(sub[1])
		b.SetOrig(sub[1])
		b.SetDest(sub[2])
		// Close the triangle
		if plane.IsCCW(coords[sub[0]], coords[sub[1]], coords[sub[2]]) {
			_ = Connect(b, a)
			return a, b.Sym()
		}
		if plane.IsCCW(coords[sub[0]], coords[sub[2]], coords[sub[1]]) {
			c := Connect(b, a)
			return c.Sym(), c
		}
		// The three points are colinear
		return a, b.Sym()
	}

	// Recursion
	lout, lins := delaunay(coords, p, sub[:len(sub)/2])
	rins, rout := delaunay(coords, p, sub[len(sub)/2:])

	// Compute the lower common tangent of L and R
loop:
	for {
		switch {
		case leftOf(coords, rins.Orig(), lins):
			lins = lins.LeftNext()
		case rightOf(coords, lins.Orig(), rins):
			rins = rins.RightPrev()
		default:
			break loop
		}
	}

	// Create a first cross edge base from rdi.Org to ldi.Org
	base := Connect(rins.Sym(), lins)
	if lins.Orig() == lout.Orig() {
		lout = base.Sym()
	}
	if rins.Orig() == rout.Orig() {
		rout = base
	}

	// Merge
	for {
		// Locate the first L point lcand.Dest to be encountered by the rising
		// bubble, and delete L edges out of basel.Dest that fail the circle test.
		lcand := base.Sym().OrigNext()
		if valid(coords, lcand, base) {
			for inCircle(coords,
				base.Dest(), base.Orig(), lcand.Dest(), lcand.OrigNext().Dest(),
			) {
				t := lcand.OrigNext()
				Delete(lcand)
				lcand = t
			}
		}
		// Symmetrically, locate the first R point to be hit, and delete R edges.
		rcand := base.OrigPrev()
		if valid(coords, rcand, base) {
			for inCircle(coords,
				base.Dest(), base.Orig(), rcand.Dest(), rcand.OrigPrev().Dest(),
			) {
				t := rcand.OrigPrev()
				Delete(rcand)
				rcand = t
			}
		}
		// If both lcand and rcand are invalid, then basel is the upper common
		// tangent
		if !valid(coords, lcand, base) && !valid(coords, rcand, base) {
			break
		}
		// the next cross edge is to be connected to either lcand.Dest or
		// rcand.Dest. If both are valid, then choose the appropriate one using the
		// InCircle test.
		if !valid(coords, lcand, base) ||
			(valid(coords, rcand, base) &&
				inCircle(coords,
					lcand.Dest(), lcand.Orig(), rcand.Orig(), rcand.Dest())) {
			base = Connect(rcand, base.Sym())
		} else {
			base = Connect(base.Sym(), lcand.Sym())
		}
	}

	return lout, rout
}

//------------------------------------------------------------------------------

func inCircle(sites []plane.Coord, a, b, c, d uint32) bool {
	return plane.InCircumcircle(sites[a], sites[b], sites[c], sites[d])
}

func rightOf(sites []plane.Coord, p uint32, e Edge) bool {
	return plane.IsCCW(sites[p], sites[e.Dest()], sites[e.Orig()])
}

func leftOf(sites []plane.Coord, p uint32, e Edge) bool {
	return plane.IsCCW(sites[p], sites[e.Orig()], sites[e.Dest()])
}

func valid(sites []plane.Coord, e, f Edge) bool {
	return rightOf(sites, e.Dest(), f)
}

//------------------------------------------------------------------------------
