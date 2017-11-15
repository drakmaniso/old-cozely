// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package core

//------------------------------------------------------------------------------

/*
#include <stdlib.h>
#include "sdl.h"
*/
import "C"

//------------------------------------------------------------------------------

import (
	"errors"
	"fmt"
	"unsafe"
)

//------------------------------------------------------------------------------

// ErrorDialog displays a dialog box.
func ErrorDialog(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)

	t := Config.Title + " - Error"
	ct := C.CString(t)
	defer C.free(unsafe.Pointer(ct))

	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	errcode := C.SDL_ShowSimpleMessageBox(
		C.SDL_MESSAGEBOX_ERROR,
		ct,
		cmsg,
		Window.window,
	)
	if errcode != 0 {
		return GetSDLError()
	}
	return nil
}

//------------------------------------------------------------------------------

// GetSDLError returns nil or the current SDL Error wrapped in a Go error.
func GetSDLError() error {
	if s := C.SDL_GetError(); s != nil {
		return errors.New(C.GoString(s))
	}
	return nil
}

//------------------------------------------------------------------------------
