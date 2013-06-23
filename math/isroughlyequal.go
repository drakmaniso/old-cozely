// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

//------------------------------------------------------------------------------

// `IsRoughlyEqual` Returns true if the absolute error between `a` and `b` is less than `epsilon`.
//
// See also `IsNearlyEqual` and `IsAlmostEqual`.
func IsRoughlyEqual(a, b float32, epsilon float32) bool {
	if a == b {
		// Shortcut, handles infinities
		return true
	} else {
		// Use absolute error
		return Abs(a-b) < epsilon
	}
}

//------------------------------------------------------------------------------
