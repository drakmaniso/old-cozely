// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/resource"
	"github.com/cozely/cozely/window"
	"github.com/cozely/cozely/x/gl"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop1 struct {
	// OpenGL Object
	pipeline *gl.Pipeline
}

// Initialization //////////////////////////////////////////////////////////////

func Example_firstTriangle() {
	defer cozely.Recover()

	err := resource.Path("testdata/")
	if err != nil {
		panic(err)
	}
	window.Events.Resize = func() {
		s := window.Size()
		gl.Viewport(0, 0, int32(s.X), int32(s.Y))
	}
	l := loop1{}
	err = cozely.Run(&l)
	if err != nil {
		panic(err)
	}
	//Output:
}

func (l *loop1) Enter() {
	// Create and configure the pipeline
	l.pipeline = gl.NewPipeline(
		gl.Shader("shader01.vert"),
		gl.Shader("shader01.frag"),
		gl.Topology(gl.Triangles),
	)
}

func (loop1) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop1) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}
}

func (loop1) Update() {
}

func (l *loop1) Render() {
	l.pipeline.Bind()
	gl.ClearColorBuffer(color.LRGBA{0.9, 0.9, 0.9, 1.0})

	gl.Draw(0, 3)
	l.pipeline.Unbind()
}
