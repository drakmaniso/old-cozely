/*
 Package plane provides 2D transforms in homogeneous coordinates.

A transform is written as the multiplication of a matrix by a column-vector.
To transform v by M and then by N, you should write N⋅M⋅v:
	M := Translation(Vec2(10, 15))
	N := Rotation(3.14)
	v := Vec3{1, 2, 1}
	w := Apply(N.Times(M), v)

See package space for more detailed explanations.
*/
package plane
