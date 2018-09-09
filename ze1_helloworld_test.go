package cozely_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/pixel"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop struct{}

// Initialization //////////////////////////////////////////////////////////////

func Example_helloWorld() {
	defer cozely.Recover()

	pixel.SetResolution(320, 200)

	err := cozely.Run(loop{})
	if err != nil {
		panic(err)
	}
	// Output:
}

func (loop) Enter() {
}

func (loop) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop) React() {
}

func (loop) Update() {
}

func (loop) Render() {
	pixel.Clear(1)

	cur := pixel.Cursor{
		Color: 7,
		Position: pixel.XY{16, 32},
	}
	cur.Print("Hello, World!")
}
