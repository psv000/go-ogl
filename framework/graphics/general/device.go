package general

type (
	// Device ...
	Device interface {
		Init() error
		CompileProgram(vertex, fragment string) (uint32, error)
	}
)
