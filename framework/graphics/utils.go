package graphics

import (
	"framework/graphics/general"
	"framework/graphics/ogl"
	"framework/graphics/primitives"
	"framework/graphics/program"
	"github.com/pkg/errors"
)

func loadFromV3f(dev general.Device, data interface{}, indices []uint32) (*primitives.Mesh, error) {
	const (
		defaultVert = "assets/programs/test.vert"
		defaultFrag = "assets/programs/test.frag"
	)

	p, err := dev.CompileProgram(defaultVert, defaultFrag)
	if err != nil {
		return nil, err
	}

	var args = []program.Arg{
		{Name: program.UModelName, Typ: program.Mat4Uniform, Dst: program.ModelDst},
		{Name: program.UViewName, Typ: program.Mat4Uniform, Dst: program.ViewDst},
		{Name: program.UProjectionName, Typ: program.Mat4Uniform, Dst: program.ProjectionDst},
	}

	uniforms := program.NewUniforms(dev, p, args)
	oglMesh := ogl.NewMesh(dev, p, uniforms)

	content, converted := data.([]ogl.V3f)
	if !converted {
		return nil, errors.New("data type does not match to v3f type")
	}

	oglMesh.Load(dev, content, indices)
	n := primitives.NewNode()
	m := primitives.NewMesh(n, oglMesh)
	return m, nil
}

func loadFromV3fC4b(dev general.Device, data interface{}, indices []uint32) (*primitives.Mesh, error) {
	const (
		defaultVert = "assets/programs/col.vert"
		defaultFrag = "assets/programs/col.frag"
	)

	p, err := dev.CompileProgram(defaultVert, defaultFrag)
	if err != nil {
		return nil, err
	}

	var args = []program.Arg{
		{Name: program.UModelName, Typ: program.Mat4Uniform, Dst: program.ModelDst},
		{Name: program.UViewName, Typ: program.Mat4Uniform, Dst: program.ViewDst},
		{Name: program.UProjectionName, Typ: program.Mat4Uniform, Dst: program.ProjectionDst},
	}

	uniforms := program.NewUniforms(dev, p, args)
	oglMesh := ogl.NewMesh(dev, p, uniforms)

	content, converted := data.([]ogl.V3fC4b)
	if !converted {
		return nil, errors.New("data type does not match to v3f type")
	}

	oglMesh.Load(dev, content, indices)
	n := primitives.NewNode()
	m := primitives.NewMesh(n, oglMesh)
	return m, nil
}
