// Copyright (c) 2017-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package atlas

import (
	"fmt"
)

//------------------------------------------------------------------------------

// An Image represents a rectangle of pixels that will be packed into an atlas.
type Image interface {
	Size() (width, height int16)
	Put(bin int16, x, y int16)
	Paint(dest interface{}) error
}

//------------------------------------------------------------------------------

type region struct {
	first, second *region
	x, y          int16
	w, h          int16
	img           Image
}

//------------------------------------------------------------------------------

func (n *region) String() string {
	if n.first == nil {
		if n.img != nil {
			w, h := n.img.Size()
			return fmt.Sprintf("%dx%d", w, h)
		}
		return "nil"
	}

	return "{ " + n.first.String() + ", " + n.second.String() + " }"
}

//------------------------------------------------------------------------------

func (n *region) insert(p Image, bin int16) *region {
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

	w, h := p.Size()

	if n.w < w || n.h < h {
		// Too small
		return nil
	}

	if n.w == w && n.h == h {
		// It's a match!
		n.img = p
		p.Put(bin, n.x, n.y)
		return n
	}

	// Split the leaf

	n.first = new(region)
	n.second = new(region)

	if n.w-w > n.h-h {
		n.first.x, n.first.y = n.x, n.y
		n.first.w, n.first.h = w, n.h

		n.second.x, n.second.y = n.x+w, n.y
		n.second.w, n.second.h = n.w-w, n.h

	} else {
		n.first.x, n.first.y = n.x, n.y
		n.first.w, n.first.h = n.w, h

		n.second.x, n.second.y = n.x, n.y+h
		n.second.w, n.second.h = n.w, n.h-h

	}

	return n.first.insert(p, bin)
}

//------------------------------------------------------------------------------

func (n *region) paint(bin int16, dest interface{}) error {
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
