// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
)

//------------------------------------------------------------------------------

const charHeight = 11

//------------------------------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		log.Printf("Error: no filename given\n")
		return
	}
	filename := os.Args[1]
	r, err := os.Open(filename)
	if err != nil {
		log.Printf("Error while opening file: %s\n", err)
		return
	}
	defer r.Close()
	img, _, err := image.Decode(r)
	if err != nil {
		log.Printf("Error while decoding file: %s\n", err)
		return
	}
	bounds := img.Bounds()
	if bounds.Size().X != 8*16 || bounds.Size().Y != charHeight*16 {
		log.Printf("Image has the wron size: (%d, %d) instead of (%d, %d)\n", bounds.Size().X, bounds.Size().Y, 16*8, 16*charHeight)
		return
	}

	for a := 0; a < 256; a++ {
		cx := a & 0x0F
		cy := a >> 4
		for y := 0; y < charHeight; y++ {
			var val byte
			line := make([]rune, 8)
			for x := 0; x < 8; x++ {
				r, _, _, _ := img.At(x+cx*8, y+cy*charHeight).RGBA()
				if r != 0 {
					val |= 1 << uint(7-x)
					line[x] = '#'
				} else {
					line[x] = '.'
				}
			}
			fmt.Printf("  0x%02x, // %s\n", val, string(line))
		}
		fmt.Printf("\n")
	}
}

//------------------------------------------------------------------------------
