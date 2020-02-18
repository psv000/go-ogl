package core

import (
	"time"

	"framework/graphics/materials"
	"framework/graphics/primitives"
	"framework/graphics/scene"
)

const (
	// Started ...
	Started EventType = iota
	// Stopped ...
	Stopped
)

type (
	// EventType ...
	EventType int
	// Event ...
	Event struct {
		// Et ...
		Et EventType
		// Args ...
		Args interface{}
	}

	// Application ...
	Application interface {
		OnEvent(ev Event)
		Update(dt time.Duration)
		ResizeView(w, h int)
	}

	// ServicePack ...
	ServicePack struct {
		Rs ReadySteady

		Pl primitives.Loader
		Ml materials.Loader

		Set scene.Set
	}
)
