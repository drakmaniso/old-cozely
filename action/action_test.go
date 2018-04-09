// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action_test

//------------------------------------------------------------------------------

import (
	"fmt"
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/action"
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var (
	screen = pixel.NewCanvas(pixel.Zoom(3))
	cursor = pixel.Cursor{Canvas: screen}
)

//------------------------------------------------------------------------------

var (
	InMenuUp    = action.NewBool("Menu Up")
	InMenuDown  = action.NewBool("Menu Down")
	InMenuStart = action.NewBool("Menu Start")
	InGameUp    = action.NewBool("Up")
	InGameDown  = action.NewBool("Down")
	InGamePause = action.NewBool("Pause")
	QuitAction  = action.NewBool("Quit")
)

var (
	InMenu = action.NewContext("Menu",
		InMenuUp, InMenuDown, InMenuStart,
		QuitAction)

	InGame = action.NewContext("Game",
		InGameUp, InGameDown, InGamePause,
		QuitAction)
)

var (
	Bindings = map[string]map[string][]string{
		"Menu": {
			"Menu Up":    {"Up"},
			"Menu Down":  {"Down"},
			"Menu Start": {"Space"},
			"Quit":       {"Escape"},
		},
		"Game": {
			"Up":    {"W"},
			"Down":  {"S"},
			"Pause": {"Enter"},
			"Quit":  {"Escape"},
		},
	}
)

//------------------------------------------------------------------------------

func TestAction(t *testing.T) {
	err := action.LoadBindings(Bindings)
	if err != nil {
		glam.ShowError(err)
		return
	}

	err = glam.Run(loop{})
	if err != nil {
		glam.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (loop) Enter() error {
	palette.Index(1).SetColour(colour.LRGB{1, .95, .9})
	InMenu.Activate()
	return nil
}

//------------------------------------------------------------------------------

var dx, dy int32
var px, py int32
var left, middle, right, extra1, extra2 bool

var menuup, menudown, menustart, quit bool
var gameup, gamedown, gamepause bool

func (loop) React() error {
	dx, dy = mouse.Delta()
	px, py = mouse.Position()
	left = mouse.IsPressed(mouse.Left)
	middle = mouse.IsPressed(mouse.Middle)
	right = mouse.IsPressed(mouse.Right)
	extra1 = mouse.IsPressed(mouse.Extra1)
	extra2 = mouse.IsPressed(mouse.Extra2)

	if InMenuStart.JustPressed() {
		println(" Just Pressed: Menu Start")
	}
	if InMenuStart.JustReleased() {
		println("Just Released: Menu Start")
		InGame.Activate()
	}
	if InGamePause.JustPressed() {
		println(" Just Pressed: Game Pause")
		InMenu.Activate()
	}
	if InGamePause.JustReleased() {
		println("Just Released: Game Pause")
	}

	menuup = InMenuUp.Pressed()
	menudown = InMenuDown.Pressed()
	menustart = InMenuStart.Pressed()
	quit = QuitAction.Pressed()
	gameup = InGameUp.Pressed()
	gamedown = InGameDown.Pressed()
	gamepause = InGamePause.Pressed()

	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw() error {
	screen.Clear(0)

	cursor.Locate(2, 12)

	cursor.Printf("   mouse delta:%+6d,%+6d\n", dx, dy)
	cursor.Printf("mouse position:%6d,%6d\n", px, py)

	cursor.Printf(" mouse buttons: ")
	if left {
		cursor.Print("LEFT ")
	} else {
		cursor.Print("left ")
	}
	if middle {
		cursor.Print("MIDDLE ")
	} else {
		cursor.Print("middle ")
	}
	if right {
		cursor.Print("RIGHT ")
	} else {
		cursor.Print("right ")
	}
	if extra1 {
		cursor.Print("EXTRA1 ")
	} else {
		cursor.Print("extra1 ")
	}
	if extra2 {
		cursor.Print("EXTRA2\n")
	} else {
		cursor.Print("extra2\n")
	}

	if InMenu.Active() {
		cursor.Printf("  MENU ACTIONS: ")
	} else {
		cursor.Printf("  menu actions: ")
	}
	if menuup {
		cursor.Print("UP ")
	} else {
		cursor.Print("up ")
	}
	if menudown {
		cursor.Print("DOWN ")
	} else {
		cursor.Print("down ")
	}
	if menustart {
		cursor.Print("START ")
	} else {
		cursor.Print("start ")
	}
	cursor.Println(" ")

	if InGame.Active() {
		cursor.Printf("  GAME ACTIONS: ")
	} else {
		cursor.Printf("  game actions: ")
	}
	if gameup {
		cursor.Print("UP ")
	} else {
		cursor.Print("up ")
	}
	if gamedown {
		cursor.Print("DOWN ")
	} else {
		cursor.Print("down ")
	}
	if gamepause {
		cursor.Print("PAUSE ")
	} else {
		cursor.Print("pause ")
	}
	cursor.Println(" ")

	cursor.Printf("       actions: ")
	if quit {
		cursor.Print("QUIT ")
	} else {
		cursor.Print("quit ")
	}

	screen.Display()
	return nil
}

//------------------------------------------------------------------------------

var relative = false

func (loop) KeyDown(l key.Label, p key.Position) {
	if l == key.LabelSpace {
		relative = !relative
		mouse.SetRelativeMode(relative)
	}
	if l == key.LabelEscape {
		glam.Stop()
	}
	fmt.Printf("%v: Key Down: %v %v\n", glam.GameTime(), l, p)
}

func (loop) MouseWheel(dx, dy int32) {
	fmt.Printf("%v: mouse wheel: %+d,%+d\n", glam.GameTime(), dx, dy)
}

//------------------------------------------------------------------------------

func (loop) Show() {
	fmt.Printf("%v: show\n", glam.GameTime())
}

func (loop) Hide() {
	fmt.Printf("%v: hide\n", glam.GameTime())
}

func (loop) Resize() {
	s := glam.WindowSize()
	fmt.Printf("%v: resize %dx%d\n", glam.GameTime(), s.X, s.Y)
}

func (loop) Focus() {
	fmt.Printf("%v: focus\n", glam.GameTime())
}

func (loop) Unfocus() {
	fmt.Printf("%v: unfocus\n", glam.GameTime())
}

func (loop) Quit() {
	fmt.Printf("%v: quit\n", glam.GameTime())
	glam.Stop()
}

//------------------------------------------------------------------------------
