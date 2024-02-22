package main

import (
	// "github.com/kmjayadeep/x/pkg/pomo"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/pomo"
	"github.com/rwxrob/help"
)

func main() {
  pomo.Duration = "30m"
  pomo.Interval = ""
  Cmd.Run()
}

var Cmd = &Z.Cmd{
  Name: "x",
  Summary: "JD's bonzai command tree",
  Commands: []*Z.Cmd{
    help.Cmd,
    pomo.Cmd,
  },
}
