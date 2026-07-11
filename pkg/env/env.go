package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "env",
	Short: "commands for environment variables",
	RunE:  printData,
}

var dataCmd = &cobra.Command{
	Use:     "data",
	Aliases: []string{"all"},
	Short:   "print environment data to stdout",
	Args:    cobra.NoArgs,
	RunE:    printData,
}

func printData(_ *cobra.Command, _ []string) error {
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}
	return nil
}

var getCmd = &cobra.Command{
	Use:   "get NAME",
	Short: "print specified environment variable to stdout",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		v := os.Getenv(args[0])
		if v == "" {
			v = os.Getenv(strings.ToUpper(args[0]))
		}
		_, err := fmt.Print(v)
		return err
	},
}

func init() { Cmd.AddCommand(getCmd, dataCmd) }
