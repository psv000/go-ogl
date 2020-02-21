package textures

import (
	"image"
)

type (
	// TexManager ...
	TexManager interface {
		NewTex(img image.Image) Texture
		DelTex(tex Texture)
	}

	Texture struct {
		Gl interface{}
	}
)
