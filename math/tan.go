// Based on code from the Cephes Mathematical Library
// http://www.moshier.net/#Cephes
// or http://www.netlib.org/cephes/single.tgz
// Copyright Stephen L. Moshier
// From the README:
//    Some software in this archive may be from the book _Methods and
// Programs for Mathematical Functions_ (Prentice-Hall or Simon & Schuster
// International, 1989) or from the Cephes Mathematical Library, a
// commercial product. In either event, it is copyrighted by the author.
// What you see here may be used freely but it comes with no support or
// guarantee.

package math

//------------------------------------------------------------------------------

// `Tan` returns the tangent of `x`.
func Tan(x float32) float32 {
	// (comment from the original source)
	// DESCRIPTION:
	//
	// Returns the circular tangent of the radian argument x.
	//
	// Range reduction is modulo pi/4.  A polynomial approximation
	// is employed in the basic interval [0, pi/4].
	//
	// ACCURACY:
	//
	//                      Relative error:
	// arithmetic   domain     # trials      peak         rms
	//    IEEE     +-4096        100000     3.3e-7      4.5e-8
	//
	// ERROR MESSAGES:
	//
	//   message         condition          value returned
	// tanf total loss   x > 2^24              0.0

	const (
		DP1  = float32(0.78515625)
		DP2  = float32(2.4187564849853515625e-4)
		DP3  = float32(3.77489497744594108e-8)
		FOPI = float32(1.27323954473516) // 4/pi
	)

	// Make argument positive but save the sign.
	sign := float32(1)
	if x < 0 {
		x = -x
		sign = -1
	}

	j := uint64(x * FOPI) // Integer part of `x/(Pi/4)`, as integer for tests on the phase angle.
	y := float32(j)       // Integer part of `x/(Pi/4)`, as float.

	// Map zeros and singularities to origin.
	if j&1 == 1 {
		j += 1
		y += 1
	}

	z := ((x - y*DP1) - y*DP2) - y*DP3
	zz := z * z

	if x > 1e-4 {
		// 1.7e-8 relative error in [-pi/4, +pi/4]
		const (
			COEFA = 9.38540185543E-3
			COEFB = 3.11992232697E-3
			COEFC = 2.44301354525E-2
			COEFD = 5.34112807005E-2
			COEFE = 1.33387994085E-1
			COEFF = 3.33331568548E-1
		)
		y = (((((COEFA*zz+COEFB)*zz+COEFC)*zz+COEFD)*zz+COEFE)*zz+COEFF)*zz*z + z
	} else {
		y = z
	}

	if j&2 == 2 {
		y = -1 / y
	}

	return sign * y
}

//------------------------------------------------------------------------------
