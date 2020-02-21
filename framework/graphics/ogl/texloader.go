package ogl

import (
	"framework/graphics/general"
	"framework/graphics/textures"
	"framework/graphics/utils"
	"image"
)

type (
	// OGLTexLoader ...
	OGLTexLoader struct {
	}
)

const (
	redComponent = iota
	greenComponent
	blueComponent
	alphaComponent
	colorComponents
)

func (tl *OGLTexLoader) NewTex(device general.Device) textures.Texture {
	oglDev, converted := device.(*OpenGL)
	utils.CheckConversion(converted)

	tex := textures.Texture{}
	tex.Gl = oglDev.NewTex()
	return tex
}

func (tl *OGLTexLoader) DelTex(tex textures.Texture, device general.Device) {
	glTex, converted := tex.Gl.(Tex2D)
	utils.CheckConversion(converted)
	oglDev, converted := device.(*OpenGL)
	utils.CheckConversion(converted)

	oglDev.DelTex(glTex)
}

func (tl *OGLTexLoader) GlLoad(tex textures.Texture, img image.Image, device general.Device) {
	glTex, converted := tex.Gl.(Tex2D)
	utils.CheckConversion(converted)
	oglDev, converted := device.(*OpenGL)
	utils.CheckConversion(converted)

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	ox, oy := bounds.Min.X, bounds.Min.Y
	bpp := 32
	length := (w - ox) * (h - oy)

	data := make([]byte, length*bpp)
	for i := ox; i < w; i++ {
		for j := oy; j < h; j++ {
			r, g, b, a := img.At(i, j).RGBA()
			position := i*w + j

			data[position*colorComponents+redComponent] = byte(r)
			data[position*colorComponents+greenComponent] = byte(g)
			data[position*colorComponents+blueComponent] = byte(b)
			data[position*colorComponents+alphaComponent] = byte(a)
		}
	}
	oglDev.UpdTex(glTex, 0, int32(w), int32(h), 0, data)
}
