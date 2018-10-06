package main

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var (
	resolution = pixel.XY{640, 360}
	gridsize   = pixel.XY{35, 21}
	cellsize   = pixel.XY{16, 16}
	origin     pixel.XY
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	defer cozely.Recover()
	pixel.SetResolution(pixel.XY{resolution.X, resolution.Y})
	err := cozely.Run(loop{})
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

func (loop) Enter() {
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
	pixel.Box(9, 0, 0, 3,
		origin.PlusS(4), origin.Plus(resolution).MinusS(4))
	for x := int16(0); x < gridsize.X; x++ {
		for y := int16(0); y < gridsize.Y; y++ {
			offset := resolution.Minus(gridsize.Times(16)).Slash(2)
			p := pixel.XY{x, y}.TimesXY(cellsize).Plus(offset)
			pixel.Box(10, 0, 0, 2,
				origin.Plus(p), origin.Plus(p).PlusS(15))
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
