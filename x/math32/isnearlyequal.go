// Based on code from http://floating-point-gui.de/errors/comparison/
// Copyright Published at floating-point-gui.de
// under the Creative Commons Attribution License (BY)
// http://creativecommons.org/licenses/by/3.0/

package math32

//------------------------------------------------------------------------------

// IsNearlyEqual Returns true if the relative error between a and b is less
// than epsilon.
//
// Handles special cases: zero, infinites, denormals.
//
// See also IsAlmostEqual and IsRoughlyEqual.
func IsNearlyEqual(a, b float32, epsilon float32) bool {
	// Source: http://floating-point-gui.de/errors/comparison/

	absA := Abs(a)
	absB := Abs(b)
	diff := Abs(a - b)

	if a == b {
		// Shortcut, handles infinities.
		return true
	}
	if a == 0 || b == 0 || diff < SmallestNormalFloat32 {
		// a or b is zero or both are extremely close to it.
		// Relative error is less meaningful here.
		return diff < epsilon*SmallestNormalFloat32
	}
	// Use relative error.
	// Note in the original source, absA+absB was used instead of largest
	largest := absA
	if absB > absA {
		largest = absB
	}
	return diff/largest < epsilon
}

//------------------------------------------------------------------------------
