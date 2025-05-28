package clip

import (
	"fmt"
	"io"
	"os"

	"github.com/atotto/clipboard"
	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var Cmd = &Z.Cmd{
	Name:     `clip`,
	Summary:  `Manage Clipboard`,
	Commands: []*Z.Cmd{help.Cmd, copyCmd, pasteCmd},
}

var copyCmd = &Z.Cmd{
	Name:     `copy`,
	Summary:  `Copy to clipboard`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		out, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		if err := clipboard.WriteAll(string(out)); err != nil {
			return err
		}
		return nil
	},
}

var pasteCmd = &Z.Cmd{
	Name:     `paste`,
	Summary:  `Paste from clipboard`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		out, err := clipboard.ReadAll()
		if err != nil {
			return err
		}
		fmt.Print(out)
		return nil
	},
}
