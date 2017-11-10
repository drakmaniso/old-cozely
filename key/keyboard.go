// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package key

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

// IsPressed returns true if the corresponding key position is currently
// held down.
func IsPressed(pos Position) bool {
	return internal.KeyState[pos]
}

// LabelOf returns the key label at the specified position in the current
// layout.
func LabelOf(pos Position) Label {
	return internal.KeyLabelOf(pos)
}

// SearchPositionOf searches the current position of label in the current
// layout.
func SearchPositionOf(l Label) Position {
	return internal.KeySearchPositionOf(l)
}

//------------------------------------------------------------------------------
