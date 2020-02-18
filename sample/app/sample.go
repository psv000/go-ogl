package app

import (
	"framework/mth"
	"math"
	"time"

	"framework/core"
	"framework/graphics/materials"
	"framework/graphics/primitives"
	"framework/graphics/scene"
	"github.com/sirupsen/logrus"
)

type (
	// Sample ...
	Sample struct {
		pl primitives.Loader
		ml materials.Loader

		set scene.Set

		cam              *scene.Camera
		xcam, ycam, zcam float64

		xl, yl, zl float32

		mg *primitives.MeshGroup
		light, cube, monkey *primitives.Mesh
	}
)

// NewSample ...
func NewSample() *Sample {
	return &Sample{
		xcam: -3,
		ycam: -3,
		zcam: -10.,
	}
}

// OnEvent ...
func (s *Sample) OnEvent(ev core.Event) {
	switch ev.Et {
	case core.Started:
		sp, ok := ev.Args.(core.ServicePack)
		if !ok {
			logrus.Panic("invalid service pack")
		}
		s.Start(sp)
	case core.Stopped:
	}
}

func (s *Sample) Start(sp core.ServicePack) {
	s.pl = sp.Pl
	s.ml = sp.Ml
	s.set = sp.Set

	sp.Rs.PreloadAtlases([]materials.AtlasPath{})

	s.cam = scene.NewCamera(scene.PerspectiveCT)
	s.set.SetCamera(s.cam)

	s.cam.Relocate(s.xcam, s.ycam, s.zcam)

	s.xl, s.yl, s.zl = -1.5, 0, 0

	var err error
	s.light, err = s.pl.LoadMeshFromFile("assets/models/cube.obj")
	if err != nil {
		logrus.Panic(err)
	}
	s.light.Node().Locate(float64(s.xl), float64(s.yl), float64(s.zl))
	s.light.Node().Resize(6, 6, 6)
	s.light.Node().Scale(0.2, 0.2, 0.2)

	s.cube, err = s.pl.LoadMeshFromFile("assets/models/cube.obj")
	if err != nil {
		logrus.Panic(err)
	}
	s.cube.SetColor(mth.Vec4f32{1., 0., 0, 1.})
	s.cube.Node().Locate(1.8, 1.2, 2.1)

	s.monkey, err = s.pl.LoadMeshFromFile("assets/models/monkey-head.obj")
	if err != nil {
		logrus.Panic(err)
	}
	s.monkey.SetColor(mth.Vec4f32{0., 1., 0., 1.})

	s.mg, err = s.pl.NewMeshGroup()
	if err != nil {
		logrus.Panic(err)
	}
	s.mg.Meshes = append(s.mg.Meshes, s.monkey)
	s.mg.Meshes = append(s.mg.Meshes, s.cube)
	s.mg.Meshes = append(s.mg.Meshes, s.light)

	s.mg.LightSources = append(s.mg.LightSources, primitives.LightSource{
		Pos: mth.Vec3f32{s.xl, s.yl, s.zl},
		Col: mth.Vec3f32{1., 1., 1.},
	})
}

// ResizeView ...
func (s *Sample) ResizeView(w, h int) {
	s.cam.SetView(w, h).Relocate(s.xcam, s.ycam, s.zcam)
	s.cam.Rotation(rad(-5), rad(8), 0)
}

var ms float64

// Update ...
func (s *Sample) Update(dt time.Duration) {
	ms += 0.00016
	var x, y, z float64
	//x = (math.Sin(ms*100.) + math.Cos(ms*100.)) * 2.
	y = (math.Sin(ms*50.) + math.Cos(ms*50.)) * 2.
	//z = (math.Sin(ms*25.) + math.Cos(ms*25.)) * 2.

	var sx, sy, sz = 1., 1., 1.
	sc := math.Sin(ms*100.)/2. + 2
	sx, sy, sz = sc, sc, sc

	var px, py, pz float64
	px = math.Sin(ms*100.)*0.5 - 0.5
	py = math.Sin(ms*100.)*0.5 - 0.5
	pz = math.Sin(ms*100.)*0.5 - 2.1

	_, _, _ = x, y, z
	_, _, _ = sx, sy, sz
	_, _, _ = px, py, pz

	s.cube.Node().Rotate(x, y, z)
	s.cube.Node().Locate(1.8, 1.2, pz)
	s.monkey.Node().Locate(0, py, 0)

}

func rad(x float64) float64 {
	return x * math.Pi / 180.
}
