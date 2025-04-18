package internal

import (
	"github.com/hetue/axure/internal/internal"
	"github.com/hetue/boot"
)

func New(params internal.Steps) []boot.Step {
	return []boot.Step{
		params.Build,
	}
}
