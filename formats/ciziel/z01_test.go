package ciziel_test

import (
	"github.com/cozely/cozely/formats/ciziel"
	"os"
	"testing"
)

func TestParse1(t *testing.T) {
	f, err := os.Open("testdata/init.czl")
	if err != nil {
		t.Error(err)
		return
	}
	_ = ciziel.Parse(f)
}
