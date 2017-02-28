// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mtx

import (
	micro "github.com/drakmaniso/glam/internal/microtext"
	"strconv"
)

//------------------------------------------------------------------------------

var curX, curY int

func Cursor() (x, y int) {
	return curX, curY
}

func Move(x, y int) {
	sx, sy := micro.Size()
	if x <= 0 {
		curX = sx + x
		if curX < 0 {
			curX = 0
		}
	} else {
		curX = x
		if curX >= sx {
			curX = sx - 1
		}
	}
	if y <= 0 {
		curY = sy + y
		if curY < 0 {
			curY = 0
		}
	} else {
		curY = y
		if curY >= sy {
			curY = sy - 1
		}
	}
}

func Step(x, y int) {
	sx, sy := micro.Size()
	switch {
	case curX+x < 0:
		curX = 0
	case curX+x >= sx:
		curX = sx - 1
	default:
		curX += x
	}
	switch {
	case curY+y < 0:
		curY = 0
	case curY+y >= sy:
		curY = sy - 1
	default:
		curY += y
	}
}

//------------------------------------------------------------------------------

var hibit = byte(0)

func Invert(i bool) {
	if i {
		hibit = 0200
	} else {
		hibit = 0
	}
}

func ToggleInvert() {
	if hibit == 0 {
		hibit = 0200
	} else {
		hibit = 0
	}
}

//------------------------------------------------------------------------------

func Clear(c byte) {
	for i := range micro.Text {
		micro.Text[i] = c
	}
	curX, curY = 0, 0
	micro.TextUpdated = true
}

//------------------------------------------------------------------------------

func Peek(x, y int) byte {
	sx, sy := micro.Size()
	x %= sx
	y %= sy
	return micro.Text[x+y*sy]
}

func Poke(x, y int, value byte) {
	sx, sy := micro.Size()
	x %= sx
	y %= sy
	ov := micro.Text[x+y*sy]
	if value != ov {
		micro.Text[x+y*sy] = value
		micro.TextUpdated = true
	}
}

//------------------------------------------------------------------------------

func pString(s string) {
	sx, sy := micro.Size()

	for i := range s {
		c := s[i]

		if c < ' ' {
			switch c {
			case '\a':
				ToggleInvert()
				continue

			case '\n':
				curX = 0 //TODO
				curY++
				continue
			}

			c = '\177'

		} else if c > '~' {
			c = '\177'
		}

		if curX < sx && curY < sy {
			micro.Text[curX+curY*sx] = c | hibit
		}
		curX++
	}
	// Sanitize cursor position
	if curX >= sx {
		curX = sx - 1
	}
	if curY >= sy {
		curY = sy - 1
	}
}

//------------------------------------------------------------------------------

func Print(things ...interface{}) {
	for _, v := range things {
		switch v := v.(type) {
		case string:
			pString(v)
		case int:
			pInt(v)
		case uint:
			pUint(v)
		case int32:
			pInt32(v)
		case uint32:
			pUint32(v)
		case int64:
			pInt64(v)
		case uint64:
			pUint64(v)
		case float32:
			pFloat32(v)
		case float64:
			pFloat64(v)
		case bool:
			pBool(v)
		case stringer:
			pString(v.String())
		case goStringer:
			pString(v.GoString())
		}
	}
}

var precision = -1

func SetPrecision(p int) {
	precision = p
}

//------------------------------------------------------------------------------

type stringer interface {
	String() string
}

type goStringer interface {
	GoString() string
}

func pInt(value int) {
	pString(strconv.Itoa(value))
}

func pInt32(value int32) {
	pString(strconv.FormatInt(int64(value), 10))
}

func pInt64(value int64) {
	pString(strconv.FormatInt(value, 10))
}

func pUint(value uint) {
	pString(strconv.FormatUint(uint64(value), 10))
}

func pUint32(value uint32) {
	pString(strconv.FormatUint(uint64(value), 10))
}

func pUint64(value uint64) {
	pString(strconv.FormatUint(value, 10))
}

func pBool(value bool) {
	pString(strconv.FormatBool(value))
}

func pFloat64(value float64) {
	pString(strconv.FormatFloat(value, 'f', precision, 64))
}

func pFloat32(value float32) {
	pString(strconv.FormatFloat(float64(value), 'f', precision, 32))
}

//------------------------------------------------------------------------------
