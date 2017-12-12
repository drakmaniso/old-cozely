// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

/*
Package colour provides types and functions to manipulate colors.

It is compatible with the standard library "image/color" package, but is
designed to work well with GPUs.

The first difference is that it is based on float32 (instead of uint32), which
can be used directly in shaders and graphics API such as OpenGL.

It also makes an explicit distinction between linear and sRGB color spaces, in
addition to the distinction between alpha-premultiplied and
alpha-postmultiplied.
*/
package colour
