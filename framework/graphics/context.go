package graphics

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/pkg/errors"
)

type (
	viewEventsListener interface {
		OnViewResize(w, h int)
	}
	// WindowContext ...
	WindowContext struct {
		window *glfw.Window
		wl     viewEventsListener
	}
)

// NewWindowContext ...
func NewWindowContext() *WindowContext {
	return &WindowContext{}
}

// Serve ...
func (wc *WindowContext) Serve(args ...interface{}) error {
	if err := glfw.Init(); err != nil {
		return errors.Wrap(err, glfwContextInfoTag)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var width, height int
	var converted bool
	width, converted = args[0].(int)
	if !converted {
		return errors.New("graphics: window context: invalid width argument")
	}
	height, converted = args[1].(int)
	if !converted {
		return errors.New("graphics: window context: invalid height argument")
	}

	var title string
	if str, converted := args[2].(string); !converted {
		title = str
	}
	var err error
	wc.window, err = glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return errors.Wrap(err, glfwContextInfoTag)
	}
	wc.window.MakeContextCurrent()

	wc.wl, converted = args[3].(viewEventsListener)
	if !converted {
		return errors.New("graphics: window context: invalid win events listener argument")
	}
	wc.window.SetFramebufferSizeCallback(func(win *glfw.Window, w, h int) {
		wc.wl.OnViewResize(w, h)
	})
	return nil
}

// Stop ...
func (wc *WindowContext) Stop() error {
	glfw.Terminate()
	return nil
}

// Update ...
func (wc *WindowContext) Update() {
	glfw.PollEvents()
	wc.window.SwapBuffers()
}

// ShouldClose ...
func (wc *WindowContext) ShouldClose() bool {
	return wc.window.ShouldClose()
}

// Resize ...
func (wc *WindowContext) Resize(w, h int) {
	wc.window.SetSize(w, h)
	wc.wl.OnViewResize(w, h)
}
