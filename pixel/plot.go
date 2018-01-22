// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/palette"
)

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Picture(p Picture, x, y int16) {
	s.appendCommand(cmdPicture, 4, 1)
	s.parameters = append(s.parameters, int16(p), s.origin.X+x, s.origin.Y+y)
}

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Point(c palette.Index, x, y int16) {
	s.appendCommand(cmdPoint, 3, 1)
	s.parameters = append(s.parameters, int16(c), s.origin.X+x, s.origin.Y+y)
}

func (s *ScreenCanvas) PointList(c palette.Index, pts ...Coord) {
	if len(pts) < 1 {
		return
	}
	s.appendCommand(cmdPointList, 3, uint32(len(pts)))
	s.parameters = append(s.parameters, int16(c))
	for _, p := range pts {
		s.parameters = append(s.parameters, s.origin.X+p.X, s.origin.Y+p.Y)
	}
}

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Line(c palette.Index, x1, y1, x2, y2 int16) {
	s.appendCommand(cmdLine, 4, 1)
	s.parameters = append(s.parameters, int16(c), s.origin.X+x1, s.origin.Y+y1, s.origin.X+x2, s.origin.Y+y2)
}

//------------------------------------------------------------------------------
