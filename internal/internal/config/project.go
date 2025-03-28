package config

import (
	"github.com/pangum/pangu"
)

type Project struct {
	Dir string `json:"dir,omitempty"`
}

func newProject(config *pangu.Config) (project *Project, err error) {
	project = new(Project)
	err = config.Build().Get(project)

	return
}
