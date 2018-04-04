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
	v := make([]uint32, len(sites), len(sites))
	for i := range v {
		v[i] = uint32(i)
	}
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
	//TODO: remove duplicates
	fmt.Print("[")
	for _, vv := range v {
		fmt.Print(sites[vv], " ")
	}
	fmt.Println("]")

	//
	p := NewPool(uint32(len(sites) * 10)) //TODO: correct size

	l, _ := delaunay(sites, p, v)

	return l
}

//------------------------------------------------------------------------------

func delaunay(sites []plane.Coord, p *Pool, sub []uint32) (l, r Edge) {
	fmt.Println("-> delaunay ", sub)
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
		if plane.IsCCW(sites[sub[0]], sites[sub[1]], sites[sub[2]]) {
			_ = Connect(b, a)
			return a, b.Sym()
		}
		if plane.IsCCW(sites[sub[0]], sites[sub[2]], sites[sub[1]]) {
			c := Connect(b, a)
			return c.Sym(), c
		}
		// The three points are colinear
		return a, b.Sym()
	}

	//
	ldo, ldi := delaunay(sites, p, sub[:len(sub)/2])
	rdi, rdo := delaunay(sites, p, sub[len(sub)/2:])

	// Compute the lower common tangent of L and R
	loop: for {
		switch {
		case leftOf(sites, sites[rdi.Orig()], ldi):
			ldi = ldi.LeftNext()
		case rightOf(sites, sites[ldi.Orig()], rdi):
			rdi = rdi.RightPrev()
		default:
			break loop
		}
	}

	// Create a first cross edge basel from rdi.Org to ldi.Org
	basel := Connect(rdi.Sym(), ldi)
	if ldi.Orig() == ldo.Orig() {
		ldo = basel.Sym()
	}
	if rdi.Orig() == rdo.Orig() {
		rdo = basel
	}

	// Merge
	for {
		// Locate the first L point lcand.Dest to be encountered by the rising
		// bubble, and delete L edges out of basel.Dest that fail the circle test.
		lcand := basel.Sym().OrigNext()
		if valid(sites, lcand, basel) {
			for plane.InCircumcircle(
				sites[basel.Dest()],
				sites[basel.Orig()],
				sites[lcand.Dest()],
				sites[lcand.OrigNext().Dest()],
			) {
				t := lcand.OrigNext()
				Delete(lcand)
				lcand = t
			}
		}
		// Symmetrically, locate the first R point to be hit, and delete R edges.
		rcand := basel.OrigPrev()
		if valid(sites, rcand, basel) {
			for plane.InCircumcircle(
				sites[basel.Dest()],
				sites[basel.Orig()],
				sites[rcand.Dest()],
				sites[rcand.OrigPrev().Dest()],
			) {
				t := rcand.OrigPrev()
				Delete(rcand)
				rcand = t
			}
		}
		// If both lcand and rcand are invalid, then basel is the upper common
		// tangent
		if !valid(sites, lcand, basel) && !valid(sites, rcand, basel) {
			break
		}
		// the next cross edge is to be connected to either lcand.Dest or
		// rcand.Dest. If both are valid, then choose the appropriate one using the
		// InCircle test.
		if !valid(sites, lcand, basel) ||
			(valid(sites, rcand, basel) &&
				plane.InCircumcircle(
					sites[lcand.Dest()],
					sites[lcand.Orig()],
					sites[rcand.Orig()],
					sites[rcand.Dest()])) {
			basel = Connect(rcand, basel.Sym())
		} else {
			basel = Connect(basel.Sym(), lcand.Sym())
		}
	}

	return ldo, rdo
}

//------------------------------------------------------------------------------

func rightOf(sites []plane.Coord, p plane.Coord, e Edge) bool {
	return plane.IsCCW(p, sites[e.Dest()], sites[e.Orig()])
}

func leftOf(sites []plane.Coord, p plane.Coord, e Edge) bool {
	return plane.IsCCW(p, sites[e.Orig()], sites[e.Dest()])
}

func valid(sites []plane.Coord, e, f Edge) bool {
	return rightOf(sites, sites[e.Dest()], f)
}

//------------------------------------------------------------------------------
