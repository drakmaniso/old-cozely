// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package vector

import "github.com/cozely/cozely/color"

////////////////////////////////////////////////////////////////////////////////

// func Point(c color.Colour, x, y int16) {
// 	appendCommand(cmdPoint, 3, 1)
// 	c8 := color.SRGBA8Of(c)
// 	rg := uint16(c8.R)<<8 | uint16(c8.G)
// 	ba := uint16(c8.B)<<8 | uint16(c8.A)
// 	parameters = append(parameters, int16(rg), int16(ba), x, y)
// }

// func PointList(c color.Colour, pts ...Coord) {
// 	if len(pts) < 1 {
// 		return
// 	}
// 	appendCommand(cmdPointList, 3, uint32(len(pts)))
// 	c8 := color.SRGBA8Of(c)
// 	rg := uint16(c8.R)<<8 | uint16(c8.G)
// 	ba := uint16(c8.B)<<8 | uint16(c8.A)
// 	parameters = append(parameters, int16(rg), int16(ba))
// 	for _, p := range pts {
// 		parameters = append(parameters, p.X, p.Y)
// 	}
// }

////////////////////////////////////////////////////////////////////////////////

func Line(c color.Color, x1, y1, x2, y2 int16) {
	appendCommand(cmdLineAA, 4, 1)
	c8 := color.SRGBA8of(c)
	rg := uint16(c8.R)<<8 | uint16(c8.G)
	ba := uint16(c8.B)<<8 | uint16(c8.A)
	parameters = append(parameters, int16(rg), int16(ba), x1, y1, x2, y2)
}

////////////////////////////////////////////////////////////////////////////////
