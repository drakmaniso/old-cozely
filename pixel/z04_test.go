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

	fontNames []string
	fonts     []pixel.FontID

	font          int
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

	a.fontNames = []string{
		"Monozela10 (builtin)",
		"fonts/tinela9",
		"fonts/simpela10",
		"fonts/cozela10",
		"fonts/monozela10",
		"fonts/simpela12",
		"fonts/cozela12",
		"fonts/chaotela12",
	}
	a.fonts = []pixel.FontID{
		pixel.Monozela10,
	}
	for i := 1; i < len(a.fontNames); i++ {
		a.fonts = append(a.fonts, pixel.Font(a.fontNames[i]))
	}
	a.font = 0

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
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
	}

	if input.MenuRight.Pushed() {
		a.font++
		if a.font >= len(a.fonts) {
			a.font = len(a.fonts) - 1
		}
	}
	if input.MenuLeft.Pushed() {
		a.font--
		if a.font < 0 {
			a.font = 0
		}
	}

	if input.MenuSelect.Pushed() {
		if a.show[0] == a.text[0] {
			a.show = a.code
			a.bg = 1
			a.fg = 7
		} else {
			a.show = a.text
			a.bg = 7
			a.fg = 1
		}
		a.interline = 0
	}
}

func (a *loop4) Update() {
	if input.MenuUp.Pressed() {
		a.line--
		if a.line < 0 {
			a.line = 0
		}
	}
	if input.MenuDown.Pressed() {
		a.line++
		if a.line > len(a.show)-1 {
			a.line = len(a.show) - 1
		}
	}
}

func (a *loop4) Render() {
	pixel.Clear(a.bg)

	cur := pixel.Cursor{}
	cur.Color = a.fg
	cur.Font = a.fonts[a.font]
	cur.Position = pixel.XY{16, cur.Font.Height() + 2}
	cur.Margin = cur.Position.X
	cur.LetterSpacing = a.letterspacing
	// u.Interline = fntInterline

	y := cur.Position.Y

	l := a.line
	for ; l < len(a.show) && y < pixel.Resolution().Y; l++ {
		cur.Println(a.show[l])
		y = cur.Position.Y
	}

	cur.Font = a.fonts[1]
	cur.Position = pixel.XY{pixel.Resolution().X - 128, cur.Font.Height() + 3}
	cur.Printf("Font: %s", a.fontNames[a.font])
	cur.Position = pixel.XY{pixel.Resolution().X - 64, pixel.Resolution().Y - 3}
	cur.Printf("Line: %d - %d", a.line, l-1)
}
