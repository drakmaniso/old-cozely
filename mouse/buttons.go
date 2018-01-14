// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mouse

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

// A Button on the mouse
type Button = internal.MouseButton

// Button constants
const (
	Left   Button = 1
	Middle Button = 2
	Right  Button = 3
	Extra1 Button = 4
	Extra2 Button = 5
)

// IsPressed returns true if a specific button is held down.
func IsPressed(b Button) bool {
	var m uint32 = 1 << (b - 1)
	return internal.MouseButtons&m != 0
}

//------------------------------------------------------------------------------
