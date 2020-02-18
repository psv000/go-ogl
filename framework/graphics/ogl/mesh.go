package ogl

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

type (
	// Mesh ...
	Mesh struct {
		vao, vbo, ibo, program uint32
		indicesLen             int32
		uniforms               []Uniform

		dev *OpenGL
	}
)

// NewMesh ...
func NewMesh(device interface{}, program uint32, uniforms []Uniform) *Mesh {
	oglDevice, converted := device.(*OpenGL)
	if !converted {
		logrus.Panic(openGLDeviceInfoTag + ": invalid ogl device")
	}
	m := &Mesh{
		vao:      oglDevice.NewVAO(),
		vbo:      oglDevice.NewVBO(),
		ibo:      oglDevice.NewIBO(),
		program:  program,
		uniforms: uniforms,
		dev:      oglDevice,
	}
	runtime.SetFinalizer(m, (*Mesh).free)
	return m
}

func (m *Mesh) free() {
	DelMesh(m)
}

// DelMesh ...
func DelMesh(m *Mesh) {
	m.dev.DelVAO(m.vao)
	m.dev.DelVBO(m.vbo)
	m.dev.DelIBO(m.ibo)
}

// Load ...
func (m *Mesh) Load(device interface{}, content interface{}, indices []uint32) {
	oglDevice, converted := device.(*OpenGL)
	if !converted {
		logrus.Panic(openGLDeviceInfoTag + ": invalid ogl device")
	}

	m.indicesLen = int32(len(indices))

	var size int
	var opts []VertexAttrOpt
	switch v := content.(type) {
	case []V3f:
		size = len(v) * V3fSize
		opts = append(opts, VertexAttrOpt{attr: VAPosition, typ: VDTFloat, size: 3})
	case []V3fC4b:
		size = len(v) * V3fC4bSize
		opts = append(opts, VertexAttrOpt{attr: VAPosition, typ: VDTFloat, size: 3})
		opts = append(opts, VertexAttrOpt{attr: VAColor, typ: VDTByte, size: 4})
	case []V3fUV2f:
		size = len(v) * V3fUV2fSize
		opts = append(opts, VertexAttrOpt{attr: VAPosition, typ: VDTFloat, size: 3})
		opts = append(opts, VertexAttrOpt{attr: VATexture, typ: VDTFloat, size: 3})
	case []V3fUV2fN3f:
		size = len(v) * V3fUV2fN3fSize
		opts = append(opts, VertexAttrOpt{attr: VAPosition, typ: VDTFloat, size: 3})
		opts = append(opts, VertexAttrOpt{attr: VATexture, typ: VDTFloat, size: 2})
		opts = append(opts, VertexAttrOpt{attr: VANormal, typ: VDTFloat, size: 3})
	}

	oglDevice.UpdVBO(m.vbo, content, size)
	oglDevice.UpdIBO(m.ibo, indices)
	oglDevice.UpdVAO(m.vao, m.vbo, m.ibo, opts)
}
