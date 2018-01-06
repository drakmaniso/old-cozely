// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package ecs

//------------------------------------------------------------------------------

type Entity uint32

const (
	None  Entity = 0
	First Entity = 1
)

const Size = 1000

//------------------------------------------------------------------------------

func New(m Mask) Entity {
	for e := First; e < Size; e++ {
		if components[e] == 0 {
			components[e] = m
			return e
		}
	}
	return None
}

//------------------------------------------------------------------------------

func Last() Entity {
	return Size - 1
}

//------------------------------------------------------------------------------

func (e Entity) Has(m Mask) bool {
	return components[e]&m != 0
}

func (e Entity) Add(m Mask) {
	components[e] |= m
}

//------------------------------------------------------------------------------

func (e Entity) Delete() {
	components[e] = 0
}

//------------------------------------------------------------------------------
