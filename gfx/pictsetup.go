// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/png" // Activate PNG support
	"os"
	"path"
	"strings"

	"github.com/drakmaniso/carol/internal/core"
	"github.com/drakmaniso/carol/internal/gpu"
)

//------------------------------------------------------------------------------

func init() {
	c := core.Hook{
		Callback: postSetupHook,
		Context:  "in picture package setup",
	}
	core.PostSetupHooks = append(core.PostSetupHooks, c)
}

func postSetupHook() error {
	err := scan()
	if err != nil {
		return core.Error("while scanning images", err)
	}

	p := core.Path + picturesPath

	f, err := os.Open(p)
	if err != nil {
		return core.Error("while opening images directory", err)
	}
	defer f.Close()

	dn, err := f.Readdirnames(0)
	if err != nil {
		return core.Error("while reading images directory", err)
	}

	data := make([]uint8, totalSize, totalSize)

	addr := uint32(0)
	for _, n := range dn {
		if path.Ext(n) == ".png" {
			s, err := load(picturesPath, n, data, addr)
			if err != nil {
				return err
			}
			addr += s
		}
	}

	gpu.CreatePictureBuffer(data)

	core.Debug.Printf("Loaded %d pictures: %v", len(pictures), pictures)

	return nil
}

//------------------------------------------------------------------------------

func scan() error {
	p := core.Path + picturesPath

	f, err := os.Open(p)
	if err != nil {
		return core.Error("while opening images directory", err)
	}
	defer f.Close()

	dn, err := f.Readdirnames(0)
	if err != nil {
		return core.Error("while reading images directory", err)
	}

	totalSize = uint64(0)
	nb := 0
	for _, n := range dn {
		if path.Ext(n) == ".png" {
			s, err := getSize(picturesPath, n)
			if err != nil {
				return err
			}
			totalSize += s
			nb++
		}
	}

	core.Debug.Printf("Scanned %d pictures: %d bytes (%.1f Mb)", nb, totalSize, float64(totalSize)/(1024.0*1024.0))

	return nil
}

var totalSize uint64

func getSize(dir, filename string) (uint64, error) {
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

func load(dir, filename string, data []uint8, address uint32) (uint32, error) {
	r, err := os.Open(dir + filename)
	if err != nil {
		return 0, core.Error(`opening picture file "`+filename+`"`, err)
	}
	defer r.Close()

	img, _, err := image.Decode(r)
	if err != nil {
		return 0, core.Error("decoding picture file", err)
	}

	pimg, ok := img.(*image.Paletted)
	if !ok {
		return 0, errors.New(`picture file "` + filename + `" not in indexed color format.`)
	}

	// Register the picture

	p := Picture{
		address: address,
		width:   int16(pimg.Rect.Max.X - pimg.Rect.Min.X),
		height:  int16(pimg.Rect.Max.Y - pimg.Rect.Min.Y),
	}
	n := strings.TrimSuffix(filename, ".png")
	pictures[n] = p

	core.Debug.Printf("Add picture '%s': %d == %d", n, len(pimg.Pix), p.width*p.height)

	s := copy(data[address:], pimg.Pix)
	if s != len(pimg.Pix) {
		return 0, fmt.Errorf(`unable to load full data for picture "%s"`, filename)
	}

	return uint32(p.width * p.height), nil
}

//------------------------------------------------------------------------------

const picturesPath = "graphics/pictures/"

//------------------------------------------------------------------------------