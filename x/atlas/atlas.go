// Copyright (c) 2017-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package atlas

import (
	"sort"
)

/*
The current implementation is rather naive. The objective is to later improve it
with ideas from the PhD Thesis of Andrea Lodi, "Algorithms for Two Dimensional
Bin Packing and Assignment Problems":

  http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.98.3502&rep=rep1&type=pdf

TODO:

- add a Flip method to Image, to be able to make all images horizontal
- find a way to compute the "touching perimeter" score
- pre-allocate the bins
*/

////////////////////////////////////////////////////////////////////////////////

// An Atlas contains the mapping information to pack a set of images into an
// array of bigger textures (called bins).
type Atlas struct {
	width, height int16
	bins          []region
	ideal         int
}

type (
	SizeFn  func(rect uint32) (int16, int16)
	PutFn   func(rect uint32, bin int16, x, y int16)
	PaintFn func(rect uint32, dest interface{}) error
)

////////////////////////////////////////////////////////////////////////////////

// New returns a new (empty) Atlas. The width and height describe the shape of
// the bins.
func New(width, height int16) *Atlas {
	return &Atlas{
		width:  width,
		height: height,
		bins:   []region{},
	}
}

////////////////////////////////////////////////////////////////////////////////

// BinSize returns the width and height of the bins managed by the atlas.
func (a *Atlas) BinSize() (width, height int16) {
	return a.width, a.height
}

// BinCount returns the number of bins currently in the atlas.
func (a *Atlas) BinCount() int16 {
	return int16(len(a.bins))
}

// Unused returns the number of unused pixels (i.e. not allocated to any image)
// in the atlas.
func (a *Atlas) Unused() int {
	return len(a.bins)*int(a.width)*int(a.height) - a.ideal
}

////////////////////////////////////////////////////////////////////////////////

// Pack fits all the rectangles in the atlas. New bins are added when needed. It
// calls the Put method of each image with the corresponding mapping
// information.
func (a *Atlas) Pack(rectangles []uint32, size SizeFn, put PutFn) {
	sort.Slice(rectangles, func(i, j int) bool {
		wi, hi := size(rectangles[i])
		wj, hj := size(rectangles[j])
		return wi*2+hi*2 > wj*2+hj*2
	})

	for i := range rectangles {
		w, h := size(rectangles[i])
		a.ideal += int(w) * int(h)
		done := false
		for j := range a.bins {
			n := a.bins[j].insert(rectangles[i], int16(j), size, put)
			if n != nil {
				done = true
				break
			}
		}
		if !done {
			a.bins = append(
				a.bins,
				region{w: a.width, h: a.height, rect: norect},
			)
			n := a.bins[len(a.bins)-1].insert(rectangles[i], int16(len(a.bins)-1), size, put)
			if n != nil {

			} else {
				print("!") //TODO:
			}
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

// Paint iterates on all images mapped to the specified bin, and call their own
// Paint method.
func (a *Atlas) Paint(bin int16, dest interface{}, paint PaintFn) error {
	return a.bins[bin].paint(bin, dest, paint)
}

////////////////////////////////////////////////////////////////////////////////
