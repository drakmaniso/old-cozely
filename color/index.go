// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

////////////////////////////////////////////////////////////////////////////////

// An Index identifies a color inside the palette.
type Index uint8

// Transparent is a reserved index in every palettes.
const Transparent Index = 0

////////////////////////////////////////////////////////////////////////////////

// Color returns the color corresponding to a color index.
func (c Index) Color() LRGBA {
	return colours[c]
}

////////////////////////////////////////////////////////////////////////////////
