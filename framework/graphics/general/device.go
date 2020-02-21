package general

import (
	"framework/graphics/program"
)

type (
	// Device ...
	Device interface {
		Init() error
		CompileProgram(vertex, fragment string) (uint32, error)
		CreateUniforms(programID uint32, args []program.Arg) []program.Uniform

		NewTexLoader() GlTexLoader
	}
)
