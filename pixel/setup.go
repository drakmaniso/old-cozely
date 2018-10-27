package pixel

import (
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/resource"
	"github.com/cozely/cozely/x/atlas"
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
}

func setup() error {
	// Create texture atlas for pictures (boxes and fonts glyphs)

	pictures.atlas = atlas.New(1024, 1024)

	l := make([]uint32, len(pictures.name))
	for i := range pictures.name {
		l[i] = uint32(i)
	}

	pictures.atlas.Pack(l, pictSize, pictPut)

	internal.Debug.Printf(
		"Packed %d pictures in %d bins (%.1fkB unused)\n",
		len(l),
		pictures.atlas.BinCount(),
		float64(pictures.atlas.Unused())/1024.0,
	)

	err := renderer.setup()

	pictures.image = pictures.image[:0]
	pictures.lut = pictures.lut[:0]
	fonts.image = fonts.image[:0]
	fonts.lut = fonts.lut[:0]

	return err
}

func cleanup() error {
	// Pictures
	pictures.dictionary = map[string]PictureID{}
	pictures.name = pictures.name[:0]
	pictures.atlas = nil
	pictures.mapping = pictures.mapping[:0]
	pictures.corners = pictures.corners[:0]
	pictures.topleft = pictures.topleft[:0]
	pictures.bottomright = pictures.bottomright[:0]

	// Boxes
	boxes.dictionary = map[string]BoxID{}

	// Fonts
	fonts.dictionary = map[string]FontID{}
	fonts.name = fonts.name[:0]
	fonts.height = fonts.height[:0]
	fonts.baseline = fonts.baseline[:0]
	fonts.first = fonts.first[:0]

	return renderer.cleanup()
}

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
