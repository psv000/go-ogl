package materials

type (
	// Loader ...
	Loader interface {
		Get(name string) (Material, bool)
	}
)
