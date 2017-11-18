// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"errors"

	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/picture"
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
	return nil
}

func (loop) Update() error {
	return nil
}

func (loop) Draw(_, _ float64) error {
	p, ok := picture.Named("logo")
	if !ok {
		return errors.New("picture not found")
	}
	// p.Paint(10, 10)
	_ = p
	return nil
}

//------------------------------------------------------------------------------
