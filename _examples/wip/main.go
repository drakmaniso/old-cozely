// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"log"
	"strings"

	"github.com/drakmaniso/glam"
	. "github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/gfx"
)

//------------------------------------------------------------------------------

var pipeline gfx.Pipeline
var vbo gfx.Buffer

//------------------------------------------------------------------------------

type vertex struct {
	position Vec4 `layout:"0"`
	// color    IVec3   `layout:"1"`
	// alpha    float32 `layout:"2"`
}

func main() {
	g := &game{}
	glam.Handler = g

	vs := strings.NewReader(`
		#version 420 core
		layout(location = 0) in vec4 pos;
		void main(void)
		{
			const float Pi = 3.14;
			const float r = 0.75;
			const vec4 v[3] = vec4[3](
				vec4(r*sin(0),       r*cos(0),       0.5, 1.0),
				vec4(r*sin(-Pi*2/3), r*cos(-Pi*2/3), 0.5, 1.0),
				vec4(r*sin(-Pi*4/3), r*cos(-Pi*4/3), 0.5, 1.0)
			);
			//gl_Position = v[gl_VertexID];
			gl_Position = pos;
		}	
	`)

	fs := strings.NewReader(`
		#version 420 core
		out vec4 color;
		void main(void)
		{
			color = vec4(0.84, 0.00, 0.44, 1.0);
		}	
	`)

	_ = pipeline.CompileShaders(vs, fs)

	if err := pipeline.DefineAttributes(0, vertex{}); err != nil {
		log.Print("ERROR: ", err)
	}

	v := []vertex{
		{Vec4{-0.25, -0.25, 0.5, 1.0}},
		{Vec4{0.25, -0.25, 0.5, 1.0}},
		{Vec4{0.25, 0.25, 0.5, 1.0}},
	}
	if err := vbo.CreateFrom(v); err != nil {
		log.Print(err)
	}

	if err := glam.Run(); err != nil {
		log.Print(err)
	}
}

//------------------------------------------------------------------------------

type game struct{}

func (g *game) Update() {
}

func (g *game) Draw() {
	pipeline.Bind()
	pipeline.BindAttributes(0, &vbo)
	gfx.Draw(gfx.Triangles, 0, 3)
}

//------------------------------------------------------------------------------
