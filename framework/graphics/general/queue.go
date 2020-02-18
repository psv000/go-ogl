package general

type (
	// Queue is a gl queue
	Queue interface {
		Process()
		Flush()
	}
)
