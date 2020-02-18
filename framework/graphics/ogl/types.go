package ogl

import (
	"framework/mth"
	"unsafe"
)

type (
	// VertexAttribute ...
	VertexAttribute int
	// VertexDataType ...
	VertexDataType int

	// VertexAttrOpt ...
	VertexAttrOpt struct {
		attr VertexAttribute
		typ  VertexDataType
		size int32
	}

	// V3f ...
	V3f struct {
		Point mth.Vec3f32
	}

	// V3fC4b ...
	V3fC4b struct {
		Point mth.Vec3f32
		Color mth.Vec4b
	}

	// V3fUV2f ...
	V3fUV2f struct {
		Point mth.Vec3f32
		UVS   mth.Vec2f32
	}

	// V3fUV2fN3f ...
	V3fUV2fN3f struct {
		Point mth.Vec3f32
		UVS   mth.Vec2f32
		Norm  mth.Vec3f32
	}
)

type (
	// DrawArgs ...
	DrawArgs struct {
		// Vao ...
		Vao uint32
		// IndLen ...
		IndLen int32
	}

	// CommandType ...
	CommandType int

	// Command ...
	Command struct {
		// Ct ...
		Ct CommandType
		// Args ...
		Args interface{}
	}

	// UniformType ...
	UniformType int

	// Destination ...
	Destination int

	// ProgramArg ...
	ProgramArg struct {
		// Name ...
		Name string
		// Typ ...
		Typ UniformType
		// Dst ...
		Dst Destination
	}
)

var (
	Float32Size = int(unsafe.Sizeof(float32(0)))
	Float64Size = int(unsafe.Sizeof(float64(0)))
	Int8Size    = int(unsafe.Sizeof(int8(0)))
	Int16Size   = int(unsafe.Sizeof(int16(0)))
	Int32Size   = int(unsafe.Sizeof(int32(0)))
	Int64Size   = int(unsafe.Sizeof(int64(0)))
	UInt8Size   = int(unsafe.Sizeof(uint8(0)))
	UInt16Size  = int(unsafe.Sizeof(uint16(0)))
	UInt32Size  = int(unsafe.Sizeof(uint32(0)))
	UInt64Size  = int(unsafe.Sizeof(uint64(0)))
	ByteSize    = int(unsafe.Sizeof(byte(0)))

	V3fSize        = int(unsafe.Sizeof(V3f{}))
	V3fC4bSize     = int(unsafe.Sizeof(V3fC4b{}))
	V3fUV2fSize    = int(unsafe.Sizeof(V3fUV2f{}))
	V3fUV2fN3fSize = int(unsafe.Sizeof(V3fUV2fN3f{}))
)
