// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

type Bool uint32

var (
	bools []bool
	boolNames []string
	boolDefaults [][]string
)

func NewBool(name string, defaults ...string) Bool {
	l := len(bools)
	if l >= maxID {
		//TODO: set error
		return Bool(0)
	}
	
	bools = append(bools, false)
	boolNames = append(boolNames, name)
	boolDefaults = append(boolDefaults, defaults)

	return Bool(l)
}

func (b Bool) Active() bool {
	return bools[b]
}
