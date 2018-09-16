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
	quitAct      = input.Button("Quit")
	inventoryAct = input.Button("Inventory")
	optionsAct   = input.Button("Options")
	closeAct     = input.Button("Close Menu")
	instCloseAct = input.Button("Instant Close Menu")
	jumpAct      = input.Button("Jump")
	openAct      = input.Button("Open Menu")
	instOpenAct  = input.Button("Instant Open Menu")
	triggerAct   = input.HalfAxis("Trigger")
	positionAct  = input.DualAxis("Position")
	cursorAct    = input.Cursor("Cursor")
	deltaAct     = input.Delta("Delta")
)

var (
	inMenu = input.Context("Menu", quitAct, closeAct, instCloseAct,
		inventoryAct, optionsAct, triggerAct, positionAct, cursorAct, deltaAct)

	inGame = input.Context("Game", quitAct, openAct, instOpenAct,
		inventoryAct, jumpAct, triggerAct, positionAct, cursorAct, deltaAct)
)

////////////////////////////////////////////////////////////////////////////////

func TestTest1(t *testing.T) {
	defer cozely.Recover()

	pixel.SetZoom(3)

	err := cozely.Run(loop1{})
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop1 struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop1) Enter() {
	inMenu.Activate()
}

func (loop1) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop1) React() {
	if jumpAct.Pushed() {
		println(" Just Pressed: *JUMP*")
	}
	if jumpAct.Released() {
		println("Just Released: (jump)")
	}

	if closeAct.Released() {
		inGame.Activate()
		input.GrabMouse(true)
	}
	if openAct.Released() {
		inMenu.Activate()
		input.GrabMouse(false)
	}

	if instCloseAct.Pushed() {
		inGame.Activate()
		input.GrabMouse(true)
	}
	if instOpenAct.Pushed() {
		inMenu.Activate()
		input.GrabMouse(false)
	}

	if quitAct.Pushed() {
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

	cur.Locate(0, pixel.XY{2, 12})
	cur.Style(7, pixel.Monozela10)

	cur.Println()
	changecolor(&cur, false)

	changecolor(&cur, inMenu.ActiveOn(1))
	cur.Printf("  Menu: ")
	changecolor(&cur, optionsAct.Pressed())
	cur.Print("Options(O/L.C.) ")
	changecolor(&cur, closeAct.Pressed())
	cur.Print("CloseMenu(ENTER) ")
	changecolor(&cur, instCloseAct.Pressed())
	cur.Print("InstantCloseMenu(MOUSE RIGHT) ")
	cur.Println(" ")

	changecolor(&cur, inGame.ActiveOn(1))
	cur.Printf("  Game: ")
	changecolor(&cur, jumpAct.Pressed())
	cur.Print("Jump(SPACE/L.C.) ")
	changecolor(&cur, openAct.Pressed())
	cur.Print("OpenMenu(ENTER) ")
	changecolor(&cur, instOpenAct.Pressed())
	cur.Print("InstantOpenMenu(MOUSE RIGHT) ")
	cur.Println(" ")

	changecolor(&cur, false)
	cur.Printf("  Both: ")
	changecolor(&cur, inventoryAct.Pressed())
	cur.Println("Inventory(I/TAB) ")

	changecolor(&cur, false)
	cur.Println()
	cur.Printf(" Trigger = % 12.6f\n", triggerAct.Value())
	p := positionAct.XY()
	cur.Printf("Position = % 12.6f, % 12.6f\n", p.X, p.Y)
	c := cursorAct.XY()
	cur.Printf("  Cursor = % 12d, % 12d", c.X, c.Y)
	if input.MouseGrabbed() {
		changecolor(&cur, true)
		cur.Printf(" (mouse GRABBED)\n")
	} else {
		changecolor(&cur, false)
		cur.Printf(" (mouse not grabbed)\n")
	}
	d := deltaAct.XY()
	cur.Printf("   Delta = %+12.6f, %+12.6f\n", d.X, d.Y)

	changecolor(&cur, false)
	cur.Locate(0, pixel.XY{2, pixel.Resolution().Y - 14})
	dv := input.CurrentDevice()
	cur.Printf("Current device: %d = %s", dv, dv.Name())

	//TODO:
	pixel.MouseCursor.Paint(0, pixel.XYof(c))
}

func changecolor(cur *pixel.Cursor, p bool) {
	if p {
		cur.Style(14, pixel.Monozela10)
	} else {
		cur.Style(7, pixel.Monozela10)
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
