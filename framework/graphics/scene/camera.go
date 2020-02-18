package scene

import (
	"framework/mth"
	"math"

	"github.com/westphae/quaternion"
)

const (
	fov   = 45. * math.Pi / 180.
	zNear = 0.1
	zFar  = 1000.
)

const (
	// PerspectiveCT ...
	PerspectiveCT = iota
	// OrthoCT ...
	OrthoCT
)

type (
	// CamType ...
	CamType int

	// Camera ...
	Camera struct {
		view       mth.Mat4f
		projection mth.Mat4f

		typ CamType
	}
)

// NewCamera ...
func NewCamera(typ CamType) *Camera {
	return &Camera{
		view:       mth.NewUnitMat4f(),
		projection: mth.NewUnitMat4f(),
		typ:        typ,
	}
}

// SetView ...
func (c *Camera) SetView(w, h int) *Camera {
	switch c.typ {
	case PerspectiveCT:
		aspect := float64(w) / float64(h)
		c.projection = mth.PerspectiveRH(fov, aspect, zNear, zFar)
	case OrthoCT:
		aspect := float64(w) / float64(h)
		c.projection = mth.OrthoRH(aspect*-0.5, aspect*0.5, -0.5, 0.5, zNear, zFar)
	}
	return c
}

// Relocate ...
func (c *Camera) Relocate(x, y, z float64) *Camera {
	c.view.Set(3, 0, x).
		Set(3, 1, y).
		Set(3, 2, z)
	return c
}

// Magnification ...
func (c *Camera) Magnification(m float64) *Camera {
	scaling := mth.NewUnitMat4f()
	scaling.Set(0, 0, m).
		Set(1, 1, m).
		Set(2, 2, m)
	c.view.Mult(scaling)
	return c
}

// Rotation ...
func (c *Camera) Rotation(rx, ry, rz float64) *Camera {
	q := quaternion.Quaternion{W: 1., X: rx, Y: ry, Z: rz}
	rm := q.RotMat()
	c.view.Rotate(rm)
	return c
}

// View ...
func (c *Camera) View() mth.Mat4f {
	return c.view
}

// Projection ...
func (c *Camera) Projection() mth.Mat4f {
	return c.projection
}
