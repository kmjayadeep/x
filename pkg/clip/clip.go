package clip

import (
	"fmt"
	"io"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "clip",
	Short: "manage clipboard",
}

var CopyCmd = &cobra.Command{
	Use:     "copy",
	Aliases: []string{"c"},
	Short:   "copy stdin to clipboard",
	Args:    cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
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

var PasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "paste from clipboard",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		out, err := clipboard.ReadAll()
		if err != nil {
			return err
		}
		fmt.Print(out)
		return nil
	},
}

func init() { Cmd.AddCommand(CopyCmd, PasteCmd) }
