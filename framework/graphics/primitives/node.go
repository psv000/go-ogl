package primitives

import (
	"framework/mth"

	"github.com/westphae/quaternion"
)

type (
	// Node ...
	Node struct {
		px, py, pz    float64 // position
		scx, scy, scz float64 // scaling
		rx, ry, rz    float64 // rotation rad
		sx, sy, sz    float64 // size
		ax, ay, az    float64 // anchor

		ppx, ppy, ppz float64 // pivot
		pcx, pcy, pcz float64 // parent size

		scaleFactor float64
	}
)

// NewNode ...
func NewNode() *Node {
	return &Node{
		scx: 1., scy: 1., scz: 1.,
		ax: 0.5, ay: 0.5, az: 0.5,
		ppx: 0.5, ppy: 0.5, ppz: 0.5,
		pcx: 0.5, pcy: 0.5, pcz: 0.5,
		scaleFactor: 1.,
	}
}

// Locate ...
func (n *Node) Locate(x, y, z float64) {
	n.px = x
	n.py = y
	n.pz = z
}

// Scale ...
func (n *Node) Scale(x, y, z float64) {
	n.scx = x
	n.scy = y
	n.scz = z
}

// Rotate ...
func (n *Node) Rotate(x, y, z float64) {
	n.rx = x
	n.ry = y
	n.rz = z
}

// Anchor ...
func (n *Node) Anchor(x, y, z float64) {
	n.ax = x
	n.ay = y
	n.az = z
}

// Pivot ...
func (n *Node) Pivot(x, y, z float64) {
	n.ppx = x
	n.ppy = y
	n.ppz = z
}

// Resize ...
func (n *Node) Resize(x, y, z float64) {
	n.sx = x
	n.sy = y
	n.sz = z
}

// ParentSize ...
func (n *Node) ParentSize(x, y, z float64) {
	n.pcx = x
	n.pcy = y
	n.pcz = z
}

// Update ...
func (n *Node) Update() mth.Mat4f {
	xpos := n.px
	ypos := n.py
	zpos := n.pz

	if n.ppx == 1. {
		xpos *= -1.
	}
	if n.ppy == 1. {
		ypos *= -1.
	}
	if n.ppz == 1. {
		zpos *= -1.
	}

	m := mth.NewUnitMat4f()
	m = scale(m, n.scx, n.scy, n.scz)
	m = rotate(m, n.rx, n.ry, n.rz)
	m = relocate(
		m,
		xpos, ypos, zpos,
		n.sx, n.sy, n.sz,
		n.pcx, n.pcy, n.pcz,
		n.scx, n.scy, n.scz,
		n.ax, n.ay, n.az,
		n.ppx, n.ppy, n.ppz,
		n.scaleFactor)

	return m
}

func relocate(m mth.Mat4f,
	x, y, z float64, // movement
	cx, cy, cz float64, // node size
	pcx, pcy, pcz float64, // parent node size
	sx, sy, sz float64, // scaling
	ax, ay, az float64, // anchor point
	ppx, ppy, ppz float64, // pivot point
	factor float64) mth.Mat4f { // scale factor
	m.Set(3, 1, factor*(y-cy*sy*ay+pcy*ppy)).
		Set(3, 0, factor*(x-cx*sx*ax+pcx*ppx)).
		Set(3, 2, factor*(z-cz*sz*az+pcz*ppz))
	return m
}

func scale(m mth.Mat4f, x, y, z float64) mth.Mat4f {
	n := mth.NewUnitMat4f()
	n.Set(0, 0, x).
		Set(1, 1, y).
		Set(2, 2, z)
	return m.Mult(n)
}

func rotate(m mth.Mat4f, rx, ry, rz float64) mth.Mat4f {
	q := quaternion.Quaternion{W: 1., X: rx, Y: ry, Z: rz}
	rm := q.RotMat()
	m.Rotate(rm)
	return m
}
