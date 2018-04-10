// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

type Device uint32

const noDevice = Device(maxID)

const (
	Any              Device = 0
	KeyboardAndMouse Device = iota
)

const maxDevices = 16
