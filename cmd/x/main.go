package main

import (
	"github.com/kmjayadeep/x/pkg/clip"
	"github.com/kmjayadeep/x/pkg/date"
	"github.com/kmjayadeep/x/pkg/env"
	"github.com/kmjayadeep/x/pkg/git"
	"github.com/kmjayadeep/x/pkg/net"
	"github.com/kmjayadeep/x/pkg/notes"
	"github.com/kmjayadeep/x/pkg/pomo"
	"github.com/kmjayadeep/x/pkg/weather"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

func main() {
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
		net.Cmd,  // Network utilities
		clip.Cmd, // Clipboard - copy and paste
		notes.Cmd,
		pomo.Cmd,
		date.Cmd,
	},
	Shortcuts: Z.ArgMap{
		// Git
		"pull":  {"git", "pull"},
		"push":  {"git", "push"},
		"pushf": {"git", "push", "force"},

		// Weather
		"weat": {"weather", "basic"},

		// Network
		"ip": {"net", "ip"},

		// Copy & Paste
		"c": {"clip", "copy"},
		"v": {"clip", "paste"},

		// Date
		"d":  {"date", "min"},
		"dt": {"date", "datetime"},
		"df": {"date", "full"},
	},
}
