// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

import (
	"errors"
	"image"
	stdcolor "image/color"
	"math"

	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/resource"
)

////////////////////////////////////////////////////////////////////////////////

// Color can convert itself to alpha-premultipled RGBA as 32 bit floats, in
// both linear and standard (sRGB) color spaces.
type Color interface {
	// Linear returns the red, green, blue and alpha values in linear color space.
	// The red, gren and blue values have been alpha-premultiplied in linear
	// space. Each value ranges within [0, 1] and can be used directly by GPU
	// shaders.
	Linear() (r, g, b, a float32)

	// Standard returns the red, green, blue and alpha values in standard (sRGB)
	// color space. The red, gren and blue values have been alpha-premultiplied in
	// linear space. Each value ranges within [0, 1].
	Standard() (r, g, b, a float32)
}

////////////////////////////////////////////////////////////////////////////////

func linearOf(c float32) float32 {
	if c <= 0.04045 {
		return c / 12.92
	}
	return float32(math.Pow(float64(c+0.055)/(1+0.055), 2.4))
}

func standardOf(c float32) float32 {
	if c <= 0.0031308 {
		return 12.92 * c
	}
	return (1+0.055)*float32(math.Pow(float64(c), 1/2.4)) - 0.055
}

////////////////////////////////////////////////////////////////////////////////

// ColorsFrom returns a new Palette created from the file at the specified
// path.
func ColorsFrom(path string) ([]Color, error) {
	var pal = []Color{}

	f, err := resource.Open(path + ".png")
	if err != nil {
		return pal, errors.New("unable to open file for palette " + path)
	}
	defer f.Close() //TODO: error handling
	cf, _, err := image.DecodeConfig(f)
	if err != nil {
		return pal, errors.New("unable to decode file for palette " + path)
	}

	p, ok := cf.ColorModel.(stdcolor.Palette)
	if !ok {
		return pal, errors.New("image file not paletted for palette " + path)
	}

	for i := range p {
		r, g, b, al := p[i].RGBA()
		if i > 255 {
			return pal, errors.New("too many colors for palette " + path)
		}
		c := SRGBA{
			R: float32(r) / float32(0xFFFF),
			G: float32(g) / float32(0xFFFF),
			B: float32(b) / float32(0xFFFF),
			A: float32(al) / float32(0xFFFF),
		}
		//TODO: append name
		pal = append(pal, c)
	}

	internal.Debug.Printf("Loaded color palette (%d entries) from %s", len(p), path)

	return pal, nil
}
