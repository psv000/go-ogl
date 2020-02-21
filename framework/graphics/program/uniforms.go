package program

import (
	"framework/graphics/utils"
	"framework/mth"
)

type (
	uniformFabric interface {
		CreateUniforms(programID uint32, args []Arg) []Uniform
	}

	// Uniforms ...
	Uniform struct {
		Loc int32
		Dst Destination
		Typ Type
		Arg interface{}
	}
)

// NewUniform ...
func NewUniforms(fabric uniformFabric, programID uint32, args []Arg) []Uniform {
	return fabric.CreateUniforms(programID, args)
}

// ApplyUniform ...
func (u *Uniform) ApplyUniform(typ Type, arg interface{}) {
	u.Typ = typ
	u.Arg = arg
}

// Unpack ...
func (u *Uniform) Unpack() (int32, Type) {
	return u.Loc, u.Typ
}

// ArgFlt ...
func (u *Uniform) ArgFlt() float64 {
	val, ok := u.Arg.(float64)
	utils.CheckConversion(ok)
	return val
}

// ArgInt ...
func (u *Uniform) ArgInt() int {
	val, ok := u.Arg.(int)
	utils.CheckConversion(ok)
	return val
}

// ArgV4 ...
func (u *Uniform) ArgV4() mth.Vec4f32 {
	val, ok := u.Arg.(mth.Vec4f32)
	utils.CheckConversion(ok)
	return val
}

// ArgV3 ...
func (u *Uniform) ArgV3() mth.Vec3f32 {
	val, ok := u.Arg.(mth.Vec3f32)
	utils.CheckConversion(ok)
	return val
}

// ArgV2 ...
func (u *Uniform) ArgV2() mth.Vec2f32 {
	val, ok := u.Arg.(mth.Vec2f32)
	utils.CheckConversion(ok)
	return val
}

// ArgM4 ...
func (u *Uniform) ArgM4() mth.Mat4f {
	val, ok := u.Arg.(mth.Mat4f)
	utils.CheckConversion(ok)
	return val
}
