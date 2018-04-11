// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

////////////////////////////////////////////////////////////////////////////////

// makeError returns nil if err is nil, or a wrapped error otherwise.
func makeError(context string, err error) error {
	if err == nil {
		return nil
	}
	return wrappedError{context, err}
}

type wrappedError struct {
	Context string
	Err     error
}

func (e wrappedError) Error() string {
	return e.Context + ": " + e.Err.Error()
}

////////////////////////////////////////////////////////////////////////////////

// Err returns the first OpenGL error since the previous call to Err().
func Err() error {
	err := stickyErr
	stickyErr = nil
	return err
}

func setErr(context string, err error) {
	if stickyErr == nil {
		stickyErr = makeError(context, err)
	}
	debug.Printf("gfx error: %s", makeError(context, err))
}

var stickyErr error

////////////////////////////////////////////////////////////////////////////////
