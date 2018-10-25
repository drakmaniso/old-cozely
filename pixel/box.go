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

// BoxID is the ID to handle stretchable image resources (also known as
// nine-patch or nine-slice images).
type BoxID PictureID

var boxes = struct {
	dictionary map[string]BoxID
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

// NewBox creates a new box from an image.
//
// Must be called *before* running the framework.
func NewBox(name string, m *image.Paletted, l *color.LUT, top, bottom, left, right uint8) BoxID {
	_, ok := pictures.dictionary[name]
	if ok && name != "" {
		setErr(errors.New(`new box: name "` + name + `" already taken`))
		return BoxID(NoPicture)
	}

	p := NewPicture(name, m, l)
	pictures.mapping[p].topbottom = int16(top)<<8 | int16(bottom)
	pictures.mapping[p].leftright = int16(left)<<8 | int16(right)
	if name != "" {
		boxes.dictionary[name] = BoxID(p)
	}
	return BoxID(p)
}

////////////////////////////////////////////////////////////////////////////////

// loadBox is the resource handler for boxes.
func loadBox(name string, tags []string, ext string, r io.Reader) error {
	if ext != "png" {
		return errors.New(`load box "` + name + `": format "` + ext + `" not supported`)
	}
	println("Loading box", name)
	m, _, err := image.Decode(r)
	if err != nil {
		return internal.Wrap("decoding ", err)
	}
	mm, ok := m.(*image.Paletted)
	if !ok {
		return errors.New("impossible to load box " + name + ": image is not in indexed color format")
	}
	b := mm.Bounds()
	var left, right, top, bottom int
	for x := 1; x < b.Dx()-1; x++ {
		if mm.Pix[mm.PixOffset(x, 0)] != uint8(color.Transparent) {
			break
		}
		left++
	}
	for x := b.Dx() - 2; x > 0; x-- {
		if mm.Pix[mm.PixOffset(x, 0)] != uint8(color.Transparent) {
			break
		}
		right++
	}
	for y := 1; y < b.Dy()-1; y++ {
		if mm.Pix[mm.PixOffset(0, y)] != uint8(color.Transparent) {
			break
		}
		top++
	}
	for y := b.Dy() - 2; y > 0; y-- {
		if mm.Pix[mm.PixOffset(0, y)] != uint8(color.Transparent) {
			break
		}
		bottom++
	}
	//TODO: check top, left, bottom, right < 256
	b.Min.X++
	b.Min.Y++
	b.Max.X--
	b.Max.Y--
	mmm, ok := mm.SubImage(b).(*image.Paletted)
	if !ok {
		return errors.New("unexpected subimage") //TODO:
	}
	NewBox(
		name, mmm, nil,
		uint8(top), uint8(bottom),
		uint8(left), uint8(right),
	)
	return nil
}

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
