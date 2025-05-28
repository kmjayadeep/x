package date

import (
	"fmt"
	"time"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var Cmd = &Z.Cmd{
	Name:     `date`,
	Short:  `Date util commands`,
	Cmds: []*Z.Cmd{help.Cmd, dateCmd, dateTimeCmd, dateFull, dateHeadCmd},
}

var dateCmd = &Z.Cmd{
	Name:     `min`,
	Short:  `Display date in YYY-MM-DD format`,
	Cmds: []*Z.Cmd{help.Cmd},
	Do: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006/01/02")
		fmt.Println(d)
		return nil
	},
}

var dateFull = &Z.Cmd{
	Name:     `full`,
	Short:  `Display date in YYY MMM DD format`,
	Cmds: []*Z.Cmd{help.Cmd},
	Do: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006 Jan 02")
		fmt.Println(d)
		return nil
	},
}

var dateTimeCmd = &Z.Cmd{
	Name:     `datetime`,
	Short:  `Display date and time in YYY-MM-DD HH:MM format`,
	Cmds: []*Z.Cmd{help.Cmd},
	Do: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006/01/02 03:04PM")
		fmt.Println(d)
		return nil
	},
}

var dateHeadCmd = &Z.Cmd{
	Name:     `head`,
	Short:  `Display date in a human readable form for headings`,
	Cmds: []*Z.Cmd{help.Cmd},
	Do: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006 January 02, Monday")
		fmt.Println(d)
		return nil
	},
}
