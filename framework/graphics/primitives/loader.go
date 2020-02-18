package primitives

import "framework/graphics"

type (
	// Loader ...
	Loader interface {
		LoadMeshFromFile(filepath string) (*Mesh, error)
		LoadMeshFromData(data interface{}, indices []uint32) (*Mesh, error)

		NewMeshGroup(vert, frag string) *graphics.MeshGroup
		DelMeshGroup(mg *graphics.MeshGroup)
	}
)
