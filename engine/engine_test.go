// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine_test

//------------------------------------------------------------------------------

import (
	"log"
	"os"
	"testing"

	"github.com/drakmaniso/glam/engine"
	"github.com/drakmaniso/glam/key"
)

//------------------------------------------------------------------------------

func TestMain(m *testing.M) {
	engine.HandleQuit(func() { log.Print("*** Bye! ***") })
	engine.HandleKeyDown(
		func(l key.Label, p key.Position, time uint32) {
			log.Print("Key Down: ", l, p, time)
		},
	)
	engine.HandleKeyUp(
		func(l key.Label, p key.Position, time uint32) {
			log.Print("Key Up: ", l, p, time)
		},
	)
	err = engine.Run()
	os.Exit(m.Run())
}

//------------------------------------------------------------------------------

var err error

func TestEngine_Run(t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

//------------------------------------------------------------------------------
