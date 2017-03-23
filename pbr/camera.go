// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pbr

//------------------------------------------------------------------------------

import (
	"math"

	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/internal"
	math32 "github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/space"
)

//------------------------------------------------------------------------------

//TODO: use quaternion or rotation matrix directly

type Camera struct {
	buffer struct {
		ProjectionView space.Matrix
		CameraPosition space.Coord
		CameraExposure float32
	}

	projection space.Matrix
	view       space.Matrix

	aspectRatio float32
	fieldOfView float32
	near, far   float32

	current struct {
		position         space.Coord
		yaw, pitch, roll float32
	}

	previous struct {
		position         space.Coord
		yaw, pitch, roll float32
	}

	ubo gfx.UniformBuffer
}

//------------------------------------------------------------------------------

func NewCamera() *Camera {
	var c Camera
	c.ubo = gfx.NewUniformBuffer(&c.buffer, gfx.DynamicStorage)
	c.SetExposure(16.0, 1.0/125.0, 100.0)
	c.SetFieldOfView(math.Pi/4, 0.001, 1000.0)

	return &c
}

//------------------------------------------------------------------------------

func (c *Camera) SetFieldOfView(fov float32, near, far float32) {
	c.fieldOfView = fov
	c.near, c.far = near, far
	c.WindowResized()
}

//------------------------------------------------------------------------------

func (c *Camera) WindowResized() {
	sx, sy := float32(internal.Window.Width), float32(internal.Window.Height)
	r := sx / sy
	c.projection = space.Perspective(c.fieldOfView, r, c.near, c.far)
}

//------------------------------------------------------------------------------

func (c *Camera) SetPosition(p space.Coord) {
	c.current.position = p
}

func (c *Camera) Move(forward, lateral, vertical float32) {
	c.current.position.X += lateral*math32.Cos(c.current.yaw) - forward*math32.Sin(c.current.yaw)
	c.current.position.Z += lateral*math32.Sin(c.current.yaw) + forward*math32.Cos(c.current.yaw)
	c.current.position.Y += vertical
}

func (c *Camera) Position() space.Coord {
	return c.current.position
}

//------------------------------------------------------------------------------

func (c *Camera) SetOrientation(yaw, pitch, roll float32) {
	c.current.yaw = yaw
	c.current.pitch = pitch
	c.current.roll = roll
}

func (c *Camera) Rotate(yaw, pitch, roll float32) {
	c.current.yaw += yaw
	for c.current.yaw > math32.Pi {
		c.current.yaw -= 2 * math32.Pi
		// Need to update previous state too, because of interpolation
		c.previous.yaw -= 2 * math32.Pi
	}
	for c.current.yaw < -math32.Pi {
		c.current.yaw += 2 * math32.Pi
		// Need to update previous state too, because of interpolation
		c.previous.yaw += 2 * math32.Pi
	}

	c.current.pitch += pitch
	switch {
	case c.current.pitch < -math.Pi/2:
		c.current.pitch = -math.Pi / 2
	case c.current.pitch > +math.Pi/2:
		c.current.pitch = +math.Pi / 2
	}

	c.current.roll += roll
	for c.current.roll > math32.Pi {
		c.current.roll -= 2 * math32.Pi
		// Need to update previous state too, because of interpolation
		c.previous.roll -= 2 * math32.Pi
	}
	for c.current.roll < -math32.Pi {
		c.current.roll += 2 * math32.Pi
		// Need to update previous state too, because of interpolation
		c.previous.roll += 2 * math32.Pi
	}
}

func (c *Camera) Orientation() (yaw, pitch, roll float32) {
	return c.current.yaw, c.current.pitch, c.current.roll
}

//------------------------------------------------------------------------------

func (c *Camera) NextState() {
	c.previous = c.current
}

//------------------------------------------------------------------------------

func (c *Camera) SetExposure(aperture, shutterTime, sensitivity float64) {
	// See "Moving Frostbite to Physically Based Rendering", Lagarde, de Rousiers (SIGGRAPH 2014)
	ev100 := math.Log2((aperture * aperture) / shutterTime * 100.0 / sensitivity)
	maxLum := 1.2 * math.Pow(2.0, ev100)
	c.buffer.CameraExposure = float32(1.0 / maxLum)
}

//------------------------------------------------------------------------------

func (c *Camera) Bind() {
	c.updateView()
	c.buffer.ProjectionView = c.projection.Times(c.view)
	c.ubo.SubData(&c.buffer, 0)
	c.ubo.Bind(0)
}

//------------------------------------------------------------------------------

func (c *Camera) updateView() {
	a := float32(internal.DrawInterpolation)
	pos := c.previous.position.Times(1 - a).Plus(c.current.position.Times(a))
	yaw := c.previous.yaw*(1-a) + c.current.yaw*a
	pitch := c.previous.pitch*(1-a) + c.current.pitch*a
	roll := c.previous.roll*(1-a) + c.current.roll*a

	c.view = space.EulerZXY(pitch, yaw, roll)
	c.view = c.view.Times(space.Translation(pos.Inverse()))

	c.buffer.CameraPosition = pos
}

//------------------------------------------------------------------------------
