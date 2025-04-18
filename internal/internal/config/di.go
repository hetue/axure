package config

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newProject, // 项目
	).Build().Apply()
}
