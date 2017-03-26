// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/drakmaniso/glam/formats/obj"
	"github.com/drakmaniso/glam"
)

//------------------------------------------------------------------------------

func main() {
	path := filepath.Dir(os.Args[0]) + "/"
	f, err := os.Open(path + "../shared/cube.obj")
	if err != nil {
		glam.ShowError("opening obj file", err)
		return
	}
	obj.Parse(f, builder{})
}

//------------------------------------------------------------------------------

type builder struct{ obj.DefaultBuilder }

func (op builder) Comment(txt string) error {
	fmt.Printf("%s\n", txt)
	return nil
}

func (op builder) O(name string) error {
	fmt.Printf("o %s\n", name)
	return nil
}

func (op builder) G(names ...string) error {
	fmt.Printf("g")
	for _, n := range names {
		fmt.Printf(" %s", n)
	}
	fmt.Printf("\n")
	return nil
}

func (op builder) S(group int) error {
	fmt.Printf("s %d\n", group)
	return nil
}

func (op builder) LOD(level int) error {
	fmt.Printf("lod %d\n", level)
	return nil
}

func (op builder) V(coords ...float32) error {
	fmt.Printf("v")
	for _, v := range coords {
		fmt.Printf(" %f", v)
	}
	fmt.Printf("\n")
	return nil
}

func (op builder) VT(coords ...float32) error {
	fmt.Printf("vt")
	for _, v := range coords {
		fmt.Printf(" %f", v)
	}
	fmt.Printf("\n")
	return nil
}

func (op builder) VN(coords ...float32) error {
	fmt.Printf("vn")
	for _, v := range coords {
		fmt.Printf(" %f", v)
	}
	fmt.Printf("\n")
	return nil
}

func (op builder) F(verts ...obj.Indices) error {
	fmt.Printf("f")
	for _, v := range verts {
		switch {
		case v.TexCoord != 0 && v.Normal != 0:
			fmt.Printf(" %d/%d/%d", v.Vertex, v.TexCoord, v.Normal)
		case v.TexCoord != 0:
			fmt.Printf(" %d/%d", v.Vertex, v.TexCoord)
		case v.Normal != 0:
			fmt.Printf(" %d//%d", v.Vertex, v.Normal)
		default:
			fmt.Printf(" %d", v.Vertex)
		}
	}
	fmt.Printf("\n")
	return nil
}

func (op builder) P(verts ...obj.Indices) error {
	fmt.Printf("p")
	for _, v := range verts {
		switch {
		case v.TexCoord != 0 && v.Normal != 0:
			fmt.Printf(" %d/%d/%d", v.Vertex, v.TexCoord, v.Normal)
		case v.TexCoord != 0:
			fmt.Printf(" %d/%d", v.Vertex, v.TexCoord)
		case v.Normal != 0:
			fmt.Printf(" %d//%d", v.Vertex, v.Normal)
		default:
			fmt.Printf(" %d", v.Vertex)
		}
	}
	fmt.Printf("\n")
	return nil
}

func (op builder) L(verts ...obj.Indices) error {
	fmt.Printf("l")
	for _, v := range verts {
		switch {
		case v.TexCoord != 0 && v.Normal != 0:
			fmt.Printf(" %d/%d/%d", v.Vertex, v.TexCoord, v.Normal)
		case v.TexCoord != 0:
			fmt.Printf(" %d/%d", v.Vertex, v.TexCoord)
		case v.Normal != 0:
			fmt.Printf(" %d//%d", v.Vertex, v.Normal)
		default:
			fmt.Printf(" %d", v.Vertex)
		}
	}
	fmt.Printf("\n")
	return nil
}

func (op builder) MtlLib(names ...string) error {
	fmt.Printf("mtllib")
	for _, n := range names {
		fmt.Printf(" %s", n)
	}
	fmt.Printf("\n")
	return nil
}

func (op builder) UseMtl(name string) error {
	fmt.Printf("usemtl %s\n", name)
	return nil
}

func (op builder) MapLib(names ...string) error {
	fmt.Printf("maplib")
	for _, n := range names {
		fmt.Printf(" %s", n)
	}
	fmt.Printf("\n")
	return nil
}

func (op builder) UseMap(name string) error {
	fmt.Printf("usemap %s\n", name)
	return nil
}

func (op builder) NotSupported(txt string) error {
	fmt.Printf("NOT SUPPORTED: %s\n", txt)
	return nil
}

//------------------------------------------------------------------------------
