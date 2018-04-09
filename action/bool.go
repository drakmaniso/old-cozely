// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

type Bool uint32

var boolean struct {
	names  []string
	values []bool
}

func NewBool(name string) Bool {
	l := len(boolean.names)
	if l >= maxID {
		//TODO: set error
		return Bool(maxID)
	}

	boolean.names = append(boolean.names, name)
	boolean.values = append(boolean.values, false)

	return Bool(l)
}

func (b Bool) Active() bool {
	return boolean.values[b]
}

func (b Bool) Origins() []Origin {
	return nil
}
