package utils

type (
	// Server ...
	Server interface {
		Serve(args ...interface{}) error
		Stop() error
	}

	// Renew ...
	Renew interface {
		Update()
	}
)
