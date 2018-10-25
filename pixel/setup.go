package pixel

import (
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/resource"
)

////////////////////////////////////////////////////////////////////////////////

func init() {
	internal.PixelSetup = setup
	internal.PixelCleanup = cleanup
	internal.PixelRender = renderer.render

	err := resource.Pack(builtins)
	if err != nil {
		setErr(err)
	}
	resource.Handle("picture", loadPicture)
	resource.Handle("box", loadBox)
	resource.Handle("font", loadFont)
	Font("builtins/default")
}

func setup() error {
	return renderer.setup()
}

func cleanup() error {
	// Pictures
	pictures.dictionary = map[string]PictureID{}
	pictures.atlas = nil
	pictures.mapping = pictures.mapping[:0]

	// Boxes
	boxes.dictionary = map[string]BoxID{}

	// Fonts
	fonts.dictionary = map[string]FontID{}
	fonts.name = fonts.name[:0]
	fonts.height = fonts.height[:0]
	fonts.baseline = fonts.baseline[:0]
	fonts.basecolor = fonts.basecolor[:0]
	fonts.first = fonts.first[:0]
	fonts.image = fonts.image[:0]
	fonts.lut = fonts.lut[:0]

	return renderer.cleanup()
}

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
