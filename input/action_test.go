// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input_test

////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/palette"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var (
	InventoryAction        = input.Bool("Inventory")
	OptionsAction          = input.Bool("Options")
	CloseMenuAction        = input.Bool("Close Menu")
	InstantCloseMenuAction = input.Bool("Instant Close Menu")
	JumpAction             = input.Bool("Jump")
	OpenMenuAction         = input.Bool("Open Menu")
	InstantOpenMenuAction  = input.Bool("Instant Open Menu")
)

var (
	InMenu = input.Context("Menu",
		CloseMenuAction, InstantCloseMenuAction, InventoryAction, OptionsAction)

	InGame = input.Context("Game",
		OpenMenuAction, InstantOpenMenuAction, InventoryAction, JumpAction)
)

var (
	Bindings = input.Bindings{
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

////////////////////////////////////////////////////////////////////////////////

var (
	screen = pixel.Canvas(pixel.Zoom(3))
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

var hidden bool
var mousepos, mousedelta coord.CR
var openmenu, closemenu, instopenmenu, instclosemenu, inventory, options, jump bool

////////////////////////////////////////////////////////////////////////////////

func TestAction(t *testing.T) {
	err := cozely.Run(loop{})
	if err != nil {
		cozely.ShowError(err)
		return
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop) Enter() error {
	err := input.Load(Bindings)
	if err != nil {
		return err
	}

	palette.Load("MSX")
	InMenu.Activate(0)
	return nil
}

func (loop) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop) React() error {
	mousepos = input.Cursor.Position()
	mousedelta = input.Cursor.Delta()

	if JumpAction.JustPressed(1) {
		println(" Just Pressed: *JUMP*")
	}
	if JumpAction.JustReleased(1) {
		println("Just Released: (jump)")
	}

	if CloseMenuAction.JustReleased(1) {
		InGame.Activate(1)
		input.Cursor.Hide()
	}
	if OpenMenuAction.JustReleased(1) {
		InMenu.Activate(1)
		input.Cursor.Show()
	}

	if InstantCloseMenuAction.JustPressed(1) {
		InGame.Activate(1)
	}
	if InstantOpenMenuAction.JustPressed(1) {
		InMenu.Activate(1)
	}
	hidden = input.Cursor.Hidden()

	openmenu = OpenMenuAction.Pressed(1)
	closemenu = CloseMenuAction.Pressed(1)
	instopenmenu = InstantOpenMenuAction.Pressed(1)
	instclosemenu = InstantCloseMenuAction.Pressed(1)
	inventory = InventoryAction.Pressed(1)
	options = OptionsAction.Pressed(1)
	jump = JumpAction.Pressed(1)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() error {
	screen.Clear(0)

	cursor.Locate(2, 12)
	cursor.Color = DarkBlue - 1

	cursor.Printf("cursor position:%6d,%6d\n", mousepos.C, mousepos.R)
	cursor.Printf("   cursor delta:%+6d,%+6d\n", mousedelta.C, mousedelta.R)
	cursor.Printf("     visibility:   ")
	if hidden {
		color(true)
		cursor.Printf("HIDDEN\n")
	} else {
		color(false)
		cursor.Printf("shown\n")
	}

	cursor.Println()
	color(false)

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

func color(p bool) {
	if p {
		cursor.Color = LightGreen - 1
	} else {
		cursor.Color = DarkBlue - 1
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Show() {
	fmt.Printf("%v: show\n", cozely.GameTime())
}

func (loop) Hide() {
	fmt.Printf("%v: hide\n", cozely.GameTime())
}

func (loop) Resize() {
	s := cozely.WindowSize()
	fmt.Printf("%v: resize %dx%d\n", cozely.GameTime(), s.C, s.R)
}

func (loop) Focus() {
	fmt.Printf("%v: focus\n", cozely.GameTime())
}

func (loop) Unfocus() {
	fmt.Printf("%v: unfocus\n", cozely.GameTime())
}

func (loop) Quit() {
	fmt.Printf("%v: quit\n", cozely.GameTime())
	cozely.Stop()
}

////////////////////////////////////////////////////////////////////////////////
