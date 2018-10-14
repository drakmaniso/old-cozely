package obj_test

import (
	"fmt"
	"os"

	"github.com/cozely/cozely/formats/obj"
)

func Example_objParser() {
	f, err := os.Open("testdata/cube.obj")
	if err != nil {
		fmt.Println(err)
		return
	}
	obj.Parse(f, builder{})
	//Output:
	// # Blender v2.78 (sub 0) OBJ File: ''
	// # www.blender.org
	// o cube
	// v -1.000000 -1.000000 1.000000
	// v -1.000000 1.000000 1.000000
	// v -1.000000 -1.000000 -1.000000
	// v -1.000000 1.000000 -1.000000
	// v 1.000000 -1.000000 1.000000
	// v 1.000000 1.000000 1.000000
	// v 1.000000 -1.000000 -1.000000
	// v 1.000000 1.000000 -1.000000
	// usemtl 0
	// s 0
	// f 5 6 2 1
	// usemtl 1
	// f 3 4 8 7
	// usemtl 2
	// f 8 4 2 6
	// usemtl 3
	// f 3 7 5 1
	// usemtl 4
	// f 1 2 4 3
	// usemtl 5
	// f 7 8 6 5
}

////////////////////////////////////////////////////////////////////////////////

type builder struct{ obj.DefaultBuilder }

func (builder) Comment(txt string) error {
	fmt.Printf("%s\n", txt)
	return nil
}

func (builder) O(name string) error {
	fmt.Printf("o %s\n", name)
	return nil
}

func (builder) G(names ...string) error {
	fmt.Printf("g")
	for _, n := range names {
		fmt.Printf(" %s", n)
	}
	fmt.Printf("\n")
	return nil
}

func (builder) S(group int) error {
	fmt.Printf("s %d\n", group)
	return nil
}

func (builder) LOD(level int) error {
	fmt.Printf("lod %d\n", level)
	return nil
}

func (builder) V(coords ...float32) error {
	fmt.Printf("v")
	for _, v := range coords {
		fmt.Printf(" %f", v)
	}
	fmt.Printf("\n")
	return nil
}

func (builder) VT(coords ...float32) error {
	fmt.Printf("vt")
	for _, v := range coords {
		fmt.Printf(" %f", v)
	}
	fmt.Printf("\n")
	return nil
}

func (builder) VN(coords ...float32) error {
	fmt.Printf("vn")
	for _, v := range coords {
		fmt.Printf(" %f", v)
	}
	fmt.Printf("\n")
	return nil
}

func (builder) F(verts ...obj.Indices) error {
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

func (builder) P(verts ...obj.Indices) error {
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

func (builder) L(verts ...obj.Indices) error {
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

func (builder) MtlLib(names ...string) error {
	fmt.Printf("mtllib")
	for _, n := range names {
		fmt.Printf(" %s", n)
	}
	fmt.Printf("\n")
	return nil
}

func (builder) UseMtl(name string) error {
	fmt.Printf("usemtl %s\n", name)
	return nil
}

func (builder) MapLib(names ...string) error {
	fmt.Printf("maplib")
	for _, n := range names {
		fmt.Printf(" %s", n)
	}
	fmt.Printf("\n")
	return nil
}

func (builder) UseMap(name string) error {
	fmt.Printf("usemap %s\n", name)
	return nil
}

func (builder) NotSupported(txt string) error {
	fmt.Printf("NOT SUPPORTED: %s\n", txt)
	return nil
}

// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).
