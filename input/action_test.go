// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input_test

//------------------------------------------------------------------------------

import (
	"fmt"
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/input"
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

const (
	Transparent palette.Index = iota
	Black
	MediumGreen
	LightGreen
	DarkBlue
	LightBlue
	DarkRed
	Cyan
	MediumRed
	LightRed
	DarkYellow
	LightYellow
	DarkGreen
	Magenta
	Gray
	White
)

//------------------------------------------------------------------------------

var (
	InventoryAction        = input.NewBool("Inventory")
	OptionsAction          = input.NewBool("Options")
	CloseMenuAction        = input.NewBool("Close Menu")
	InstantCloseMenuAction = input.NewBool("Instant Close Menu")
	JumpAction             = input.NewBool("Jump")
	OpenMenuAction         = input.NewBool("Open Menu")
	InstantOpenMenuAction  = input.NewBool("Instant Open Menu")
)

var (
	InMenu = input.NewContext("Menu",
		CloseMenuAction, InstantCloseMenuAction, InventoryAction, OptionsAction)

	InGame = input.NewContext("Game",
		OpenMenuAction, InstantOpenMenuAction, InventoryAction, JumpAction)
)

var (
	Bindings = map[string]map[string][]string{
		"Menu": {
			"Close Menu":         {"Escape"},
			"Instant Close Menu": {"Mouse Right", "Enter"},
			"Inventory":          {"I"},
			"Options":            {"O", "Mouse Left"},
		},
		"Game": {
			"Open Menu":         {"Escape"},
			"Instant Open Menu": {"Mouse Right", "Enter"},
			"Inventory":         {"Tab"},
			"Jump":              {"Space", "Mouse Left"},
		},
	}
)

//------------------------------------------------------------------------------

func TestAction(t *testing.T) {
	err := glam.Run(loop{})
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
	err := input.LoadBindings(Bindings)
	if err != nil {
		return err
	}

	palette.Load("MSX")
	InMenu.Activate(0)
	return nil
}

//------------------------------------------------------------------------------

var dx, dy int32
var px, py int32
var left, middle, right, extra1, extra2 bool

var openmenu, closemenu, instopenmenu, instclosemenu, inventory, options, jump bool

func (loop) React() error {
	dx, dy = mouse.Delta()
	px, py = mouse.Position()
	left = mouse.IsPressed(mouse.Left)
	middle = mouse.IsPressed(mouse.Middle)
	right = mouse.IsPressed(mouse.Right)
	extra1 = mouse.IsPressed(mouse.Extra1)
	extra2 = mouse.IsPressed(mouse.Extra2)

	if CloseMenuAction.JustPressed(1) {
		println(" Just Pressed: CLOSE")
	}
	if CloseMenuAction.JustReleased(1) {
		println("Just Released: close")
		InGame.Activate(1)
	}
	if OpenMenuAction.JustPressed(1) {
		println(" Just Pressed: OPEN")
	}
	if OpenMenuAction.JustReleased(1) {
		println("Just Released: open")
		InMenu.Activate(1)
	}

	if InstantCloseMenuAction.JustPressed(1) {
		println(" Just Pressed: INSTANT CLOSE")
		InGame.Activate(1)
	}
	if InstantCloseMenuAction.JustReleased(1) {
		println("Just Released: instant close")
	}
	if InstantOpenMenuAction.JustPressed(1) {
		println(" Just Pressed: INSTANT OPEN")
		InMenu.Activate(1)
	}
	if InstantOpenMenuAction.JustReleased(1) {
		println("Just Released: instant open")
	}

	openmenu = OpenMenuAction.Pressed(1)
	closemenu = CloseMenuAction.Pressed(1)
	instopenmenu = InstantOpenMenuAction.Pressed(1)
	instclosemenu = InstantCloseMenuAction.Pressed(1)
	inventory = InventoryAction.Pressed(1)
	options = OptionsAction.Pressed(1)
	jump = JumpAction.Pressed(1)

	return nil
}

//------------------------------------------------------------------------------

func color(p bool) {
	if p {
		cursor.Color = LightGreen - 1
	} else {
		cursor.Color = DarkBlue - 1
	}
}

func (loop) Draw() error {
	screen.Clear(0)

	cursor.Locate(2, 12)
	cursor.Color = DarkBlue - 1

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

	color(InMenu.Active(1))
	cursor.Printf("  Menu: ")
	color(options)
	cursor.Print("Options(O/L.C.) ")
	color(closemenu)
	cursor.Print("CloseMenu(ESC) ")
	color(instclosemenu)
	cursor.Print("InstantCloseMenu(ENTER/R.C.) ")
	cursor.Println(" ")

	color(InGame.Active(1))
	cursor.Printf("  Game: ")
	color(jump)
	cursor.Print("Jump(SPACE/L.C.) ")
	color(openmenu)
	cursor.Print("OpenMenu(ESC) ")
	color(instopenmenu)
	cursor.Print("InstantOpenMenu(ENTER/R.C.) ")
	cursor.Println(" ")

	color(false)
	cursor.Printf("  Both: ")
	color(inventory)
	cursor.Print("Inventory(I/TAB) ")

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
