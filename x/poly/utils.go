// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

////////////////////////////////////////////////////////////////////////////////

func clamp(v, min, max float64) float64 {
	switch {
	case v < min:
		return min
	case v > max:
		return max
	default:
		return v
	}
}

func saturate(v float64) float64 {
	return clamp(v, 0.0, 1.0)
}

////////////////////////////////////////////////////////////////////////////////
