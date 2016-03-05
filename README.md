Glam: a minimalist game engine in Go
====================================


**WARNING: This is a work in progress, and far from complete.**


Implemented Features
--------------------

- Vectors and matrices.
- Efficient single-precision math.
- Some noise functions (Perlin, Simplex...).


Package geom
------------

    import "github.com/drakmaniso/glam/geom"

This package provides vectors and matrices, and their associated operations.

- The names mirrors the GLSL types: Vec2, Vec3, Vec4, Mat3, Mat4, IVec3...
- All types are pure values: there's no heap allocation, and no hidden data.
- All types have the same memory layout than their corresponding GLSL types.
- Most methods are inlined by the compiler.


Package math
------------

    import "github.com/drakmaniso/glam/math"

This package aims to provide *fast* float32 math functions, using assembly
when appropriate.


Package noise
-------------

    import "github.com/drakmaniso/glam/noise"

The Perlin and Simplex noise functions are adapted from
["Simplex Noise Demystified"](http://www.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf)
by Stefan Gustavson (code in the public domain).


Author
------

Laurent Moussault <moussault.laurent@gmail.com>
