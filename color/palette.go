// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

////////////////////////////////////////////////////////////////////////////////

// An Index in the palette.
type Index uint8

// Reserved color indices.
const (
	Transparent = Index(0)
	Black       = Index(251)
	DarkGray    = Index(252)
	MidGray     = Index(253)
	LightGray   = Index(254)
	White       = Index(255)
)

////////////////////////////////////////////////////////////////////////////////

var (
	colors [256]LRGBA
	used   [256]bool
	dirty  bool
)

func init() {
	Clear()
}

var debugColor = LRGBA{1, 0, 1, 1}

////////////////////////////////////////////////////////////////////////////////

// Clear initialisez the palette.
func Clear() {
	for j := range colors {
		i := Index(j)
		switch i {
		case Transparent:
			set(i, LRGBA{0, 0, 0, 0})
		case Black:
			set(i, SRGBA{0, 0, 0, 1})
		case DarkGray:
			set(i, SRGBA{0.25, 0.25, 0.25, 1})
		case MidGray:
			set(i, SRGBA{0.5, 0.5, 0.5, 1})
		case LightGray:
			set(i, SRGBA{0.75, 0.75, 0.75, 1})
		case White:
			set(i, SRGBA{1, 1, 1, 1})
		default:
			set(i, nil)
		}
		dirty = true
	}
}

////////////////////////////////////////////////////////////////////////////////

// Set changes the color associated with an index.
//
// Note that the modified palette will be used for every drawing command of the
// current frame, even those issued before the call to this function. In other
// words, you cannot modify the palette in the middle of a frame.
func Set(i Index, c Color) {
	if i > Transparent && i < Black {
		set(i, c)
		dirty = true //TODO: finer-grained palette upload?
	}
}

func set(i Index, c Color) {
	if c == nil {
		colors[i] = debugColor
		used[i] = false
		return
	}
	colors[i] = LRGBAof(c)
	used[i] = true
}

// Add finds the first unused index in the palette and adds a new color. It
// returns the found index, or 0 if the palette is full.
func Add(c Color) Index {
	i := Index(1)
	for ; i < Black; i++ {
		if !used[i] {
			set(i, c)
			return i
		}
	}
	return Index(0)
}

////////////////////////////////////////////////////////////////////////////////

// Find returns the first color index associated with specific LRGBA
// values. If there isn't any color with these values in the palette, index 0 is
// returned.
func Find(c Color) Index {
	lc := LRGBAof(c)
	for i, pc := range colors {
		if i == 0 || !used[i] {
			continue
		}
		if pc == lc {
			return Index(i)
		}
	}

	return Index(0)
}

//TODO: search by color proximity

////////////////////////////////////////////////////////////////////////////////

// At return the color associated with the index, or nil if the index is not
// used.
func At(i Index) Color {
	if !used[i] {
		return nil
	}
	return colors[i]
}

// LRGBAat returns the color corresponding to a specific index.
func LRGBAat(i Index) (LRGBA, bool) {
	if !used[i] {
		return LRGBA{}, false
	}
	return colors[i], true
}
