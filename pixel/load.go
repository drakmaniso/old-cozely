// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
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

	"github.com/drakmaniso/glam/x/atlas"
	"github.com/drakmaniso/glam/x/gl"
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

var (
	rgbaFiles   []atlas.Image
	rgbaAtlas   *atlas.Atlas
	rgbaTexture gl.TextureArray2D

	indexedFiles   []atlas.Image
	indexedAtlas   *atlas.Atlas
	indexedTexture gl.TextureArray2D
)

//------------------------------------------------------------------------------

var picturesPath string

func init() {
	picturesPath = filepath.Join(internal.FilePath, "graphics")
}

//------------------------------------------------------------------------------

func loadAllPictures() error {
	// Scan all pictures
	err := filepath.Walk(picturesPath, scan)
	switch {
	case os.IsNotExist(err):
		return nil
	case err != nil:
		return internal.Error("while scanning images", err)
	}

	// Pack them into atlases
	indexedAtlas = atlas.New(1024, 1024)
	rgbaAtlas = atlas.New(1024, 1024)

	indexedAtlas.Pack(indexedFiles)
	rgbaAtlas.Pack(rgbaFiles)

	{
		iu := indexedAtlas.Unused()
		internal.Debug.Printf(
			"Packed %d indexed images in %d bins: %d unused pixels (%d kb, %d Mb)\n",
			len(indexedFiles),
			indexedAtlas.BinCount(),
			iu, iu/1024, iu/(1024*1024),
		)
		ru := rgbaAtlas.Unused()
		internal.Debug.Printf(
			"Packed %d RGBA images in %d bins: %d unused pixels (%d kb, %d Mb)\n",
			len(rgbaFiles),
			rgbaAtlas.BinCount(),
			ru, 4*ru/1024, 4*ru/(1024*1024),
		)
	}

	// Create the indexed texture atlas
	w, h := indexedAtlas.BinSize()
	indexedTexture = gl.NewTextureArray2D(1, gl.R8UI, int32(w), int32(h), int32(indexedAtlas.BinCount()))
	for i := int16(0); i < indexedAtlas.BinCount(); i++ {
		m := image.NewPaletted(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{int(w), int(h)},
		},
			color.Palette{},
		)

		err := indexedAtlas.Paint(i, m)
		if err != nil {
			return err
		}

		indexedTexture.SubImage(0, 0, 0, int32(i), m)
	}
	indexedTexture.Bind(1)

	// Create the RGBA texture atlas
	w, h = rgbaAtlas.BinSize()
	rgbaTexture = gl.NewTextureArray2D(1, gl.SRGBA8, int32(w), int32(h), int32(rgbaAtlas.BinCount()))
	for i := int16(0); i < rgbaAtlas.BinCount(); i++ {
		m := image.NewNRGBA(image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{int(w), int(h)},
		})

		err := rgbaAtlas.Paint(i, m)
		if err != nil {
			return err
		}

		rgbaTexture.SubImage(0, 0, 0, int32(i), m)
	}
	rgbaTexture.Bind(2)

	internal.Debug.Printf("Loaded %d pictures.", len(pictures))

	return nil
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

	fp, err := filepath.Rel(picturesPath, path)
	if err != nil {
		return err
	}
	n := strings.TrimSuffix(fp, filepath.Ext(fp))
	n = filepath.ToSlash(n)
	//TODO: check for width and height overflow
	w, h := int16(conf.Width), int16(conf.Height)

	switch conf.ColorModel {

	case color.RGBAModel, color.NRGBAModel, color.GrayModel,
		color.Gray16Model, color.RGBA64Model, color.NRGBA64Model:
		newPicture(n, FullColor, w, h)
		rgbaFiles = append(rgbaFiles, imgfile{name: n, path: path})

	case color.AlphaModel, color.Alpha16Model:
		return errors.New(`image "` + path + `" color model (16-bit alpha) not yet supported.`)

	default:
		_, ok := conf.ColorModel.(color.Palette)
		if ok {
			newPicture(n, Indexed, w, h)
			indexedFiles = append(indexedFiles, imgfile{name: n, path: path})

		} else {
			return errors.New(`image "` + path + `" color model not recognized.`)
		}
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
	pm, _, err := image.Decode(pf)
	if err != nil {
		return err
	}

	switch dm := dest.(type) {

	case *image.NRGBA:
		for y := 0; y < int(ph); y++ {
			for x := 0; x < int(pw); x++ {
				c := pm.At(x, y)
				dm.Set(int(px)+x, int(py)+y, c)
			}
		}

	case *image.Paletted:
		pmp := pm.(*image.Paletted)
		pal, ok := pmp.ColorModel().(color.Palette)
		if !ok {
			return errors.New("unable to access color palette for image")
		}
		for y := 0; y < int(ph); y++ {
			for x := 0; x < int(pw); x++ {
				w := dm.Bounds().Dx()
				ci := pmp.Pix[x+int(pw)*y]
				if internal.Config.PaletteAuto {
					// Convert image color index to index into current palette
					r, g, b, a := pal[ci].RGBA()
					cc := colour.SRGBA{
						float32(r) / float32(0xFFFF),
						float32(g) / float32(0xFFFF),
						float32(b) / float32(0xFFFF),
						float32(a) / float32(0xFFFF),
					}
					ci = uint8(requestColor(cc))
				}
				dm.Pix[int(px)+x+w*(int(py)+y)] = uint8(ci)
			}
		}

	default:
		return errors.New("unexpected argument to imgfile paint method")
	}

	return nil
}

//------------------------------------------------------------------------------
