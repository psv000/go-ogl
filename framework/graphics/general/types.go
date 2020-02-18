package general

import "framework/mth"

type (
	// Mesh ...
	Mesh interface {
		Load(device interface{}, content interface{}, indices []uint32)
		Update(t mth.Mat4f) mth.Mat4f
	}
)
