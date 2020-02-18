package resources

import "framework/mth"

const (
	// VertInd ...
	VertInd = iota
	// UVSInd ...
	UVSInd
	// NormInd ...
	NormInd
	// IndLen ...
	IndLen
)

type (
	// Obj3DModel ...
	Obj3DModel struct {
		Vert []mth.Vec3f64
		UVS  []mth.Vec2f64
		Norm []mth.Vec3f64

		Ind [][IndLen]uint32
	}
)
