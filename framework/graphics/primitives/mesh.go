package primitives

import (
	"framework/mth"
)

type (
	// Mesh ...
	Mesh struct {
		node *Node
		back interface{}

		color mth.Vec4f32
	}
)

// NewMesh ...
func NewMesh(node *Node, back interface{}) *Mesh {
	return &Mesh{node: node, back: back, color: mth.Vec4f32{1., 1., 1., 1.}}
}

// Gl ...
func (m *Mesh) Gl() interface{} {
	return m.back
}

// Node ...
func (m *Mesh) Node() *Node {
	return m.node
}

// SetColor ...
func (m *Mesh) SetColor(color mth.Vec4f32) {
	m.color = color
}

// Color ...
func (m *Mesh) Color() mth.Vec4f32 {
	return m.color
}
