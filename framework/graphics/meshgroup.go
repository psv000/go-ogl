package graphics

import (
	"framework/graphics/ogl"
	"framework/graphics/primitives"
)

type (
	// MeshGroup ...
	MeshGroup struct {
		// Meshes ...
		Meshes []*primitives.Mesh
		// LightSources ...
		LightSources []LightSource
		// GPUPack ...
		GPUPack ogl.ProgramPack
	}
)