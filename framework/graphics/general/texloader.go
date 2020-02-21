package general

import (
	"image"

	"framework/graphics/textures"
)

type (
	GlTexLoader interface {
		NewTex(device Device) textures.Texture
		DelTex(tex textures.Texture, device Device)
		GlLoad(tex textures.Texture, img image.Image, device Device)
	}
)
