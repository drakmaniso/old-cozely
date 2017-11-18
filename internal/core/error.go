// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package core

//------------------------------------------------------------------------------

type logger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type nolog struct{}

func (nolog) Print(v ...interface{})                 {}
func (nolog) Println(v ...interface{})               {}
func (nolog) Printf(format string, v ...interface{}) {}

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
	msg := "- " + e.context + ":\n"
	a := e.err
	for b, ok := a.(wrappedError); ok; {
		msg += "- " + b.context + ":\n"
		a = b.err
		b, ok = a.(wrappedError)
	}
	return msg + a.Error()
}

//------------------------------------------------------------------------------
