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
