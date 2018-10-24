package pixel

import (
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/resource"
)

////////////////////////////////////////////////////////////////////////////////

// Built-in pictures
const (
	noPicture PictureID = iota
	MouseCursor
	Rectangle
	FilledRectangle
	RectangleR1
	FilledRectangleR1
	RectangleR2
	FilledRectangleR2
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
	Picture("builtins/pictures/nopicture")
	Picture("builtins/pictures/cursor")
	Picture("builtins/pictures/rectangle")
	Picture("builtins/pictures/filled_rectangle")
	Picture("builtins/pictures/rectangle_r1")
	Picture("builtins/pictures/filled_rectangle_r1")
	Picture("builtins/pictures/rectangle_r2")
	Picture("builtins/pictures/filled_rectangle_r2")
	Font("builtins/fonts/monozela10")
}

func setup() error {
	return renderer.setup()
}

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
