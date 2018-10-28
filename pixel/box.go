package pixel

import (
	"errors"
	"image"
	_ "image/png" // png support
	"io"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

// BoxID identifies a stretchable image resources (also known as nine-patch or
// nine-slice images).
type BoxID PictureID

var boxes = struct {
	dictionary map[string]BoxID
	borders    []int16
}{
	dictionary: map[string]BoxID{},
}

////////////////////////////////////////////////////////////////////////////////

// Box returns the box ID corresponding to a name.
//
// Should only be called when the framework is running (since resources are
// loaded when the framework starts).
func Box(name string) BoxID {
	return boxes.dictionary[name]
}

////////////////////////////////////////////////////////////////////////////////

// newBox creates a new box from an image.
//
// Must be called *before* running the framework.
func newBox(name string, m *image.Paletted, l *color.LUT) BoxID {
	_, ok := boxes.dictionary[name]
	if ok && name != "" {
		setErr(errors.New(`new box "` + name + `": name already taken`))
		return BoxID(NoPicture)
	}

	b := BoxID(newPicture(name, m, l))
	if name != "" {
		boxes.dictionary[name] = b
	}
	return b
}

////////////////////////////////////////////////////////////////////////////////

// loadBox is the resource handler for boxes.
func loadBox(name string, tags []string, ext string, r io.Reader) error {
	if ext != "png" {
		return errors.New(`load box "` + name + `": format "` + ext + `" not supported`)
	}

	m, _, err := image.Decode(r)
	if err != nil {
		return internal.Wrap("decoding ", err)
	}
	mm, ok := m.(*image.Paletted)
	if !ok {
		return errors.New("impossible to load box " + name + ": image is not in indexed color format")
	}

	// Check the optional tags
	meta := true // always on
	for _, t := range tags {
		switch t {
		case "meta":
			// ignore, already on
		default:
			setErr(errors.New(`load box "` + name + `": invalid tag`))
		}
	}

	// Corners and anchors
	cornTL, cornBR := XY{}, XY{}
	borTL, borBR := XY{}, XY{}
	if meta {
		b := mm.Bounds()
		// Corners
		for x := 1; x < b.Dx()-1; x++ {
			if mm.Pix[mm.PixOffset(x, 0)] != uint8(color.Transparent) {
				break
			}
			cornTL.X++
		}
		for x := b.Dx() - 2; x > 0; x-- {
			if mm.Pix[mm.PixOffset(x, 0)] != uint8(color.Transparent) {
				break
			}
			cornBR.X++
		}
		for y := 1; y < b.Dy()-1; y++ {
			if mm.Pix[mm.PixOffset(0, y)] != uint8(color.Transparent) {
				break
			}
			cornTL.Y++
		}
		for y := b.Dy() - 2; y > 0; y-- {
			if mm.Pix[mm.PixOffset(0, y)] != uint8(color.Transparent) {
				break
			}
			cornBR.Y++
		}
		// Anchors
		for x := 1; x < b.Dx()-1; x++ {
			if mm.Pix[mm.PixOffset(x, b.Dy()-1)] != uint8(color.Transparent) {
				break
			}
			borTL.X++
		}
		for x := b.Dx() - 2; x > 0; x-- {
			if mm.Pix[mm.PixOffset(x, b.Dy()-1)] != uint8(color.Transparent) {
				break
			}
			borBR.X++
		}
		for y := 1; y < b.Dy()-1; y++ {
			if mm.Pix[mm.PixOffset(b.Dx()-1, y)] != uint8(color.Transparent) {
				break
			}
			borTL.Y++
		}
		for y := b.Dy() - 2; y > 0; y-- {
			if mm.Pix[mm.PixOffset(b.Dy()-1, y)] != uint8(color.Transparent) {
				break
			}
			borBR.Y++
		}
		//TODO: check top, left, bottom, right < 256
		b.Min.X++
		b.Min.Y++
		b.Max.X--
		b.Max.Y--
		mm, ok = mm.SubImage(b).(*image.Paletted)
		if !ok {
			return errors.New("unexpected subimage") //TODO:
		}
	}

	b := newBox(name, mm, nil)
	b.setCorners(cornTL, cornBR)
	pictures.topleft[b] = borTL
	pictures.bottomright[b] = borBR
	println(name, borTL.X, borTL.Y, borBR.X, borBR.Y)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// Corners of the box (i.e. the fixed parts of the image, as opposed to the
// stretchable center).
func (b BoxID) Corners() (topLeft, bottomRight XY) {
	topLeft.Y = pictures.corners[b] >> 12
	topLeft.X = (pictures.corners[b] >> 8) & 0xF
	bottomRight.Y = (pictures.corners[b] >> 4) & 0xF
	bottomRight.X = pictures.corners[b] & 0xF
	return topLeft, bottomRight
}

func (b BoxID) setCorners(topLeft, bottomRight XY) error {
	if topLeft.Y < 0 || topLeft.Y > 15 ||
		bottomRight.Y < 0 || bottomRight.Y > 15 ||
		topLeft.X < 0 || topLeft.X > 15 ||
		bottomRight.X < 0 || bottomRight.X > 15 {
		return errors.New(`new box "` + pictures.name[b] + `": invalid borders`)
	}
	pictures.corners[b] = int16(topLeft.Y<<12 | bottomRight.Y<<8 | topLeft.X<<4 | bottomRight.X)
	return nil
}

// Anchors of the box (i.e. the starting and ending points of the box content in
// the image).
func (b BoxID) Anchors() (topLeft, bottomRight XY) {
	return pictures.topleft[b], pictures.bottomright[b]
}

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
