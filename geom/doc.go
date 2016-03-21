/*
Package geom provides vectors and matrices, and their associated operations.

All types defined in this package use the same memory layout than the
corresponding GLSL type. They are pure values (no hidden data).

The notation also tries to be as close to GLSL as possible: literals use the
same component order, function names are similar, and component access for
matrices is identical: m[2][3] means the same thing in Go than in GLSL.

Transformation Matrices

As is usual in GLSL, to transform a vector use left-multiplication by a
matrix:
	T := Translation4(10, 15, 2)
	v := Vec4{1, 2, 3}
	vTrans := T.Transform(v)

When writing literals, remember to use the transpose of the mathematical
notation. In other words the following mathematical notation:
	⎡ a11  a12  a13 ⎤
	⎢ a21  a22  a23 ⎥
	⎣ a31  a32  a33 ⎦
Translates to:
    m := Mat3{
		{a11, a21, a31},
		{a12, a22, a32},
		{a13, a23, a33},
    }

Note that the same inversion happens with indices: the last component
of the first column is written a31 in math but accessed with m[0][2] in Go
(and GLSL).

Finally, although all methods returns their result by value, they take their
receiver and parameters by reference, for efficiency. They are never modified.

Note: Some describes this convention as "column-major". This can be confusing,
because it depends on the meaning assigned to the indices of 2D arrays. E.g.
Wikipedia describes C as row-major, but in doing so assumes that arrays are
accessed with a[row][col]. OpenGL, defined as column-major, uses (inGLSL) the
exact same data structure and notation than C, but assumes arrays are accessed
with a[col][row]. What really matters is the underlying memory layout, and the
order of matrix-vector multiplication. Yet another way to describe the
situation is to say that vectors are treated as column-vectors.
*/
package geom
