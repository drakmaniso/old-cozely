// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package carol

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal/core"
)

//------------------------------------------------------------------------------

// Error returns nil if err is nil, or a wrapped error otherwise.
func Error(context string, err error) error {
	if err == nil {
		return nil
	}
	return core.WrappedError{context, err}
}

// ShowError shows an error and its context to the user. In debug mode, it only
// prints to the standard error output, otherwise it also brings a dialog box.
func ShowError(context string, err error) {
	e := Error(context, err)
	core.Log.Printf("ERROR: %s", e)
	if !core.Config.Debug {
		err2 := core.ErrorDialog("ERROR: %s", e)
		if err2 != nil {
			core.Log.Printf("ERROR opening dialog:\n%s", err2)
		}
	}
}

// Log logs a formated message.
func Log(format string, v ...interface{}) {
	core.Log.Printf(format, v...)
}

//------------------------------------------------------------------------------
