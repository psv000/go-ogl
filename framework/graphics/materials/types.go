package materials

type (
	// MaterialMap ...
	MaterialMap struct {
		// X, Y, Z, W left-bottom and right-top borders of material
		X, Y, Z, W float64
		// Name ...
		Name string
	}
	// AtlasMapping ...
	AtlasMapping struct {
		// Materials ...
		Materials []MaterialMap
		// Name ...
		Name string
	}

	// Material ...
	Material struct {
		// ID ...
		ID string
		// TexID ...
		TexID uint32
		// X Y Z W ...
		X, Y, Z, W float64
	}

	// AtlasPath ...
	AtlasPath struct {
		image   string
		mapping string
	}
)
