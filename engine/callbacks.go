// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine

import (
	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
)

//------------------------------------------------------------------------------

// HandleUpdate sets the callback for Update events.
func HandleUpdate(callback func()) {
	handleUpdate = callback
}

var handleUpdate = func() {}

// HandleDraw sets the callback for Draw events.
func HandleDraw(callback func()) {
	handleDraw = callback
}

var handleDraw = func() {}

//------------------------------------------------------------------------------

// HandleQuit sets the callback for Quit events.
func HandleQuit(callback func()) {
	handleQuit = callback
}

var handleQuit = func() {}

//------------------------------------------------------------------------------

// HandleKeyDown sets the callback for KeyDown events.
func HandleKeyDown(callback func(l key.Label, p key.Position, time uint32)) {
	handleKeyDown = callback
}

var handleKeyDown = func(l key.Label, p key.Position, time uint32) {}

// HandleKeyUp sets the callback for KeyUp events.
func HandleKeyUp(callback func(l key.Label, p key.Position, time uint32)) {
	handleKeyUp = callback
}

var handleKeyUp = func(l key.Label, p key.Position, time uint32) {}

//------------------------------------------------------------------------------

// HandleMouseMotion sets the callback for MouseMotion events.
func HandleMouseMotion(callback func(rel geom.IVec2, pos geom.IVec2, b mouse.ButtonState, time uint32)) {
	handleMouseMotion = callback
}

var handleMouseMotion = func(rel geom.IVec2, pos geom.IVec2, b mouse.ButtonState, time uint32) {}

// HandleMouseButtonDown sets the callback for MouseButtonDown events.
func HandleMouseButtonDown(callback func(b mouse.Button, clicks int, pos geom.IVec2, time uint32)) {
	handleMouseButtonDown = callback
}

var handleMouseButtonDown = func(b mouse.Button, clicks int, pos geom.IVec2, time uint32) {}

// HandleMouseButtonUp sets the callback for MouseButtonUp events.
func HandleMouseButtonUp(callback func(b mouse.Button, clicks int, pos geom.IVec2, time uint32)) {
	handleMouseButtonUp = callback
}

var handleMouseButtonUp = func(b mouse.Button, clicks int, pos geom.IVec2, time uint32) {}

// HandleMouseWheel sets the callback for MouseWheel events.
func HandleMouseWheel(callback func(w geom.IVec2, time uint32)) {
	handleMouseWheel = callback
}

var handleMouseWheel = func(w geom.IVec2, time uint32) {}

//------------------------------------------------------------------------------
