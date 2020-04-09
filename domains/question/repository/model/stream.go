package model

type (
	// Stream represents how stream model looks in question domain
	Stream struct {
		UUID      string
		Questions []Question
	}
)
