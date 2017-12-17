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

//------------------------------------------------------------------------------

// An Atlas contains the mapping information to pack a set of images into an
// array of bigger textures (called bins).
type Atlas struct {
	width, height int16
	bins          []region
	ideal         int
}

//------------------------------------------------------------------------------

// New returns a new (empty) Atlas. The width and height describe the shape of
// the bins.
func New(width, height int16) *Atlas {
	return &Atlas{
		width:  width,
		height: height,
		bins:   []region{},
	}
}

//------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------

// Pack fits all the source images in the atlas. New bins are added when needed.
// It calls the Put method of each image with the corresponding mapping
// information.
func (a *Atlas) Pack(sources []Image) {
	sort.Sort(byPerimeter(sources))

	for i := range sources {
		a.ideal += int(sources[i].Width()) * int(sources[i].Height())
		done := false
		for j := range a.bins {
			n := a.bins[j].insert(sources[i], int16(j))
			if n != nil {
				done = true
				break
			}
		}
		if !done {
			a.bins = append(
				a.bins,
				region{w: a.width, h: a.height},
			)
			n := a.bins[len(a.bins)-1].insert(sources[i], int16(len(a.bins)-1))
			if n != nil {

			} else {
				print("!")
			}
		}
	}
}

//------------------------------------------------------------------------------

type byPerimeter []Image

func (bp byPerimeter) Len() int {
	return len(bp)
}

func (bp byPerimeter) Swap(i, j int) {
	bp[i], bp[j] = bp[j], bp[i]
}

func (bp byPerimeter) Less(i, j int) bool {
	return bp[i].Width()*2+bp[i].Height()*2 > bp[j].Width()*2+bp[j].Height()*2
}

//------------------------------------------------------------------------------

// Paint iterates on all images mapped to the specified bin, and call their own
// Paint method.
func (a *Atlas) Paint(bin int16, dest interface{}) error {
	return a.bins[bin].paint(bin, dest)
}

//------------------------------------------------------------------------------
