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
	"strings"

	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/x/atlas"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

var (
	pictFiles []atlas.Image
	pictAtlas *atlas.Atlas
)

var autoPalette = true

//------------------------------------------------------------------------------

// Load creates a picture for each image file found in the provided path (which
// must be a directory).
func Load(path string) error {
	if internal.Running {
		return errors.New("loading graphics while running not implemented")
	}

	path = filepath.FromSlash(internal.Path + path)
	path, err := filepath.EvalSymlinks(path)
	if err != nil {
		return internal.Error("in path while loading graphics", err)
	}

	fi, err := os.Stat(path)
	if err != nil {
		return internal.Error("in path info while loading graphics", err)
	}
	if !fi.IsDir() {
		return errors.New("path for loading graphics is not a directory")
	}

	// Scan all pictures

	err = filepath.Walk(path, scan)
	switch {
	case os.IsNotExist(err):
		return nil
	case err != nil:
		return internal.Error("while scanning images", err)
	}

	// Pack them into atlases

	pictAtlas.Pack(pictFiles)

	iu := pictAtlas.Unused()
	internal.Debug.Printf(
		"Packed %d indexed images in %d bins: %d unused pixels (%d kb, %d Mb)\n",
		len(pictFiles),
		pictAtlas.BinCount(),
		iu, iu/1024, iu/(1024*1024),
	)

	return gl.Err()
}

//------------------------------------------------------------------------------

func scan(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
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

	n := filepath.Base(path)
	if err != nil {
		return err
	}
	n = strings.TrimSuffix(n, filepath.Ext(n))
	n = filepath.ToSlash(n)
	//TODO: check for width and height overflow
	w, h := int16(conf.Width), int16(conf.Height)

	_, ok := conf.ColorModel.(color.Palette)
	if ok {
		newPicture(n, w, h)
		pictFiles = append(pictFiles, imgfile{name: n, path: path})

	} else {
		internal.Debug.Println(`ignoring image: "` + path + `" (color model not supported)`)
		return nil
	}

	return nil
}

//------------------------------------------------------------------------------

type imgfile struct {
	name string
	path string
}

func (im imgfile) Size() (width, height int16) {
	s := pictures[im.name].Size()
	return s.X, s.Y
}

func (im imgfile) Put(bin int16, x, y int16) {
	pictures[im.name].mapTo(bin, x, y)
}

func (im imgfile) Paint(dest interface{}) error {
	p := pictures[im.name]
	_, px, py, pw, ph := p.getMap()

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
	pal, ok := sm.ColorModel().(color.Palette)
	if !ok {
		return errors.New("unable to access src color palette in imgfile paint method")
	}
	dm, ok := dest.(*image.Paletted)
	if !ok {
		return errors.New("unexpected dest argument to imgfile paint method")
	}
	for y := 0; y < int(ph); y++ {
		for x := 0; x < int(pw); x++ {
			w := dm.Bounds().Dx()
			ci := sm.Pix[x+int(pw)*y]
			if autoPalette {
				// Convert image color index to index into current palette
				r, g, b, a := pal[ci].RGBA()
				cc := colour.SRGBA{
					float32(r) / float32(0xFFFF),
					float32(g) / float32(0xFFFF),
					float32(b) / float32(0xFFFF),
					float32(a) / float32(0xFFFF),
				}
				ci = uint8(palette.Request(cc))
			}
			dm.Pix[int(px)+x+w*(int(py)+y)] = uint8(ci)
		}
	}

	return nil
}

//------------------------------------------------------------------------------
