// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"fmt"
	"reflect"

	"strconv"

	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/geom"
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

static inline void VertexBuffer(GLuint vao, GLuint binding, GLuint buffer, GLintptr offset, GLsizei stride) {
	glVertexArrayVertexBuffer(vao, binding, buffer, offset, stride);
}
*/
import "C"

//------------------------------------------------------------------------------

// VertexFormat prepares everything the pipeline needs to use a
// vertex buffer of a specific format, and assign a binding index to it.
//
// The format must be a struct with layout tags.
func (p *Pipeline) VertexFormat(binding uint32, format interface{}) error {
	f := reflect.TypeOf(format)
	if f.Kind() != reflect.Struct {
		return fmt.Errorf("attributes binding format must be a struct, not a %s", f.Kind())
	}

	p.attribStride[binding] = f.Size()

	for i := 0; i < f.NumField(); i++ {
		a := f.Field(i)
		al := a.Tag.Get("layout")
		if al == "" {
			continue
		}
		ali, err := strconv.Atoi(al)
		if err != nil {
			return fmt.Errorf("invalid layout for attributes binding: %q", al)
		}
		//TODO: check that ali is in range
		at := a.Type
		var as int32
		ao := a.Offset
		var ate C.GLenum
		switch {
		// Float32
		case at.ConvertibleTo(float32Type):
			as = 1
			ate = C.GL_FLOAT
		case at.ConvertibleTo(vec4Type), at.ConvertibleTo(rgbaType):
			as = 4
			ate = C.GL_FLOAT
		case at.ConvertibleTo(vec3Type), at.ConvertibleTo(rgbType):
			as = 3
			ate = C.GL_FLOAT
		case at.ConvertibleTo(vec2Type):
			as = 2
			ate = C.GL_FLOAT
		// Int32
		case at.ConvertibleTo(int32Type):
			as = 1
			ate = C.GL_INT
		case at.ConvertibleTo(ivec4Type):
			as = 4
			ate = C.GL_INT
		case at.ConvertibleTo(ivec3Type):
			as = 3
			ate = C.GL_INT
		case at.ConvertibleTo(ivec2Type):
			as = 2
			ate = C.GL_INT
		}

		C.VertexAttribute(
			p.vao,
			C.GLuint(ali),
			C.GLuint(0), //TODO
			C.GLint(as),
			ate,
			C.GLboolean(0), //TODO
			C.GLuint(ao),
		)
	}
	return nil
}

var (
	float32Type = reflect.TypeOf(float32(0))
	vec4Type    = reflect.TypeOf(geom.Vec4{})
	vec3Type    = reflect.TypeOf(geom.Vec3{})
	vec2Type    = reflect.TypeOf(geom.Vec2{})
	int32Type   = reflect.TypeOf(int32(0))
	ivec4Type   = reflect.TypeOf(geom.IVec4{})
	ivec3Type   = reflect.TypeOf(geom.IVec3{})
	ivec2Type   = reflect.TypeOf(geom.IVec2{})
	rgbType     = reflect.TypeOf(color.RGB{})
	rgbaType    = reflect.TypeOf(color.RGBA{})
)

//------------------------------------------------------------------------------

// VertexBuffer binds a buffer to a vertex buffer binding index.
//
// The buffer should use the same struct type than the one used in the
// corresponding call to VertexBufferFormat.
func (p *Pipeline) VertexBuffer(binding uint32, b Buffer, offset uintptr) {
	C.VertexBuffer(p.vao, C.GLuint(binding), b.buffer, C.GLintptr(offset), C.GLsizei(p.attribStride[binding]))
}

//------------------------------------------------------------------------------
