// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"errors"

	"github.com/drakmaniso/carol"
)

//------------------------------------------------------------------------------

func main() {
	err := carol.Run(loop{})
	if err != nil {
		carol.ShowError("in game loop", err)
		return
	}
}

//------------------------------------------------------------------------------

type loop struct {
	carol.Handlers
}

func (loop) Setup() error {
	return errors.New("plop")
}

func (loop) Update() error {
	return nil
}

func (loop) Draw(_, _ float64) error {
	return nil
}

//------------------------------------------------------------------------------
