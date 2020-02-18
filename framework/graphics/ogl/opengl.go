package ogl

import (
	"framework/graphics/program"
	"framework/graphics/utils"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	strTerminator = "\x00"
)

type (
	// OpenGL ...
	OpenGL struct {
	}
)

// NewOpenGL ...
func NewOpenGL() *OpenGL {
	return &OpenGL{}
}

// Init ...
func (ogl *OpenGL) Init() error {
	if err := gl.Init(); err != nil {
		return errors.Wrap(err, openGLDeviceInfoTag)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	logrus.Infof("openGL version %s", version)

	gl.Enable(gl.DEPTH_TEST)
	//gl.Enable(gl.CULL_FACE)
	return nil
}

// ClrScr ...
func (ogl *OpenGL) ClrScr() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// CompileProgram ...
func (ogl *OpenGL) CompileProgram(vertex, fragment string) (uint32, error) {
	vSource, err := readFile(vertex)
	if err != nil {
		return 0, errors.Wrap(err, openGLDeviceInfoTag)
	}
	vertexShader, err := compileShader(vSource+strTerminator, gl.VERTEX_SHADER)
	if err != nil {
		return 0, errors.Wrap(err, openGLDeviceInfoTag)
	}

	fSource, err := readFile(fragment)
	if err != nil {
		return 0, errors.Wrap(err, openGLDeviceInfoTag)
	}
	fragmentShader, err := compileShader(fSource+strTerminator, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, errors.Wrap(err, openGLDeviceInfoTag)
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	return program, nil
}

// CreateUniforms ...
func (ogl *OpenGL) CreateUniforms(programID uint32, args []program.Arg) []program.Uniform {
	var uniforms []program.Uniform
	for _, arg := range args {
		loc := gl.GetUniformLocation(programID, gl.Str(arg.Name+strTerminator))
		uniforms = append(uniforms, program.Uniform{
			Loc: loc,
			Dst: arg.Dst,
			Typ: arg.Typ,
		})
	}
	return uniforms
}

// ApplyProgram ...
func (ogl *OpenGL) ApplyUniform(u program.Uniform) {
	loc, typ := u.Unpack()
	switch typ {
	case program.FltUniform:
		val := u.ArgFlt()
		gl.Uniform1f(loc, float32(val))
	case program.IntUniform:
		val := u.ArgInt()
		gl.Uniform1i(loc, int32(val))
	case program.Vec2Uniform:
		val := u.ArgV2()
		v1, v2 := val[0], val[1]
		gl.Uniform2f(loc, v1, v2)
	case program.Vec3Uniform:
		val := u.ArgV3()
		v1, v2, v3 := val[0], val[1], val[2]
		gl.Uniform3f(loc, v1, v2, v3)
	case program.Vec4Uniform:
		val := u.ArgV4()
		v1, v2, v3, v4 := val[0], val[1], val[2], val[3]
		gl.Uniform4f(loc, v1, v2, v3, v4)
	case program.Mat4Uniform:
		const matValCount = 16
		val := u.ArgM4()
		var glMat [matValCount]float32
		for i, v := range val.Values() {
			glMat[i] = v
		}
		gl.UniformMatrix4fv(u.Loc, 1, false, &glMat[0])
	//case program.Tex2DUniform:
	//	val := u.ArgTex2D()
	//	gl.ActiveTexture(gl.TEXTURE0 + uint32(val.lev))
	//	gl.BindTexture(gl.TEXTURE_2D, val.id)
	//	gl.Uniform1i(u.Loc, val.lev)
	}
}

// ApplyProgram ...
func (ogl *OpenGL) ApplyProgram(program uint32) {
	gl.UseProgram(program)
}

// NewVBO ...
func (ogl *OpenGL) NewVBO() uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	return vbo
}

// NewVAO ...
func (ogl *OpenGL) NewVAO() uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	return vao
}

// NewIBO ...
func (ogl *OpenGL) NewIBO() uint32 {
	var ibo uint32
	gl.GenBuffers(1, &ibo)
	return ibo
}

// DelVBO ...
func (ogl *OpenGL) DelVBO(vbo uint32) {
	gl.DeleteBuffers(1, &vbo)
}

// DelVAO ...
func (ogl *OpenGL) DelVAO(vao uint32) {
	gl.DeleteVertexArrays(1, &vao)
}

// DelIBO ...
func (ogl *OpenGL) DelIBO(ibo uint32) {
	gl.DeleteBuffers(1, &ibo)
}

// UpdVBO ...
func (ogl *OpenGL) UpdVBO(vbo uint32, data interface{}, size int) {
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(data), gl.STATIC_DRAW)
	utils.CheckOGLError()
}

// UpdIBO ...
func (ogl *OpenGL) UpdIBO(ibo uint32, indices []uint32) {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, UInt32Size*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)
	utils.CheckOGLError()
}

// UpdVAO ...
func (ogl *OpenGL) UpdVAO(vao, vbo, ibo uint32, opts []VertexAttrOpt) {
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	var stride int32
	for _, opt := range opts {
		switch opt.typ {
		case VDTUnsignedShort:
			stride += int32(UInt16Size) * opt.size
		case VDTByte:
			stride += int32(ByteSize) * opt.size
		case VDTFloat:
			stride += int32(Float32Size) * opt.size
		default:
			logrus.Panic(openGLDeviceInfoTag + ": unhandled opt type")
		}
	}

	var typ uint32
	var offset int
	var norm bool
	for _, opt := range opts {
		optOffset := offset
		norm = false

		switch opt.typ {
		case VDTUnsignedShort:
			typ = gl.UNSIGNED_SHORT
			offset += int(int32(UInt16Size) * opt.size)
			norm = true
		case VDTByte:
			typ = gl.UNSIGNED_BYTE
			offset += int(int32(ByteSize) * opt.size)
			norm = true
		case VDTFloat:
			typ = gl.FLOAT
			offset += int(int32(Float32Size) * opt.size)
			norm = false
		default:
			logrus.Panic(openGLDeviceInfoTag + ": unhandled opt type")
		}

		gl.VertexAttribPointer(uint32(opt.attr), opt.size, typ, norm, stride, gl.PtrOffset(optOffset))
		gl.EnableVertexAttribArray(uint32(opt.attr))
		utils.CheckOGLError()
	}

	if ibo != 0 {
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	}
}

// NewTex ...
func (ogl *OpenGL) NewTex() uint32 {
	var id uint32
	gl.GenTextures(1, &id)
	gl.BindTexture(gl.TEXTURE_2D, id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	return id
}

// UpdTex ...
func (ogl *OpenGL) UpdTex(id uint32, level, glInternal, width, height, border int32, glFormat, xtype uint32, data interface{}) {
	gl.BindTexture(gl.TEXTURE_2D, id)
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	gl.TexImage2D(gl.TEXTURE_2D,
		level,
		glInternal,
		width, height, border,
		glFormat, xtype,
		gl.Ptr(data.(uint8))) // nolint: govet
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

// DelTex ...
func (ogl *OpenGL) DelTex(id uint32) {
	gl.DeleteTextures(1, &id)
}

// Draw ...
func (ogl *OpenGL) Draw(vao uint32, indLen int32) {
	gl.BindVertexArray(vao)
	gl.DrawElements(gl.TRIANGLES, indLen, gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.BindVertexArray(0)
}

// DrawArray ...
func (ogl *OpenGL) DrawArray(vao uint32, dataLen int32) {
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, dataLen)
}
