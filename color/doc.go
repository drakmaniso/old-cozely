// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

/*
Package color provides types and functions to manipulate colors.

It is compatible with the standard library "image/color" package, but is
designed to work well with GPUs.

The first difference is that it is based on float32 (instead of uint32), which
can be used directly in shaders and graphics API such as OpenGL.

It also makes an explicit distinction between linear and standard (sRGB) color
spaces, in addition to the distinction between alpha-premultiplied and
alpha-postmultiplied already done in the standard library.

An alpha-premultiplied color component has been scaled by alpha, so has valid
values within [0, alpha]. It is the most useful choice for alpha-blending. For
an explanation, see:
https://blogs.msdn.microsoft.com/shawnhar/2009/11/06/premultiplied-alpha/
*/
package color
