// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/space"
)

//------------------------------------------------------------------------------

func SetupMeshBuffers(m Meshes) error {
	faceSSBO.Delete()
	faceSSBO = gfx.NewStorageBuffer(
		uintptr(len(m.Faces)*8),
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

// func (m *Meshes) AddObj(filename string) (MeshID, error) {
// 	// return 0, 0, nil
// }

//------------------------------------------------------------------------------

type Face struct {
	Material uint16
	Faces    [3]uint16
}

//------------------------------------------------------------------------------

type Instance struct {
	model      space.Matrix `layout:"0" divisor:"1"`
	BaseVertex uint32       `layout:"1"`
}

//------------------------------------------------------------------------------
