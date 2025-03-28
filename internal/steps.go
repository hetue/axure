package internal

import (
	"github.com/hetue/axure/internal/internal"
	"github.com/hetue/core"
)

func New(params internal.Steps) []core.Step {
	return []core.Step{
		params.Build,
	}
}
