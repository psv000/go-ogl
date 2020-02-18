// +build debug

package ogl

import ggl "github.com/go-gl/gl/v4.1-core/gl"

func checkConversion(c bool) {
	if !c {
		panic(renderQueueInfoTag + ": invalid conversion")
	}
}

func checkOGLError() {
	err := ggl.GetError()
	if err != 0 {
		panic(renderQueueInfoTag + string(err))
	}
}
