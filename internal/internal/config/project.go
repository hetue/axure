package config

import (
	"github.com/harluo/boot"
)

type Project struct {
	Dir string `json:"dir,omitempty"`
}

func newProject(config *boot.Config) (project *Project, err error) {
	project = new(Project)
	err = config.Build().Get(project)

	return
}
