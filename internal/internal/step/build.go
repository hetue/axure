package step

import (
	"context"

	"github.com/goexl/gfx"
	"github.com/hetue/axure/internal/internal/config"
	"github.com/hetue/core"
)

var _ core.Step = (*Build)(nil)

type Build struct {
	config  *config.Project
	command *core.Command
}

func newBuild(config *config.Project, command *core.Command) *Build {
	return &Build{
		config:  config,
		command: command,
	}
}

func (b *Build) Name() string {
	return "编译"
}

func (b *Build) Runnable() (runnable bool) { // 只在有文件时执行
	if _, exist := gfx.Exists().Build().Check(); exist {
		runnable = true
	}

	return
}

func (b *Build) Retryable() bool { // 不重试
	return false
}

func (b *Build) Asyncable() bool { // 不异步
	return false
}

func (b *Build) Run(ctx *context.Context) (err error) {
	if _, ee := b.command.New("axure").Context(*ctx).Dir(b.config.Dir).Build().Exec(); nil != ee {
		err = ee
	}

	return
}
