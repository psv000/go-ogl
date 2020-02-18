package graphics

import (
	"framework/graphics/primitives"
)

type (
	// MeshGroup ...
	MeshGroup struct {
		Meshes []*primitives.Mesh
		LightSources []LightSource
	}
)