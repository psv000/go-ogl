package primitives

type (
	// Loader ...
	Loader interface {
		LoadMeshFromFile(filepath string) (*Mesh, error)
		LoadMeshFromData(data interface{}, indices []uint32) (*Mesh, error)

		NewMeshGroup(vert, frag string) *MeshGroup
		DelMeshGroup(mg *MeshGroup)
	}
)
