// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"strings"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

var vertexShader = `
#version 450 core

out gl_PerVertex {
	vec4 gl_Position;
};

void main(void)
{
	const vec4 triangle[3] = vec4[3](
		vec4(0, 0.65, 0.5, 1),
		vec4(-0.65, -0.475, 0.5, 1),
		vec4(0.65, -0.475, 0.5, 1)
	);
	gl_Position = triangle[gl_VertexID];
}
`

var fragmentShader = `
#version 450 core

out vec4 color;

void main(void)
{
	color = vec4(0.3, 0.1, 0.6, 1.0);
}
`

//------------------------------------------------------------------------------

func main() {
	err := setup()
	if err != nil {
		glam.ErrorDialog(err)
		return
	}

	// Run the main Loop
	glam.Loop = looper{}
	err = glam.Run()
	if err != nil {
		glam.ErrorDialog(err)
	}
}

//------------------------------------------------------------------------------

// OpenGL pipeline object
var (
	pipeline gfx.Pipeline
)

//------------------------------------------------------------------------------

func setup() error {
	// Setup the pipeline
	pipeline = gfx.NewPipeline(
		gfx.VertexShader(strings.NewReader(vertexShader)),
		gfx.FragmentShader(strings.NewReader(fragmentShader)),
	)

	return gfx.Err()
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update() {
}

func (l looper) Draw() {
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.9, 0.9, 1.0})
	pipeline.Bind()
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
