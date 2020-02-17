package materials

import "image"

type (
	// TexManager ...
	TexManager interface {
		NewTex(img image.Image) uint32
		DelTex(id uint32)
	}
)
