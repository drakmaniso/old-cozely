// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/palette"
)

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Picture(p Picture, x, y int16) {
	s.appendCommand(cmdPicture, 4, 1)
	s.parameters = append(s.parameters, int16(p), x, y)
}

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Point(c palette.Index, x, y int16) {
	s.appendCommand(cmdPoint, 3, 1)
	s.parameters = append(s.parameters, int16(c), x, y)
}

func (s *ScreenCanvas) PointList(c palette.Index, pts ...Coord) {
	if len(pts) < 1 {
		return
	}
	s.appendCommand(cmdPointList, 3, uint32(len(pts)))
	s.parameters = append(s.parameters, int16(c))
	for _, p := range pts {
		s.parameters = append(s.parameters, p.X, p.Y)
	}
}

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Line(c palette.Index, x1, y1, x2, y2 int16) {
	s.appendCommand(cmdLine, 4, 1)
	s.parameters = append(s.parameters, int16(c), x1, y1, x2, y2)
}

//------------------------------------------------------------------------------
