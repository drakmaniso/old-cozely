// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

import (
	"fmt"
	"os"
	"strconv"

	"github.com/drakmaniso/cozely/formats/obj"
	"github.com/drakmaniso/cozely/space"
	"github.com/drakmaniso/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

func SetupMeshBuffers(m Meshes) error {
	faceSSBO.Delete()
	faceSSBO = gl.NewStorageBuffer(
		uintptr(len(m.Faces)*8),
		gl.DynamicStorage,
	)
	faceSSBO.SubData(m.Faces, 0)

	vertexSSBO.Delete()
	vertexSSBO = gl.NewStorageBuffer(
		uintptr(len(m.Vertices)*12),
		gl.DynamicStorage,
	)
	vertexSSBO.SubData(m.Vertices, 0)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func BindMeshBuffers() {
	faceSSBO.Bind(0)
	vertexSSBO.Bind(1)
}

////////////////////////////////////////////////////////////////////////////////

func DeleteMeshBuffers() {
	faceSSBO.Delete()
	vertexSSBO.Delete()
}

////////////////////////////////////////////////////////////////////////////////

var faceSSBO gl.StorageBuffer

var vertexSSBO gl.StorageBuffer

////////////////////////////////////////////////////////////////////////////////

type Meshes struct {
	Faces    []Face
	Vertices []space.Coord
}

type MeshID struct {
	FaceID   uint32
	VertexID uint32
}

func (m *Meshes) AddObj(filename string) (MeshID, error) {
	mid := MeshID{FaceID: uint32(len(m.Faces)), VertexID: uint32(len(m.Vertices))}

	f, err := os.Open(filename)
	if err != nil {
		return MeshID{}, err //TODO: error wrapping
	}
	b := builder{
		meshes: m,
	}
	obj.Parse(f, &b)

	//TODO: debug log
	fmt.Printf(
		"Loaded %s: \n    vertices: %d, faces: %d\n",
		filename,
		len(m.Vertices)-int(mid.VertexID),
		len(m.Faces)-int(mid.FaceID),
	)

	return mid, nil
}

type builder struct {
	obj.DefaultBuilder
	meshes  *Meshes
	currMat byte
}

func (b *builder) V(coords ...float32) error {
	if len(coords) < 3 {
		return nil //TODO:error handling
	}

	v := space.Coord{coords[0], coords[1], coords[2]}
	b.meshes.Vertices = append(b.meshes.Vertices, v)

	return nil
}

func (b *builder) F(verts ...obj.Indices) error {
	if len(verts) < 3 {
		return nil //TODO:error handling
	}

	f := Face{}
	// f := Face{
	// 	Material: b.currMat,
	// 	Faces: [4]uint16{
	// 		uint16(verts[0].Vertex - 1),
	// 		uint16(verts[1].Vertex - 1),
	// 		uint16(verts[2].Vertex - 1),
	// 	},
	// }
	if len(verts) >= 4 {
		f.MakeFace(
			b.currMat,
			uint32(verts[0].Vertex-1),
			uint32(verts[1].Vertex-1),
			uint32(verts[2].Vertex-1),
			uint32(verts[3].Vertex-1),
		)
		// f.Faces[3] = uint16(verts[3].Vertex - 1)
	} else {
		f.MakeFace(
			b.currMat,
			uint32(verts[0].Vertex-1),
			uint32(verts[1].Vertex-1),
			uint32(verts[2].Vertex-1),
			uint32(verts[2].Vertex-1),
		)
		// f.Faces[3] = f.Faces[2]
	}
	//TODO: handle negative indices
	b.meshes.Faces = append(b.meshes.Faces, f)

	return nil
}

func (b *builder) UseMtl(name string) error {
	v, _ := strconv.Atoi(name)
	b.currMat = byte(v)

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type Face struct {
	MathiVert0Vert1 uint32
	MatloVert2Vert3 uint32
}

func (f *Face) MakeFace(material byte, a, b, c, d uint32) {
	m1 := uint32(material&0xF0) >> 4
	m2 := uint32(material & 0x0F)
	a &= 0x3FFF
	b &= 0x3FFF
	c &= 0x3FFF
	d &= 0x3FFF
	f.MathiVert0Vert1 = (m1 << 28) | (a << 14) | b
	f.MatloVert2Vert3 = (m2 << 28) | (c << 14) | d
}

////////////////////////////////////////////////////////////////////////////////

type Instance struct {
	model      space.Matrix `layout:"0" divisor:"1"`
	BaseVertex uint32       `layout:"1"`
}

////////////////////////////////////////////////////////////////////////////////
