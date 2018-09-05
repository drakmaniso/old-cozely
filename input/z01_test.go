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
	quit                   = input.Digital("Quit")
	InventoryAction        = input.Digital("Inventory")
	OptionsAction          = input.Digital("Options")
	CloseMenuAction        = input.Digital("Close Menu")
	InstantCloseMenuAction = input.Digital("Instant Close Menu")
	JumpAction             = input.Digital("Jump")
	OpenMenuAction         = input.Digital("Open Menu")
	InstantOpenMenuAction  = input.Digital("Instant Open Menu")
	trigger                = input.Unipolar("Trigger")
	position               = input.Analog("Position")
	cursor                 = input.Cursor("Cursor")
	delta                  = input.Delta("Delta")
)

var (
	InMenu = input.Context("Menu", quit,
		CloseMenuAction, InstantCloseMenuAction, InventoryAction, OptionsAction,
		trigger, position, cursor, delta)

	InGame = input.Context("Game", quit,
		OpenMenuAction, InstantOpenMenuAction, InventoryAction, JumpAction,
		trigger, position, cursor, delta)
)

var (
	Bindings = input.Bindings{
		"Menu": {
			"Quit":               {"Escape", "Button Back"},
			"Close Menu":         {"Enter", "Button Start"},
			"Instant Close Menu": {"Mouse Right", "Button B"},
			"Inventory":          {"I", "Button Y", "Mouse Scroll Up"},
			"Options":            {"O", "Mouse Left"},
			"Trigger":            {"Left Trigger", "Right Trigger", "T", "Button X"},
			"Position":           {"Mouse", "Left Stick", "Right Stick"},
			"Cursor":             {"Mouse", "Left Stick", "Right Stick"},
			"Delta":              {"Mouse", "Left Stick", "Right Stick"},
		},
		"Game": {
			"Quit":              {"Escape", "Button Back"},
			"Open Menu":         {"Enter", "Button Start"},
			"Instant Open Menu": {"Mouse Right", "Button B"},
			"Inventory":         {"Tab", "Button Y"},
			"Jump":              {"Space", "Mouse Left", "Button A"},
			"Trigger":           {"Right Trigger"},
			"Position":          {"Mouse", "Right Stick"},
			"Cursor":            {"Mouse", "Right Stick"},
			"Delta":             {"Mouse", "Right Stick"},
		},
	}
)

////////////////////////////////////////////////////////////////////////////////

var hidden bool
var openmenu, closemenu, instopenmenu, instclosemenu, inventory, options, jump bool

var triggerval float32
var positionval, cursorval, deltaval coord.XY

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	defer cozely.Recover()

	pixel.SetZoom(3)

	input.Load(Bindings)
	err := cozely.Run(loop1{})
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop1 struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Enter() {
	msx.Palette.Activate()
	InMenu.Activate(0)
}

func (loop1) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) React() {
	if JumpAction.Started(0) {
		println(" Just Pressed: *JUMP*")
	}
	if JumpAction.Stopped(0) {
		println("Just Released: (jump)")
	}

	if CloseMenuAction.Stopped(0) {
		InGame.Activate(0)
		input.GrabMouse(true)
	}
	if OpenMenuAction.Stopped(0) {
		InMenu.Activate(0)
		input.GrabMouse(false)
	}

	if InstantCloseMenuAction.Started(0) {
		InGame.Activate(0)
		input.GrabMouse(true)
	}
	if InstantOpenMenuAction.Started(0) {
		InMenu.Activate(0)
		input.GrabMouse(false)
	}

	openmenu = OpenMenuAction.Ongoing(0)
	closemenu = CloseMenuAction.Ongoing(0)
	instopenmenu = InstantOpenMenuAction.Ongoing(0)
	instclosemenu = InstantCloseMenuAction.Ongoing(0)
	inventory = InventoryAction.Ongoing(0)
	options = OptionsAction.Ongoing(0)
	jump = JumpAction.Ongoing(0)

	triggerval = trigger.Value(0)
	positionval = position.XY(0)
	cursorval = cursor.XY(0)
	deltaval = delta.XY(0)

	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Update() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Render() {
	pixel.Clear(msx.DarkBlue)

	pixel.Locate(coord.CR{2, 12})
	pixel.Text(msx.White, pixel.Monozela10)

	pixel.Println()
	changecolor(false)

	changecolor(InMenu.Active(1))
	pixel.Printf("  Menu: ")
	changecolor(options)
	pixel.Print("Options(O/L.C.) ")
	changecolor(closemenu)
	pixel.Print("CloseMenu(ENTER) ")
	changecolor(instclosemenu)
	pixel.Print("InstantCloseMenu(MOUSE RIGHT) ")
	pixel.Println(" ")

	changecolor(InGame.Active(1))
	pixel.Printf("  Game: ")
	changecolor(jump)
	pixel.Print("Jump(SPACE/L.C.) ")
	changecolor(openmenu)
	pixel.Print("OpenMenu(ENTER) ")
	changecolor(instopenmenu)
	pixel.Print("InstantOpenMenu(MOUSE RIGHT) ")
	pixel.Println(" ")

	changecolor(false)
	pixel.Printf("  Both: ")
	changecolor(inventory)
	pixel.Println("Inventory(I/TAB) ")

	changecolor(false)
	pixel.Println()
	pixel.Printf(" Trigger = % 12.6f\n", triggerval)
	pixel.Printf("Position = % 12.6f, % 12.6f\n", positionval.X, positionval.Y)
	pixel.Printf("  Cursor = % 12.6f, % 12.6f", cursorval.X, cursorval.Y)
	if input.MouseGrabbed() {
		changecolor(true)
		pixel.Printf(" (mouse GRABBED)\n")
	} else {
		changecolor(false)
		pixel.Printf(" (mouse not grabbed)\n")
	}
	pixel.Printf("   Delta = %+12.6f, %+12.6f\n", deltaval.X, deltaval.Y)

	pixel.Paint(pixel.MouseCursor, pixel.ToCanvas(cursorval.CR()))
}

func changecolor(p bool) {
	if p {
		pixel.Text(msx.LightRed, pixel.Monozela10)
	} else {
		pixel.Text(msx.White, pixel.Monozela10)
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
	cozely.Stop(nil)
}

////////////////////////////////////////////////////////////////////////////////
