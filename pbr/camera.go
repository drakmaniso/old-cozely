// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pbr

//------------------------------------------------------------------------------

import (
	"math"

	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/space"
)

//------------------------------------------------------------------------------

//TODO: use quaternion or rotation matrix directly

type Camera struct {
	ubo gfx.UniformBuffer

	buffer struct {
		ProjectionView space.Matrix
		CameraPosition space.Coord
		CameraExposure float32
	}

	projection  space.Matrix
	aspectRatio float32
	fieldOfView float32
	near, far   float32

	view, currView, prevView space.Matrix
}

//------------------------------------------------------------------------------

func NewCamera() *Camera {
	var c Camera
	c.ubo = gfx.NewUniformBuffer(&c.buffer, gfx.DynamicStorage)
	c.SetExposure(16.0, 1.0/125.0, 100.0)
	c.SetFieldOfView(math.Pi/4, 0.001, 1000.0)
	c.currView = space.Identity()

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
	s := plane.Coord{
		float32(internal.Window.Width),
		float32(internal.Window.Height),
	}
	r := s.X / s.Y
	c.projection = space.Perspective(c.fieldOfView, r, c.near, c.far)
}

//------------------------------------------------------------------------------

func (c *Camera) SetExposure(aperture, shutterTime, sensitivity float64) {
	// See "Moving Frostbite to Physically Based Rendering", Lagarde, de Rousiers (SIGGRAPH 2014)
	ev100 := math.Log2((aperture * aperture) / shutterTime * 100.0 / sensitivity)
	maxLum := 1.2 * math.Pow(2.0, ev100)
	c.buffer.CameraExposure = float32(1.0 / maxLum)
}

//------------------------------------------------------------------------------

func (c *Camera) SetPosition(p space.Coord) {
	c.currView[3][0] = 0
	c.currView[3][1] = 0
	c.currView[3][2] = 0
	t := space.Apply(c.currView, p.Opposite().Homogen())
	c.currView[3][0] = t.X
	c.currView[3][1] = t.Y
	c.currView[3][2] = t.Z
}

func (c *Camera) Position() space.Coord {
	r := c.currView
	r[3][0] = 0
	r[3][1] = 0
	r[3][2] = 0
	t := space.Homogen{-c.currView[3][0], -c.currView[3][1], -c.currView[3][2], 1.0}
	return space.Apply(r, t).Coord()
}

func (c *Camera) Move(forward, lateral, vertical float32) {
	t := space.Coord{c.currView[3][0], c.currView[3][1], c.currView[3][2]}
	t = t.Minus(space.Coord{lateral, vertical, forward})
	c.currView[3][0] = t.X
	c.currView[3][1] = t.Y
	c.currView[3][2] = t.Z
}

//------------------------------------------------------------------------------

func (c *Camera) SetOrientation(yaw, pitch, roll float32) {
	p := c.Position()
	c.currView = space.EulerZXY(-pitch, -yaw, -roll)
	c.SetPosition(p)
}

func (c *Camera) Yaw(a float32) {
	r := space.EulerXYZ(0, a, 0)
	c.currView = r.Times(c.currView)
}

func (c *Camera) Pitch(a float32) {
	r := space.EulerXYZ(a, 0, 0)
	c.currView = r.Times(c.currView)
}

func (c *Camera) Roll(a float32) {
	r := space.EulerXYZ(0, 0, a)
	c.currView = r.Times(c.currView)
}

func (c *Camera) Rotate(angle float32, axis space.Coord) {
	r := c.currView
	r[3][0] = 0
	r[3][1] = 0
	r[3][2] = 0
	axis = space.Apply(r, axis.Opposite().Homogen()).Coord()
	rr := space.Rotation(angle, axis)
	c.currView = rr.Times(c.currView)
}

func (c *Camera) RotateFP(yaw, pitch float32) {
	ry := space.EulerXYZ(0, yaw, 0)
	rp := space.EulerXYZ(pitch, 0, 0)
	c.currView = ry.Times(c.currView.Times(rp))
}

// func (c *Camera) Rotate(yaw, pitch, roll float32) {
// 	r := space.EulerZXY(pitch, yaw, roll)
// 	c.currView = r.Times(c.currView)
// c.current.yaw += yaw
// for c.current.yaw > math32.Pi {
// 	c.current.yaw -= 2 * math32.Pi
// 	// Need to update previous state too, because of interpolation
// 	c.previous.yaw -= 2 * math32.Pi
// }
// for c.current.yaw < -math32.Pi {
// 	c.current.yaw += 2 * math32.Pi
// 	// Need to update previous state too, because of interpolation
// 	c.previous.yaw += 2 * math32.Pi
// }

// c.current.pitch += pitch
// switch {
// case c.current.pitch < -math.Pi/2:
// 	c.current.pitch = -math.Pi / 2
// case c.current.pitch > +math.Pi/2:
// 	c.current.pitch = +math.Pi / 2
// }

// c.current.roll += roll
// for c.current.roll > math32.Pi {
// 	c.current.roll -= 2 * math32.Pi
// 	// Need to update previous state too, because of interpolation
// 	c.previous.roll -= 2 * math32.Pi
// }
// for c.current.roll < -math32.Pi {
// 	c.current.roll += 2 * math32.Pi
// 	// Need to update previous state too, because of interpolation
// 	c.previous.roll += 2 * math32.Pi
// }
// }

func (c *Camera) Orientation() (yaw, pitch, roll float32) {
	return 0, 0, 0 // c.current.yaw, c.current.pitch, c.current.roll
}

//------------------------------------------------------------------------------

func (c *Camera) NextState() {
	c.prevView = c.currView
}

//------------------------------------------------------------------------------

func (c *Camera) Bind() {
	c.updateView()
	c.buffer.ProjectionView = c.projection.Times(c.currView)
	c.ubo.SubData(&c.buffer, 0)
	c.ubo.Bind(0)
}

//------------------------------------------------------------------------------

func (c *Camera) updateView() {
	// a := float32(internal.DrawInterpolation)
	// pos := c.previous.position.Times(1 - a).Plus(c.current.position.Times(a))
	// yaw := c.previous.yaw*(1-a) + c.current.yaw*a
	// pitch := c.previous.pitch*(1-a) + c.current.pitch*a
	// roll := c.previous.roll*(1-a) + c.current.roll*a

	// c.view = space.EulerZXY(pitch, yaw, roll)
	// c.view = c.view.Times(space.Translation(pos.Opposite()))

	c.buffer.CameraPosition = c.Position()
}

//------------------------------------------------------------------------------
