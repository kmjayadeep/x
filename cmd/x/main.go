package main

import (
	"github.com/kmjayadeep/x/pkg/clip"
	"github.com/kmjayadeep/x/pkg/date"
	"github.com/kmjayadeep/x/pkg/env"
	"github.com/kmjayadeep/x/pkg/git"
	"github.com/kmjayadeep/x/pkg/kubeseal"
	"github.com/kmjayadeep/x/pkg/net"
	"github.com/kmjayadeep/x/pkg/notes"
	"github.com/kmjayadeep/x/pkg/pomo"
	"github.com/kmjayadeep/x/pkg/weather"
	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

func main() {
	Cmd.Exec()
}

var Cmd = &Z.Cmd{
	Name:    "x",
	Short: "bonzai command tree by JD",
	Cmds: []*Z.Cmd{
		help.Cmd,
		pomo.Cmd,
		git.Cmd,
		weather.Cmd,
		env.Cmd,
		net.Cmd,  // Network utilities
		clip.Cmd, clip.CopyCmd, // Clipboard - copy and paste
		notes.Cmd,
		pomo.Cmd,
		date.Cmd, date.DateHeadCmd,
		kubeseal.Cmd,
	},
}
