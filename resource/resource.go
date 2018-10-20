package resource

import (
	"io"
	"os"
)

////////////////////////////////////////////////////////////////////////////////

func Exist(name string) bool {
	for _, s := range sources {
		if s.exist(name) {
			return true
		}
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////

func Open(name string) (io.ReadCloser, error) {
	for _, s := range sources {
		f, err := s.open(name)
		if !os.IsNotExist(err) {
			return f, err
		}
	}
	return nil, os.ErrNotExist
}

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
