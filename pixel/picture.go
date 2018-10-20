package pixel

import (
	"encoding/json"
	"errors"
	"image"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/resource"
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

var pictures struct {
	atlas   *atlas.Atlas
	path    []string
	mapping []mapping
	image   []*image.Paletted
	lut     []*color.LUT
}

type mapping struct {
	bin       int16
	x, y      int16
	w, h      int16
	leftright int16
	topbottom int16
}

////////////////////////////////////////////////////////////////////////////////

// Picture declares a new picture and returns its ID.
func Picture(path string) PictureID {
	return picture(path, nil, nil)
}

func picture(path string, m *image.Paletted, l *color.LUT) PictureID {
	if internal.Running {
		setErr(errors.New("pixel picture declaration: declarations must happen before starting the framework"))
		return noPicture
	}

	if len(pictures.mapping) >= maxPictureID {
		setErr(errors.New("pixel picture declaration: too many pictures"))
		return noPicture
	}

	pictures.path = append(pictures.path, path)
	pictures.image = append(pictures.image, m)
	pictures.lut = append(pictures.lut, l)
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
	var err error

	conf := struct {
		TopBorder    int8
		BottomBorder int8
		LeftBorder   int8
		RightBorder  int8
	}{}

	if pictures.image[p] == nil {
		// Load the image file

		if resource.Exist(pictures.path[p] + ".json") {
			f, err := resource.Open(pictures.path[p] + ".json")
			if err != nil {
				return internal.Wrap("opening "+pictures.path[p]+".json", err)
			}
			defer f.Close()
			d := json.NewDecoder(f)
			if err := d.Decode(&conf); err != nil {
				return internal.Wrap("decoding "+pictures.path[p]+".json", err)
			}
		}

		switch {
		case resource.Exist(pictures.path[p] + ".9.png"):
			// Load nine-patch image
			f, err := resource.Open(pictures.path[p] + ".9.png")
			if err != nil {
				return internal.Wrap("opening "+pictures.path[p]+".9.png", err)
			}
			defer f.Close()
			m, _, err := image.Decode(f)
			if err != nil {
				return internal.Wrap("decoding "+pictures.path[p]+".9.png", err)
			}
			var ok bool
			mm, ok := m.(*image.Paletted)
			if !ok {
				return errors.New("impossible to load " + pictures.path[p] + ".9.png: image is not in indexed color format")
			}
			r := mm.Bounds()
			for x := 1; x < r.Dx()-1; x++ {
				if mm.Pix[mm.PixOffset(x, 0)] != uint8(color.Transparent) {
					break
				}
				conf.LeftBorder++
			}
			for x := r.Dx() - 2; x > 0; x-- {
				if mm.Pix[mm.PixOffset(x, 0)] != uint8(color.Transparent) {
					break
				}
				conf.RightBorder++
			}
			for y := 1; y < r.Dy()-1; y++ {
				if mm.Pix[mm.PixOffset(0, y)] != uint8(color.Transparent) {
					break
				}
				conf.TopBorder++
			}
			for y := r.Dy() - 2; y > 0; y-- {
				if mm.Pix[mm.PixOffset(0, y)] != uint8(color.Transparent) {
					break
				}
				conf.BottomBorder++
			}
			r.Min.X++
			r.Min.Y++
			r.Max.X--
			r.Max.Y--
			pictures.image[p], ok = mm.SubImage(r).(*image.Paletted)
			if !ok {
				return errors.New("unexpected subimage") //TODO:
			}

		case resource.Exist(pictures.path[p] + ".png"):
			// Load simple image
			f, err := resource.Open(pictures.path[p] + ".png")
			if err != nil {
				return internal.Wrap("opening "+pictures.path[p]+".png", err)
			}
			defer f.Close()
			m, _, err := image.Decode(f)
			if err != nil {
				return internal.Wrap("decoding ", err)
			}
			var ok bool
			pictures.image[p], ok = m.(*image.Paletted)
			if !ok {
				return errors.New("impossible to load " + pictures.path[p] + ".png: image is not in indexed color format")
			}

		default:
			return errors.New("impossible to load " + pictures.path[p] + ": resource not found")
		}
	}

	if pictures.lut[p] == nil {
		// Construct the LUT
		//TODO: handle different modes (strict, flexible, lenient...)
		var a int
		pictures.lut[p], a, err = color.ToMaster(pictures.image[p])
		if a != 0 {
			internal.Debug.Printf("Warning: %d new colors in picture "+pictures.path[p], a)
		}
		if err != nil {
			return internal.Wrap("loading picture "+pictures.path[p], err)
		}
	}

	w, h := int16(pictures.image[p].Bounds().Dx()), int16(pictures.image[p].Bounds().Dy())
	if w > 0x7FFF || h > 0x7FFF {
		return errors.New("unable to load image " + pictures.path[p] + ": too large")
	}
	pictures.mapping[p].w, pictures.mapping[p].h = w, h
	pictures.mapping[p].topbottom = int16(conf.TopBorder)<<8 | int16(conf.BottomBorder)
	pictures.mapping[p].leftright = int16(conf.LeftBorder)<<8 | int16(conf.RightBorder)

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
			dm.Pix[int(px)+x+w*(int(py)+y)] = uint8(pictures.lut[rect][ci])
		}
	}

	return nil
}

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
