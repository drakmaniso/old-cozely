// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"errors"
	"image"
	"image/color"
	_ "image/png" // Activate PNG support
	"os"
	"path"
	"strings"

	"github.com/drakmaniso/carol/internal/core"
	"github.com/drakmaniso/carol/pixel"
)

//------------------------------------------------------------------------------

type PictureFormat struct {
	size   pixel.Coord
	number int
}

type Picture struct {
	format uint16 // index in PictureFormats
	index  uint16 // index in the OpenGL texture array
}

//------------------------------------------------------------------------------

func ScanPictures() error {
	p := core.Path + "data/images/"

	f, err := os.Open(p)
	if err != nil {
		return core.Error("while opening images directory", err)
	}
	defer f.Close()

	dn, err := f.Readdirnames(0)
	if err != nil {
		return core.Error("while reading images directory", err)
	}

	for _, n := range dn {
		if path.Ext(n) == ".png" {
			err = registerPicture("data/images/", n)
			if err != nil {
				return err
			}
		}
	}

	core.Debug.Printf("Registered %d formats: %v", len(PictureFormats), PictureFormats)
	core.Debug.Printf("Registered %d pictures: %v", len(Pictures), Pictures)

	return nil
}

//------------------------------------------------------------------------------

func registerPicture(dir, filename string) error {
	r, err := os.Open(dir + filename)
	if err != nil {
		return core.Error(`opening picture file "`+filename+`"`, err)
	}
	defer r.Close()

	conf, _, err := image.DecodeConfig(r)
	if err != nil {
		return core.Error("decoding picture file", err)
	}

	s := pixel.Coord{X: int64(conf.Width), Y: int64(conf.Height)}

	_, ok := conf.ColorModel.(color.Palette)
	if !ok {
		return errors.New(`picture file "` + filename + `" not in indexed color format.`)
	}

	// Find or create the picture format

	f := -1
	for i := 0; i < len(PictureFormats); i++ {
		if PictureFormats[i].size == s {
			f = i
			break
		}
	}
	if f == -1 {
		pf := PictureFormat{
			size:   s,
			number: 0,
		}
		PictureFormats = append(PictureFormats, pf)
		f = len(PictureFormats) - 1
	}

	// Register the picture

	p := Picture{
		format: uint16(f),
		index:  uint16(PictureFormats[f].number),
	}
	PictureFormats[f].number++
	n := strings.TrimSuffix(filename, ".png")
	Pictures[n] = p

	return nil
}

//------------------------------------------------------------------------------
