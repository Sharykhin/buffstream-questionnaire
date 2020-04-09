package model

type (
	// Stream represents stream model in question domain
	Stream struct {
		UUID      string
		Questions []Question
	}

	Streams map[string][]Question
)
