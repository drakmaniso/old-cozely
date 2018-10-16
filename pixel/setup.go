package pixel

import (
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PixelSetup = setup
	internal.PixelCleanup = renderer.cleanup
	internal.PixelRender = renderer.render

	font("(builtin Monozela10)", &monozela10, &color.Identity)
	picture("(builtin)", &mousecursor, &color.Identity)
	picture("(builtin cursor)", &mousecursor, &color.Identity)
}

func setup() error {
	return renderer.setup()
}

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
