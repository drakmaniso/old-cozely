package engine_test

//------------------------------------------------------------------------------

import (
	"log"
	"testing"

	"github.com/drakmaniso/glam/engine"
)

//------------------------------------------------------------------------------

func TestEngine_Run(t *testing.T) {
	err := engine.Run()
	if err != nil {
		log.Panic(err)
	}
}

//------------------------------------------------------------------------------
