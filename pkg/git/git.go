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
	Commands: []*Z.Cmd{help.Cmd, pushCmd, pullCmd},
}

var pushCmd = &Z.Cmd{
	Name:     `push`,
	Summary:  `Push current branch to git`,
	Usage:    `[help|force]`,
	MaxArgs:  1,
	Params:   []string{"force"},
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		force := ""
		if len(args) > 0 && args[0] == "force" {
			force = "-f"
		}
		branch := Z.Out("git", "rev-parse", "--abbrev-ref", "HEAD")
		if force != "" {
			if err := Z.Exec("git", "push", "origin", strings.TrimSpace(branch), force); err != nil {
				return err
			}
		} else {
			if err := Z.Exec("git", "push", "origin", strings.TrimSpace(branch)); err != nil {
				return err
			}
		}
		commit := Z.Out("git", "log", "--oneline", "HEAD^..HEAD")
		fmt.Printf("pushed to %s\nlast commit:%s\n", branch, commit)
		return nil
	},
}

var pullCmd = &Z.Cmd{
	Name:     `pull`,
	Summary:  `Pull current branch from git`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		branch := Z.Out("git", "rev-parse", "--abbrev-ref", "HEAD")
		if err := Z.Exec("git", "pull", "origin", strings.TrimSpace(branch)); err != nil {
			return err
		}
		return nil
	},
}
