// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

import (
	"fmt"
	"reflect"

	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

type Buffer struct {
	internal internal.Buffer
}

//------------------------------------------------------------------------------

func (b *Buffer) CreateFrom(data interface{}, f bufferFlags) error {
	s, p, err := sizeAndPointerOf(data)
	if err != nil {
		return err
	}
	b.internal.CreateFrom(s, p, uint32(f))
	return nil
}

func (b *Buffer) UpdateWith(data interface{}, offset uintptr) error {
	s, p, err := sizeAndPointerOf(data)
	if err != nil {
		return err
	}
	b.internal.UpdateWith(offset, s, p)
	return nil
}

func sizeAndPointerOf(data interface{}) (size uintptr, ptr uintptr, err error) {
	var s uintptr
	var p uintptr
	v := reflect.ValueOf(data)
	k := v.Type().Kind()
	switch k {
	case reflect.Slice:
		l := v.Len()
		if l == 0 {
			return 0, 0, fmt.Errorf("buffer data cannot be an empty slice")
		}
		p = v.Pointer()
		s = uintptr(l) * v.Index(0).Type().Size()
	case reflect.Ptr:
		p = v.Pointer()
		s = v.Elem().Type().Size()
	default:
		return 0, 0, fmt.Errorf("buffer data must be a slice or a pointer, not a %s", reflect.TypeOf(data).Kind())
	}
	return s, p, nil
}

//------------------------------------------------------------------------------

type bufferFlags uint32

// Flags for buffer creation.
const (
	MapRead        bufferFlags = 0x0001
	MapWrite       bufferFlags = 0x0002
	MapPersistent  bufferFlags = 0x0040
	MapCoherent    bufferFlags = 0x0080
	DynamicStorage bufferFlags = 0x0100
)

//------------------------------------------------------------------------------
