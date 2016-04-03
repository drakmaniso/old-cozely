// This code is adapted from:
// http://devmaster.net/forums/topic/4648-fast-and-accurate-sinecosine/
// Copyright Nicolas Capens

package math

//------------------------------------------------------------------------------

// FastCos returns an approximation of sin(x).
//
// Max absolute error in range [-Pi, Pi]: less than 1e-3
//
// Faster than Cos.
func FastCos(x float32) float32

func fastCos(x float32) float32 {
	const (
		PISLASHTWO = Pi / 2
		TWOPI      = 2 * Pi
		B          = 4 / Pi
		C          = -4 / (Pi * Pi)
	)

	x += PISLASHTWO
	if x > Pi {
		x -= TWOPI
	}

	y := B*x + C*x*Abs(x)

	//     const (
	//		P = 0.225
	//		Q = 0.775
	// 	)
	const (
		P = 0.224008178776
		Q = 0.775991821224
	)

	//y = P * (y * Abs(y) - y) + y
	y = Q*y + P*y*Abs(y)

	return y
}

//------------------------------------------------------------------------------
