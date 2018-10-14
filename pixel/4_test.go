package pixel_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/pico8"
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

func (l *loop4) declare() {
	pixel.SetZoom(2)

	color.Load(&pico8.Palette)
	l.bg = pico8.White
	l.fg = pico8.DarkBlue

	l.fontNames = []string{
		"Monozela10 (builtin)",
		"fonts/tinela9",
		"fonts/simpela10",
		"fonts/cozela10",
		"fonts/monozela10",
		"fonts/simpela12",
		"fonts/cozela12",
		"fonts/chaotela12",
	}
	l.fonts = []pixel.FontID{
		pixel.Monozela10,
	}
	for i := 1; i < len(l.fontNames); i++ {
		l.fonts = append(l.fonts, pixel.Font(l.fontNames[i]))
	}
	l.font = 0

	l.interline = int16(18)
	l.letterspacing = int16(0)
}

func (l *loop4) Enter() {
	f, err := os.Open(cozely.Path() + "frankenstein.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		l.text = append(l.text, s.Text())
	}
	f, err = os.Open(cozely.Path() + "sourcecode.txt")
	if err != nil {
		panic(err)
	}
	s = bufio.NewScanner(f)
	for s.Scan() {
		l.code = append(l.code, s.Text())
	}
	l.show = l.text
}

func (loop4) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (l *loop4) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}

	if input.Right.Pressed() {
		l.font++
		if l.font >= len(l.fonts) {
			l.font = len(l.fonts) - 1
		}
	}
	if input.Left.Pressed() {
		l.font--
		if l.font < 0 {
			l.font = 0
		}
	}

	if input.Select.Pressed() {
		if l.show[0] == l.text[0] {
			l.show = l.code
			l.bg, l.fg = l.fg, l.bg
		} else {
			l.show = l.text
			l.bg, l.fg = l.fg, l.bg
		}
		l.interline = 0
	}
}

func (l *loop4) Update() {
	if input.Up.Ongoing() {
		l.line--
		if l.line < 0 {
			l.line = 0
		}
	}
	if input.Down.Ongoing() {
		l.line++
		if l.line > len(l.show)-1 {
			l.line = len(l.show) - 1
		}
	}
}

func (l *loop4) Render() {
	pixel.Clear(l.bg)

	cur := pixel.Cursor{}
	cur.Color = l.fg
	cur.Font = l.fonts[l.font]
	cur.Position = pixel.XY{16, cur.Font.Height() + 2}
	cur.Margin = cur.Position.X
	cur.LetterSpacing = l.letterspacing
	// u.Interline = fntInterline

	y := cur.Position.Y

	li := l.line
	for ; li < len(l.show) && y < pixel.Resolution().Y; li++ {
		cur.Println(l.show[li])
		y = cur.Position.Y
	}

	cur.Font = l.fonts[1]
	cur.Position = pixel.XY{pixel.Resolution().X - 128, cur.Font.Height() + 3}
	cur.Printf("Font: %s", l.fontNames[l.font])
	cur.Position = pixel.XY{pixel.Resolution().X - 64, pixel.Resolution().Y - 3}
	cur.Printf("Line: %d - %d", l.line, li-1)
}

// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
