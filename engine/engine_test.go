package engine_test

//------------------------------------------------------------------------------

import (
	"os"
	"testing"

	"github.com/drakmaniso/glam/engine"
)

//------------------------------------------------------------------------------

func TestMain(m *testing.M) {
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
