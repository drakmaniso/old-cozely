// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol"
)

//------------------------------------------------------------------------------

func main() {
	err := carol.Setup()
	if err != nil {
		carol.ShowError("setting up carol", err)
		return
	}

	carol.Loop(loop{})

	err = carol.Run()
	if err != nil {
		carol.ShowError("running", err)
		return
	}
}

//------------------------------------------------------------------------------

type loop struct {
	carol.Handlers
}

func (loop) Update() {
}

func (loop) Draw(_, _ float64) {
}

//------------------------------------------------------------------------------
