package git

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var filtersCmd = &cobra.Command{
	Use:   "filter",
	Short: "filters useful for git",
}

var tfFilter = &cobra.Command{
	Use:   "tf",
	Short: "wrap terraform plan in markdown",
	Long:  "Wrap terraform plan input in a <details> tag with a summary.",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
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

var codeFilter = &cobra.Command{
	Use:   "code [language]",
	Short: "wrap input in a Markdown code block",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
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

func init() { filtersCmd.AddCommand(tfFilter, codeFilter) }
