// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input_test

////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color/palettes/msx"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quit                   = input.Bool("Quit")
	InventoryAction        = input.Bool("Inventory")
	OptionsAction          = input.Bool("Options")
	CloseMenuAction        = input.Bool("Close Menu")
	InstantCloseMenuAction = input.Bool("Instant Close Menu")
	JumpAction             = input.Bool("Jump")
	OpenMenuAction         = input.Bool("Open Menu")
	InstantOpenMenuAction  = input.Bool("Instant Open Menu")
)

var (
	InMenu = input.Context("Menu", quit,
		CloseMenuAction, InstantCloseMenuAction, InventoryAction, OptionsAction)

	InGame = input.Context("Game", quit,
		OpenMenuAction, InstantOpenMenuAction, InventoryAction, JumpAction)
)

var (
	Bindings = input.Bindings{
		"Menu": {
			"Quit":               {"Escape"},
			"Close Menu":         {"Space"},
			"Instant Close Menu": {"Mouse Right", "Enter"},
			"Inventory":          {"I"},
			"Options":            {"O", "Mouse Left"},
		},
		"Game": {
			"Quit":              {"Escape"},
			"Open Menu":         {"Space"},
			"Instant Open Menu": {"Mouse Right", "Enter"},
			"Inventory":         {"Tab"},
			"Jump":              {"Space", "Mouse Left"},
		},
	}
)

////////////////////////////////////////////////////////////////////////////////

var (
	canvas1 = pixel.Canvas(pixel.Zoom(3))
)

var hidden bool
var mousepos, mousedelta coord.CR
var openmenu, closemenu, instopenmenu, instclosemenu, inventory, options, jump bool

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	err := cozely.Run(loop1{})
	if err != nil {
		cozely.ShowError(err)
		return
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop1 struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Enter() error {
	msx.Palette.Activate()
	Bindings.Load()

	InMenu.Activate(0)
	return input.Err()
}

func (loop1) Leave() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop1) React() error {
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

	if quit.JustPressed(1) {
		cozely.Stop()
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Update() error { return nil }

////////////////////////////////////////////////////////////////////////////////

func (loop1) Render() error {
	canvas1.Clear(msx.DarkBlue)

	canvas1.Locate(0, coord.CR{2, 12})
	canvas1.Text(msx.White-1, pixel.Monozela10)

	canvas1.Printf("screen position:%6d,%6d\n", mousepos.C, mousepos.R)
	canvas1.Printf("   screen delta:%+6d,%+6d\n", mousedelta.C, mousedelta.R)
	canvas1.Printf("     visibility:   ")
	if hidden {
		changecolor(true)
		canvas1.Printf("HIDDEN\n")
	} else {
		changecolor(false)
		canvas1.Printf("shown\n")
	}

	canvas1.Println()
	changecolor(false)

	changecolor(InMenu.Active(1))
	canvas1.Printf("  Menu: ")
	changecolor(options)
	canvas1.Print("Options(O/L.C.) ")
	changecolor(closemenu)
	canvas1.Print("CloseMenu(ESC) ")
	changecolor(instclosemenu)
	canvas1.Print("InstantCloseMenu(ENTER/R.C.) ")
	canvas1.Println(" ")

	changecolor(InGame.Active(1))
	canvas1.Printf("  Game: ")
	changecolor(jump)
	canvas1.Print("Jump(SPACE/L.C.) ")
	changecolor(openmenu)
	canvas1.Print("OpenMenu(ESC) ")
	changecolor(instopenmenu)
	canvas1.Print("InstantOpenMenu(ENTER/R.C.) ")
	canvas1.Println(" ")

	changecolor(false)
	canvas1.Printf("  Both: ")
	changecolor(inventory)
	canvas1.Print("Inventory(I/TAB) ")

	canvas1.Display()
	return nil
}

func changecolor(p bool) {
	if p {
		canvas1.Text(msx.LightRed-1, pixel.Monozela10)
	} else {
		canvas1.Text(msx.White-1, pixel.Monozela10)
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Show() {
	fmt.Printf("%v: show\n", cozely.GameTime())
}

func (loop1) Hide() {
	fmt.Printf("%v: hide\n", cozely.GameTime())
}

func (loop1) Resize() {
	s := cozely.WindowSize()
	fmt.Printf("%v: resize %dx%d\n", cozely.GameTime(), s.C, s.R)
}

func (loop1) Focus() {
	fmt.Printf("%v: focus\n", cozely.GameTime())
}

func (loop1) Unfocus() {
	fmt.Printf("%v: unfocus\n", cozely.GameTime())
}

func (loop1) Quit() {
	fmt.Printf("%v: quit\n", cozely.GameTime())
	cozely.Stop()
}

////////////////////////////////////////////////////////////////////////////////
