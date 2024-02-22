package git

import (
	"fmt"
	"strings"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `git`,
	Summary:  `git extensions`,
	Commands: []*Z.Cmd{help.Cmd, pushCmd},
}

var pushCmd = &Z.Cmd{
	Name:     `push`,
	Summary:  `Push current branch to git`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		branch := Z.Out("git", "rev-parse", "--abbrev-ref", "HEAD")
    if err := Z.Exec("git", "push", "origin", strings.TrimSpace(branch)); err != nil {
      return err
    }
		commit := Z.Out("git", "log", "--oneline", "HEAD^..HEAD")
		fmt.Printf("pushed to %s\nlast commit:%s\n", branch, commit)
		return nil
	},
}
