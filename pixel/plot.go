// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/palette"
)

//------------------------------------------------------------------------------

func Point(c palette.Index, x, y int16) {
	appendCommand(cmdPoint, 3, 1)
	parameters = append(parameters, int16(c), x, y)
}

func PointList(c palette.Index, pts ...Coord) {
	if len(pts) < 1 {
		return
	}
	appendCommand(cmdPointList, 3, uint32(len(pts)))
	parameters = append(parameters, int16(c))
	for _, p := range pts {
		parameters = append(parameters, p.X, p.Y)
	}
}

//------------------------------------------------------------------------------

func Line(c palette.Index, x1, y1, x2, y2 int16) {
	appendCommand(cmdLine, 4, 1)
	parameters = append(parameters, int16(c), x1, y1, x2, y2)
}

//------------------------------------------------------------------------------
