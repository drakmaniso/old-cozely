// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package color

import (
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

var stickyErr error

func init() {
	internal.ColorErr = func() error {
		return stickyErr
	}
}

////////////////////////////////////////////////////////////////////////////////

// Err returns the first unchecked error of the package since last call to the
// function. The error is then considered checked, and further calls to Err will
// return nil until the next error occurs.
//
// Note: errors occuring while there already is an unchecked error will not be
// recorded. However, if the debug mode is active, all errors will be logged.
func Err() error {
	err := stickyErr
	stickyErr = nil
	return err
}

func setErr(err error) {
	if stickyErr == nil {
		stickyErr = err
	}
	internal.Debug.Printf("*** ERROR in package color ***\n%s", err)
}

