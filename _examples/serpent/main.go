package main

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/plane"
)

////////////////////////////////////////////////////////////////////////////////

var (
	canvas = pixel.Canvas(pixel.TargetResolution(resolution.C, resolution.R))
)

var (
	resolution = plane.CR{640, 360}
	gridsize   = plane.CR{35, 21}
	cellsize   = plane.CR{16, 16}
	origin     plane.CR
)

const (
	Transparent palette.Index = iota
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
	err := cozely.Run(loop{})
	if err != nil {
		cozely.ShowError(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

func (loop) Enter() error {
	palette.Load("C64")
	return nil
}

func (loop) Leave() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) React() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() error {
	canvas.Clear(0)
	origin = canvas.Size().Minus(resolution).Slash(2)
	canvas.Box(Red, 0, 3, 0, origin.Pluss(4, 4), origin.Plus(resolution).Minuss(4, 4))
	for c := int16(0); c < gridsize.C; c++ {
		for r := int16(0); r < gridsize.R; r++ {
			offset := resolution.Minus(gridsize.Times(16)).Slash(2)
			p := plane.CR{c, r}.Timescw(cellsize).Plus(offset)
			canvas.Box(Yellow, 0, 2, 0, origin.Plus(p), origin.Plus(p).Pluss(15, 15))
		}
	}
	canvas.Display()
	return nil
}

////////////////////////////////////////////////////////////////////////////////
