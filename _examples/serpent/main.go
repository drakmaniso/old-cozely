package main

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/palettes/c64"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var (
	resolution = pixel.XY{640, 360}
	gridsize   = pixel.XY{35, 21}
	cellsize   = pixel.XY{16, 16}
	origin     pixel.XY
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
	pixel.SetResolution(resolution.X, resolution.Y)
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
	pixel.Clear(0)
	origin = pixel.Resolution().Minus(resolution).Slash(2)
	pixel.Box(c64.Red, c64.Transparent, 3,
		origin.Pluss(4), origin.Plus(resolution).Minuss(4))
	for x := int16(0); x < gridsize.X; x++ {
		for y := int16(0); y < gridsize.Y; y++ {
			offset := resolution.Minus(gridsize.Times(16)).Slash(2)
			p := pixel.XY{x, y}.Timescr(cellsize).Plus(offset)
			pixel.Box(c64.Yellow, c64.Transparent, 2,
				origin.Plus(p), origin.Plus(p).Pluss(15))
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
