// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/drakmaniso/cozely"
	"github.com/drakmaniso/cozely/input"
	"github.com/drakmaniso/cozely/palette"
	"github.com/drakmaniso/cozely/pixel"
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
	fntInterline     = int16(18)
	fntLetterSpacing = int16(0)
)

var (
	fntText []string
	fntCode []string
	fntShow []string
	fntLine int
)

////////////////////////////////////////////////////////////////////////////////

func TestFont_load(t *testing.T) {
	do(func() {
		err := cozely.Run(fntLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////

type fntLoop struct{}

////////////////////////////////////////////////////////////////////////////////

func (fntLoop) Enter() error {
	input.Load(testBindings)
	testContext.Activate(1)

	f, err := os.Open(cozely.Path() + "frankenstein.txt")
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		fntText = append(fntText, s.Text())
	}
	f, err = os.Open(cozely.Path() + "sourcecode.txt")
	if err != nil {
		return err
	}
	s = bufio.NewScanner(f)
	for s.Scan() {
		fntCode = append(fntCode, s.Text())
	}
	fntShow = fntText
	palette.Load("C64")
	curBg = palette.Find("white")
	curFg = palette.Find("black")
	cursor.Color = curFg - 1
	return nil
}

func (fntLoop) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (fntLoop) React() error {
	if quit.JustPressed(1) {
		cozely.Stop()
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (fntLoop) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (fntLoop) Render() error {
	curScreen.Clear(curBg)

	cursor.Locate(16, font.Height()+2)

	cursor.Font = font
	cursor.Spacing = fntLetterSpacing
	// cursor.Interline = fntInterline

	y := cursor.Position.Y

	for l := fntLine; l < len(fntShow) && y < curScreen.Size().Y; l++ {
		cursor.Println(fntShow[l])
		y = cursor.Position.Y
	}

	cursor.Locate(curScreen.Size().X-96, 16)
	cursor.Printf("Line %d", fntLine)

	curScreen.Display()
	return nil
}

////////////////////////////////////////////////////////////////////////////////

//TODO:
// func (fl fntLoop) KeyDown(l key.Keyabel, p key.Position) {
// 	switch l {
// 	case key.LabelSpace:
// 		if fntShow[0] == fntText[0] {
// 			fntShow = fntCode
// 			curBg = palette.Find("black")
// 			curFg = palette.Find("green")
// 		} else {
// 			fntShow = fntText
// 			curBg = palette.Find("white")
// 			curFg = palette.Find("black")
// 		}
// 		cursor.Color = curFg - 1
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
// 	fntLine -= int(dy)
// 	if fntLine < 0 {
// 		fntLine = 0
// 	} else if fntLine > len(fntShow)-1 {
// 		fntLine = len(fntShow) - 1
// 	}
// }

////////////////////////////////////////////////////////////////////////////////
