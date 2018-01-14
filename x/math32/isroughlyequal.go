// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math32

//------------------------------------------------------------------------------

// IsRoughlyEqual Returns true if the absolute error between a and b is less
// than epsilon.
//
// See also IsNearlyEqual and IsAlmostEqual.
func IsRoughlyEqual(a, b float32, epsilon float32) bool {
	if a == b {
		// Shortcut, handles infinities
		return true
	}
	// Use absolute error
	return Abs(a-b) < epsilon
}

//------------------------------------------------------------------------------
