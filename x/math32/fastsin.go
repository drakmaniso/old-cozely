// This code is adapted from:
// http://devmaster.net/forums/topic/4648-fast-and-accurate-sinecosine/
// Copyright Nicolas Capens

package math32

////////////////////////////////////////////////////////////////////////////////

// FastSin returns an approximation of sin(x).
//
// Max absolute error in range [-Pi, Pi]: less than 1e-3
//
// Faster than Sin.
func FastSin(x float32) float32

func fastSin(x float32) float32 {
	const (
		B = 4 / Pi
		C = -4 / (Pi * Pi)
	)

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

// func FastSin(x float32) float32 {
// 	const (
// 		P = 0.225
// 		A = 16 * 0.47434165
// 		B = (1 - P) / 0.47434165
// 	)
//
//     y := x / (2 * Pi)
//
//     y = y - Floor(y + 0.5)  // y in range -0.5..0.5
//
//     y = A * y * (0.5 - Abs(y))
//
//     return y * (B + Abs(y))
//
// }

////////////////////////////////////////////////////////////////////////////////
