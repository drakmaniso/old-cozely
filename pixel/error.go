// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/cozely/internal"
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

func setErr(context string, err error) {
	if stickyErr == nil {
		stickyErr = internal.Error(context, err)
	}
	internal.Debug.Printf("pixel error: %s", internal.Error(context, err))
}

//------------------------------------------------------------------------------
