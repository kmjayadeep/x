package date

import (
	"fmt"
	"time"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
)

var Cmd = &Z.Cmd{
	Name:     `date`,
	Short:  `date util commands`,
	Alias: `d`,
	Def: dateCmd,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden(), dateCmd, dateTimeCmd, dateFull, DateHeadCmd},
}

var dateCmd = &Z.Cmd{
	Name:     `min`,
	Short:  `display date in YYY-MM-DD format`,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden()},
	Do: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006/01/02")
		fmt.Println(d)
		return nil
	},
}

var dateFull = &Z.Cmd{
	Name:     `full`,
	Short:  `display date in YYY MMM DD format`,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden()},
	Do: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006 Jan 02")
		fmt.Println(d)
		return nil
	},
}

var dateTimeCmd = &Z.Cmd{
	Name:     `datetime`,
	Short:  `display date and time in YYY-MM-DD HH:MM format`,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden()},
	Do: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006/01/02 03:04PM")
		fmt.Println(d)
		return nil
	},
}

var DateHeadCmd = &Z.Cmd{
	Name:     `datehead`,
	Short:  `display date in a human readable form for headings`,
	Alias: `dh`,
	Cmds: []*Z.Cmd{help.Cmd.AsHidden()},
	Do: func(_ *Z.Cmd, args ...string) error {
		d := time.Now().Format("2006 January 02, Monday")
		fmt.Println(d)
		return nil
	},
}
