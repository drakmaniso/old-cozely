// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"fmt"
	"image"
	stdcolor "image/color"
	_ "image/png" // activate png support
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) != 3 {
		log.Fatal("Usage: font2go <font_file.png> <output.go>")
	}

	n := os.Args[1]
	f, err := os.Open(n)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	mm, ok := m.(*image.Paletted)
	if !ok {
		log.Fatal("font image file should be in indexed color format")
	}
	_, ok = mm.ColorModel().(stdcolor.Palette)
	if !ok {
		log.Fatal("unable to retrieve source image palette")
	}

	o, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	on := strings.TrimSuffix(filepath.Base(os.Args[2]), ".go")

	fmt.Fprintf(o, "package whatever\n\n")
	fmt.Fprintf(o, "import \"image\"\n\n")
	fmt.Fprintf(o, "var %s = image.Paletted{\n", on)
	fmt.Fprintf(
		o, "\tRect: image.Rectangle{Max: image.Point{%d, %d}},\n",
		mm.Bounds().Dx(),
		mm.Bounds().Dy(),
	)
	fmt.Fprintf(o, "\tStride: %d,\n", mm.Bounds().Dx())
	fmt.Fprint(o, "\tPix: []uint8{")
	for i, c := range mm.Pix {
		if i%24 == 0 {
			fmt.Fprint(o, "\n\t\t")
		}
		fmt.Fprintf(o, "%d, ", c)
	}
	fmt.Fprint(o, "\n\t},")
	fmt.Fprint(o, "\n}\n")

}
