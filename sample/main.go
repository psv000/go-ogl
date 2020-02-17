package main

import (
	"runtime"

	"sample/app"

	"framework/core"
	"framework/graphics"
	"framework/graphics/ogl"

	"github.com/sirupsen/logrus"
)

const (
	defaultWinWidth  = 800
	defaultWinHeight = 600

	sampleName = "sample"
)

func terminate(err error) {
	if err != nil {
		logrus.Panic(err)
	}
}

func main() {
	runtime.LockOSThread()

	ep := core.NewEntryPoint()
	windowContext := graphics.NewWindowContext()
	renderer := graphics.NewRenderer()

	terminate(windowContext.Serve(defaultWinWidth, defaultWinHeight, sampleName, ep))

	glDevice := ogl.NewOpenGL()
	glQueue := ogl.NewQueue(glDevice)
	glArtist := ogl.NewArtist(glQueue)

	terminate(renderer.Serve(glDevice, glQueue, glArtist))
	terminate(ep.Serve(renderer, windowContext, defaultWinWidth, defaultWinHeight, app.NewSample())) // inf loop

	terminate(ep.Stop())
	terminate(renderer.Stop())
	terminate(windowContext.Stop())
}
