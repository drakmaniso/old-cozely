// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import "github.com/drakmaniso/glam/internal"

//------------------------------------------------------------------------------

// Err returns the first glam error since the previous call to Err().
func Err() error {
	err := stickyErr
	stickyErr = nil
	return err
}

//------------------------------------------------------------------------------

func setErr(err error) {
	// TODO: use two different functions and a *func variable
	if stickyErr == nil {
		internal.Log("gfx error: %s", err)
		stickyErr = err
	} else {
		internal.Log("unchecked gfx error: %s", err)
	}
}

var stickyErr error

//------------------------------------------------------------------------------
