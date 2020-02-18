package ogl

import (
	"framework/graphics/general"
	"framework/mth"
	"github.com/sirupsen/logrus"
)

type (
	// Uniforms ...
	Uniform struct {
		loc int32
		dst Destination
		typ UniformType
		arg interface{}
	}
)

// NewUniform ...
func NewUniforms(device general.Device, programID uint32, args []ProgramArg) []Uniform {
	oglDevice, converted := device.(*OpenGL)
	if !converted {
		logrus.Panic(openGLDeviceInfoTag + ": invalid ogl device")
	}
	return oglDevice.CreateUniforms(programID, args)
}

// ApplyUniform ...
func (u *Uniform) ApplyUniform(typ UniformType, arg interface{}) {
	u.typ = typ
	u.arg = arg
}

// Unpack ...
func (u *Uniform) Unpack() (int32, UniformType) {
	return u.loc, u.typ
}

// ArgFlt ...
func (u *Uniform) ArgFlt() float64 {
	val, ok := u.arg.(float64)
	checkConversion(ok)
	return val
}

// ArgInt ...
func (u *Uniform) ArgInt() int {
	val, ok := u.arg.(int)
	checkConversion(ok)
	return val
}

// ArgV4 ...
func (u *Uniform) ArgV4() mth.Vec4f32 {
	val, ok := u.arg.(mth.Vec4f32)
	checkConversion(ok)
	return val
}

// ArgV3 ...
func (u *Uniform) ArgV3() mth.Vec3f32 {
	val, ok := u.arg.(mth.Vec3f32)
	checkConversion(ok)
	return val
}

// ArgV2 ...
func (u *Uniform) ArgV2() mth.Vec2f32 {
	val, ok := u.arg.(mth.Vec2f32)
	checkConversion(ok)
	return val
}

// ArgM4 ...
func (u *Uniform) ArgM4() mth.Mat4f {
	val, ok := u.arg.(mth.Mat4f)
	checkConversion(ok)
	return val
}

// ArgTex2D ...
func (u *Uniform) ArgTex2D() Tex2D {
	val, ok := u.arg.(Tex2D)
	checkConversion(ok)
	return val
}
