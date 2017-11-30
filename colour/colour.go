// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package colour

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

// HSLA is a color defined by its hue, saturation and luminance components,
// with alpha.
type HSLA struct {
	H float32
	S float32
	L float32
	A float32
}

// HSL is a color defined by its hue, saturation and luminance components.
type HSL struct {
	H float32
	S float32
	L float32
}

// HCLA is a color defined by its hue, chroma and luminance components,
// with alpha.
type HCLA struct {
	H float32
	C float32
	L float32
	A float32
}

// HCL is a color defined by its hue, chroma and luminance components.
type HCL struct {
	H float32
	C float32
	L float32
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

func (c RGB) Times(s float32) RGB {
	return RGB{c.R * s, c.G * s, c.B * s}
}

//------------------------------------------------------------------------------
