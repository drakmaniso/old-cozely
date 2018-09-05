// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package coord

////////////////////////////////////////////////////////////////////////////////

// Coordinates represents any three-dimensional vector.
type Coordinates interface {
	// Cartesian returns the cartesian coordinates of the vector.
	Cartesian() (x, y, z float32)
}
