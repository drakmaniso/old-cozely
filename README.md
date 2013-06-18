GLaM: OpenGL Mathematics for Go
===============================

Features
--------

GLaM is a [Go](http://golang.org/) package providing mathematical types and 
operations for use with OpenGL.

- Type names mirroring GLSL types: Vec2, Vec3, Vec4, Mat3, Mat4, IVec3...
- All types are pure values: there's no heap allocation, and no hidden data.
- All types have the same memory layout than their corresponding GLSL types.
- Most methods are inlined by the compiler.
- Efficient single-precision math (using assembly when appropriate).
- Several noise functions (Perlin, Simplex, ...).


Package GLaM Math
-----------------

This package provides single-precision (i.e. float32) math functions.


Package GLaM Noise
------------------

The Perlin and Simplex noise functions are adapted from
["Simplex Noise Demystified"](http://www.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf)
by Stefan Gustavson (code in the public domain).


Author
------

Laurent Moussault <moussault.laurent@gmail.com>