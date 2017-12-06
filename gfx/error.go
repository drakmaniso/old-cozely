// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal/core"
)

//------------------------------------------------------------------------------

var stickyErr error

// Err returns the first unchecked error of package gfx, and considers it
// checked.
func Err() error {
	err := stickyErr
	stickyErr = nil
	return err
}

func setErr(context string, err error) {
	// TODO: use two different functions and a *func variable
	if stickyErr == nil {
		stickyErr = core.Error(context, err)
	} else {
		if core.Config.Debug {
			core.Log.Printf("gfx unchecked error: %s", core.Error(context, err))
		}
	}
}

//------------------------------------------------------------------------------
