package git

import (
	"fmt"
	"io"
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var filtersCmd = &Z.Cmd{
	Name:     `filter`,
	Summary:  `Filters useful for git`,
	Commands: []*Z.Cmd{help.Cmd, tfFilter, codeFilter},
}

var tfFilter = &Z.Cmd{
	Name:     `tf`,
	Summary:  `Wrap terraform plan input around <detail> tag with summary`,
	MaxArgs:  0,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		fmt.Println(
			"<details><summary>Terraform Plan</summary>\n" +
				"```hcl\n" +
				string(stdin) +
				"\n```\n" +
				"</details>")
		return nil
	},
}

var codeFilter = &Z.Cmd{
	Name:     `code`,
	Summary:  `Wrap input around markcode code syntax`,
	MaxArgs:  1,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
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
