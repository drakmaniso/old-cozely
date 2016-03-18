// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine

import (
	"time"

	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
)

//------------------------------------------------------------------------------

// Handler implements the game loop.
var Handler interface {
	Update()
	Draw()
	Quit()
}

//------------------------------------------------------------------------------

var defaultHandler DefaultHandler

// DefaultHandler provides a default implementation for all handlers.
type DefaultHandler struct{}

func (dh DefaultHandler) Update() {}
func (dh DefaultHandler) Draw()   {}
func (dh DefaultHandler) Quit()   { Stop() }

func (dh DefaultHandler) KeyDown(l key.Label, p key.Position, timestamp time.Duration) {}
func (dh DefaultHandler) KeyUp(l key.Label, p key.Position, timestamp time.Duration)   {}

func (dh DefaultHandler) MouseMotion(rel geom.IVec2, pos geom.IVec2, timestamp time.Duration) {}
func (dh DefaultHandler) MouseButtonDown(b mouse.Button, clicks int, timestamp time.Duration) {}
func (dh DefaultHandler) MouseButtonUp(b mouse.Button, clicks int, timestamp time.Duration)   {}
func (dh DefaultHandler) MouseWheel(w geom.IVec2, timestamp time.Duration)                    {}

func (dh DefaultHandler) WindowShown(timestamp time.Duration)                 {}
func (dh DefaultHandler) WindowHidden(timestamp time.Duration)                {}
func (dh DefaultHandler) WindowResized(s geom.IVec2, timestamp time.Duration) {}
func (dh DefaultHandler) WindowMinimized(timestamp time.Duration)             {}
func (dh DefaultHandler) WindowMaximized(timestamp time.Duration)             {}
func (dh DefaultHandler) WindowRestored(timestamp time.Duration)              {}
func (dh DefaultHandler) WindowMouseEnter(timestamp time.Duration)            {}
func (dh DefaultHandler) WindowMouseLeave(timestamp time.Duration)            {}
func (dh DefaultHandler) WindowFocusGained(timestamp time.Duration)           {}
func (dh DefaultHandler) WindowFocusLost(timestamp time.Duration)             {}

//------------------------------------------------------------------------------
