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

// NewBox creates a new box from an image.
//
// Must be called *before* running the framework.
func NewBox(name string, m *image.Paletted, l *color.LUT, top, bottom, left, right int) BoxID {
	if internal.Running {
		setErr(errors.New("pixel box: declarations must happen before starting the framework"))
		return BoxID(NoPicture)
	}

	_, ok := pictures.dictionary[name]
	if ok && name != "" {
		setErr(errors.New(`new box "` + name + `": name already taken`))
		return BoxID(NoPicture)
	}

	if top < 0 || top > 15 ||
		bottom < 0 || bottom > 15 ||
		left < 0 || left > 15 ||
		right < 0 || right > 15 {
		setErr(errors.New(`new box "` + name + `": invalid borders`))
		return BoxID(NoPicture)
	}

	p := NewPicture(name, m, l)
	pictures.border[p] = int16(top<<12 | bottom<<8 | left<<4 | right)
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
	NewBox(name, mmm, nil, top, bottom, left, right)
	return nil
}

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
