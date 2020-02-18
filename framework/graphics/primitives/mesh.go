package primitives

import (
	"framework/graphics/general"
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

// Draw ...
func (m *Mesh) Draw(qa general.Artist, t mth.Mat4f) {
	qa.DrawMesh(m.back, m.node.Update(t), m.color)
}

// Node ...
func (m *Mesh) Node() *Node {
	return m.node
}

// SetColor ...
func (m *Mesh) SetColor(color mth.Vec4f32) {
	m.color = color
}
