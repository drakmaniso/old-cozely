// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package cozely

import (
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// Error returns nil if err is nil, or a wrapped error otherwise.
func Error(context string, err error) error {
	if err == nil {
		return nil
	}
	return internal.WrappedError{context, err}
}

// ShowError shows an error to the user. In debug mode, it only prints to the
// standard error output, otherwise it also brings a dialog box.
func ShowError(e error) {
	internal.Log.Printf("ERROR: %s", e)
	if !internal.Config.Debug {
		err2 := internal.ErrorDialog("ERROR: %s", e)
		if err2 != nil {
			internal.Log.Printf("ERROR opening dialog:\n%s", err2)
		}
	}
}

// Log logs a formated message.
func Log(format string, v ...interface{}) {
	internal.Log.Printf(format, v...)
}

////////////////////////////////////////////////////////////////////////////////
