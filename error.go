// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package carol

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

// Error returns nil if err is nil, or a wrapped error otherwise.
func Error(context string, err error) error {
	if err == nil {
		return nil
	}
	return wrappedError{context, err}
}

type wrappedError struct {
	context string
	err     error
}

func (e wrappedError) Error() string {
	msg := "- " + e.context + ",\n"
	a := e.err
	for b, ok := a.(wrappedError); ok; {
		msg += "- " + b.context + ",\n"
		a = b.err
		b, ok = a.(wrappedError)
	}
	return msg + a.Error()
}

// ShowError shows an error and its context to the user. In debug mode, it only
// prints to the standard error output, otherwise it also brings a dialog box.
func ShowError(context string, err error) {
	e := Error(context, err)
	internal.Log("ERROR:\n%s", e)
	if !internal.Config.Debug {
		err2 := internal.ErrorDialog("ERROR:\n%s", e)
		if err2 != nil {
			internal.Log("ERROR opwning dialog:\n%s", err2)
		}
	}
}

// Log logs a formated message.
func Log(format string, v ...interface{}) {
	internal.Log(format, v...)
}

//------------------------------------------------------------------------------
