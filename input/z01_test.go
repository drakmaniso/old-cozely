// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input_test

////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/window"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quitAct      = input.Digital("Quit")
	inventoryAct = input.Digital("Inventory")
	optionsAct   = input.Digital("Options")
	closeAct     = input.Digital("Close Menu")
	instCloseAct = input.Digital("Instant Close Menu")
	jumpAct      = input.Digital("Jump")
	openAct      = input.Digital("Open Menu")
	instOpenAct  = input.Digital("Instant Open Menu")
	triggerAct   = input.Unipolar("Trigger")
	positionAct  = input.Analog("Position")
	cursorAct    = input.Cursor("Cursor")
	deltaAct     = input.Delta("Delta")
)

var (
	inMenu = input.Context("Menu", quitAct, closeAct, instCloseAct,
		inventoryAct, optionsAct, triggerAct, positionAct, cursorAct, deltaAct)

	inGame = input.Context("Game", quitAct, openAct, instOpenAct,
		inventoryAct, jumpAct, triggerAct, positionAct, cursorAct, deltaAct)
)

var (
	bindings = input.Bindings{
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

func TestTest1(t *testing.T) {
	defer cozely.Recover()

	pixel.SetZoom(3)

	input.Load(bindings)
	err := cozely.Run(loop1{})
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop1 struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Enter() {
	inMenu.Activate(0)
}

func (loop1) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) React() {
	if jumpAct.Started(0) {
		println(" Just Pressed: *JUMP*")
	}
	if jumpAct.Stopped(0) {
		println("Just Released: (jump)")
	}

	if closeAct.Stopped(0) {
		inGame.Activate(0)
		input.GrabMouse(true)
	}
	if openAct.Stopped(0) {
		inMenu.Activate(0)
		input.GrabMouse(false)
	}

	if instCloseAct.Started(0) {
		inGame.Activate(0)
		input.GrabMouse(true)
	}
	if instOpenAct.Started(0) {
		inMenu.Activate(0)
		input.GrabMouse(false)
	}

	if quitAct.Started(0) {
		cozely.Stop(nil)
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Update() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Render() {
	pixel.Clear(1)

	cur := pixel.Cursor{}

	cur.Locate(pixel.XY{2, 12})
	cur.Text(7, pixel.Monozela10)

	cur.Println()
	changecolor(&cur, false)

	changecolor(&cur, inMenu.Active(1))
	cur.Printf("  Menu: ")
	changecolor(&cur, optionsAct.Ongoing(0))
	cur.Print("Options(O/L.C.) ")
	changecolor(&cur, closeAct.Ongoing(0))
	cur.Print("CloseMenu(ENTER) ")
	changecolor(&cur, instCloseAct.Ongoing(0))
	cur.Print("InstantCloseMenu(MOUSE RIGHT) ")
	cur.Println(" ")

	changecolor(&cur, inGame.Active(1))
	cur.Printf("  Game: ")
	changecolor(&cur, jumpAct.Ongoing(0))
	cur.Print("Jump(SPACE/L.C.) ")
	changecolor(&cur, openAct.Ongoing(0))
	cur.Print("OpenMenu(ENTER) ")
	changecolor(&cur, instOpenAct.Ongoing(0))
	cur.Print("InstantOpenMenu(MOUSE RIGHT) ")
	cur.Println(" ")

	changecolor(&cur, false)
	cur.Printf("  Both: ")
	changecolor(&cur, inventoryAct.Ongoing(0))
	cur.Println("Inventory(I/TAB) ")

	changecolor(&cur, false)
	cur.Println()
	cur.Printf(" Trigger = % 12.6f\n", triggerAct.Value(0))
	p := positionAct.XY(0)
	cur.Printf("Position = % 12.6f, % 12.6f\n", p.X, p.Y)
	c := cursorAct.XY(0)
	cur.Printf("  Cursor = % 12.6f, % 12.6f", c.X, c.Y)
	if input.MouseGrabbed() {
		changecolor(&cur, true)
		cur.Printf(" (mouse GRABBED)\n")
	} else {
		changecolor(&cur, false)
		cur.Printf(" (mouse not grabbed)\n")
	}
	d := deltaAct.XY(0)
	cur.Printf("   Delta = %+12.6f, %+12.6f\n", d.X, d.Y)

	//TODO:
	pixel.MouseCursor.Paint(0, pixel.ToCanvas(window.XYof(c)))
}

func changecolor(cur *pixel.Cursor, p bool) {
	if p {
		cur.Text(14, pixel.Monozela10)
	} else {
		cur.Text(7, pixel.Monozela10)
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
	s := window.Size()
	fmt.Printf("%v: resize %dx%d\n", cozely.GameTime(), s.X, s.Y)
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
