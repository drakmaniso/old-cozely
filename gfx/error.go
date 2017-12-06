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

var setErr = func(context string, err error) {
	if stickyErr == nil {
		stickyErr = core.Error(context, err)
	}
}

func init() {
	h := core.Hook{
		Callback: hookStickyErr,
		Context:  "while setting up gfx sticky error",
	}
	core.PreSetupHooks = append(core.PreSetupHooks, h)
}

func hookStickyErr() error {
	if core.Config.Debug {
		setErr = func(context string, err error) {
			// TODO: use two different functions and a *func variable
			if stickyErr == nil {
				stickyErr = core.Error(context, err)
			}
			core.Log.Printf("gfx error: %s", core.Error(context, err))
		}
	}

	return nil
}

//------------------------------------------------------------------------------
