package main

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/palettes/c64"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var (
	canvas = pixel.Canvas(pixel.Resolution(resolution.C, resolution.R))
)

var (
	resolution = coord.CR{640, 360}
	gridsize   = coord.CR{35, 21}
	cellsize   = coord.CR{16, 16}
	origin     coord.CR
)

const (
	Transparent color.Index = iota
	Black
	White
	Red
	Cyan
	Violet
	Green
	Blue
	Yellow
	Orange
	Brown
	LightRed
	DarkGrey
	Grey
	LightGreen
	LightBlue
	LightGrey
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	defer cozely.Recover()
	err := cozely.Run(loop{})
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

func (loop) Enter() {
	c64.Palette.Activate()
}

func (loop) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop) React() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() {
	canvas.Clear(0)
	origin = canvas.Size().Minus(resolution).Slash(2)
	canvas.Box(c64.Red, c64.Transparent, 3, 0,
		origin.Pluss(4), origin.Plus(resolution).Minuss(4))
	for c := int16(0); c < gridsize.C; c++ {
		for r := int16(0); r < gridsize.R; r++ {
			offset := resolution.Minus(gridsize.Times(16)).Slash(2)
			p := coord.CR{c, r}.Timescr(cellsize).Plus(offset)
			canvas.Box(c64.Yellow, c64.Transparent, 2, 0,
				origin.Plus(p), origin.Plus(p).Pluss(15))
		}
	}
	canvas.Display()
}

////////////////////////////////////////////////////////////////////////////////
