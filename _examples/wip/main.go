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
	"github.com/drakmaniso/glam/math"
)

//------------------------------------------------------------------------------

var pipeline gfx.Pipeline
var vbo gfx.Buffer

type vertex struct {
	position Vec3 `layout:"0"`
	color    Vec3 `layout:"1"`
}

//------------------------------------------------------------------------------

func main() {
	g := &game{}
	glam.Handler = g

	vs := strings.NewReader(`
		#version 420 core
		layout(location = 0) in vec3 pos;
		layout(location = 1) in vec3 col;
		layout(location = 0) out vec3 fs_col;
		void main(void)
		{
			gl_Position = vec4(pos, 1);
			fs_col = col;
		}	
	`)

	fs := strings.NewReader(`
		#version 420 core
		layout(location = 0) in vec3 vs_col;
		out vec4 color;
		void main(void)
		{
			color = vec4(vs_col, 1);
		}	
	`)

	if err := pipeline.CompileShaders(vs, fs); err != nil {
		log.Fatal(err)
	}

	if err := pipeline.DefineAttributes(0, vertex{}); err != nil {
		log.Fatal(err)
	}

	r := float32(0.75)
	v := []vertex{
		{
			Vec3{r * math.Sin(0), r * math.Cos(0), 0.5},
			Vec3{0.3, 0, 0.8},
		},
		{
			Vec3{r * math.Sin(-math.Pi*2/3), r * math.Cos(-math.Pi*2/3), 0.5},
			Vec3{0.8, 0.3, 0},
		},
		{
			Vec3{r * math.Sin(-math.Pi*4/3), r * math.Cos(-math.Pi*4/3), 0.5},
			Vec3{0, 0.6, 0.2},
		},
	}
	if err := vbo.CreateFrom(v); err != nil {
		log.Fatal(err)
	}

	if err := glam.Run(); err != nil {
		log.Fatal(err)
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
