// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

//------------------------------------------------------------------------------

// EmptyLoop implements a default, empty game loop.
//
// It's an empty struct intended to be embedded in the user-defined GameLoop.
type EmptyLoop struct{}

//------------------------------------------------------------------------------

// Enter does nothing
func (EmptyLoop) Enter() error { return nil }

// React does nothing
func (EmptyLoop) React() error { return nil }

// Update does nothing
func (EmptyLoop) Update() error { return nil }

// Draw does nothing
func (EmptyLoop) Draw() error { return nil }

// Leave does nothing
func (EmptyLoop) Leave() error { return nil }

//------------------------------------------------------------------------------

// Resize does nothing.
func (EmptyLoop) Resize() {}

// Hide does nothing.
func (EmptyLoop) Hide() {}

// Show does nothing.
func (EmptyLoop) Show() {}

// Focus does nothing.
func (EmptyLoop) Focus() {}

// Unfocus does nothing.
func (EmptyLoop) Unfocus() {}

// Quit requests the game loop to stop.
func (EmptyLoop) Quit() {
	QuitRequested = true
}

//------------------------------------------------------------------------------
