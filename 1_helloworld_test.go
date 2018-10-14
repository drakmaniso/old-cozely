package cozely_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

func Example_helloWorld() {
	cozely.Run(loop{})
	//Output:
}

type loop struct{}

// Enter the game loop
func (loop) Enter() {}

// Leave the game loop
func (loop) Leave() {}

// React to user inputs
func (loop) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}
}

var cursor = pixel.Cursor{}

// Update the game state
func (loop) Update() {
	cursor.Color = color.White
	if int(cozely.GameTime()*10)%3 == 0 {
		cursor.Color = color.MidGray
	}
}

// Render the game state
func (loop) Render() {
	pixel.Clear(color.DarkGray)
	cursor.Position = pixel.XY{8, 16}
	cursor.Print("hello, world!")
}
