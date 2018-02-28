// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/drakmaniso/glam/key"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var (
	tinela9    = pixel.NewFont("fonts/tinela9")
	monozela10 = pixel.NewFont("fonts/monozela10")
	simpela10  = pixel.NewFont("fonts/simpela10")
	simpela12  = pixel.NewFont("fonts/simpela12")
	cozela10   = pixel.NewFont("fonts/cozela10")
	cozela12   = pixel.NewFont("fonts/cozela12")
	chaotela12  = pixel.NewFont("fonts/chaotela12")
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

//------------------------------------------------------------------------------

func TestFont_load(t *testing.T) {
	do(func() {
		err := glam.Run(fntLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type fntLoop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (fntLoop) Enter() error {
	f, err := os.Open(glam.Path() + "frankenstein_full.txt")
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		fntText = append(fntText, s.Text())
	}
	f, err = os.Open(glam.Path() + "sourcecode.txt")
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
	cursor.ColorShift(curFg - 1)
	return nil
}

//------------------------------------------------------------------------------

func (fntLoop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (fntLoop) Draw() error {
	curScreen.Clear(curBg)

	cursor.Locate(16, font.Height()+2, 0)

	cursor.Font(font)
	cursor.LetterSpacing(fntLetterSpacing)
	// cursor.Interline(fntInterline)

	_, y, _ := cursor.Position()

	for l := fntLine; l < len(fntShow) && y < curScreen.Size().Y; l++ {
		cursor.Println(fntShow[l])
		_, y, _ = cursor.Position()
	}

	cursor.Locate(curScreen.Size().X-96, 16, 0)
	cursor.Printf("Line %d", fntLine)

	curScreen.Display()
	return nil
}

//------------------------------------------------------------------------------

func (fl fntLoop) KeyDown(l key.Label, p key.Position) {
	switch l {
	case key.LabelSpace:
		if fntShow[0] == fntText[0] {
			fntShow = fntCode
			curBg = palette.Find("black")
			curFg = palette.Find("green")
		} else {
			fntShow = fntText
			curBg = palette.Find("white")
			curFg = palette.Find("black")
		}
		cursor.ColorShift(curFg - 1)
		fntLine = 0
	case key.Label1:
		font = tinela9
	case key.Label2:
		font = monozela10
	case key.Label3:
		font = simpela10
	case key.Label4:
		font = cozela10
	case key.Label5:
		font = simpela12
	case key.Label6:
		font = cozela12
	case key.Label7:
		font = chaotela12
	case key.Label0:
		font = pixel.Font(0)
	case key.LabelKPDivide:
		fntLetterSpacing--
	case key.LabelKPMultiply:
		fntLetterSpacing++
	case key.LabelKPMinus:
		fntInterline--
	case key.LabelKPPlus:
		fntInterline++
	case key.LabelPageDown:
		fntLine += 40
		if fntLine > len(fntShow) - 1 {
			fntLine = len(fntShow) - 1
		}
	case key.LabelPageUp:
		fntLine -= 40
		if fntLine < 0 {
			fntLine = 0
		}
	default:
		fl.Handlers.KeyDown(l, p)
	}
}

//------------------------------------------------------------------------------

func (fntLoop) MouseWheel(_, dy int32) {
	fntLine -= int(dy)
	if fntLine < 0 {
		fntLine = 0
	} else if fntLine > len(fntShow)-1 {
		fntLine = len(fntShow) - 1
	}
}

//------------------------------------------------------------------------------
