package env

import (
	"fmt"
	"os"
	"strings"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
	"github.com/rwxrob/bonzai/term"
)

var Cmd = &Z.Cmd{
	Name:     `env`,
	Short:  `commands for environment variables`,
	Cmds: []*Z.Cmd{getCmd, dataCmd, help.Cmd},
	Def: dataCmd,
}

var dataCmd = &Z.Cmd{
	Name:     `data`,
	Alias:  `all`,
	Short:  `print environment data to stdout`,
	Cmds: []*Z.Cmd{help.Cmd},
	Do: func(_ *Z.Cmd, _ ...string) error {
		for _, pair := range os.Environ() {
			fmt.Println(pair)
		}
		return nil
	},
}

var getCmd = &Z.Cmd{
	Name:     `get`,
	Usage:    `(help|NAME)`,
	Short:  `print specified environment variable to stdout`,
	Cmds: []*Z.Cmd{help.Cmd},
	NumArgs:  1,
	Do: func(_ *Z.Cmd, args ...string) error {
		v := os.Getenv(args[0])
		if v == "" {
			v = os.Getenv(strings.ToUpper(args[0]))
		}
		_, err := term.Print(v)
		return err
	},
}
