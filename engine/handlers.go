// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine

import (
	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
)

//------------------------------------------------------------------------------

// HandleLoop sets the handler for the game loop.
func HandleLoop(h LoopHandler) {
	loopHandler = h
}

var loopHandler LoopHandler = defaultHandler

// A LoopHandler implements the game loop.
type LoopHandler interface {
	Update()
	Draw()
	Quit()
}

// HandleKey sets the handler for Key events.
func HandleKey(h KeyHandler) {
	keyHandler = h
}

var keyHandler KeyHandler = defaultHandler

// A KeyHandler reacts to key events.
type KeyHandler interface {
	KeyDown(l key.Label, p key.Position, time uint32)
	KeyUp(l key.Label, p key.Position, time uint32)
}

// HandleMouse sets the handler for Mouse events.
func HandleMouse(h MouseHandler) {
	mouseHandler = h
}

var mouseHandler MouseHandler = defaultHandler

// A MouseHandler reacts to mouse events. 
type MouseHandler interface {
	MouseMotion(rel geom.IVec2, pos geom.IVec2, b mouse.ButtonState, time uint32)
	MouseButtonDown(b mouse.Button, clicks int, pos geom.IVec2, time uint32)
	MouseButtonUp(b mouse.Button, clicks int, pos geom.IVec2, time uint32)
	MouseWheel(w geom.IVec2, time uint32)
}

//------------------------------------------------------------------------------

var defaultHandler DefaultHandler

// DefaultHandler provides a default implementation for all handlers.
type DefaultHandler struct{}

func (dh DefaultHandler) Update() {}
func (dh DefaultHandler) Draw()   {}
func (dh DefaultHandler) Quit()   {}

func (dh DefaultHandler) KeyDown(l key.Label, p key.Position, time uint32) {}
func (dh DefaultHandler) KeyUp(l key.Label, p key.Position, time uint32)   {}

func (dh DefaultHandler) MouseMotion(rel geom.IVec2, pos geom.IVec2, b mouse.ButtonState, time uint32) {
}
func (dh DefaultHandler) MouseButtonDown(b mouse.Button, clicks int, pos geom.IVec2, time uint32) {}
func (dh DefaultHandler) MouseButtonUp(b mouse.Button, clicks int, pos geom.IVec2, time uint32)   {}
func (dh DefaultHandler) MouseWheel(w geom.IVec2, time uint32)                                    {}

//------------------------------------------------------------------------------
