// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"image"
	stdcolor "image/color"
	_ "image/png" // Activate PNG support
	"os"
	"path/filepath"

	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/atlas"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

var pictAtlas *atlas.Atlas

////////////////////////////////////////////////////////////////////////////////

func loadAssets() error {
	if internal.Running {
		return errors.New("loading graphics while running not implemented")
	}

	// Scan all pictures

	prects := []uint32{}

	for i := range picturePaths {
		err := PictureID(i).scan(&prects)
		if err != nil {
			//TODO: sticky error instead?
			return err
		}
	}

	// Pack them into a texture atlas

	pictAtlas.Pack(prects, pictSize, pictPut)

	internal.Debug.Printf(
		"Packed %d pictures in %d bins (%.1fkB unused)\n",
		len(prects),
		pictAtlas.BinCount(),
		float64(pictAtlas.Unused())/1024.0,
	)

	// Load all fonts

	frects := []uint32{}

	for i := range fonts {
		err := FontID(i).load(&frects)
		if err != nil {
			//TODO: sticky error instead?
			return err
		}
	}

	// Pack them into the atlas

	fntAtlas.Pack(frects, fntSize, fntPut)

	internal.Debug.Printf(
		"Packed %d fonts in %d bins (%.1fkB unused)\n",
		len(fonts),
		fntAtlas.BinCount(),
		float64(fntAtlas.Unused())/1024.0,
	)

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func (p PictureID) scan(prects *[]uint32) error {
	switch p {
	case noPicture:
		return nil
	case MouseCursor:
		w, h := int16(mousecursor.Bounds().Dx()), int16(mousecursor.Bounds().Dy())
		pictureMap[p].w, pictureMap[p].h = w, h
		*prects = append(*prects, uint32(p))
		return nil
	}

	n := picturePaths[p]
	//TODO: support other image formats?
	path := filepath.FromSlash(internal.Path + n + ".png")
	path, err := filepath.EvalSymlinks(path)
	if err != nil {
		return internal.Wrap("in path while scanning picture", err)
	}

	f, err := os.Open(path)
	if err != nil {
		return internal.Wrap(`while opening image "`+path+`"`, err)
	}
	defer f.Close() //TODO: error handling

	conf, _, err := image.DecodeConfig(f)
	switch err {
	case nil:
	case image.ErrFormat:
		return nil
	default:
		return internal.Wrap("decoding picture file", err)
	}

	//TODO: check for width and height overflow
	w, h := int16(conf.Width), int16(conf.Height)

	_, ok := conf.ColorModel.(stdcolor.Palette)
	if !ok {
		return errors.New("image format not supported")
	}

	pictureMap[p].w, pictureMap[p].h = w, h
	*prects = append(*prects, uint32(p))

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func pictSize(rect uint32) (width, height int16) {
	s := PictureID(rect).Size()
	return s.C, s.R
}

func pictPut(rect uint32, bin int16, x, y int16) {
	pictureMap[PictureID(rect)].bin = bin
	pictureMap[PictureID(rect)].x, pictureMap[PictureID(rect)].y = x, y
}

func pictPaint(rect uint32, dest interface{}) error {
	px, py := pictureMap[PictureID(rect)].x, pictureMap[PictureID(rect)].y
	pw, ph := pictureMap[PictureID(rect)].w, pictureMap[PictureID(rect)].h

	var src image.Image
	if PictureID(rect) == MouseCursor {
		src = &mousecursor
	} else {
		fp := filepath.FromSlash(internal.Path + picturePaths[PictureID(rect)] + ".png")
		f, err := os.Open(fp)
		if err != nil {
			return err
		}
		defer f.Close()
		src, _, err = image.Decode(f)
		if err != nil {
			return err
		}
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

////////////////////////////////////////////////////////////////////////////////
