package main

import (
	"github.com/hetue/axure/internal"
	"github.com/hetue/boot"
)

func main() {
	bootstrap := boot.New()
	bootstrap.Build().Boot(internal.New)
}
