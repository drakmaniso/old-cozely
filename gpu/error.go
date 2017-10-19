// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gpu

//------------------------------------------------------------------------------

import "github.com/drakmaniso/carol/internal"

//------------------------------------------------------------------------------

// Err returns the first carol error since the previous call to Err().
func Err() error {
	err := stickyErr
	stickyErr = nil
	return err
}

//------------------------------------------------------------------------------

func setErr(context string, err error) {
	// TODO: use two different functions and a *func variable
	if stickyErr == nil {
		stickyErr = internal.Error(context, err)
	} else {
		if internal.Config.Debug {
			internal.Log("gpu unchecked error:\n%s", internal.Error(context, err))
		}
	}
}

var stickyErr error

//------------------------------------------------------------------------------
