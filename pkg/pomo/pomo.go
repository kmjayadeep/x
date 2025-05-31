package pomo

import (
	"time"

	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
	"github.com/rwxrob/bonzai/to"
	"github.com/rwxrob/bonzai/term"
	"github.com/rwxrob/bonzai/dtime"
	"github.com/rwxrob/bonzai/run"
	"github.com/rwxrob/bonzai/persisters/inprops"
)

var props = inprops.NewUserCache("x", "pomo.props")

var (
	Duration   = "30m"
	Warn       = "1m"
	Prefix     = "🍅"
	PrefixWarn = "💢"
	WarnTime   = 5 * time.Minute
)

var Cmd = &Z.Cmd{
	Name: `pomo`,
	Cmds: []*Z.Cmd{
		printCmd,
		help.Cmd,
		initCmd, startCmd, stopCmd,
	},
	Def: printCmd,
}

var initCmd = &Z.Cmd{
	Name:     `init`,
	Short:  `initialize pomo`,
	Cmds: []*Z.Cmd{help.Cmd},

	Do: func(x *Z.Cmd, _ ...string) error {
		return nil
	},
}

var printCmd = &Z.Cmd{
	Name:     `print`,
	Alias:  `show|p`,
	Cmds: []*Z.Cmd{help.Cmd},
	Short:  `print pomo status`,

	Do: func(x *Z.Cmd, _ ...string) error {

		started := props.Get("started")
		if started == "" {
			return nil
		}

		endt, err := time.Parse(time.RFC3339, started)
		if err != nil {
			return err
		}

		sec := time.Second
		left := endt.Sub(time.Now()).Round(sec)
		prefix := Prefix

		if left < WarnTime && left%(sec*2) == 0 {
			prefix = PrefixWarn
		}

		if left > 0 {
			term.Printf("%v%v", prefix, to.StopWatch(left))
			return nil
		}

		term.Printf("%v%v", prefix, "Pomo up!")

		notified := props.Get("notified")

		if notified == "" {
			if err := run.Exec("notify-send", "-u", "critical", "Pomo time up"); err != nil {
				return err
			}
			props.Set("notified", "1")
			return nil
		}

		return nil
	},
}

var startCmd = &Z.Cmd{
	Name:     `start`,
	Usage:    `[help|hour|DURATION]`,
	Cmds: []*Z.Cmd{help.Cmd},
	// Params:   []string{`hour`},
	MaxArgs:  1,

	Do: func(x *Z.Cmd, args ...string) error {
		if len(args) > 0 {
			if args[0] == `hour` {
				t := time.Now()
				args[0] = dtime.Until(dtime.NextHourOf, &t).String()
			}
			props.Set("duration", args[0])
		}
		s := props.Get("duration")
		if s == "" {
			s = Duration
		}
		dur, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		started := time.Now().Add(dur).Format(time.RFC3339)
		props.Set("notified","")
		props.Set("started", started)
		return nil
	},
}

var stopCmd = &Z.Cmd{
	Name:     `stop`,
	Cmds: []*Z.Cmd{help.Cmd},
	Short:  `stop pomo clock`,

	Do: func(x *Z.Cmd, args ...string) error {
		props.Set("started","")
		props.Set("notified","")
		return nil
	},
}
