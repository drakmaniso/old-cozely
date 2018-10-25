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

	// Load all fonts

	// c := len(pictures.name)

	// internal.Debug.Print("Loading fonts...")

	// for i := range fonts.name {
	// 	err := FontID(i).load()
	// 	if err != nil {
	// 		//TODO: sticky error instead?
	// 		return err
	// 	}
	// }

	// internal.Debug.Printf(
	// 	"Loaded %d fonts (%d glyphs)\n",
	// 	len(fonts.name),
	// 	len(pictures.name)-c,
	// )

	// Load all pictures

	// internal.Debug.Print("Loading pictures...")

	// for i := range pictures.name {
	// 	err := PictureID(i).load()
	// 	if err != nil {
	// 		//TODO: sticky error instead?
	// 		return err
	// 	}
	// }

	// Pack them into a texture atlas

	prects := make([]uint32, len(pictures.name))
	for i := range pictures.name {
		prects[i] = uint32(i)
	}

	pictures.atlas.Pack(prects, pictSize, pictPut)

	internal.Debug.Printf(
		"Packed %d pictures in %d bins (%.1fkB unused)\n",
		len(prects),
		pictures.atlas.BinCount(),
		float64(pictures.atlas.Unused())/1024.0,
	)

	return nil
}

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
