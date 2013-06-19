// Copyright Published at floating-point-gui.de
// under the Creative Commons Attribution License (BY)
// http://creativecommons.org/licenses/by/3.0/

package math

//------------------------------------------------------------------------------

// Returns true if the relative error between a and b is less than epsilon.
//
// Handles special cases: zero, infinites, denormals.
func NearlyEqual(a, b float32, epsilon float32) bool {
	// Source: http://floating-point-gui.de/errors/comparison/

	absA := a
	if absA < 0 {
		absA = -absA
	}
	absB := b
	if absB < 0 {
		absB = -absB
	}
	diff := a - b
	if diff < 0 {
		diff = -diff
	}

	if a == b {
		// Shortcut, handles infinities.
		return true
	} else if a == 0 || b == 0 || diff < SmallestNormalFloat32 {
		// a or b is zero or both are extremely close to it.
		// Relative error is less meaningful here.
		return diff < (epsilon * SmallestNormalFloat32)
	} else {
		// Use relative error.
		return diff/(absA+absB) < epsilon
	}
}

//------------------------------------------------------------------------------
