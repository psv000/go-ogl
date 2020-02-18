package general

import (
	"framework/graphics/primitives"
	"framework/mth"
)

type (
	// Artist ...
	Artist interface {
		ClearScreen()
		DrawMesh(glMesh interface{}, model, view, projection mth.Mat4f, color mth.Vec4f32)
		DrawMeshGroup(group *primitives.MeshGroup, view, projection mth.Mat4f)
	}
)
