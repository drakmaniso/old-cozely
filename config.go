// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import (
	"errors"

	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// An Option represents a configuration option used to change some parameters of
// the framework: see Configure.
type Option = func() error

//------------------------------------------------------------------------------

// Configure the framework.
func Configure(c ...Option) error {
	for _, f := range c {
		err := f()
		if err != nil {
			return nil
		}
	}
	return nil
}

//------------------------------------------------------------------------------

// Title sets the title of the game.
func Title(t string) Option {
	return func() error {
		internal.Title = t
		if internal.Running {
			internal.SetWindowTitle(internal.Title)
		}
		return nil
	}
}

//------------------------------------------------------------------------------

// UpdateStep sets the time step between calls to Update.
func UpdateStep(t float64) Option {
	return func() error {
		internal.UpdateStep = t
		return nil
	}
}

//------------------------------------------------------------------------------

// Multisample activate multisampling for the game window. Note that this is
// currently incompatible with the pixel package.
func Multisample(s int32) Option {
	return func() error {
		if internal.Running {
			return errors.New("cannot change window multisampling while running")
		}
		internal.Window.Multisample = s
		return nil
	}
}

//------------------------------------------------------------------------------
