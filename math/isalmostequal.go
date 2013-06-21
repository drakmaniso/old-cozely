// Code partly based on:
// http://randomascii.wordpress.com/2012/02/25/comparing-floating-point-numbers-2012-edition/
// Copyright Bruce Dawson
// No restrictions. Also no warranty. Use it for what you will.
// I’d recommend a link back to the article just so people will understand how it works.

package math

//------------------------------------------------------------------------------

// IsAlmostEqual returns true if the difference between a and b in ULPs
// (Unit in the Last Place) is less than ulps.
func IsAlmostEqual(a, b float32, ulps uint32) bool {

	diff := Abs(a - b)

	if a == b {
		// Shortcut, handles infinities.
		return true
	} else if a == 0 || b == 0 || diff < SmallestNormalFloat32 {
		// a or b is zero or both are extremely close to it.
		// Relative error is less meaningful here.
		return diff < float32(ulps)*SmallestNonzeroFloat32
	} else {
		ua := Float32bits(a)
		ub := Float32bits(b)

		// Different signs means they do not match.
		if ua>>31 != ub>>31 {
			return false
		}

		// Find the difference in ULPs.
		var ulpsDiff uint32
		if ua > ub {
			ulpsDiff = ua - ub
		} else {
			ulpsDiff = ub - ua
		}

		return ulpsDiff < ulps
	}
}

// Original algorithm:
func isAlmostEqual(a, b float32, ulps uint32) bool {
	largest := a
	if b > a {
		largest = b
	}

	// Check if the numbers are really close -- needed
	// when comparing numbers near zero.
	if Abs(a-b) < largest*EpsilonFloat32 {
		return true
	}

	ua := Float32bits(a)
	ub := Float32bits(b)

	// Different signs means they do not match.
	if ua>>31 != ub>>31 {
		return false
	}

	// Find the difference in ULPs.
	var ulpsDiff uint32
	if ua > ub {
		ulpsDiff = ua - ub
	} else {
		ulpsDiff = ub - ua
	}

	return ulpsDiff < ulps
}

//------------------------------------------------------------------------------
