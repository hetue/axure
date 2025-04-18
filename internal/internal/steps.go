package internal

import (
	"github.com/harluo/di"
	"github.com/hetue/axure/internal/internal/step"
)

type Steps struct {
	di.Get

	Build *step.Build
}
