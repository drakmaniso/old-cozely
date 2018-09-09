// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

type loop4 struct {
	fg, bg color.Index

	tinela9, monozela10, simpela10, simpela12,
	cozela10, cozela12, chaotela12, font pixel.FontID

	interline     int16
	letterspacing int16

	text []string
	code []string
	show []string
	line int
}

////////////////////////////////////////////////////////////////////////////////

func TestTest4(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop4{}
		l.declare()

		input.Load(bindings)
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (a *loop4) declare() {
	pixel.SetZoom(2)
	//TODO:
	a.bg = 7
	a.fg = 1

	a.tinela9 = pixel.Font("fonts/tinela9")
	a.monozela10 = pixel.Font("fonts/monozela10")
	a.simpela10 = pixel.Font("fonts/simpela10")
	a.simpela12 = pixel.Font("fonts/simpela12")
	a.cozela10 = pixel.Font("fonts/cozela10")
	a.cozela12 = pixel.Font("fonts/cozela12")
	a.chaotela12 = pixel.Font("fonts/chaotela12")
	a.font = a.monozela10

	a.interline = int16(18)
	a.letterspacing = int16(0)
}

func (a *loop4) Enter() {
	f, err := os.Open(cozely.Path() + "frankenstein.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		a.text = append(a.text, s.Text())
	}
	f, err = os.Open(cozely.Path() + "sourcecode.txt")
	if err != nil {
		panic(err)
	}
	s = bufio.NewScanner(f)
	for s.Scan() {
		a.code = append(a.code, s.Text())
	}
	a.show = a.text
}

func (loop4) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (a *loop4) React() {
	if scrollup.Started(0) {
		a.line--
		if a.line < 0 {
			a.line = 0
		}
	}
	if scrolldown.Started(0) {
		a.line++
		if a.line > len(a.show)-1 {
			a.line = len(a.show) - 1
		}
	}

	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop4) Update() {
}

func (a *loop4) Render() {
	pixel.Clear(a.bg)

	cur := pixel.Cursor{}

	cur.Color = a.fg
	cur.Locate(pixel.XY{16, a.font.Height() + 2})

	cur.Font = a.font
	cur.LetterSpacing = a.letterspacing
	// u.Interline = fntInterline

	y := cur.Position.Y

	for l := a.line; l < len(a.show) && y < pixel.Resolution().Y; l++ {
		cur.Println(a.show[l])
		y = cur.Position.Y
	}

	cur.Locate(pixel.XY{pixel.Resolution().X - 96, 16})
	cur.Printf("Line %d", a.line)
}

//TODO:
// func (fl fntLoop) KeyDown(l key.Keyabel, p key.Position) {
// 	switch l {
// 	case key.LabelSpace:
// 		if fntShow[0] == fntText[0] {
// 			fntShow = fntCode
// 			curBg = color.Find("black")
// 			curFg = color.Find("green")
// 		} else {
// 			fntShow = fntText
// 			curBg = color.Find("white")
// 			curFg = color.Find("black")
// 		}
// 		curScreen.Color = curFg - 1
// 		fntLine = 0
// 	case key.Label1:
// 		font = tinela9
// 	case key.Label2:
// 		font = monozela10
// 	case key.Label3:
// 		font = simpela10
// 	case key.Label4:
// 		font = cozela10
// 	case key.Label5:
// 		font = simpela12
// 	case key.Label6:
// 		font = cozela12
// 	case key.Label7:
// 		font = chaotela12
// 	case key.Label0:
// 		font = pixel.Font(0)
// 	case key.LabelKPDivide:
// 		fntLetterSpacing--
// 	case key.LabelKPMultiply:
// 		fntLetterSpacing++
// 	case key.LabelKPMinus:
// 		fntInterline--
// 	case key.LabelKPPlus:
// 		fntInterline++
// 	case key.LabelPageDown:
// 		fntLine += 40
// 		if fntLine > len(fntShow)-1 {
// 			fntLine = len(fntShow) - 1
// 		}
// 	case key.LabelPageUp:
// 		fntLine -= 40
// 		if fntLine < 0 {
// 			fntLine = 0
// 		}
// 	default:
// 		fl.EmptyLoop.KeyDown(l, p)
// 	}
// }

// func (fntLoop) MouseWheel(_, dy int32) {
// }
