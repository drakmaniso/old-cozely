package ciziel_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/cozely/cozely/formats/ciziel"
)

func TestParse1(t *testing.T) {
	f, err := os.Open("testdata/inputbindings.czl")
	if err != nil {
		t.Error(err)
		return
	}
	d := ciziel.Parse(f)
	fmt.Println(d.String())
}
