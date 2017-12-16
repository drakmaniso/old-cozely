package atlas

import (
	"fmt"
)

//------------------------------------------------------------------------------

// An Image represents a rectangle of pixels that will be packed into an atlas.
type Image interface {
	Width() int16
	Height() int16
	Put(bin int16, x, y int16)
	Paint(dest interface{}) error
}

//------------------------------------------------------------------------------

type node struct {
	first, second *node
	x, y          int16
	w, h          int16
	img           Image
}

//------------------------------------------------------------------------------

func (n *node) String() string {
	if n.first == nil {
		if n.img != nil {
			return fmt.Sprintf("%dx%d", n.img.Width(), n.img.Height())
		}
		return "nil"
	}

	return "{ " + n.first.String() + ", " + n.second.String() + " }"
}

//------------------------------------------------------------------------------

func (n *node) insert(p Image, bin int16) *node {
	// If already split, recurse

	if n.first != nil {
		f := n.first.insert(p, bin)
		if f != nil {
			return f
		}

		return n.second.insert(p, bin)
	}

	// It's a leaf

	if n.img != nil {
		// Already filled
		return nil
	}

	if n.w < p.Width() || n.h < p.Height() {
		// Too small
		return nil
	}

	if n.w == p.Width() && n.h == p.Height() {
		// It's a match!
		n.img = p
		p.Put(bin, n.x, n.y)
		return n
	}

	// Split the leaf

	n.first = new(node)
	n.second = new(node)

	if n.w-p.Width() > n.h-p.Height() {
		n.first.x, n.first.y = n.x, n.y
		n.first.w, n.first.h = p.Width(), n.h

		n.second.x, n.second.y = n.x+p.Width(), n.y
		n.second.w, n.second.h = n.w-p.Width(), n.h

	} else {
		n.first.x, n.first.y = n.x, n.y
		n.first.w, n.first.h = n.w, p.Height()

		n.second.x, n.second.y = n.x, n.y+p.Height()
		n.second.w, n.second.h = n.w, n.h-p.Height()

	}

	return n.first.insert(p, bin)
}

//------------------------------------------------------------------------------

func (n *node) paint(bin int16, dest interface{}) error {
	if n.img != nil {
		return n.img.Paint(dest)
	}

	if n.first != nil {
		err := n.first.paint(bin, dest)
		if err != nil {
			return err
		}

		return n.second.paint(bin, dest)
	}

	return nil
}

//------------------------------------------------------------------------------
