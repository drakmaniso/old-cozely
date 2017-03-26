// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

import (
	"log"
	"os"
)

//------------------------------------------------------------------------------

var logger = log.New(os.Stderr, "glam: ", log.Ltime)

// Log logs a formated message.
func Log(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

// DebugLog logs a formated message if Debug mode is enabled.
func DebugLog(format string, v ...interface{}) {
	if Config.Debug {
		logger.Printf(format, v...)
	}
}

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
	msg := e.context + ":\n\t"
	a := e.err
	for b, ok := a.(wrappedError); ok; {
		msg += b.context + ":\n\t"
		a = b.err
		b, ok = a.(wrappedError)
	}
	return msg + a.Error()
}

//------------------------------------------------------------------------------
