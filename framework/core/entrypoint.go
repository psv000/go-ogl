package core

import (
	"time"

	"framework/graphics"
	"framework/graphics/materials"

	"github.com/pkg/errors"
)

type (
	// ReadySteady ...
	ReadySteady interface {
		PreloadAtlases(files []materials.AtlasPath) error
	}

	// EntryPoint ...
	EntryPoint struct {
		renderer *graphics.Renderer
		window   *graphics.WindowContext

		materials *materials.Manager

		app Application

		lastUpdate time.Time
	}
)

// NewEntryPoint ...
func NewEntryPoint() *EntryPoint {
	return &EntryPoint{}
}

// Serve ...
func (ep *EntryPoint) Serve(args ...interface{}) error {
	var converted bool
	ep.renderer, converted = args[0].(*graphics.Renderer)
	if !converted {
		return errors.New("core: entry point: invalid renderer")
	}
	ep.window, converted = args[1].(*graphics.WindowContext)
	if !converted {
		return errors.New("core: entry point: invalid window context")
	}
	winWidth, converted := args[2].(int)
	if !converted {
		return errors.New("core: entry point: invalid win width arg")
	}
	winHeight, converted := args[3].(int)
	if !converted {
		return errors.New("core: entry point: invalid win height arg")
	}

	ep.app, converted = args[4].(Application)
	if !converted {
		return errors.New("core: entry point: invalid application delegate")
	}

	ep.materials = materials.NewManager(ep.renderer)

	ep.app.OnEvent(Event{Et: Started, Args: ServicePack{
		Rs:  ep,
		Pl:  ep.renderer,
		Ml:  ep.materials,
		Set: ep.renderer,
	}})

	ep.window.Resize(winWidth, winHeight)

	ep.loop()
	return nil
}

// Stop ...
func (ep *EntryPoint) Stop() error {
	ep.app.OnEvent(Event{Et: Stopped})
	return nil
}

func (ep *EntryPoint) OnViewResize(w, h int) {
	ep.renderer.ResizeView(w, h)
	ep.app.ResizeView(w, h)
}

func (ep *EntryPoint) PreloadAtlases(files []materials.AtlasPath) error {
	for _, file := range files {
		if err := ep.materials.LoadAtlas(file); err != nil {
			return errors.Wrap(err, entrypointInfoTag)
		}
	}
	return nil
}

func (ep *EntryPoint) loop() {
	ep.lastUpdate = time.Now()
	for !ep.window.ShouldClose() {
		ep.app.Update(time.Since(ep.lastUpdate))
		ep.renderer.Update()
		ep.window.Update()

		ep.lastUpdate = time.Now()
	}
}
