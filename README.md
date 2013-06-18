GLaM: OpenGL Mathematics for Go
===============================

Features
--------

GLaM is a [Go](http://golang.org/) package providing mathematical types and 
operations for use with OpenGL. It is written with game development in mind,
so the focus is on speed over portability or accuracy.

**NOTE: this is a work in progress.**

- Vectors and matrices.
- Efficient single-precision math.
- Several noise functions (Perlin, Simplex, ...).


Package GLam
------------

    import "github.com/drakmaniso/glam"

This package provides vectors and matrices, and their associated operations.
 
- The names mirrors the GLSL types: Vec2, Vec3, Vec4, Mat3, Mat4, IVec3...
- All types are pure values: there's no heap allocation, and no hidden data.
- All types have the same memory layout than their corresponding GLSL types.
- Most methods are inlined by the compiler.


Package GLaM Math
-----------------

    import "github.com/drakmaniso/glam/math"

This package aims to provide *fast* float32 math functions, using assembly 
when appropriate.


Package GLaM Noise
------------------

    import "github.com/drakmaniso/glam/noise"

The Perlin and Simplex noise functions are adapted from
["Simplex Noise Demystified"](http://www.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf)
by Stefan Gustavson (code in the public domain).


Author
------

Laurent Moussault <moussault.laurent@gmail.com>