package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newProject, // 项目
	).Build().Apply()
}
