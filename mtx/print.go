// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mtx

import (
	micro "github.com/drakmaniso/glam/internal/microtext"
	"strconv"
)

//------------------------------------------------------------------------------

func Size() (x, y int) {
	return micro.Size()
}

func Clamp(x, y int) (int, int) {
	sx, sy := micro.Size()

	if x < 0 {
		x += sx
		if x < 0 {
			x = 0
		}
	}
	if x >= sx {
		x = sx - 1
	}

	if y < 0 {
		y += sy
		if y < 0 {
			y = 0
		}
	}
	if y >= sy {
		y = sy - 1
	}

	return x, y
}

//------------------------------------------------------------------------------

var colour = byte(0)

func ReverseVideo(i bool) {
	if i {
		colour = 0x80
	} else {
		colour = 0
	}
}

//------------------------------------------------------------------------------

func Clear() {
	for i := range micro.Text {
		micro.Text[i] = '\x00'
	}
	micro.TextUpdated = true
}

//------------------------------------------------------------------------------

func Peek(x, y int) byte {
	sx, _ := micro.Size()
	x, y = Clamp(x, y)
	return micro.Text[x+y*sx]
}

func Poke(x, y int, value byte) {
	sx, _ := micro.Size()
	x, y = Clamp(x, y)
	ov := micro.Text[x+y*sx]
	if value != ov {
		micro.Text[x+y*sx] = value
		micro.TextUpdated = true
	}
}

//------------------------------------------------------------------------------

func Print(x, y int, things ...interface{}) (int, int) {
	origX, origY = x, y
	x, y = Clamp(x, y)
	for _, v := range things {
		switch v := v.(type) {
		case string:
			x, y = pString(x, y, v)
		case int:
			x, y = pString(x, y, strconv.Itoa(v))
		case uint:
			x, y = pString(x, y, strconv.FormatUint(uint64(v), 10))
		case int32:
			x, y = pString(x, y, strconv.FormatInt(int64(v), 10))
		case uint32:
			x, y = pString(x, y, strconv.FormatUint(uint64(v), 10))
		case int64:
			x, y = pString(x, y, strconv.FormatInt(v, 10))
		case uint64:
			x, y = pString(x, y, strconv.FormatUint(v, 10))
		case float32:
			x, y = pString(x, y, strconv.FormatFloat(float64(v), 'f', precision, 32))
		case float64:
			x, y = pString(x, y, strconv.FormatFloat(v, 'f', precision, 64))
		case bool:
			x, y = pString(x, y, strconv.FormatBool(v))
		case stringer:
			x, y = pString(x, y, v.String())
		case goStringer:
			x, y = pString(x, y, v.GoString())
		}
	}

	return x, y
}

var origX, origY int

var precision = -1

func SetPrecision(p int) {
	precision = p
}

//------------------------------------------------------------------------------

func pString(x, y int, s string) (int, int) {
	sx, sy := micro.Size()

	for i := range s {
		c := s[i]

		if c < ' ' {
			switch c {
			case '\a':
				if colour != 0 {
					colour = 0
				} else {
					colour = 0x80
				}
				continue

			case '\n':
				x = origX
				y++
				continue
			}

			c = '\177'

		} else if c > '~' {
			c = '\177'
		}

		if x < sx && y < sy {
			micro.Text[x+y*sx] = c | colour
		}
		x++
	}
	// Sanitize cursor position
	if x >= sx {
		x = sx - 1
	}
	if y >= sy {
		y = sy - 1
	}

	micro.TextUpdated = true

	return x, y
}

//------------------------------------------------------------------------------

type stringer interface {
	String() string
}

type goStringer interface {
	GoString() string
}

//------------------------------------------------------------------------------
