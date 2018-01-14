// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pbr

//------------------------------------------------------------------------------

import (
	"math"

	"github.com/drakmaniso/carol/core/gl"
	"github.com/drakmaniso/carol/core/math32"
	"github.com/drakmaniso/carol/internal"
	"github.com/drakmaniso/carol/plane"
	"github.com/drakmaniso/carol/space"
)

//------------------------------------------------------------------------------

//TODO: add another camera class using quaternions

// A PlanarCamera is a camera designed to move on a plane. Rotation is locked
// around the "up" axis.
type PlanarCamera struct {
	// Uniform buffer

	ubo gl.UniformBuffer

	buffer struct {
		ProjectionView space.Matrix
		CameraPosition space.Coord
		CameraExposure float32
	}

	ready bool

	// Projection matrix

	projection space.Matrix

	aspectRatio float32
	fieldOfView float32
	near, far   float32

	// View matrix

	view space.Matrix

	focus            space.Coord
	distance         float32
	yaw, pitch, roll float32
}

//------------------------------------------------------------------------------

// NewPlanarCamera returns a new camera.
func NewPlanarCamera() *PlanarCamera {
	var c PlanarCamera
	c.ubo = gl.NewUniformBuffer(&c.buffer, gl.DynamicStorage)
	c.SetExposure(16.0, 1.0/125.0, 100.0)
	c.SetFieldOfView(math.Pi/4, 0.001, 1000.0)

	return &c
}

//------------------------------------------------------------------------------

func (c *PlanarCamera) View() space.Matrix {
	if !c.ready {
		c.prepare()
	}
	return c.view
}

//------------------------------------------------------------------------------

func (c *PlanarCamera) SetFieldOfView(fov float32, near, far float32) {
	c.ready = false

	c.fieldOfView = fov
	c.near, c.far = near, far
	c.WindowResized()
}

//------------------------------------------------------------------------------

func (c *PlanarCamera) WindowResized() {
	s := plane.Coord{
		float32(internal.Window.Width),
		float32(internal.Window.Height),
	}
	r := s.X / s.Y
	c.projection = space.Perspective(c.fieldOfView, r, c.near, c.far)
}

//------------------------------------------------------------------------------

func (c *PlanarCamera) SetFocus(p space.Coord) {
	c.ready = false

	c.focus = p
}

func (c *PlanarCamera) Focus() space.Coord {
	return c.focus
}

func (c *PlanarCamera) Move(forward, lateral, vertical float32) {
	c.ready = false

	cos := math32.Cos(c.yaw)
	sin := math32.Sin(c.yaw)
	c.focus.X += lateral*cos - forward*sin
	c.focus.Z += lateral*sin + forward*cos
	c.focus.Y += vertical
}

//------------------------------------------------------------------------------

func (c *PlanarCamera) SetDistance(d float32) {
	c.distance = d
}

func (c *PlanarCamera) ChangeDistance(d float32) {
	c.ready = false

	c.distance += d
}

func (c *PlanarCamera) Distance() float32 {
	return c.distance
}

//------------------------------------------------------------------------------

func (c *PlanarCamera) SetOrientation(yaw, pitch, roll float32) {
	c.ready = false

	c.yaw = yaw
	c.pitch = pitch
	c.roll = roll
}

func (c *PlanarCamera) Rotate(yaw, pitch, roll float32) {
	c.ready = false

	if c.roll != 0 {
		cos := math32.Cos(c.roll)
		sin := math32.Sin(c.roll)
		yaw, pitch = cos*yaw-sin*pitch, sin*yaw+cos*pitch
	}

	c.yaw += yaw
	for c.yaw > math32.Pi {
		c.yaw -= 2 * math32.Pi
	}
	for c.yaw < -math32.Pi {
		c.yaw += 2 * math32.Pi
	}

	c.pitch += pitch
	switch {
	case c.pitch < -math.Pi/2:
		c.pitch = -math.Pi / 2
	case c.pitch > +math.Pi/2:
		c.pitch = +math.Pi / 2
	}

	c.roll += roll
	for c.roll > math32.Pi {
		c.roll -= 2 * math32.Pi
	}
	for c.roll < -math32.Pi {
		c.roll += 2 * math32.Pi
	}
}

func (c *PlanarCamera) Orientation() (yaw, pitch, roll float32) {
	c.ready = false
	return c.yaw, c.pitch, c.roll
}

//------------------------------------------------------------------------------

func (c *PlanarCamera) SetExposure(aperture, shutterTime, sensitivity float64) {
	// See "Moving Frostbite to Physically Based Rendering", Lagarde, de Rousiers (SIGGRAPH 2014)
	ev100 := math.Log2((aperture * aperture) / shutterTime * 100.0 / sensitivity)
	maxLum := 1.2 * math.Pow(2.0, ev100)
	c.buffer.CameraExposure = float32(1.0 / maxLum)
}

//------------------------------------------------------------------------------

func (c *PlanarCamera) Bind() {
	if !c.ready {
		c.prepare()
	}
	c.ubo.SubData(&c.buffer, 0)
	c.ubo.Bind(0)
}

//------------------------------------------------------------------------------

func (c *PlanarCamera) prepare() {
	// Compute the view and projection matrices
	r := space.EulerZXY(c.pitch, c.yaw, c.roll)
	c.view = space.Translation(space.Coord{0, 0, -c.distance})
	c.view = c.view.Times(r)
	c.view = c.view.Times(space.Translation(c.focus.Opposite()))

	c.buffer.ProjectionView = c.projection.Times(c.view)

	// Compute the focus point position
	d := space.Apply(r.Transpose(), space.Homogen{0, 0, c.distance, 1}).Coord()
	c.buffer.CameraPosition = c.focus.Plus(d)

	c.ready = true
}

//------------------------------------------------------------------------------
