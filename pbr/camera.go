// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pbr

//------------------------------------------------------------------------------

import "math"

//------------------------------------------------------------------------------

func Exposure(aperture, shutterTime, sensitivity float64) float64 {
	// See "Moving Frostbite to Physically Based Rendering", Lagarde, de Rousiers (SIGGRAPH 2014)
	ev100 := math.Log2((aperture * aperture) / shutterTime * 100.0 / sensitivity)
	maxLum := 1.2 * math.Pow(2.0, ev100)
	return 1.0 / maxLum
}

//------------------------------------------------------------------------------
