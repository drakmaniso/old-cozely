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

// Cos returns the sinus of x.
func Cos(x float32) float32 {
	// (comment from the original source)
	// DESCRIPTION:
	//
	// Returns the circular cotangent of the radian argument x.
	// A common routine computes either the tangent or cotangent.
	//
	//
	//
	// ACCURACY:
	//
	//                      Relative error:
	// arithmetic   domain     # trials      peak         rms
	//    IEEE     +-4096        100000     3.0e-7      4.5e-8
	//
	//
	// ERROR MESSAGES:
	//
	//   message         condition          value returned
	// cot total loss   x > 2^24                0.0
	// cot singularity  x = 0                  MAXNUMF

	const (
		FOPI = 1.27323954473516 // 4/Pi
		DP1  = 0.78515625
		DP2  = 2.4187564849853515625e-4
		DP3  = 3.77489497744594108e-8
	)

	const (
		SINCOF1 = -1.9515295891E-4
		SINCOF2 = 8.3321608736E-3
		SINCOF3 = -1.6666654611E-1
	)

	const (
		COSCOF1 = 2.443315711809948E-5
		COSCOF2 = -1.388731625493765E-3
		COSCOF3 = 4.166664568298827E-2
	)

	sign := float32(1)

	// Make argument positive
	if x < 0 {
		x = -x
	}

	// Integer part of `x/(Pi/4)`
	j := uint32(FOPI * x)
	y := float32(j)

	if j&1 != 0 {
		j++
		y++
	}

	j &= 7 // Octant modulo 360°

	if j > 3 {
		j -= 4
		sign = -sign
	}

	if j > 1 {
		sign = -sign
	}

	// Extended precision modular arithmetic
	x = ((x - y*DP1) - y*DP2) - y*DP3

	z := x * x

	if j == 1 || j == 2 {
		y = ((SINCOF1*z+SINCOF2)*z+SINCOF3)*z*x + x
	} else {
		y = ((COSCOF1*z+COSCOF2)*z+COSCOF3)*z*z - 0.5*z + 1
	}

	return sign * y
}

//------------------------------------------------------------------------------
