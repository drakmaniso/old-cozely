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
)

//------------------------------------------------------------------------------

type Picture struct {
	address uint32
	width   uint16
	height  uint16
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

	totalPictureSize = uint64(0)
	nb := 0
	for _, n := range dn {
		if path.Ext(n) == ".png" {
			s, err := getPictureSize("data/images/", n)
			if err != nil {
				return err
			}
			totalPictureSize += s
			nb++
		}
	}

	core.Debug.Printf("Scanned %d pictures: %d bytes (%.1f Mb)", nb, totalPictureSize, float64(totalPictureSize)/(1024.0*1024.0))

	return nil
}

var totalPictureSize uint64

func getPictureSize(dir, filename string) (uint64, error) {
	r, err := os.Open(dir + filename)
	if err != nil {
		return 0, core.Error(`opening picture file "`+filename+`"`, err)
	}
	defer r.Close()

	conf, _, err := image.DecodeConfig(r)
	if err != nil {
		return 0, core.Error("decoding picture file", err)
	}

	_, ok := conf.ColorModel.(color.Palette)
	if !ok {
		return 0, errors.New(`picture file "` + filename + `" not in indexed color format.`)
	}

	return uint64(conf.Width) * uint64(conf.Height), nil
}

//------------------------------------------------------------------------------

func LoadPictures() error {
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

	addr := uint32(0)
	for _, n := range dn {
		if path.Ext(n) == ".png" {
			s, err := loadPicture("data/images/", n, addr)
			if err != nil {
				return err
			}
			addr += s
		}
	}

	core.Debug.Printf("Loaded %d pictures: %v", len(Pictures), Pictures)

	return nil
}

func loadPicture(dir, filename string, address uint32) (uint32, error) {
	r, err := os.Open(dir + filename)
	if err != nil {
		return 0, core.Error(`opening picture file "`+filename+`"`, err)
	}
	defer r.Close()

	conf, _, err := image.DecodeConfig(r)
	if err != nil {
		return 0, core.Error("decoding picture file", err)
	}

	_, ok := conf.ColorModel.(color.Palette)
	if !ok {
		return 0, errors.New(`picture file "` + filename + `" not in indexed color format.`)
	}

	// Register the picture

	p := Picture{
		address: address,
		width:   uint16(conf.Width),
		height:  uint16(conf.Height),
	}
	n := strings.TrimSuffix(filename, ".png")
	Pictures[n] = p

	return uint32(p.width * p.height), nil
}

//------------------------------------------------------------------------------
