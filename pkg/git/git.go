package git

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "git",
	Short: "git extensions",
}

func init() { Cmd.AddCommand(filtersCmd) }
