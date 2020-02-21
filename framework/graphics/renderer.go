package graphics

import (
	"framework/graphics/program"
	"framework/graphics/textures"
	"framework/mth"
	"framework/resources"
	"image"

	"framework/graphics/general"
	"framework/graphics/ogl"
	"framework/graphics/primitives"
	"framework/graphics/scene"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var unitMat = mth.NewUnitMat4f()

type (
	// Renderer ...
	Renderer struct {
		dev    general.Device
		queue  general.Queue
		artist general.Artist

		camera *scene.Camera

		tl general.GlTexLoader

		groups []*primitives.MeshGroup
	}
)

// NewRenderer ...
func NewRenderer() *Renderer {
	return &Renderer{}
}

// Serve ...
func (r *Renderer) Serve(args ...interface{}) error {
	var ok bool
	r.dev, ok = args[0].(general.Device)
	if !ok {
		return errors.New(rendererInfoTag + ": invalid gl device")
	}
	r.queue, ok = args[1].(general.Queue)
	if !ok {
		return errors.New(rendererInfoTag + ": invalid gl queue")
	}
	r.artist, ok = args[2].(general.Artist)
	if !ok {
		return errors.New(rendererInfoTag + ": invalid gl artist")
	}

	if err := r.dev.Init(); err != nil {
		return errors.Wrap(err, rendererInfoTag)
	}

	r.tl = r.dev.NewTexLoader()

	logrus.Info("renderer started")
	return nil
}

// Stop ...
func (r *Renderer) Stop() error {
	logrus.Info("renderer stopped")
	return nil
}

// ResizeView ...
func (r *Renderer) ResizeView(w, h int) {

}

// Update ...
func (r *Renderer) Update() {
	r.artist.ClearScreen()

	var view, projection = unitMat, unitMat
	if r.camera != nil {
		view, projection = r.camera.View(), r.camera.Projection()
	}

	for _, g := range r.groups {
		r.artist.DrawMeshGroup(g, view, projection)
	}

	r.queue.Process()
	r.queue.Flush()
}

// NewMeshGroup ...
func (r *Renderer) NewMeshGroup() (*primitives.MeshGroup, error) {
	p, err := r.dev.CompileProgram("assets/programs/lighting.vert", "assets/programs/lighting.frag")
	if err != nil {
		return nil, err
	}
	mg := &primitives.MeshGroup{}
	var (
		mArgs = []program.Arg{
			{Name: program.UModelName, Typ: program.Mat4Uniform, Dst: program.ModelDst},
			{Name: program.UViewName, Typ: program.Mat4Uniform, Dst: program.ViewDst},
			{Name: program.UProjectionName, Typ: program.Mat4Uniform, Dst: program.ProjectionDst},
			{Name: program.UColorName, Typ: program.Vec4Uniform, Dst: program.ColorDst},
		}
		lArgs = []program.Arg{
			{Name: program.ULightColorName, Typ: program.Vec3Uniform, Dst: program.LightColorDst},
			{Name: program.ULightPositionName, Typ: program.Vec3Uniform, Dst: program.LightPosDst},
		}
	)
	mUniforms := program.NewUniforms(r.dev, p, mArgs)
	lUniforms := program.NewUniforms(r.dev, p, lArgs)
	mg.GPUPack = program.Pack{
		ID: p,
		Uniforms: map[program.UniformCat][]program.Uniform{
			program.UCMesh:  mUniforms,
			program.UCLight: lUniforms,
		},
	}

	r.groups = append(r.groups, mg)
	return mg, nil
}

// DelMeshGroup ...
func (r *Renderer) DelMeshGroup(mg *primitives.MeshGroup) {
	var toDel int = -1
	for i, ptr := range r.groups {
		if ptr == mg {
			toDel = i
		}
	}
	if toDel >= 0 {
		newlen := len(r.groups) - 1
		r.groups[toDel], r.groups[newlen] = r.groups[newlen], r.groups[toDel]
		r.groups = r.groups[:newlen]
	}
}

// LoadMesh ...
func (r *Renderer) LoadMeshFromFile(filepath string) (*primitives.Mesh, error) {
	p, err := r.dev.CompileProgram("assets/programs/lighting.vert", "assets/programs/lighting.frag")
	if err != nil {
		return nil, err
	}

	var args = []program.Arg{
		{Name: program.UModelName, Typ: program.Mat4Uniform, Dst: program.ModelDst},
		{Name: program.UViewName, Typ: program.Mat4Uniform, Dst: program.ViewDst},
		{Name: program.UProjectionName, Typ: program.Mat4Uniform, Dst: program.ProjectionDst},
		{Name: program.ULightColorName, Typ: program.Vec3Uniform, Dst: program.LightColorDst},
		{Name: program.ULightPositionName, Typ: program.Vec3Uniform, Dst: program.LightPosDst},
		{Name: program.UColorName, Typ: program.Vec4Uniform, Dst: program.ColorDst},
	}

	uniforms := program.NewUniforms(r.dev, p, args)
	oglMesh := ogl.NewMesh(r.dev, p, uniforms)

	model, err := resources.Load3ModelObj(filepath)
	if err != nil {
		return nil, errors.Wrap(err, rendererInfoTag)
	}

	content, indices := resources.ConvertObjToOGL(model)

	oglMesh.Load(r.dev, content, indices)
	n := primitives.NewNode()
	m := primitives.NewMesh(n, oglMesh)

	return m, nil
}

func (r *Renderer) LoadMeshFromData(data interface{}, indices []uint32) (*primitives.Mesh, error) {
	switch data.(type) {
	case []ogl.V3f:
		m, err := loadFromV3f(r.dev, data, indices)
		if err != nil {
			return nil, errors.Wrap(err, rendererInfoTag)
		}
		return m, nil
	case []ogl.V3fC4b:
		m, err := loadFromV3fC4b(r.dev, data, indices)
		if err != nil {
			return nil, errors.Wrap(err, rendererInfoTag)
		}
		return m, nil
	default:
		return nil, errors.New(rendererInfoTag + ": invalid mesh data type")
	}
	return nil, errors.New(rendererInfoTag + ": something went wrong")
}

// SetCamera ...
func (r *Renderer) SetCamera(c *scene.Camera) {
	r.camera = c
}

// NewTex ...
func (r *Renderer) NewTex(img image.Image) textures.Texture {
	tex := r.tl.NewTex(r.dev)
	r.tl.GlLoad(tex, img, r.dev)
	return tex
}

func (r *Renderer) DelTex(tex textures.Texture) {
	r.tl.DelTex(tex, r.dev)
}
