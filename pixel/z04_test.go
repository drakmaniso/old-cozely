// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var (
	tinela9    = pixel.Font("fonts/tinela9")
	monozela10 = pixel.Font("fonts/monozela10")
	simpela10  = pixel.Font("fonts/simpela10")
	simpela12  = pixel.Font("fonts/simpela12")
	cozela10   = pixel.Font("fonts/cozela10")
	cozela12   = pixel.Font("fonts/cozela12")
	chaotela12 = pixel.Font("fonts/chaotela12")
	font       = monozela10
)

var (
	interline     = int16(18)
	letterspacing = int16(0)
)

var (
	text []string
	code []string
	show []string
	line int
)

type loop4 struct{}

////////////////////////////////////////////////////////////////////////////////

func TestTest4(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		err := cozely.Run(loop4{})
		if err != nil {
			t.Error(err)
		}
	})
}

func (loop4) Enter() {
	input.Load(bindings)
	context.Activate(0)

	f, err := os.Open(cozely.Path() + "frankenstein.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		text = append(text, s.Text())
	}
	f, err = os.Open(cozely.Path() + "sourcecode.txt")
	if err != nil {
		panic(err)
	}
	s = bufio.NewScanner(f)
	for s.Scan() {
		code = append(code, s.Text())
	}
	show = text
	palette3.Activate()
}

func (loop4) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop4) React() {
	if scrollup.Started(0) {
		line--
		if line < 0 {
			line = 0
		}
	}
	if scrolldown.Started(0) {
		line++
		if line > len(show)-1 {
			line = len(show) - 1
		}
	}

	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop4) Update() {
}

func (loop4) Render() {
	canvas3.Clear(bg3)

	canvas3.Cursor().Color = fg3
	canvas3.Locate(0, coord.CR{16, font.Height() + 2})

	canvas3.Cursor().Font = font
	canvas3.Cursor().LetterSpacing = letterspacing
	// curScreen.Interline = fntInterline

	y := canvas3.Cursor().Position.R

	for l := line; l < len(show) && y < canvas3.Size().R; l++ {
		canvas3.Println(show[l])
		y = canvas3.Cursor().Position.R
	}

	canvas3.Locate(0, coord.CR{canvas3.Size().C - 96, 16})
	canvas3.Printf("Line %d", line)

	canvas3.Display()
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
