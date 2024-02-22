package main

import (
	"github.com/kmjayadeep/x/pkg/git"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/pomo"
)

func main() {
	pomo.Duration = "30m"
	pomo.Interval = ""
	Cmd.Run()
}

var Cmd = &Z.Cmd{
	Name:    "x",
	Summary: "JD's bonzai command tree",
	Commands: []*Z.Cmd{
		help.Cmd,
		pomo.Cmd,
		git.Cmd,
	},
}
