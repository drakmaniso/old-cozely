// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

//------------------------------------------------------------------------------

// RGBA is a color defined by its red, green and blue components, with alpha.
type RGBA struct {
	R float32
	G float32
	B float32
	A float32
}

// RGB is a color defined by its red, green and blue components.
type RGB struct {
	R float32
	G float32
	B float32
}

// HSVA is a color defined by its hue, saturation and value components,
// with alpha.
type HSVA struct {
	H float32
	S float32
	V float32
	A float32
}

// HSV is a color defined by its hue, saturation and value components.
type HSV struct {
	H float32
	S float32
	V float32
}

// RGBA8 is a color defined by its red, green and blue components, with alpha.
type RGBA8 struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// RGB8 is a color defined by its red, green and blue components.
type RGB8 struct {
	R uint8
	G uint8
	B uint8
}

//------------------------------------------------------------------------------
