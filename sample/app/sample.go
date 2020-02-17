package app

import (
	"math"
	"time"

	"framework/core"
	"framework/graphics/materials"
	"framework/graphics/ogl"
	"framework/graphics/primitives"
	"framework/graphics/scene"
	"github.com/sirupsen/logrus"
)

var vertices = []ogl.V3fC4b{
	{[3]float32{-0.25, -0.25, -0.25}, [4]uint8{255, 0, 0, 255}},
	{[3]float32{0.25, -0.25, -0.25}, [4]uint8{0, 255, 0, 255}},
	{[3]float32{0.25, 0.25, -0.25}, [4]uint8{167, 76, 54, 255}},
	{[3]float32{-0.25, 0.25, -0.25}, [4]uint8{128, 187, 35, 255}},

	{[3]float32{-0.25, -0.25, 0.25}, [4]uint8{200, 21, 127, 255}},
	{[3]float32{0.25, -0.25, 0.25}, [4]uint8{31, 98, 0, 255}},
	{[3]float32{0.25, 0.25, 0.25}, [4]uint8{0, 128, 78, 255}},
	{[3]float32{-0.25, 0.25, 0.25}, [4]uint8{128, 0, 128, 255}},
}

var indices = []uint32{
	0, 1, 2, 0, 2, 3,
	1, 5, 6, 1, 6, 2,
	0, 1, 5, 0, 5, 4,
	0, 3, 7, 0, 7, 4,
	3, 2, 6, 3, 6, 7,
	4, 5, 6, 4, 6, 7,
}

type (
	// Sample ...
	Sample struct {
		pl primitives.Loader
		ml materials.Loader

		set scene.Set

		cam              *scene.Camera
		xcam, ycam, zcam float64

		mesh *primitives.Mesh
	}
)

// NewSample ...
func NewSample() *Sample {
	return &Sample{
		//xcam: -400,
		//ycam: -300,
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
	//s.cam.Rotation(rad(90.), rad(90.), 0.)

	var err error
	//s.mesh, err = s.pl.LoadMeshFromData(vertices, indices)
	//s.mesh.SetColor(mth.Vec4f32{0., 1., 0., 1.})
	s.mesh, err = s.pl.LoadMeshFromFile("assets/models/monkey-head.obj")
	//s.mesh.Node().Rotate(rad(45.), 0., 0.)

	if err != nil {
		logrus.Panic(err)
	}
}

// ResizeView ...
func (s *Sample) ResizeView(w, h int) {
	s.cam.SetView(w, h).Relocate(s.xcam, s.ycam, s.zcam)
}

var ms float64

// Update ...
func (s *Sample) Update(dt time.Duration) {
	ms += 0.00016
	var x, y, z float64
	x = (math.Sin(ms*100.) + math.Cos(ms*100.)) * 2.
	y = (math.Sin(ms*50.) + math.Cos(ms*50.)) * 2.
	x = (math.Sin(ms*25.) + math.Cos(ms*25.)) * 2.

	var sx, sy, sz = 1., 1., 1.
	sc := math.Sin(ms*100.)/2. + 2
	sx, sy, sz = sc, sc, sc

	var px, py, pz float64
	px = math.Sin(ms*100.) * 0.5

	_, _, _ = x, y, z
	_, _, _ = sx, sy, sz
	_, _, _ = px, py, pz

	s.mesh.Node().Rotate(x, y, z)
	//s.mesh.Node().Scale(sx, sy, sz)
	s.mesh.Node().Locate(px, py, pz)
}

func rad(x float64) float64 {
	return x * math.Pi / 180.
}
