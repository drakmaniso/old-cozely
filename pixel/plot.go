// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"unicode/utf8"

	"github.com/drakmaniso/glam/palette"
)

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Picture(p Picture, x, y int16) {
	s.appendCommand(cmdPicture, 4, 1, int16(p), s.origin.X+x, s.origin.Y+y)
}

//------------------------------------------------------------------------------

func (s *ScreenCanvas) Text(f Font, c palette.Index, x, y int16, txt string) int16 {
	done := int(0)
	xx, yy := s.origin.X+x, s.origin.Y+y
	for done < utf8.RuneCountInString(txt) {
		t := []int16{int16(f), int16(c), xx, yy} //TODO: remove allocation
		dx := uint16(0)
		i := 0
		for _, r := range txt {
			i++
			if i <= done {
				continue
			}
			rr := uint16(0x7F)
			if r <= 0x7F {
				rr = uint16(r)
			}
			_, _, _, rw, _ := f.getMap(rune(rr))
			rr |= dx << 7
			t = append(t, int16(rr))
			dx += uint16(rw) + 0
			done++
			if dx > 0x1FF {
				break
			}
		}
		s.appendCommand(cmdPrint, 4, uint32(len(t)-4), t...)
		xx += int16(dx)
	}

	return xx + 7 - x
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
