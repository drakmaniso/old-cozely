// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl

import (
	"fmt"
	"image"
	"reflect"
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////

/*
#include "glad.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// A Access specifies...
type Access C.GLenum

// Used in ...
const (
	ReadOnly  Access = C.GL_READ_ONLY
	WriteOnly Access = C.GL_WRITE_ONLY
	ReadWrite Access = C.GL_READ_WRITE
)

////////////////////////////////////////////////////////////////////////////////

// A TextureFormat specifies the format used to store textures in memory.
type TextureFormat C.GLenum

// Used in NewTexture1D, NewTexture2D...
const (
	R8    TextureFormat = C.GL_R8
	R16   TextureFormat = C.GL_R16
	R16F  TextureFormat = C.GL_R16F
	R32F  TextureFormat = C.GL_R32F
	R8I   TextureFormat = C.GL_R8I
	R16I  TextureFormat = C.GL_R16I
	R32I  TextureFormat = C.GL_R32I
	R8UI  TextureFormat = C.GL_R8UI
	R16UI TextureFormat = C.GL_R16UI
	R32UI TextureFormat = C.GL_R32UI

	RG8    TextureFormat = C.GL_RG8
	RG16   TextureFormat = C.GL_RG16
	RG16F  TextureFormat = C.GL_RG16F
	RG32F  TextureFormat = C.GL_RG32F
	RG8I   TextureFormat = C.GL_RG8I
	RG16I  TextureFormat = C.GL_RG16I
	RG32I  TextureFormat = C.GL_RG32I
	RG8UI  TextureFormat = C.GL_RG8UI
	RG16UI TextureFormat = C.GL_RG16UI
	RG32UI TextureFormat = C.GL_RG32UI

	RGB32F  TextureFormat = C.GL_RGB32F
	RGB32I  TextureFormat = C.GL_RGB32I
	RGB32UI TextureFormat = C.GL_RGB32UI

	RGB8 TextureFormat = C.GL_RGB8

	RGBA8    TextureFormat = C.GL_RGBA8
	RGBA16   TextureFormat = C.GL_RGBA16
	RGBA16F  TextureFormat = C.GL_RGBA16F
	RGBA32F  TextureFormat = C.GL_RGBA32F
	RGBA8I   TextureFormat = C.GL_RGBA8I
	RGBA16I  TextureFormat = C.GL_RGBA16I
	RGBA32I  TextureFormat = C.GL_RGBA32I
	RGBA8UI  TextureFormat = C.GL_RGBA8UI
	RGBA16UI TextureFormat = C.GL_RGBA16UI
	RGBA32UI TextureFormat = C.GL_RGBA32UI

	SRGBA8 TextureFormat = C.GL_SRGB8_ALPHA8
	SRGB8  TextureFormat = C.GL_SRGB8

	Depth16  TextureFormat = C.GL_DEPTH_COMPONENT16
	Depth24  TextureFormat = C.GL_DEPTH_COMPONENT24
	Depth32F TextureFormat = C.GL_DEPTH_COMPONENT32F
)

////////////////////////////////////////////////////////////////////////////////

func pointerFormatAndTypeOf(img image.Image) (p unsafe.Pointer, pformat C.GLenum, ptype C.GLenum) {
	switch img := img.(type) {
	//TODO: other formats
	case *image.RGBA:
		p = unsafe.Pointer(&img.Pix[0])
		pformat = C.GL_RGBA
		ptype = C.GL_UNSIGNED_BYTE
	case *image.NRGBA:
		p = unsafe.Pointer(&img.Pix[0])
		pformat = C.GL_RGBA
		ptype = C.GL_UNSIGNED_BYTE
	case *image.Paletted:
		p = unsafe.Pointer(&img.Pix[0])
		pformat = C.GL_RED_INTEGER
		ptype = C.GL_UNSIGNED_BYTE
	}
	return p, pformat, ptype
}

func pointerFormatAndTypeOfData(data interface{}) (p unsafe.Pointer, pformat C.GLenum, ptype C.GLenum) {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Ptr:
		p = unsafe.Pointer(v.Pointer())
		v = v.Elem()
		if v.Kind() != reflect.Array {
			fmt.Println(v.Kind())
			panic("TODO") //TODO
		}
		vv := v.Index(0)
		switch vv.Kind() {
		case reflect.Int8:
			pformat = C.GL_RED_INTEGER
			ptype = C.GL_BYTE
		case reflect.Uint8:
			pformat = C.GL_RED_INTEGER
			ptype = C.GL_UNSIGNED_BYTE
		case reflect.Int16:
			pformat = C.GL_RED_INTEGER
			ptype = C.GL_SHORT
		case reflect.Uint16:
			pformat = C.GL_RED_INTEGER
			ptype = C.GL_UNSIGNED_SHORT
		case reflect.Int32:
			pformat = C.GL_RED_INTEGER
			ptype = C.GL_INT
		case reflect.Uint32:
			pformat = C.GL_RED_INTEGER
			ptype = C.GL_UNSIGNED_INT
		}
	}
	return p, pformat, ptype
}

////////////////////////////////////////////////////////////////////////////////
