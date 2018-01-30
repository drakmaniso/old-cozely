// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/palette"
)

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Picture(p Picture, x, y int16) {
	s.appendCommand(cmdPicture, 4, 1, int16(p), s.origin.X+x, s.origin.Y+y)
}

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Point(c palette.Index, x, y int16) {
	s.appendCommand(cmdPoint, 3, 1, int16(c), s.origin.X+x, s.origin.Y+y)
}

func (s *ScreenCanvas) PointList(c palette.Index, pts ...Coord) {
	if len(pts) < 1 {
		return
	}
	prm := []int16{int16(c)} //TODO: remove alloc
	for _, p := range pts {
		prm = append(prm, s.origin.X+p.X, s.origin.Y+p.Y)
	}
	s.appendCommand(cmdPointList, 3, uint32(len(prm)/2-1), prm...)
}

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Line(c palette.Index, x1, y1, x2, y2 int16) {
	s.appendCommand(cmdLine, 4, 1, int16(c), s.origin.X+x1, s.origin.Y+y1, s.origin.X+x2, s.origin.Y+y2)
}

//------------------------------------------------------------------------------
