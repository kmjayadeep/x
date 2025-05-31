package git

import (
	"fmt"
	"io"
	"os"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var filtersCmd = &Z.Cmd{
	Name:     `filter`,
	Short:  `filters useful for git`,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden(), tfFilter, codeFilter},
}

var tfFilter = &Z.Cmd{
	Name:     `tf`,
	Short  : `wrap terraform plan in markdown`,
	Long:  `wrap terraform plan input around <detail> tag with summary`,
	MaxArgs:  0,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden()},
	Do: func(x *Z.Cmd, args ...string) error {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		fmt.Println(
			"<details><summary>Terraform Plan</summary>\n\n" +
				"```hcl\n" +
				string(stdin) +
				"\n```\n" +
				"</details>")
		return nil
	},
}

var codeFilter = &Z.Cmd{
	Name:     `code`,
	Short:  `wrap input around markcode code syntax`,
	MaxArgs:  1,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden()},
	Do: func(x *Z.Cmd, args ...string) error {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		lang := "bash"
		if len(args) > 0 {
			lang = args[0]
		}
		fmt.Println(
			"```" + lang + "\n" +
				string(stdin) +
				"\n```")
		return nil
	},
}
