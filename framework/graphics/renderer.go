package graphics

import (
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

		meshes []*primitives.Mesh
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
	var t = unitMat
	if r.camera != nil {
		t = r.camera.View().Mult(r.camera.Projection())
	}
	for _, m := range r.meshes {
		m.Draw(r.artist, t)
	}

	r.queue.Process()
	r.queue.Flush()
}

// LoadMesh ...
func (r *Renderer) LoadMeshFromFile(filepath string) (*primitives.Mesh, error) {
	p, err := r.dev.CompileProgram("assets/programs/uv.vert", "assets/programs/uv.frag")
	if err != nil {
		return nil, err
	}

	var args = []ogl.ProgramArg{
		{Name: ogl.UTransformNameDefault, Typ: ogl.Mat4Uniform, Dst: ogl.MVPDst},
		{Name: ogl.UColorNameDefault, Typ: ogl.Vec4Uniform, Dst: ogl.ColorDst},
	}

	uniforms := ogl.NewUniforms(r.dev, p, args)
	oglMesh := ogl.NewMesh(r.dev, p, uniforms)

	model, err := resources.Load3ModelObj(filepath)
	if err != nil {
		return nil, errors.Wrap(err, rendererInfoTag)
	}

	content, indices := resources.ConvertObjToOGL(model)

	oglMesh.Load(r.dev, content, indices)
	n := primitives.NewNode()
	m := primitives.NewMesh(n, oglMesh)

	r.meshes = append(r.meshes, m)
	return m, nil
}

func (r *Renderer) LoadMeshFromData(data interface{}, indices []uint32) (*primitives.Mesh, error) {
	switch data.(type) {
	case []ogl.V3f:
		m, err := loadFromV3f(r.dev, data, indices)
		if err != nil {
			return nil, errors.Wrap(err, rendererInfoTag)
		}
		r.meshes = append(r.meshes, m)
		return m, nil
	case []ogl.V3fC4b:
		m, err := loadFromV3fC4b(r.dev, data, indices)
		if err != nil {
			return nil, errors.Wrap(err, rendererInfoTag)
		}
		r.meshes = append(r.meshes, m)
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
func (r *Renderer) NewTex(img image.Image) uint32 {
	return 0
}

func (r *Renderer) DelTex(id uint32) {

}

func loadFromV3f(dev general.Device, data interface{}, indices []uint32) (*primitives.Mesh, error) {
	const (
		defaultVert = "assets/programs/test.vert"
		defaultFrag = "assets/programs/test.frag"
	)

	p, err := dev.CompileProgram(defaultVert, defaultFrag)
	if err != nil {
		return nil, err
	}

	var args = []ogl.ProgramArg{
		{Name: ogl.UTransformNameDefault, Typ: ogl.Mat4Uniform, Dst: ogl.MVPDst},
	}

	uniforms := ogl.NewUniforms(dev, p, args)
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

	var args = []ogl.ProgramArg{
		{Name: ogl.UTransformNameDefault, Typ: ogl.Mat4Uniform, Dst: ogl.MVPDst},
	}

	uniforms := ogl.NewUniforms(dev, p, args)
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
