// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

type Action interface {
	Origins() []Origin
}

const maxID = 0xFFFFFFFF
