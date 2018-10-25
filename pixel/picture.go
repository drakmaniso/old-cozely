package pixel

import (
	"errors"
	"image"
	_ "image/png" // png support
	"io"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/atlas"
)

////////////////////////////////////////////////////////////////////////////////

// PictureID is the ID to handle static image assets.
type PictureID uint16

const maxPictureID = 0xFFFF

// NoPicture is returned when a name is not found.
const NoPicture PictureID = 0

var pictures = struct {
	dictionary map[string]PictureID
	atlas      *atlas.Atlas
	name       []string
	mapping    []mapping
	image      []*image.Paletted
	lut        []*color.LUT
}{
	dictionary: map[string]PictureID{},
}

type mapping struct {
	bin       int16
	x, y      int16
	w, h      int16
	leftright int16
	topbottom int16
}

////////////////////////////////////////////////////////////////////////////////

// Picture returns the picture ID corresponding to a name.
//
// Should only be called when the framework is running (since resources are
// loaded when the framework starts).
func Picture(name string) PictureID {
	p, ok := pictures.dictionary[name]
	if ok {
		return p
	}
	return NoPicture //TODO: reserve ID instead?
}

////////////////////////////////////////////////////////////////////////////////

// NewPicture creates a picture from an image.
//
// Must be called *before* running the framework.
func NewPicture(name string, m *image.Paletted, l *color.LUT) PictureID {
	var err error

	if internal.Running {
		setErr(errors.New("pixel picture creation: declarations must happen before starting the framework"))
		return NoPicture
	}

	_, ok := pictures.dictionary[name]
	if ok && name != "" {
		setErr(errors.New(`new picture: name "` + name + `" already taken`))
		return NoPicture
	}

	if len(pictures.mapping) >= maxPictureID {
		setErr(errors.New("pixel picture creation: too many pictures"))
		return NoPicture
	}

	pictures.name = append(pictures.name, name)
	pictures.image = append(pictures.image, m)
	pictures.lut = append(pictures.lut, l)
	pictures.mapping = append(pictures.mapping, mapping{})
	p := PictureID(len(pictures.name) - 1)

	if pictures.lut[p] == nil {
		// Construct the LUT
		//TODO: handle different modes (strict, flexible, lenient...)
		var a int
		pictures.lut[p], a, err = color.ToMaster(pictures.image[p])
		if a != 0 {
			internal.Debug.Printf("WARNING: %d new colors in picture "+pictures.name[p], a)
		}
		if err != nil {
			setErr(internal.Wrap("loading picture "+pictures.name[p], err))
			return NoPicture
		}
	}

	w, h := int16(pictures.image[p].Bounds().Dx()), int16(pictures.image[p].Bounds().Dy())
	if w > 0x7FFF || h > 0x7FFF {
		setErr(errors.New("unable to load image " + pictures.name[p] + ": too large"))
		return NoPicture
	}
	pictures.mapping[p].w, pictures.mapping[p].h = w, h
	pictures.mapping[p].topbottom = 0 //int16(conf.TopBorder)<<8 | int16(conf.BottomBorder)
	pictures.mapping[p].leftright = 0 //int16(conf.LeftBorder)<<8 | int16(conf.RightBorder)

	if name != "" {
		pictures.dictionary[name] = p
	}

	return p
}

////////////////////////////////////////////////////////////////////////////////

// loadPicture is the resource handler for pictures.
func loadPicture(name string, tags []string, ext string, r io.Reader) error {
	if ext != "png" {
		return errors.New(`load picture "` + name + `": format "` + ext + `" not supported`)
	}
	println("Loading picture", name)
	m, _, err := image.Decode(r)
	if err != nil {
		return internal.Wrap("decoding ", err)
	}
	mm, ok := m.(*image.Paletted)
	if !ok {
		return errors.New("impossible to load picture " + name + ": image is not in indexed color format")
	}
	NewPicture(name, mm, nil)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// Size returns the width and height of the picture.
func (p PictureID) Size() XY {
	return XY{pictures.mapping[p].w, pictures.mapping[p].h}
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
			dm.Pix[int(px)+x+w*(int(py)+y)] = uint8(pictures.lut[rect][ci])
		}
	}

	return nil
}

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
