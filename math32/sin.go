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

package math32

//------------------------------------------------------------------------------

// Sin returns the sinus of x.
func Sin(x float32) float32 {
	// (comment from the original source)
	// DESCRIPTION:
	//
	// Range reduction is into intervals of pi/4.  The reduction
	// error is nearly eliminated by contriving an extended precision
	// modular arithmetic.
	//
	// Two polynomial approximating functions are employed.
	// Between 0 and pi/4 the sine is approximated by
	//      x  +  x**3 P(x**2).
	// Between pi/4 and pi/2 the cosine is represented as
	//      1  -  x**2 Q(x**2).
	//
	// ACCURACY:
	//
	//                      Relative error:
	// arithmetic   domain      # trials      peak       rms
	//    IEEE    -4096,+4096   100,000      1.2e-7     3.0e-8
	//    IEEE    -8192,+8192   100,000      3.0e-7     3.0e-8
	//
	// ERROR MESSAGES:
	//
	//   message           condition        value returned
	// sin total loss      x > 2^24              0.0
	//
	// Partial loss of accuracy begins to occur at x = 2^13
	// = 8192. Results may be meaningless for x >= 2^24
	// The routine as implemented flags a TLOSS error
	// for x >= 2^24 and returns 0.0.

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

	// Make argument positive but save the sign
	sign := float32(1)
	if x < 0 {
		x = -x
		sign = -1
	}

	// Integer part of `x/(Pi/4)`
	j := uint64(FOPI * x)
	y := float32(j)

	// Map zeros to origin
	if j&1 != 0 {
		j++
		y++
	}

	j &= 7 // Octant modulo 360°

	// Reflect in axis
	if j > 3 {
		sign = -sign
		j -= 4
	}

	// Extended precision modular arithmetic
	x = ((x - y*DP1) - y*DP2) - y*DP3

	z := x * x
	if j == 1 || j == 2 {
		y = ((COSCOF1*z+COSCOF2)*z+COSCOF3)*z*z - 0.5*z + 1
	} else {
		y = ((SINCOF1*z+SINCOF2)*z+SINCOF3)*z*x + x
	}

	return sign * y
}

//------------------------------------------------------------------------------
