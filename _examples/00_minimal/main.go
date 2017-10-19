// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol"
	"github.com/drakmaniso/carol/mtx"
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
	carol.DefaultHandlers
}

func (loop) Update() {
	mtx.Locate(1, 1)
	mtx.Print("hello, world\n")
}

func (loop) Draw(_, _ float64) {
}

//------------------------------------------------------------------------------
