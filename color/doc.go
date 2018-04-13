// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

/*
Package color provides types and functions to manipulate colors.

The Palette

The most important function of the package is to handle the global color
palette. This palette is used by both packages pixel and poly.

  var (
      palette = color.Palette()
      orange  = palette.Entry(color.SRGB{1, 0.6, 0})
      cyan    = palette.Entry(color.SRGB{0, 0.9, 1})
      black   = palette.Entry(color.SRGB{0, 0, 0})
  )

The Colors

Although packages pixel and poly only work with color indices, There's still the
need for a way to specify the color associated with each palette entry.

This is the role of the various RGB structs in this package. Those types are
compatible with the standard library "image/color" package, but are designed to
work well with the GPU.

The first difference is that most of them based on float32 (instead of uint32),
which can be directly passed to GPU shaders.

These types also makes an explicit distinction between linear and standard
(sRGB) color spaces, while the standard library only makes distinction between
alpha-premultiplied and alpha-postmultiplied. Both distinction are equally
important for correct color handling.

Linear and sRGB

Structs prefixed by "L" are in linear color space, while structs prefixed by "S"
are in standard (sRGB) color space.

For the importance of this distinction, see:
http://blog.johnnovak.net/2016/09/21/what-every-coder-should-know-about-gamma/

Alpha Pre-Multipled and Post-Multiplied

Structs ending with "nA" are alpha post-multplied; all others are alpha
pre-multiplied.

In an alpha-premultiplied color, the three RGB component have been scaled by
alpha; valid values are therefore within [0, alpha]. This is the most useful
choice for alpha-blending.

For the importance of alpha pre-multipled, see:
https://blogs.msdn.microsoft.com/shawnhar/2009/11/06/premultiplied-alpha/
*/
package color
