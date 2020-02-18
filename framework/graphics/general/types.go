package general

import (
	"framework/mth"

	"framework/graphics/primitives"
)

type (
	// Mesh ...
	Mesh interface {
		Load(device interface{}, content interface{}, indices []uint32)
		Update() mth.Mat4f
		Node() *primitives.Node
	}
)
