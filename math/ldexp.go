// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

//------------------------------------------------------------------------------

// Ldexp is the inverse of Frexp.
// It returns frac × 2**exp.
//
// Special cases are:
//	Ldexp(±0, exp) = ±0
//	Ldexp(±Inf, exp) = ±Inf
//	Ldexp(NaN, exp) = NaN
func Ldexp(frac float32, exp int) float32 {
	//TODO: asm?
	// special cases
	switch {
	case frac == 0:
		return frac // correctly return -0
	case IsInf(frac, 0) || IsNaN(frac):
		return frac
	}
	frac, e := normalize(frac)
	exp += e
	x := Float32bits(frac)
	exp += int(x>>shift)&mask - bias
	//TODO: convert to 32bit
	// if exp < -1074 {
	// 	return Copysign(0, frac) // underflow
	// }
	// if exp > 1023 { // overflow
	// 	if frac < 0 {
	// 		return Inf(-1)
	// 	}
	// 	return Inf(1)
	// }
	var m float32 = 1
	// if exp < -1022 { // denormal
	// 	exp += 52
	// 	m = 1.0 / (1 << 52) // 2**-52
	// }
	x &^= mask << shift
	x |= uint32(exp+bias) << shift
	return m * Float32frombits(x)
}

//------------------------------------------------------------------------------
