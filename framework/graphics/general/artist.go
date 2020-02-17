package general

import "framework/mth"

type (
	// Artist ...
	Artist interface {
		ClearScreen()
		DrawMesh(glMesh interface{}, m mth.Mat4f, color mth.Vec4f32)
	}
)
