package init

import (
	"promAgent/plugins"
	"promAgent/tasks"
)

func init() {
	tasks.Register("register", &plugins.Register{})
}
