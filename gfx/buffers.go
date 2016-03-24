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

func (b *Buffer) CreateFrom(data interface{}) error {
	//TODO: handle pointers too
	if reflect.TypeOf(data).Kind() != reflect.Slice {
		return fmt.Errorf("buffer data must be a slice, not a %s", reflect.TypeOf(data).Kind())
	}
	v := reflect.ValueOf(data)
	l := v.Len()
	if l == 0 {
		return fmt.Errorf("buffer data cannot be an empty slice")
	}
	p := v.Pointer()
	s := v.Index(0).Type().Size()
	b.internal.CreateFrom(uintptr(l)*s, p)
	return nil
}

//------------------------------------------------------------------------------
