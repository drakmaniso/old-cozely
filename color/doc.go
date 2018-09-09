// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

/*
Package color provides types and functions to manipulate colors.

The types defined here are compatible with the standard library package
"image/color", but are designed to work well with the GPU.

The first difference with the standard library is that most of them are based on
float32 (instead of uint32), so they can be directly passed to GPU shaders.

The second difference is that they make an explicit distinction between linear
and standard (sRGB) color spaces, while the standard library only makes
distinction between alpha-premultiplied and alpha-postmultiplied. Both
distinction are equally important for correct color handling.

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
