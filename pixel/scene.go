// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"unsafe"

	"github.com/cozely/cozely/internal"
	"github.com/cozely/cozely/x/gl"
)

////////////////////////////////////////////////////////////////////////////////

type SceneID uint16

const (
	maxSceneID = 0xFFFF
	noScene    = SceneID(maxSceneID)
)

var scenes struct {
	// For each scene
	commandsICBO  []gl.IndirectBuffer
	parametersTBO []gl.BufferTexture
	cursor        []TextCursor
	commands      [][]gl.DrawIndirectCommand
	parameters    [][]int16
}

////////////////////////////////////////////////////////////////////////////////

func Scene() SceneID {
	if internal.Running {
		setErr(errors.New("pixel scene declaration: declarations must happen before starting the framework"))
		return noScene
	}

	if len(scenes.commandsICBO) >= maxSceneID {
		setErr(errors.New("pixel scene declaration: too many scenes"))
		return noScene
	}

	a := SceneID(len(scenes.commandsICBO))

	scenes.commandsICBO = append(scenes.commandsICBO, gl.IndirectBuffer{})
	scenes.parametersTBO = append(scenes.parametersTBO, gl.BufferTexture{})
	scenes.cursor = append(scenes.cursor, TextCursor{})
	scenes.commands = append(scenes.commands,
		make([]gl.DrawIndirectCommand, 0, maxCommandCount))
	scenes.parameters = append(scenes.parameters,
		make([]int16, 0, maxParamCount))

	return a
}

////////////////////////////////////////////////////////////////////////////////

func (a SceneID) setup() {
	scenes.commandsICBO[a] = gl.NewIndirectBuffer(
		uintptr(cap(scenes.commands[a]))*unsafe.Sizeof(scenes.commands[a][0]),
		gl.DynamicStorage,
	)
	scenes.parametersTBO[a] = gl.NewBufferTexture(
		uintptr(cap(scenes.parameters[a]))*unsafe.Sizeof(scenes.parameters[a][0]),
		gl.R16I,
		gl.DynamicStorage,
	)
}
