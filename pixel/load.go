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

	for n, p := range pictureNames {
		err := scanPicture(n, p)
		if err != nil {
			//TODO: sticky error instead?
			return err
		}
	}

	// Pack them into a texture atlas

	pictAtlas.Pack(pictFiles)

	iu := pictAtlas.Unused()
	internal.Debug.Printf(
		"Packed %d pictures in %d bins: %d unused pixels (%d kb, %d Mb)\n",
		len(pictFiles),
		pictAtlas.BinCount(),
		iu, iu/1024, iu/(1024*1024),
	)

	// Load all fonts

	for n, f := range fontNames {
		err := loadFont(n, f)
		if err != nil {
			//TODO: sticky error instead?
			return err
		}
	}

	return gl.Err()
}

//------------------------------------------------------------------------------

func scanPicture(n string, p Picture) error {
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
	pictFiles = append(pictFiles, imgfile{name: n, path: path})

	return nil
}

//------------------------------------------------------------------------------

type imgfile struct {
	name string
	path string
}

func (im imgfile) Size() (width, height int16) {
	s := pictureNames[im.name].Size()
	return s.X, s.Y
}

func (im imgfile) Put(bin int16, x, y int16) {
	p := pictureNames[im.name]
	pictureMap[p].bin = bin
	pictureMap[p].x, pictureMap[p].y = x, y
}

func (im imgfile) Paint(dest interface{}) error {
	p := pictureNames[im.name]
	px, py := pictureMap[p].x, pictureMap[p].y
	pw, ph := pictureMap[p].w, pictureMap[p].h

	pf, err := os.Open(im.path)
	if err != nil {
		return err
	}
	defer pf.Close()
	src, _, err := image.Decode(pf)
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
