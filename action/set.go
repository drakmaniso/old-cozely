// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

type Set uint32

var set struct {
	names           []string
	booleans        [][]Bool
	booleanBindings [][]string
}

func NewSet(name string, actions map[Action][]string) Set {
	l := len(set.names)
	if l >= maxID {
		//TODO: set error
		return Set(maxID)
	}

	set.names = append(set.names, name)

	return Set(l)
}

var MenuUp = NewBool("Menu Up")

var MenuActions = NewSet("Menu", map[Action][]string{
	MenuUp: {"Up", "Dpad Up"},
})
