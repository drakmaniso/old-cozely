// Copyright (c) 2017-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package atlas

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////

type region struct {
	first, second *region
	x, y          int16
	w, h          int16
	rect          uint32
}

////////////////////////////////////////////////////////////////////////////////

const norect = uint32(0xFFFFFFFF)

////////////////////////////////////////////////////////////////////////////////

func (n *region) String() string {
	if n.first == nil {
		if n.rect != norect {
			return fmt.Sprintf("%d", n.rect)
		}
		return "norect"
	}

	return "{ " + n.first.String() + ", " + n.second.String() + " }"
}

////////////////////////////////////////////////////////////////////////////////

func (n *region) insert(rect uint32, bin int16, size SizeFn, put PutFn) *region {
	// If already split, recurse

	if n.first != nil {
		f := n.first.insert(rect, bin, size, put)
		if f != nil {
			return f
		}

		return n.second.insert(rect, bin, size, put)
	}

	// It's a leaf

	if n.rect != norect {
		// Already filled
		return nil
	}

	w, h := size(rect)

	if n.w < w || n.h < h {
		// Too small
		return nil
	}

	if n.w == w && n.h == h {
		// It's a match!
		n.rect = rect
		put(rect, bin, n.x, n.y)
		return n
	}

	// Split the leaf

	n.first = &region{rect: norect}
	n.second = &region{rect: norect}

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

	return n.first.insert(rect, bin, size, put)
}

////////////////////////////////////////////////////////////////////////////////

func (n *region) paint(bin int16, dest interface{}, paint PaintFn) error {
	if n.rect != norect {
		return paint(n.rect, dest)
	}

	if n.first != nil {
		err := n.first.paint(bin, dest, paint)
		if err != nil {
			return err
		}

		return n.second.paint(bin, dest, paint)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
