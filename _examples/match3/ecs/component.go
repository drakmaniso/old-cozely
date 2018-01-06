// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package ecs

//------------------------------------------------------------------------------

type Mask uint32

const (
	GridPosition Mask = 1 << iota
	Color
)

//------------------------------------------------------------------------------

var components [Size]Mask

//------------------------------------------------------------------------------
