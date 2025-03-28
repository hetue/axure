package main

import (
	"github.com/hetue/axure/internal"
	"github.com/hetue/core"
)

func main() {
	bootstrap := core.New()
	bootstrap.Build().Boot(internal.New)
}
