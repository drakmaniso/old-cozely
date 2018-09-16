package ciziel_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cozely/cozely/formats/ciziel"
)

func TestParse1(t *testing.T) {
	f, err := os.Open("testdata/bindings.czl")
	if err != nil {
		t.Error(err)
		return
	}
	d := ciziel.Parse(f)
	fmt.Print(d.String())
}
