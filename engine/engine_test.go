// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine_test

//------------------------------------------------------------------------------

import (
	"log"
	"os"
	"testing"

	"github.com/drakmaniso/glam/engine"
)

//------------------------------------------------------------------------------

func TestMain(m *testing.M) {
	engine.HandleQuit(func() { log.Print("*** Bye! ***") })
	engine.HandleKeyDown(func() { log.Print("---Key Down---") })
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
