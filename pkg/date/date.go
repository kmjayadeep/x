package date

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "date",
	Aliases: []string{"d"},
	Short:   "date utility commands",
	RunE:    printMin,
}

var dateCmd = &cobra.Command{
	Use:   "min",
	Short: "display date in YYYY/MM/DD format",
	Args:  cobra.NoArgs,
	RunE:  printMin,
}

func printMin(_ *cobra.Command, _ []string) error {
	d := time.Now().Format("2006/01/02")
	fmt.Println(d)
	return nil
}

var dateFull = &cobra.Command{
	Use:   "full",
	Short: "display date in YYYY MMM DD format",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		d := time.Now().Format("2006 Jan 02")
		fmt.Println(d)
		return nil
	},
}

var dateTimeCmd = &cobra.Command{
	Use:   "datetime",
	Short: "display date and time in YYYY/MM/DD HH:MM format",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		d := time.Now().Format("2006/01/02 03:04PM")
		fmt.Println(d)
		return nil
	},
}

var DateHeadCmd = &cobra.Command{
	Use:     "datehead",
	Aliases: []string{"dh"},
	Short:   "display date in a human-readable form for headings",
	Args:    cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		d := time.Now().Format("2006 January 02, Monday")
		fmt.Println(d)
		return nil
	},
}

func init() { Cmd.AddCommand(dateCmd, dateTimeCmd, dateFull, DateHeadCmd) }
