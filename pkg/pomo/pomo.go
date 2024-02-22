package pomo

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var (
	Duration   = "30m"
	Interval   = "20s"
	Warn       = "1m"
	Prefix     = "🍅"
	PrefixWarn = "💢"
)

var Cmd = &Z.Cmd{
	Name: `pomo`,
	Commands: []*Z.Cmd{
		help.Cmd,
	},
}

