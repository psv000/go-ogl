package primitives

import (
	"framework/graphics/ogl"
)

type (
	// MeshGroup ...
	MeshGroup struct {
		// Meshes ...
		Meshes []*Mesh
		// LightSources ...
		LightSources []LightSource
		// GPUPack ...
		GPUPack ogl.ProgramPack
	}
)