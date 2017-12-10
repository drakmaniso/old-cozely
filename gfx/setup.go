// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal/core"
)

//------------------------------------------------------------------------------

func init() {
	c := core.Hook{
		Callback: preSetupHook,
		Context:  "in gfx pre-Setup hook",
	}
	core.PreSetupHooks = append(core.PreSetupHooks, c)

	c = core.Hook{
		Callback: postDrawHook,
		Context:  "in gfx post-Draw hook",
	}
	core.PostDrawHooks = append(core.PostDrawHooks, c)
}

//------------------------------------------------------------------------------

func preSetupHook() error {
	createScreen(core.Config.FramebufferSize, core.Config.PixelSize)

	err := loadAllPictures()
	if err != nil {
		return err
	}

	return nil
}

//------------------------------------------------------------------------------

func postDrawHook() error {
	if colChanged {
		// gpu.UpdatePaletteBuffer(uint8(0), colours[:])
		colChanged = false
	}

	blitScreen(core.Window.Size)
	return nil
}

//------------------------------------------------------------------------------
