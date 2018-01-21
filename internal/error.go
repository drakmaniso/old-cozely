// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

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
	return WrappedError{context, err}
}

type WrappedError struct {
	Context string
	Err     error
}

func (e WrappedError) Error() string {
	return e.Context + ": " + e.Err.Error()
}

//------------------------------------------------------------------------------
