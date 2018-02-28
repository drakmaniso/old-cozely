// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"image"
	"image/color"
	_ "image/png" // Activate PNG support
	"os"
	"path/filepath"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/x/atlas"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

var (
	pictFiles []atlas.Image
	pictAtlas *atlas.Atlas
)

//------------------------------------------------------------------------------

func loadAssets() error {
	if internal.Running {
		return errors.New("loading graphics while running not implemented")
	}

	// Scan all pictures

	for i := range picturePaths {
		err := Picture(i).scan()
		if err != nil {
			//TODO: sticky error instead?
			return err
		}
	}

	// Pack them into a texture atlas

	pictAtlas.Pack(pictFiles)

	internal.Debug.Printf(
		"Packed %d pictures in %d bins (%.1fkB unused)\n",
		len(pictFiles),
		pictAtlas.BinCount(),
		float64(pictAtlas.Unused())/1024.0,
	)

	// Load all fonts

	for i := range fonts {
		err := Font(i).load()
		if err != nil {
			//TODO: sticky error instead?
			return err
		}
	}

	// Pack them into the atlas

	fntAtlas.Pack(fntFiles)

	internal.Debug.Printf(
		"Packed %d fonts in %d bins (%.1fkB unused)\n",
		len(fonts),
		fntAtlas.BinCount(),
		float64(fntAtlas.Unused())/1024.0,
	)

	return gl.Err()
}

//------------------------------------------------------------------------------

func (p Picture) scan() error {
	if p == 0 {
		return nil
	}

	n := picturePaths[p]
	//TODO: support other image formats?
	path := filepath.FromSlash(internal.Path + n + ".png")
	path, err := filepath.EvalSymlinks(path)
	if err != nil {
		return internal.Error("in path while scanning picture", err)
	}

	f, err := os.Open(path)
	if err != nil {
		return internal.Error(`while opening image "`+path+`"`, err)
	}
	defer f.Close() //TODO: error handling

	conf, _, err := image.DecodeConfig(f)
	switch err {
	case nil:
	case image.ErrFormat:
		return nil
	default:
		return internal.Error("decoding picture file", err)
	}

	//TODO: check for width and height overflow
	w, h := int16(conf.Width), int16(conf.Height)

	_, ok := conf.ColorModel.(color.Palette)
	if !ok {
		return errors.New("image format not supported")
	}

	pictureMap[p].w, pictureMap[p].h = w, h
	pictFiles = append(pictFiles, imgfile(p))

	return nil
}

//------------------------------------------------------------------------------

type imgfile Picture

func (im imgfile) Size() (width, height int16) {
	s := Picture(im).Size()
	return s.X, s.Y
}

func (im imgfile) Put(bin int16, x, y int16) {
	pictureMap[Picture(im)].bin = bin
	pictureMap[Picture(im)].x, pictureMap[Picture(im)].y = x, y
}

func (im imgfile) Paint(dest interface{}) error {
	px, py := pictureMap[Picture(im)].x, pictureMap[Picture(im)].y
	pw, ph := pictureMap[Picture(im)].w, pictureMap[Picture(im)].h

	fp := filepath.FromSlash(internal.Path + picturePaths[Picture(im)] + ".png")
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	src, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	sm, ok := src.(*image.Paletted)
	if !ok {
		return errors.New("unexpected src argument to imgfile paint method")
	}
	dm, ok := dest.(*image.Paletted)
	if !ok {
		return errors.New("unexpected dest argument to imgfile paint method")
	}
	for y := 0; y < int(ph); y++ {
		for x := 0; x < int(pw); x++ {
			w := dm.Bounds().Dx()
			ci := sm.Pix[x+int(pw)*y]
			dm.Pix[int(px)+x+w*(int(py)+y)] = uint8(ci)
		}
	}

	return nil
}

//------------------------------------------------------------------------------
