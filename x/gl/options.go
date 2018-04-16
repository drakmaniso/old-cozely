// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

////////////////////////////////////////////////////////////////////////////////

// An Option represents a configuration option used to change some parameters of
// the framework: see cozely.Configure.
type Option = func() error

var noclear = false

////////////////////////////////////////////////////////////////////////////////

// NoClear prevents the default behavior of clearing the default framebuffer at
// the start of each frame.
func NoClear() Option {
	return func() error {
		noclear = true
		return nil
	}
}
