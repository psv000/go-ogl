package primitives

import (
	"framework/graphics/program"
)

type (
	// MeshGroup ...
	MeshGroup struct {
		// Meshes ...
		Meshes []*Mesh
		// LightSources ...
		LightSources []LightSource
		// GPUPack ...
		GPUPack program.Pack
	}
)
