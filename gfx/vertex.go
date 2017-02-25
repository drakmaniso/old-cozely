// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
#include "glad.h"

void VertexAttribute(
	GLuint vao,
	GLuint index,
	GLuint binding,
	GLint size,
	GLenum type,
	GLboolean normalized,
	GLuint relativeOffset
) {
	glVertexArrayAttribFormat(vao, index, size, type, normalized, relativeOffset);
	glVertexArrayAttribBinding(vao, index, binding);
	glEnableVertexArrayAttrib(vao, index);
}

static inline void VertexArrayBindingDivisor(GLuint vao, GLuint i, GLuint d) {
  glVertexArrayBindingDivisor(vao, i, d);
}
*/
import "C"

//------------------------------------------------------------------------------

// VertexFormat prepares everything the pipeline needs to use a
// vertex buffer of a specific format, and assign a binding index to it.
//
// The format must be a struct with layout tags.
func VertexFormat(binding uint32, format interface{}) PipelineOption {
	return func(p *Pipeline) {
		p.setVertexFormat(binding, format)
	}
}

func (p *Pipeline) setVertexFormat(binding uint32, format interface{}) {
	t := reflect.TypeOf(format)
	if t.Kind() != reflect.Struct {
		setErr(fmt.Errorf("attributes binding format must be fld struct, not fld %s", t.Kind()))
		return
	}

	for i := 0; i < t.NumField(); i++ {
		fld := t.Field(i)

		// Layout tag
		layStr := fld.Tag.Get("layout")
		if layStr == "" {
			continue
		}
		lay, err := strconv.Atoi(layStr)
		if err != nil {
			setErr(fmt.Errorf("invalid layout for attributes binding: %q", layStr))
			return
		}
		//TODO: check that lay is in range

		//TODO: check that lay is in range
		typ := fld.Type
		var siz int32
		off := fld.Offset
		var typEnum C.GLenum
		switch {
		// Float32
		case typ.ConvertibleTo(float32Type):
			siz = 1
			typEnum = C.GL_FLOAT
		case typ.ConvertibleTo(vec4Type), typ.ConvertibleTo(rgbaType):
			siz = 4
			typEnum = C.GL_FLOAT
		case typ.ConvertibleTo(vec3Type), typ.ConvertibleTo(rgbType):
			siz = 3
			typEnum = C.GL_FLOAT
		case typ.ConvertibleTo(vec2Type):
			siz = 2
			typEnum = C.GL_FLOAT
		// Int32
		case typ.ConvertibleTo(int32Type):
			siz = 1
			typEnum = C.GL_INT
		case typ.ConvertibleTo(ivec4Type):
			siz = 4
			typEnum = C.GL_INT
		case typ.ConvertibleTo(ivec3Type):
			siz = 3
			typEnum = C.GL_INT
		case typ.ConvertibleTo(ivec2Type):
			siz = 2
			typEnum = C.GL_INT
		}

		C.VertexAttribute(
			p.vao,
			C.GLuint(lay),
			C.GLuint(binding),
			C.GLint(siz),
			typEnum,
			C.GLboolean(0), //TODO
			C.GLuint(off),
		)

		// Divisor Tag
		divStr := fld.Tag.Get("divisor")
		if divStr != "" {
			var div = 0
			div, err = strconv.Atoi(divStr)
			if err != nil {
				setErr(fmt.Errorf("invalid divisor for attributes binding: %q", divStr))
				return
			}
			C.VertexArrayBindingDivisor(p.vao, C.GLuint(binding), C.GLuint(div))
		}
	}
	return
}

var (
	float32Type = reflect.TypeOf(float32(0))
	vec4Type    = reflect.TypeOf(struct{ X, Y, Z, W float32 }{})
	vec3Type    = reflect.TypeOf(struct{ X, Y, Z float32 }{})
	vec2Type    = reflect.TypeOf(struct{ X, Y float32 }{})
	int32Type   = reflect.TypeOf(int32(0))
	ivec4Type   = reflect.TypeOf(struct{ X, Y, Z, W int32 }{})
	ivec3Type   = reflect.TypeOf(struct{ X, Y, Z int32 }{})
	ivec2Type   = reflect.TypeOf(struct{ X, Y int32 }{})
	rgbType     = reflect.TypeOf(struct{ R, G, B float32 }{})
	rgbaType    = reflect.TypeOf(struct{ R, G, B, A float32 }{})
)

//------------------------------------------------------------------------------
