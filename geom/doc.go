/*
Package geom provides vectors and matrices, and their associated operations.

It follows the conventions used by GLSL, by using similar type names, and
most importantly, the same memory layout (i.e. column-major for matrices).

All types are pure values: there's no heap allocation, and no hidden data.

Matrices

Since they are pure values, there is no constructors, only literals. Be aware
that they are specified and stored in column-major order, just like GLSL.
So when writing literals, remember to use the transpose of the mathematical
notation. In other words:
    m := Mat3{
		{a, b, c},
		{d, e, f},
		{g, h, i},
    }
... corresponds to the following mathematical notation:
	⎡ a  d  g ⎤
	⎢ b  e  h ⎥
	⎣ c  f  i ⎦

The same inversion happens with indices: m[column][row] corresponds to the
mathematical indices (row,column).

Although all methods returns their result by value, they take their receiver and
parameters by reference, for efficiency. They are never modified.

Vectors

When creating vector literals, if you are using the "go vet" tool, you must
specify the X,Y,Z,... keys:
	v := Vec4{X: x, Y: y, Z: z, W: w}

If you really need something shorter and cannot ignore go vet, you can write a
small local function:
	func vec4(x, y, z, w: float32) {
		return Vec4{X: x, Y: y, Z: z, W: w}
	}
	v := vec4(x, y, z, w)
*/
package geom
