package materials

import (
	"framework/graphics/general"
	"image"
)

type (
	// TexManager ...
	TexManager interface {
		NewTex(img image.Image) uint32
		DelTex(id uint32)
	}

	GlTexLoader interface {
		GlLoad(tex Texture, img image.Image, device general.Device)
	}

	Texture struct {
		Gl interface{}
	}
)
