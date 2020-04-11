package uuid

//go:generate mockery -name Generator

import (
	"github.com/nu7hatch/gouuid"
	"log"
)

type (
	// Generator is UUID generator
	Generator interface {
		NewV4() string
	}

	GouuidAdapter struct {
	}
)

func (g *GouuidAdapter) NewV4() string {
	u, err := uuid.NewV4()
	if err != nil {
		log.Panicf("failed to generate uuid v4: %v", err)
	}

	return u.String()
}

func NewGoouidAdapter() *GouuidAdapter {
	return &GouuidAdapter{}
}
