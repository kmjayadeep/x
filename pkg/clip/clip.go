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
	Short:  `manage Clipboard`,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden(), CopyCmd, PasteCmd},
}

var CopyCmd = &Z.Cmd{
	Name:     `copy`,
	Short:  `copy to clipboard (also a subdomannd of clip)`,
	Alias: `c`,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden()},
	Do: func(_ *Z.Cmd, args ...string) error {
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

var PasteCmd = &Z.Cmd{
	Name:     `paste`,
	Short:  `paste from clipboard`,
	Cmds: []*Z.Cmd{help.Cmd},
	Do: func(_ *Z.Cmd, args ...string) error {
		out, err := clipboard.ReadAll()
		if err != nil {
			return err
		}
		fmt.Print(out)
		return nil
	},
}
