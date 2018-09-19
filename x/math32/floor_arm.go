// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math32

// Floor returns the nearest integer less than or equal to x.
func Floor(x float32) float32 {
	if x == 0 || IsNaN(x) || IsInf(x, 0) {
		return x
	}
	if x < 0 {
		d, fract := Modf(-x)
		if fract != 0.0 {
			d = d + 1
		}
		return -d
	}
	d, _ := Modf(x)
	return d

}

////////////////////////////////////////////////////////////////////////////////
