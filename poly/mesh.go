// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

//------------------------------------------------------------------------------

import (
	"os"

	"strconv"

	"github.com/drakmaniso/glam/formats/obj"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/space"
)

//------------------------------------------------------------------------------

func SetupMeshBuffers(m Meshes) error {
	faceSSBO.Delete()
	faceSSBO = gfx.NewStorageBuffer(
		uintptr(len(m.Faces)*12),
		gfx.DynamicStorage,
	)
	faceSSBO.SubData(m.Faces, 0)

	vertexSSBO.Delete()
	vertexSSBO = gfx.NewStorageBuffer(
		uintptr(len(m.Vertices)*12),
		gfx.DynamicStorage,
	)
	vertexSSBO.SubData(m.Vertices, 0)

	return nil
}

//------------------------------------------------------------------------------

func BindMeshBuffers() {
	faceSSBO.Bind(0)
	vertexSSBO.Bind(1)
}

//------------------------------------------------------------------------------

func DeleteMeshBuffers() {
	faceSSBO.Delete()
	vertexSSBO.Delete()
}

//------------------------------------------------------------------------------

var faceSSBO gfx.StorageBuffer

var vertexSSBO gfx.StorageBuffer

//------------------------------------------------------------------------------

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

	return mid, nil
}

type builder struct {
	obj.DefaultBuilder
	meshes  *Meshes
	currMat uint32
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

	f := Face{
		Material: b.currMat,
		Faces: [4]uint16{
			uint16(verts[0].Vertex - 1),
			uint16(verts[1].Vertex - 1),
			uint16(verts[2].Vertex - 1),
		},
	}
	if len(verts) >= 4 {
		f.Faces[3] = uint16(verts[3].Vertex - 1)
	} else {
		f.Faces[3] = f.Faces[2]
	}
	//TODO: handle negative indices
	b.meshes.Faces = append(b.meshes.Faces, f)

	return nil
}

func (b *builder) UseMtl(name string) error {
	v, _ := strconv.Atoi(name)
	b.currMat = uint32(v)

	return nil
}

//------------------------------------------------------------------------------

type Face struct {
	Material uint32
	Faces    [4]uint16
}

//------------------------------------------------------------------------------

type Instance struct {
	model      space.Matrix `layout:"0" divisor:"1"`
	BaseVertex uint32       `layout:"1"`
}

//------------------------------------------------------------------------------
