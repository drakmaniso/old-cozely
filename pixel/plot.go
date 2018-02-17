// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/palette"
)

//------------------------------------------------------------------------------

func (s *Canvas) Picture(p Picture, x, y int16) {
	s.appendCommand(cmdPicture, 4, 1, int16(p), x, y)
}

//------------------------------------------------------------------------------

func (s *Canvas) Point(c palette.Index, x, y int16) {
	s.appendCommand(cmdPoint, 3, 1, int16(c), x, y)
}

func (s *Canvas) PointList(c palette.Index, pts ...Coord) {
	if len(pts) < 1 {
		return
	}
	prm := []int16{int16(c)} //TODO: remove alloc
	for _, p := range pts {
		prm = append(prm, p.X, p.Y)
	}
	s.appendCommand(cmdPointList, 3, uint32(len(prm)/2-1), prm...)
}

//------------------------------------------------------------------------------

func (s *Canvas) Line(c palette.Index, x1, y1, x2, y2 int16) {
	s.appendCommand(cmdLine, 4, 1, int16(c), x1, y1, x2, y2)
}

//------------------------------------------------------------------------------
