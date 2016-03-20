package gfx

import (
	"io"

	"github.com/drakmaniso/glam/internal"
)

type Pipeline struct {
	program internal.GLuint
}

func NewPipeline(
	vertexShader io.Reader,
	fragmentShader io.Reader,
) (*Pipeline, error) {
	var p Pipeline
	var err error
	p.program, err = internal.CompileShaders(vertexShader, fragmentShader)
	return &p, err
}
