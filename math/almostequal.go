// Code based on:
// http://randomascii.wordpress.com/2012/02/25/comparing-floating-point-numbers-2012-edition/
// Copyright Bruce Dawson
// No restrictions. Also no warranty. Use it for what you will. 
// I’d recommend a link back to the article just so people will understand how it works.

package math

//------------------------------------------------------------------------------

func IsAlmostEqual(a, b float32, ulps uint32) bool {
	largest := a
	if b > a {
		largest = b
	}
	if Abs(a-b) < largest*EpsilonFloat32 {
		return true
	}

	ua := Float32bits(a)
	ub := Float32bits(b)

	if ua>>31 != ub>>31 {
		return false
	}

	var ulpsDiff uint32
	if ua > ub {
		ulpsDiff = ua - ub
	} else {
		ulpsDiff = ub - ua
	}

	return ulpsDiff < ulps
}

//------------------------------------------------------------------------------
