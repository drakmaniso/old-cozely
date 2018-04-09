// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

type Action interface {
	Active() bool
}

const (
	flagActive byte = 1 << iota
)

const maxID = 0xFFFFFFFF
