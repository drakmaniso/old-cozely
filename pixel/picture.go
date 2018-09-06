// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"image"
	"os"
	"path/filepath"

	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/atlas"
)

////////////////////////////////////////////////////////////////////////////////

// PictureID is the ID to handle static image assets.
type PictureID uint16

const (
	maxPictureID = 0xFFFF
	noPicture    = PictureID(0)
)

// MouseCursor is a small picture that can be used as mouse cursor.
const MouseCursor = PictureID(1)

var pictures = struct {
	atlas   *atlas.Atlas
	path    []string
	mapping []mapping
	image   []*image.Paletted
}{
	path:    []string{"", ""},
	mapping: []mapping{{}, {}},
	image:   []*image.Paletted{nil, nil},
}

type mapping struct {
	bin  int16
	x, y int16
	w, h int16
}

////////////////////////////////////////////////////////////////////////////////

// Picture declares a new picture and returns its ID.
func Picture(path string) PictureID {
	if internal.Running {
		setErr(errors.New("pixel picture declaration: declarations must happen before starting the framework"))
		return noPicture
	}

	if len(pictures.mapping) >= maxPictureID {
		setErr(errors.New("pixel picture declaration: too many pictures"))
		return noPicture
	}

	pictures.path = append(pictures.path, path)
	pictures.image = append(pictures.image, nil)
	pictures.mapping = append(pictures.mapping, mapping{})
	return PictureID(len(pictures.path) - 1)
}

// picture declares a new picture (from an image) and returns its ID.
func picture(img *image.Paletted) PictureID {
	if internal.Running {
		setErr(errors.New("pixel picture declaration: declarations must happen before starting the framework"))
		return noPicture
	}

	if len(pictures.mapping) >= maxPictureID {
		setErr(errors.New("pixel picture declaration: too many pictures"))
		return noPicture
	}

	pictures.path = append(pictures.path, "")
	pictures.image = append(pictures.image, img)
	pictures.mapping = append(pictures.mapping, mapping{})
	return PictureID(len(pictures.path) - 1)
}

////////////////////////////////////////////////////////////////////////////////

// Size returns the width and height of the picture.
func (p PictureID) Size() XY {
	return XY{pictures.mapping[p].w, pictures.mapping[p].h}
}

////////////////////////////////////////////////////////////////////////////////

func (p PictureID) load(prects *[]uint32) error {
	switch p {
	case noPicture:
		//TODO: add mapping
		return nil
	case MouseCursor:
		w, h := int16(mousecursor.Bounds().Dx()), int16(mousecursor.Bounds().Dy())
		//TODO: add mapping
		pictures.mapping[p].w, pictures.mapping[p].h = w, h
		pictures.image[p] = &mousecursor
		*prects = append(*prects, uint32(p))
		return nil
	}

	if pictures.image[p] != nil {
		//TODO: check for width and height overflow
		w, h := int16(pictures.image[p].Bounds().Dx()),
			int16(pictures.image[p].Bounds().Dy())

		pictures.mapping[p].w, pictures.mapping[p].h = w, h
		*prects = append(*prects, uint32(p))

		return nil
	}

	n := pictures.path[p]
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

	img, _, err := image.Decode(f)
	switch err {
	case nil:
	case image.ErrFormat:
		return nil
	default:
		return internal.Wrap("decoding picture file", err)
	}

	m, ok := img.(*image.Paletted)
	if !ok {
		return errors.New("impossible to load picture " + path + " (color model not supported)")
	}

	//TODO: check for width and height overflow
	w, h := int16(m.Bounds().Dx()), int16(m.Bounds().Dy())

	pictures.mapping[p].w, pictures.mapping[p].h = w, h
	pictures.image[p] = m
	*prects = append(*prects, uint32(p))

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func pictSize(rect uint32) (width, height int16) {
	s := PictureID(rect).Size()
	return s.X, s.Y
}

func pictPut(rect uint32, bin int16, x, y int16) {
	pictures.mapping[PictureID(rect)].bin = bin
	pictures.mapping[PictureID(rect)].x, pictures.mapping[PictureID(rect)].y = x, y
}

func pictPaint(rect uint32, dest interface{}) error {
	px, py := pictures.mapping[PictureID(rect)].x, pictures.mapping[PictureID(rect)].y
	pw, ph := pictures.mapping[PictureID(rect)].w, pictures.mapping[PictureID(rect)].h

	sm := pictures.image[PictureID(rect)]
	dm, ok := dest.(*image.Paletted)
	if !ok {
		return errors.New("unexpected dest argument to imgfile paint method")
	}
	for y := 0; y < int(ph); y++ {
		for x := 0; x < int(pw); x++ {
			w := dm.Bounds().Dx()
			ci := sm.Pix[x+y*pictures.image[rect].Stride]
			dm.Pix[int(px)+x+w*(int(py)+y)] = uint8(ci)
		}
	}

	return nil
}
