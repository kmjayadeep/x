package main

import (
	"fmt"
	"os"

	"github.com/kmjayadeep/x/pkg/clip"
	"github.com/kmjayadeep/x/pkg/date"
	"github.com/kmjayadeep/x/pkg/env"
	"github.com/kmjayadeep/x/pkg/git"
	"github.com/kmjayadeep/x/pkg/kubeseal"
	"github.com/kmjayadeep/x/pkg/net"
	"github.com/kmjayadeep/x/pkg/notes"
	"github.com/kmjayadeep/x/pkg/pomo"
	"github.com/kmjayadeep/x/pkg/weather"
	"github.com/spf13/cobra"
)

func main() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var Cmd = &cobra.Command{
	Use:           "x",
	Short:         "command tree by JD",
	SilenceErrors: true,
}

func init() {
	Cmd.AddCommand(
		pomo.Cmd,
		git.Cmd,
		weather.Cmd,
		env.Cmd,
		net.Cmd,                // Network utilities
		clip.Cmd, clip.CopyCmd, // Clipboard - copy and paste
		notes.Cmd,
		date.Cmd, date.DateHeadCmd,
		kubeseal.Cmd,
	)
}
