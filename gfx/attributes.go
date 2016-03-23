// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

import (
	"fmt"
	"log"
	"reflect"

	"strconv"

	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

func (p *Pipeline) CreateAttributesBinding(binding uint32, format interface{}) error {
	f := reflect.TypeOf(format)
	if f.Kind() != reflect.Struct {
		return fmt.Errorf("attributes binding format must be a struct, not a %s", f.Kind())
	}
	p.attribStride[binding] = f.Size()
	log.Print(p.attribStride[binding])
	for i := 0; i < f.NumField(); i++ {
		a := f.Field(i)
		log.Print("*** Attribute: ", i)
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
		as := at.Size()
		ao := a.Offset
		ate := internal.GlByteEnum
		switch {
		case at.ConvertibleTo(float32Type),
			at.ConvertibleTo(vec4Type),
			at.ConvertibleTo(vec3Type),
			at.ConvertibleTo(vec2Type):
			ate = internal.GlFloatEnum
		case at.ConvertibleTo(int32Type),
			at.ConvertibleTo(ivec4Type),
			at.ConvertibleTo(ivec3Type),
			at.ConvertibleTo(ivec2Type):
			ate = internal.GlIntEnum
		}

		log.Print("        Index: ", ali)
		log.Print("         Size: ", as)
		log.Print("         Type: ", ate)
		log.Print("       Offset: ", ao)
		p.internal.CreateAttributeBinding(
			uint32(ali),
			uint32(0), //TODO
			int32(as),
			uint32(ate),
			byte(0), //TODO
			uint32(ao),
		)
	}
	return nil
}

//------------------------------------------------------------------------------

var (
	float32Type = reflect.TypeOf(float32(0))
	vec4Type    = reflect.TypeOf(geom.Vec4{})
	vec3Type    = reflect.TypeOf(geom.Vec3{})
	vec2Type    = reflect.TypeOf(geom.Vec2{})
	int32Type   = reflect.TypeOf(int32(0))
	ivec4Type   = reflect.TypeOf(geom.IVec4{})
	ivec3Type   = reflect.TypeOf(geom.IVec3{})
	ivec2Type   = reflect.TypeOf(geom.IVec2{})
)

//------------------------------------------------------------------------------
