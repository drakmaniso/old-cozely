// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

var stickyErr error

// Err returns the first unchecked error of package pixel, and considers it
// checked.
func Err() error {
	err := stickyErr
	stickyErr = nil
	return err
}

var setErr = func(context string, err error) {
	if stickyErr == nil {
		stickyErr = internal.Error(context, err)
	}
	//TODO: simplify, as in core/gl package
}

func init() {
	h := internal.Hook{
		Callback: hookStickyErr,
		Context:  "while setting up gfx sticky error",
	}
	internal.PreSetupHooks = append(internal.PreSetupHooks, h)
}

func hookStickyErr() error {
	if internal.Config.Debug {
		setErr = func(context string, err error) {
			// TODO: use two different functions and a *func variable
			if stickyErr == nil {
				stickyErr = internal.Error(context, err)
			}
			internal.Log.Printf("gfx error: %s", internal.Error(context, err))
		}
	}

	return nil
}

//------------------------------------------------------------------------------
