// +build debug

package utils

import (
	ggl "github.com/go-gl/gl/v4.1-core/gl"
)

func CheckConversion(c bool) {
	if !c {
		panic("invalid conversion")
	}
}

func CheckOGLError() {
	err := ggl.GetError()
	if err != 0 {
		panic("gl err: "+string(err))
	}
}
