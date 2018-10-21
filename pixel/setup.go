package pixel

import (
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/resource"
)

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PixelSetup = setup
	internal.PixelCleanup = renderer.cleanup
	internal.PixelRender = renderer.render

	err := resource.Pack(builtins)
	if err != nil {
		setErr(err)
	}
	font("builtins/fonts/monozela10", nil, &color.Identity)
	picture("builtins/pictures/cursor", nil, &color.Identity) //TODO:
	picture("builtins/pictures/cursor", nil, &color.Identity)
}

func setup() error {
	return renderer.setup()
}

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
