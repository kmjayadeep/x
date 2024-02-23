package main

import (
	"github.com/kmjayadeep/x/pkg/git"
	"github.com/kmjayadeep/x/pkg/env"
	"github.com/kmjayadeep/x/pkg/weather"
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
		weather.Cmd,
		env.Cmd,
	},
	Shortcuts: Z.ArgMap{
		// Git
		"pull":  {"git", "pull"},
		"push":  {"git", "push"},
		"pushf": {"git", "push", "force"},

		// Weather
		"weat": {"weather", "basic"},
	},
}
