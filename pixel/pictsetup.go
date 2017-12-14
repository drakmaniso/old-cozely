// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

//------------------------------------------------------------------------------

import (
	"errors"
	"image"
	"image/color"
	_ "image/png" // Activate PNG support
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/drakmaniso/carol/colour"

	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

var (
	rgbaPixels uint64
	rgbaData   []colour.SRGBA8
	rgbaTBO    gl.BufferTexture

	indexedPixels uint64
	indexedData   []uint8
	indexedTBO    gl.BufferTexture
)

//------------------------------------------------------------------------------

const picturesPath = "graphics/pictures/"

//------------------------------------------------------------------------------

func loadAllPictures() error {
	err := scan()
	if err != nil {
		return internal.Error("while scanning images", err)
	}

	p := internal.Path + picturesPath

	f, err := os.Open(p)
	if err != nil {
		return internal.Error("while opening images directory", err)
	}
	defer f.Close()

	dn, err := f.Readdirnames(0)
	if err != nil {
		return internal.Error("while reading images directory", err)
	}

	rgbaData = make([]colour.SRGBA8, 0, rgbaPixels)
	indexedData = make([]uint8, 0, indexedPixels)

	for _, n := range dn {
		if path.Ext(n) == ".png" {
			err := load(picturesPath + n)
			if err != nil {
				return err
			}
		}
	}

	indexedTBO = gl.NewBufferTexture(indexedData, gl.R8UI, 0)
	indexedTBO.Bind(0) //TODO: move elsewhere
	rgbaTBO = gl.NewBufferTexture(rgbaData, gl.RGBA8, 0)
	rgbaTBO.Bind(1) //TODO: move elsewhere

	internal.Debug.Printf("Loaded %d pictures: %v", len(pictures), pictures)

	return nil
}

//------------------------------------------------------------------------------

func scan() error {
	p := internal.Path + picturesPath

	f, err := os.Open(p)
	if err != nil {
		return internal.Error("while opening images directory", err)
	}
	defer f.Close()

	dn, err := f.Readdirnames(0)
	if err != nil {
		return internal.Error("while reading images directory", err)
	}

	rgbaPixels = uint64(0)
	indexedPixels = uint64(0)
	nb := 0
	for _, n := range dn {
		if path.Ext(n) == ".png" {
			err := countSize(picturesPath + n)
			if err != nil {
				return err
			}
			nb++
		}
	}

	internal.Debug.Printf("Scanned %d pictures: %d bytes (%.1f Mb)", nb, rgbaPixels, float64(rgbaPixels)/(1024.0*1024.0))

	return nil
}

func countSize(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return internal.Error(`opening picture file "`+filename+`"`, err)
	}
	defer r.Close()

	conf, _, err := image.DecodeConfig(r)
	if err != nil {
		return internal.Error("decoding picture file", err)
	}

	s := uint64(conf.Width) * uint64(conf.Height)

	switch conf.ColorModel {

	case color.RGBAModel,
		color.NRGBAModel,
		color.GrayModel,
		color.Gray16Model,
		color.RGBA64Model,
		color.NRGBA64Model:

		rgbaPixels += s

	case color.AlphaModel, color.Alpha16Model:
		return errors.New(`picture file "` + filename + `" color model (16-bit alpha) not yet supported.`)

	default:
		_, ok := conf.ColorModel.(color.Palette)
		if ok {
			indexedPixels += s

		} else {
			return errors.New(`picture file "` + filename + `" color model not recognized.`)
		}
	}

	return nil
}

//------------------------------------------------------------------------------

func load(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return internal.Error(`opening picture file "`+filename+`"`, err)
	}
	defer r.Close()

	img, _, err := image.Decode(r)
	if err != nil {
		return internal.Error("decoding picture file", err)
	}

	w := img.Bounds().Max.X - img.Bounds().Min.X
	h := img.Bounds().Max.Y - img.Bounds().Min.Y

	switch img.ColorModel() {

	case color.RGBAModel,
		color.NRGBAModel,
		color.GrayModel,
		color.Gray16Model,
		color.RGBA64Model,
		color.NRGBA64Model:

		// Copy picture data
		addr := len(rgbaData)
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				r, g, b, a := img.At(x, y).RGBA()
				rgbaData = append(
					rgbaData,
					colour.SRGBA8{
						uint8(r >> 8),
						uint8(g >> 8),
						uint8(b >> 8),
						uint8(a >> 8),
					},
				)
			}
		}
		// Register the picture
		p := Picture{
			address: uint32(addr),
			width:   int16(w),
			height:  int16(h),
			mode:    2,
		}
		n := strings.TrimSuffix(filepath.Base(filename), ".png")
		pictures[n] = p
		internal.Debug.Printf("Added picture '%s': %d x %d = %d -> %v", n, p.width, p.height, len(rgbaData)-addr, p)

	case color.AlphaModel, color.Alpha16Model:
		return errors.New(`picture file "` + filename + `" color model (8-bit or 16-bit alpha) not yet supported.`)

	default:
		pimg, ok := img.(*image.Paletted)
		if ok {
			// return errors.New(`picture file "` + filename + `" color model (indexed) not yet supported.`)
			addr := len(indexedData)
			// Register the picture
			p := Picture{
				address: uint32(addr),
				width:   int16(pimg.Rect.Max.X - pimg.Rect.Min.X),
				height:  int16(pimg.Rect.Max.Y - pimg.Rect.Min.Y),
				mode:    1,
			}
			n := strings.TrimSuffix(filepath.Base(filename), ".png")
			pictures[n] = p
			// Copy picture data
			indexedData = append(indexedData, pimg.Pix...)
			internal.Debug.Printf("Added picture '%s': %d x %d = %d -> %v", n, p.width, p.height, len(indexedData)-addr, p)

		} else {
			return errors.New(`picture file "` + filename + `" color model not recognized.`)
		}
	}

	return nil
}

//------------------------------------------------------------------------------
