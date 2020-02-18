package materials

import (
	"encoding/json"
	"image"
	"image/png"
	"os"
	"runtime"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type (
	// Manager ...
	atlas struct {
		name string
		uses int
	}

	// Manager ...
	Manager struct {
		tm TexManager

		materials map[string]Material
		atlases   map[uint32]atlas
	}
)

// NewManager ...
func NewManager(tm TexManager) *Manager {
	man := &Manager{
		materials: make(map[string]Material),
		atlases:   make(map[uint32]atlas),
		tm:        tm,
	}
	runtime.SetFinalizer(man, (*Manager).free)
	return man
}

func loadPng(filepath string) (image.Image, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close() // nolint: errcheck
	return png.Decode(f)
}

func loadMap(filepath string) (AtlasMapping, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return AtlasMapping{}, err
	}
	defer f.Close() // nolint: errcheck

	var m AtlasMapping
	err = json.NewDecoder(f).Decode(&m)
	return m, err
}

// LoadAtlas ...
func (m *Manager) LoadAtlas(filepath AtlasPath) error {
	data, err := loadPng(filepath.image)
	if err != nil {
		return errors.Wrap(err, materialManagerInfoTag)
	}
	mapping, err := loadMap(filepath.mapping)
	if err != nil {
		return errors.Wrap(err, materialManagerInfoTag)
	}

	texID := m.tm.NewTex(data)
	m.atlases[texID] = atlas{
		name: mapping.Name,
	}
	for _, mp := range mapping.Materials {
		m.materials[mp.Name] = Material{
			ID:    mp.Name,
			TexID: texID,
			X:     mp.X,
			Y:     mp.Y,
			Z:     mp.Z,
			W:     mp.W,
		}
	}
	return nil
}

// Get returns material
func (m *Manager) Get(name string) (Material, bool) {
	mat, found := m.materials[name]
	if found {
		atl, ok := m.atlases[mat.TexID]
		if !ok {
			return Material{}, false
		}
		atl.uses++
		m.atlases[mat.TexID] = atl
	}
	return mat, found
}

// Release decreases use count for material
func (m *Manager) Release(name string) {
	mat, found := m.materials[name]
	if found {
		atl, found := m.atlases[mat.TexID]
		if !found || atl.uses < 1 {
			return
		}
		atl.uses--
		m.atlases[mat.TexID] = atl
	}
}

func (m *Manager) free() {
	for texID, atl := range m.atlases {
		if atl.uses != 0 {
			logrus.Panicf(materialManagerInfoTag+": <%s> have not been released", atl.name)
		}
		m.tm.DelTex(texID)
	}
}
