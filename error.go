// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package cozely

import (
	"runtime/debug"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

var stopErr error

// Stop request the game loop to stop.
func Stop(err error) {
	if stopErr == nil {
		stopErr = err
	}
	internal.QuitRequested = true
}

////////////////////////////////////////////////////////////////////////////////

// Error returns true if there is an unchecked error in one of Cozely's
// packages.
func Error() bool {
	return internal.GLErr() != nil ||
		internal.InputErr() != nil ||
		internal.PixelErr() != nil ||
		internal.PolyErr() != nil
}

// Recover is a convenience function.  When called with defer at the start of
// the main function, it will intercept any panic, log the stack trace, and
// finally log any unchecked errors in Cozely's packages.
//
// This function exists because many errors in the framework, in addition to
// setting the package's sticky error, also returns an invalid ID; if the error
// is unchecked and the ID used, this will trigger an "index out of range"
// panic.
func Recover() {
	r := recover()
	if r != nil {
		internal.Log.Printf("*** PANIC stack trace *********************************\n")
		internal.Log.Println(string(debug.Stack()))
		internal.Log.Printf("*** end of PANIC stack trace **************************\n%s", r)

		var err error
		err = internal.InputErr()
		if err != nil {
			internal.Log.Printf("*** panic: INPUT unchecked ERROR ***\n%s", err)
		}
		err = internal.PixelErr()
		if err != nil {
			internal.Log.Printf("*** panic: PIXEL unchecked ERROR ***\n%s", err)
		}
		err = internal.GLErr()
		if err != nil {
			internal.Log.Printf("*** panic: GL unchecked ERROR ***\n%s", err)
		}
		err = internal.PolyErr()
		if err != nil {
			internal.Log.Printf("*** panic: POLY unchecked ERROR ***\n%s", err)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

// Wrap returns nil if err is nil, or a wrapped error otherwise.
func Wrap(context string, err error) error {
	if err == nil {
		return nil
	}
	return internal.WrappedError{context, err}
}

// ShowError shows an error to the user. In debug mode, it only prints to the
// standard error output, otherwise it also brings a dialog box.
//TODO:
// func ShowError(e error) {
// 	internal.Log.Printf("*** ERROR ***\n%s", e)
// 	if !internal.Config.Debug {
// 		err2 := internal.ErrorDialog("*** ERROR ***\n%s", e)
// 		if err2 != nil {
// 			internal.Log.Printf("*** ERROR opening dialog ***\n%s", err2)
// 		}
// 	}
// }

// Log logs a formated message.
//TODO:
// func Log(format string, v ...interface{}) {
// 	internal.Log.Printf(format, v...)
// }

////////////////////////////////////////////////////////////////////////////////
