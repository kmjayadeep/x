package date

import (
	"fmt"
	"time"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var Cmd = &Z.Cmd{
	Name:     `date`,
	Summary:  `Date util commands`,
	Commands: []*Z.Cmd{help.Cmd, dateCmd, dateTimeCmd, dateFull, dateHeadCmd},
}

var dateCmd = &Z.Cmd{
	Name:     `min`,
	Summary:  `Display date in YYY-MM-DD format`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006/01/02")
		fmt.Println(d)
		return nil
	},
}

var dateFull = &Z.Cmd{
	Name:     `full`,
	Summary:  `Display date in YYY MMM DD format`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006 Jan 02")
		fmt.Println(d)
		return nil
	},
}

var dateTimeCmd = &Z.Cmd{
	Name:     `datetime`,
	Summary:  `Display date and time in YYY-MM-DD HH:MM format`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006/01/02 03:04PM")
		fmt.Println(d)
		return nil
	},
}

var dateHeadCmd = &Z.Cmd{
	Name:     `head`,
	Summary:  `Display date in a human readable form for headings`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006 January 02, Monday")
		fmt.Println(d)
		return nil
	},
}
