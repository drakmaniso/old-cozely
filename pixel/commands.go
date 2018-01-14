// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

/*
type paramIndexedPicture struct {
	x int16
	y int16
}

type paramFullIndexedPicture struct {
	x                 int16
	y                 int16
	w                 int16
	h                 int16
	transform, shift  uint8
	alpha, brightness uint8
}

type paramRGBAPicture struct {
	x int16
	y int16
}

type paramFulRGBAPicturel struct {
	x                int16
	y                int16
	w                int16
	h                int16
	transform, alpha uint8
}

//------------------------------------------------------------------------------

type paramIndexedPoint struct {
	color, _ uint8
	x        int16
	y        int16
}

type paramPoint struct {
	rg uint16
	ba uint16
	x  int16
	y  int16
}

type paramIndexedLine struct {
	color, width uint8
	x1           int16
	y1           int16
	x2           int16
	y2           int16
}

type paramRGBALine struct {
	rg               uint16
	ba               uint16
	x1               int16
	y1               int16
	x2               int16
	y2               int16
	width, antialias uint8
}

type paramRGBABezier struct {
	rg               uint16
	ba               uint16
	x1               int16
	y1               int16
	x2               int16
	y2               int16
	x3               int16
	y3               int16
	x4               int16
	y4               int16
	width, antialias uint8
}

*/

//------------------------------------------------------------------------------

const (
	cmdIndexedPoint uint32 = 1 << 2
)

//------------------------------------------------------------------------------
