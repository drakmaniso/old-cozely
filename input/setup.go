// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"encoding/json"
	"os"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.InputSetup = setup
}

func setup() error {
	internal.Log.Printf("Input declarations: %d contexts and %d actions",
		len(contexts.name), len(actions.name)) //TODO

	if len(contexts.name) == 0 {
		aa := []Action{}
		for _, a := range actions.list {
			aa = append(aa, a)
		}
		c := Context("Default", aa...)
		c.Activate()
		internal.Log.Printf("(added default context with all actions)")
	}

	//TODO: always start with the first context instead?
	if len(contexts.name) == 1 {
		ContextID(0).Activate()
	}

	f, err := os.Open(internal.Path + "input.json")
	if !os.IsNotExist(err) {
		if err != nil {
			return internal.Wrap(`in configuration file "input.json" opening`, err)
		}
		d := json.NewDecoder(f)
		if err := d.Decode(&bindings); err != nil {
			return internal.Wrap(`in configuration file "input.json" parsing`, err)
		}
	}

	if len(bindings) == 0 {
		bindings = Bindings{
			"Default": {},
		}
		internal.Log.Printf("(added default bindings)")
	}

	for i, c := range contexts.name {
		m, ok := bindings[c]
		if !ok {
			m = map[string][]string{}
			bindings[c] = m
		}
		for _, a := range contexts.actions[i] {
			var b []string
			switch a {
			case Close:
				actions.name["Close"] = Close
				b, ok = m["Close"]
				if !ok || len(b) == 0 {
					m["Close"] = []string{"Escape", "Button B"}
				}
			case Select:
				actions.name["Select"] = Select
				b, ok = m["Select"]
				if !ok || len(b) == 0 {
					m["Select"] = []string{"Space", "Enter", "Button A"}
				}
			case Up:
				actions.name["Up"] = Up
				b, ok = m["Up"]
				if !ok || len(b) == 0 {
					m["Up"] = []string{"Up", "Dpad Up"} //TODO: keypad and left stick support
				}
			case Down:
				actions.name["Down"] = Down
				b, ok = m["Down"]
				if !ok || len(b) == 0 {
					m["Down"] = []string{"Down", "Dpad Down"} //TODO: keypad and left stick support
				}
			case Left:
				actions.name["Left"] = Left
				b, ok = m["Left"]
				if !ok || len(b) == 0 {
					m["Left"] = []string{"Left", "Dpad Left"} //TODO: keypad and left stick support
				}
			case Right:
				actions.name["Right"] = Right
				b, ok = m["Right"]
				if !ok || len(b) == 0 {
					m["Right"] = []string{"Right", "Dpad Right"} //TODO: keypad and left stick support
				}
			case Pointer:
				actions.name["Pointer"] = Pointer
				b, ok = m["Pointer"]
				if !ok || len(b) == 0 {
					m["Pointer"] = []string{"Mouse", "Right Stick"}
				}
			case Click:
				actions.name["Click"] = Click
				b, ok = m["Click"]
				if !ok || len(b) == 0 {
					m["Click"] = []string{"Mouse Left", "Right Trigger"}
				}
			}
		}
	}

	load()

	return nil
}
