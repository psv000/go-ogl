package ogl

const (
	graphicsPacketTag   = "graphics"
	openGLDeviceInfoTag = "OpenGL device"
	renderQueueInfoTag  = "render Queue"
)

const (
	// VAPosition ...
	VAPosition VertexAttribute = iota
	// VAColor ...
	VAColor
	// VATexture ...
	VATexture
	// VANormal ...
	VANormal
	// VAShift ...
	VAShift
	// VAGamma ...
	VAGamma
	// VAUnknown ...
	VAUnknown
)

const (
	// VDTUnsignedShort ...
	VDTUnsignedShort VertexDataType = iota
	// VDTByte ...
	VDTByte
	// VDTFloat ...
	VDTFloat
)

const (
	NoCmd CommandType = iota
	ApplyProgramCmd
	ApplyUniformsCmd
	DrawMeshCmd
	ClearContextCmd
	NewTex2DCmd
	UpdTex2DCmd

	CmdLenDefault = 1024
)

const (
	// UnkUniform ...
	UnkUniform UniformType = iota
	// FltUniform ...
	FltUniform
	// IntUniform ...
	IntUniform
	// Vec4Uniform ...
	Vec4Uniform
	// Vec3Uniform ...
	Vec3Uniform
	// Vec2Uniform ...
	Vec2Uniform
	// Mat4Uniform ...
	Mat4Uniform
	// Tex2DUniform ...
	Tex2DUniform
)

const (
	// UnkDst ...
	UnkDst Destination = iota
	// MVPDst ...
	MVPDst
	// ColorDst ...
	ColorDst
)

const (
	UTransformNameDefault = "uModelViewProjection"
	UColorNameDefault     = "uColor"
)
