package internal

import (
	"github.com/hetue/axure/internal/internal/step"
	"github.com/pangum/pangu"
)

type Steps struct {
	pangu.Get

	Build *step.Build
}
