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

	"github.com/drakmaniso/carol/core/atlas"
	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/internal"
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

type imgfile string

func (mf imgfile) Size() (width, height int16) {
	p := pictures[string(mf)]
	return p.width, p.height
}

func (mf imgfile) Put(bin int16, x, y int16) {
	p := pictures[string(mf)]
	p.bin = bin
	p.x, p.y = x, y
}

func (mf imgfile) Paint(dest interface{}) error {
	p := pictures[string(mf)]

	f, err := os.Open(filepath.Join(internal.FilePath, picturesPath, string(mf)) + ".png")
	if err != nil {
		return err
	}
	defer f.Close()
	pm, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	switch m := dest.(type) {

	case *image.NRGBA:
		for y := 0; y < int(p.height); y++ {
			for x := 0; x < int(p.width); x++ {
				c := pm.At(x, y)
				m.Set(int(p.x)+x, int(p.y)+y, c)
			}
		}

	case *image.Paletted:
		pmp := pm.(*image.Paletted)
		for y := 0; y < int(p.height); y++ {
			for x := 0; x < int(p.width); x++ {
				w := m.Bounds().Dx()
				m.Pix[int(p.x)+x+w*(int(p.y)+y)] = pmp.Pix[x+int(p.width)*y]
			}
		}

	default:
		return errors.New("unexpected argument to imgfile paint method")
	}

	return nil
}

//------------------------------------------------------------------------------

func loadAllPictures() error {
	err := filepath.Walk(picturesPath, scan)
	if err != nil {
		return internal.Error("while scanning images", err)
	}

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
			iu, 4*iu/1024, 4*iu/(1024*1024),
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
	indexedTexture.Bind(0)

	// Create the RGBA texture atlas
	w, h = rgbaAtlas.BinSize()
	rgbaTexture = gl.NewTextureArray2D(1, gl.RGBA8, int32(w), int32(h), int32(rgbaAtlas.BinCount()))
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

		// n := internal.FilePath + fmt.Sprintf("packed/%d.png", i)
		// f, err := os.Create(n)
		// if err != nil {
		// 	panic("cannot open output " + n)
		// }
		// err = png.Encode(f, m)
		// if err != nil {
		// 	panic(err)
		// }
		// err = f.Close()
		// if err != nil {
		// 	panic(err)
		// }
	}
	rgbaTexture.Bind(1)

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
	p := Picture{
		//TODO: check for overflow
		width:  int16(conf.Width),
		height: int16(conf.Height),
	}

	switch conf.ColorModel {

	case color.RGBAModel, color.NRGBAModel, color.GrayModel,
		color.Gray16Model, color.RGBA64Model, color.NRGBA64Model:
		p.mode = 2
		pictures[n] = &p
		rgbaFiles = append(rgbaFiles, imgfile(n))

	case color.AlphaModel, color.Alpha16Model:
		return errors.New(`image "` + path + `" color model (16-bit alpha) not yet supported.`)

	default:
		_, ok := conf.ColorModel.(color.Palette)
		if ok {
			p.mode = 1
			pictures[n] = &p
			indexedFiles = append(indexedFiles, imgfile(n))

		} else {
			return errors.New(`image "` + path + `" color model not recognized.`)
		}
	}

	return nil
}

//------------------------------------------------------------------------------
