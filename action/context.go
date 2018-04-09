// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

type Context uint32

var contexts struct {
	Name  []string
	Bools [][]Bool
}

func NewContext(name string, actions ...Action) Context {
	l := len(contexts.Name)
	if l >= maxID {
		//TODO: set error
		return Context(maxID)
	}

	contexts.Name = append(contexts.Name, name)

	return Context(l)
}

func (c Context) Activate() {

}
