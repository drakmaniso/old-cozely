// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import "github.com/drakmaniso/glam"
import "github.com/drakmaniso/glam/mtx"

//------------------------------------------------------------------------------

func main() {
	glam.Loop = looper{}
	err := glam.Run()
	if err != nil {
		glam.ErrorDialog(err)
	}
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update() {
	mtx.Print(2, 2, "hello, world\n")
}

func (l looper) Draw() {
}

//------------------------------------------------------------------------------
