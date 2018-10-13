// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

import "errors"

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

var debugColor = LRGBA{1, 0, 1, 1}

////////////////////////////////////////////////////////////////////////////////

func init() {
	Clear()
}

var (
	sources [256]Color
	colors  [256]LRGBA
	dirty   bool
)

////////////////////////////////////////////////////////////////////////////////

// Clear initialisez the master palette.
func Clear() {
	for j := range colors {
		i := Index(j)
		switch i {
		case Transparent:
			Set(i, LRGBA{0, 0, 0, 0})
		case Black:
			Set(i, SRGB8{0, 0, 0})
		case DarkGray:
			Set(i, SRGB8{64, 64, 64})
		case MidGray:
			Set(i, SRGB8{128, 128, 128})
		case LightGray:
			Set(i, SRGB8{196, 196, 196})
		case White:
			Set(i, SRGB8{255, 255, 255})
		default:
			Set(i, nil)
		}
	}
}

// Load clears the master palette and fill it with the specified colors. The
// first entry in the slice must be transparent (LRGBA{}).
func Load(c []Color) error {
	Clear()
	if len(c) == 0 {
		return errors.New("color.Load: invalid palette (empty slice)")
	}
	if c[0] != (LRGBA{}) {
		return errors.New("color.Load: invalid palette (first color must be trasnparent)")
	}
	if len(c) > 256 {
		return errors.New("color.Load: invalid palette (too many colors)")
	}
	for i := range c {
		Set(Index(i), c[i])
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// Set changes the color associated with an index in the master palette.
//
// Note that the modified palette will be used for every drawing command of the
// current frame, even those issued before the call to this function. In other
// words, you cannot modify the palette in the middle of a frame.
func Set(i Index, c Color) {
	switch {
	case i == Transparent:
		return

	case c == nil:
		sources[i] = nil
		colors[i] = debugColor

	default:
		sources[i] = c
		colors[i] = LRGBAof(c)
	}
	dirty = true //TODO: finer-grained palette upload?
}

// Add finds the first unused index in the master palette and adds a new color.
// It returns the found index, or 0 if the palette is full.
func Add(c Color) Index {
	i := Index(1)
	for ; i < Black; i++ {
		if sources[i] == nil {
			Set(i, c)
			return i
		}
	}
	return Index(0)
}

////////////////////////////////////////////////////////////////////////////////

// Find returns the first color index associated with specific LRGBA values in
// the master palette. If there isn't any color with these values in the
// palette, index 0 is returned.
func Find(c Color) Index {
	lc := LRGBAof(c)
	for i := range colors {
		if i == 0 || sources[i] == nil {
			continue
		}
		if c == sources[i] || lc == colors[i] {
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
	return sources[i]
}

// LRGBAat returns the color corresponding to a specific index.
func LRGBAat(i Index) (LRGBA, bool) {
	if sources[i] == nil {
		return LRGBA{}, false
	}
	return colors[i], true
}
