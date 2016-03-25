/*
Package space provides 3D transforms in homogeneous coordinates.

This package uses column-vectors, as this has become the standard mathematical
convention. To transform v by M and then by N, you should write N*M*v:
	M := Translation4(Vec3(10, 15, 2))
	N := Rotation(3.14, Vec3(0, 1, 0))
	v := Vec4{1, 2, 3}
	w := Transform(N.Times(M), v)

In order to be compatible with both GLSL and column vectors, matrices are stored
in "column-major" order. So when writing matrix literals, remember to use the
transpose of the mathematical notation. In other words:
	⎡ a11  a12  a13 ⎤
	⎢ a21  a22  a23 ⎥
	⎣ a31  a32  a33 ⎦
Translates to:
    m := Mat3{
		{a11, a21, a31},
		{a12, a22, a32},
		{a13, a23, a33},
    }

The same inversion happens with indices: the third component of the first column
is written a31 in math but accessed with m[0][2] in Go (like in GLSL).

Finally, although all methods returns their result by value, they take their
receiver and parameters by reference, for efficiency. They are never modified.
*/
package space
