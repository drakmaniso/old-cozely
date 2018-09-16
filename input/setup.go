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
		len(contexts.name), len(actions.name)-8) //TODO

	if len(contexts.name) == 0 {
		aa := []Action{}
		for _, a := range actions.name {
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
			case MenuBack:
				b, ok = m["Menu Back"]
				if !ok || len(b) == 0 {
					m["Menu Back"] = []string{"Escape", "Button B"}
				}
			case MenuSelect:
				b, ok = m["Menu Select"]
				if !ok || len(b) == 0 {
					m["Menu Select"] = []string{"Space", "Enter", "Button A"}
				}
			case MenuUp:
				b, ok = m["Menu Up"]
				if !ok || len(b) == 0 {
					m["Menu Up"] = []string{"Up", "Dpad Up"} //TODO: keypad and left stick support
				}
			case MenuDown:
				b, ok = m["Menu Down"]
				if !ok || len(b) == 0 {
					m["Menu Down"] = []string{"Down", "Dpad Down"} //TODO: keypad and left stick support
				}
			case MenuLeft:
				b, ok = m["Menu Left"]
				if !ok || len(b) == 0 {
					m["Menu Left"] = []string{"Left", "Dpad Left"} //TODO: keypad and left stick support
				}
			case MenuRight:
				b, ok = m["Menu Right"]
				if !ok || len(b) == 0 {
					m["Menu Right"] = []string{"Right", "Dpad Right"} //TODO: keypad and left stick support
				}
			case MenuPointer:
				b, ok = m["Menu Pointer"]
				if !ok || len(b) == 0 {
					m["Menu Pointer"] = []string{"Mouse", "Right Stick"}
				}
			case MenuClick:
				b, ok = m["Menu Click"]
				if !ok || len(b) == 0 {
					m["Menu Click"] = []string{"Mouse Left", "Right Trigger"}
				}
			}
		}
	}

	load()

	return nil
}
