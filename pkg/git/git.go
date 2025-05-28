package git

import (
	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var Cmd = &Z.Cmd{
	Name:     `git`,
	Short:  `git extensions`,
	Cmds: []*Z.Cmd{help.Cmd, filtersCmd},
}
