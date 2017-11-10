// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

import (
	"fmt"
	"unsafe"
)

/*
#include <stdlib.h>
#include "sdl.h"
*/
import "C"

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
