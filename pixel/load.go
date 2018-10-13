// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

func loadAssets() error {
	if internal.Running {
		return errors.New("loading graphics while running not implemented")
	}

	prects := []uint32{} //TODO: move with pictures.image

	// Load all fonts

	c := len(pictures.path)

	for i := range fonts.path {
		err := FontID(i).load(&prects)
		if err != nil {
			//TODO: sticky error instead?
			return err
		}
	}

	internal.Debug.Printf(
		"Loaded %d fonts (%d glyphs)\n",
		len(fonts.path),
		len(pictures.path)-c,
	)

	// Load all pictures

	for i := range pictures.path {
		err := PictureID(i).load(&prects)
		if err != nil {
			//TODO: sticky error instead?
			return err
		}
	}

	// Pack them into a texture atlas

	pictures.atlas.Pack(prects, pictSize, pictPut)

	internal.Debug.Printf(
		"Packed %d pictures in %d bins (%.1fkB unused)\n",
		len(prects),
		pictures.atlas.BinCount(),
		float64(pictures.atlas.Unused())/1024.0,
	)

	return nil
}
