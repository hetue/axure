package step

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newBuild,
	).Build().Apply()
}
