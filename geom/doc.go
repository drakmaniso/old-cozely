/*
Package geom provides vectors and matrices, and their associated operations.

All types defined in this package are pure values (no hidden data), and use
the same memory layout than the corresponding GLSL type.

The notation also tries to be as close to GLSL as possible: literals use the
same component order, and component access for matrices is identical: m[2][3]
means the same thing in Go than in GLSL.

Note: this package contains only exported definitions of types; it's safe to
"dot import" it if you want more concise notation.
*/
package geom
