// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import "github.com/drakmaniso/glam"

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
}

func (l looper) Draw() {
}

//------------------------------------------------------------------------------
